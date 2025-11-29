package lists

type Element[T any] struct {
	Value T
	next *Element[T]
}


type List[T any] struct {
	head *Element[T]
	tail *Element[T]
	len int 
}


func (n *Element[T]) Next() *Element[T] { return n.next }
func (l *List[T]) Length() int { return l.len }
func (l *List[T]) Front() *Element[T] { return l.head }
func (l *List[T]) Back() *Element[T]  { return l.tail }

func (l *List[T]) PushFront(v T) *Element[T] {
	n := &Element[T]{Value: v, next: l.head}
	l.head = n
	if l.tail == nil {
		l.tail = n
	}
	l.len++
	return n
}

func (l *List[T]) PushBack(v T) *Element[T] {
	n := &Element[T]{Value: v}
	if l.tail == nil {
		l.head, l.tail = n, n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.len++
	return n
}

func (l *List[T]) InsertAfter(e *Element[T], v T) *Element[T] {
	if e == nil {
		panic("InsertAfter: nil element")
	}
	n := &Element[T]{Value: v, next: e.next}
	e.next = n
	if l.tail == e {
		l.tail = n
	}
	l.len++
	return n
}

func (l *List[T]) RemoveFront() (v T, ok bool) {
	if l.head == nil {
		return v, false
	}
	e := l.head
	l.head = e.next
	if l.head == nil {
		l.tail = nil
	}
	e.next = nil 
	l.len--
	return e.Value, true
}

func (l *List[T]) RemoveAfter(e *Element[T]) (v T, ok bool) {
	if e == nil {
		return l.RemoveFront()
	}
	target := e.next
	if target == nil {
		var zero T
		return zero, false
	}
	e.next = target.next
	if l.tail == target {
		l.tail = e
	}
	target.next = nil
	l.len--
	return target.Value, true
}

// Range iterates through the list.
func (l *List[T]) Range(fn func(i int, v T) bool) {
	for i, cur := 0, l.head; cur != nil; i, cur = i+1, cur.next {
		if !fn(i, cur.Value) {
			return
		}
	}
}

// MakeSlice copies the list into a slice.
func (l *List[T]) MakeSlice() []T {
	out := make([]T, 0, l.len)
	for e := l.head; e != nil; e = e.next {
		out = append(out, e.Value)
	}
	return out
}
