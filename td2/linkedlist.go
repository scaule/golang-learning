package td2

type Linkedlist struct {
	e *Elem
}

type Elem struct {
	v    string
	next *Elem
}

func CreateLinkedList(v string) *Linkedlist {
	return &Linkedlist{
		&Elem{
			v:    v,
			next: nil,
		},
	}
}

func (l *Linkedlist) Add(v string) bool {
	// Add a new element to linked list at the end.
	return false
}

func (l *Linkedlist) Remove(i int) bool {
	// Remove element in index i.
	return false
}

func (l *Linkedlist) Print() string {
	// Print linkedlist like "1 2 3 4".
	return ""
}
