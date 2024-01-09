package queue

type Queue interface {
	Len() int
	Cap() int
	IsFull() bool
	IsEmpty() bool
	Enqueue(val any) error
	Dequeue() (any, error)
	Peek() (any, error)
}
