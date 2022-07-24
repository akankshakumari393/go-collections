package collection

// Collection interface to be implemented by all the data structure
type Collection interface {
	Size() int
	IsEmpty() bool
	Clean()
}
