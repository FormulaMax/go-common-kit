package queue

import (
	"context"
	"fmt"
	"github.com/FormulaMax/go-common-kit/internal/queue"
	"sync"
	"time"
)

type Delayable interface {
	Delay() time.Duration
}

// DelayQueue 延时队列
// 每次出队的元素必然都是已经到期的元素，即 Delay() 返回的值小于等于 0
// 延时队列本身对时间的精确度并不是很高，其时间精确度主要取决于 time.Timer
// 所以如果需要极度精确的延时队列，那么这个结构并不太适合。
// 但是如果能够容忍至多在毫秒级的误差，那么这个结构还是可以使用的
type DelayQueue[T Delayable] struct {
	q             queue.PriorityQueue[T]
	mutex         *sync.Mutex
	dequeueSignal *cond
	enqueueSignal *cond
}

func NewDelayQueue[T Delayable](size int) *DelayQueue[T] {
	m := &sync.Mutex{}
	res := &DelayQueue[T]{
		q: *queue.NewPriorityQueue[T](size, func(src, dst T) int {
			srcDelay := src.Delay()
			dstDelay := src.Delay()
			if srcDelay > dstDelay {
				return 1
			}
			if srcDelay == dstDelay {
				return 0
			}
			return -1
		}),
		mutex:         m,
		dequeueSignal: newCond(m),
		enqueueSignal: newCond(m),
	}
	return res
}

func (dq *DelayQueue[T]) Enqueue(ctx context.Context, t T) error {
	for {
		select {
		//先检测ctx过期
		case <-ctx.Done():
			return ctx.Err()
		default:

		}
		dq.mutex.Lock()
		err := dq.q.Enqueue(t)
		switch err {
		case nil:
			dq.enqueueSignal.broadcast()
			return nil
		case queue.ErrOutOfCapacity:
			signal := dq.dequeueSignal.signalCh()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-signal:
			}
		default:
			dq.mutex.Unlock()
			return fmt.Errorf("go-common-kit:延时队列入队的时候遇到未知错误 %w，请上报", err)
		}
	}
}

func (dq *DelayQueue[T]) Dequeue(ctx context.Context) (T, error) {
	var timer *time.Timer
	defer func() {
		if timer != nil {
			timer.Stop()
		}
	}()
	for {
		select {
		case <-ctx.Done():
			var t T
			return t, ctx.Err()
		default:
		}
		dq.mutex.Lock()
		val, err := dq.q.Peek()
		switch err {
		case nil:
			delay := val.Delay()
			if delay <= 0 {
				val, err = dq.q.Dequeue()
				dq.dequeueSignal.broadcast()
				// err应该不可能不为nil
				return val, err
			}
			signal := dq.enqueueSignal.signalCh()
			if timer == nil {
				timer = time.NewTimer(delay)
			} else {
				timer.Reset(delay)
			}
			select {
			case <-ctx.Done():
				var t T
				return t, ctx.Err()
			case <-timer.C:
				// 时间到了
				dq.mutex.Lock()
				// 原队头可能已经被其他协程先出队，所以再次检查队头
				val, err := dq.q.Peek()
				if err != nil || val.Delay() > 0 {
					dq.mutex.Unlock()
					continue
				}
				// 验证元素过期后将其出队
				val, err = dq.q.Dequeue()
				dq.dequeueSignal.broadcast()
				return val, err
			case <-signal:
				// 进入下一个循环，这里可能是新的元素入队，也可能是过期了
			}
		case queue.ErrEmptyQueue:
			signal := dq.enqueueSignal.signalCh()
			select {
			case <-ctx.Done():
				var t T
				return t, ctx.Err()
			case <-signal:
			}
		default:
			dq.mutex.Unlock()
			var t T
			return t, fmt.Errorf("go-common-kit:延时队列出队的时候遇到未知错误 %w，请上报", err)
		}
	}
}

type cond struct {
	signal chan struct{}
	l      sync.Locker
}

func newCond(l sync.Locker) *cond {
	return &cond{
		signal: make(chan struct{}),
		l:      l,
	}
}

// broadcast 唤醒等待者
// 如果没有人等待，那么什么也不会发生
// 必须加锁之后才能调用这个方法
// 广播之后锁会被释放，这也是为了确保用户必然是在锁范围内调用的
func (c *cond) broadcast() {
	signal := make(chan struct{})
	old := c.signal
	c.signal = signal
	c.l.Unlock()
	close(old)
}

// signalCh 返回一个 channel，用于监听广播信号
// 必须在锁范围内使用
// 调用后，锁会被释放，这也是为了确保用户必然是在锁范围内调用的
func (c *cond) signalCh() <-chan struct{} {
	res := c.signal
	c.l.Unlock()
	return res
}
