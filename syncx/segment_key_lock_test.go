package syncx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSegmentKeysLock_Lock(t *testing.T) {
	l := NewSegmentKeysLock(8)
	key1 := "key1"
	l.Lock(key1)
	// 必然加锁失败
	assert.False(t, l.TryLock(key1))
	// 读锁也失败
	assert.False(t, l.TryRLock(key1))
	key2 := "key2"
	// 加锁成功
	assert.True(t, l.TryLock(key2))
	// 解锁不会触发 panic
	defer l.Unlock(key2)

	// 释放锁
	l.Unlock(key1)
	// 此时应该预期自己可以再次加锁
	assert.True(t, l.TryLock(key1))
}

func TestNewSegmentKeysLock_RLock(t *testing.T) {
	l := NewSegmentKeysLock(8)
	key1, key2 := "key1", "key2"
	l.RLock(key1)
	// 必然加锁失败
	assert.False(t, l.TryLock(key1))
	// 读锁可以成功
	assert.True(t, l.TryRLock(key1))
	// 加锁成功
	assert.True(t, l.TryRLock(key2))
	// 解锁不会触发 panic
	defer l.RUnlock(key2)

	// 释放读锁
	l.RUnlock(key1)
	// 此时还有一个读锁没有释放
	assert.False(t, l.TryLock(key1))
	// 再次释放读锁
	l.RUnlock(key1)
	assert.True(t, l.TryLock(key1))
}
