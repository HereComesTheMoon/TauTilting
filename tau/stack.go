package tau

import "github.com/gammazero/deque"
import "sync"
import "errors"

// WrappedStack. Use deque + mutex to make it threadsafe.
type wStack[T any] struct {
    dq deque.Deque[T]
    mutex sync.Mutex
}

func newWStack[T any](size ...int) *wStack[T] {
    ws := wStack[T]{
        dq: *deque.New[T](size...),
        mutex: sync.Mutex{},
    }
    return &ws
}

func (ws *wStack[T]) push(elem T) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    ws.dq.PushBack(elem)
}

func (ws *wStack[T]) pushFront(elem T) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    ws.dq.PushFront(elem)
}

func (ws *wStack[T]) pushMany(elems *[]T) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    for _, elem := range(*elems) {
        ws.dq.PushBack(elem)
    }
}

func (ws *wStack[T]) pop() (T, error) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    if ws.dq.Len() == 0 {
        return *new(T), errors.New("Stack is empty.")
    } else {
        return ws.dq.PopBack(), nil
    }
}

func (ws *wStack[T]) popFront() (T, error) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    if ws.dq.Len() == 0 {
        return *new(T), errors.New("Stack is empty.")
    } else {
        return ws.dq.PopFront(), nil
    }
}

func (ws *wStack[T]) empty() bool {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    return ws.dq.Len() == 0
}

