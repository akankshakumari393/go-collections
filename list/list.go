package list

import (
	"github.com/akankshakumari393/go-collections/collection"
)

type List interface {
	collection.Collection
	// append is used to add an element at the end of the List.
	Append(objects ...interface{})
	// addTo is used to add an element at a specific index in the List
	AddTo(index int, object interface{})
	// Set is used update the element at a given position
	Set(index int, object interface{})
	// Remove an object from the List. If there are multiple such objects, then the first occurrence of the object is removed.
	Remove(Object interface{})
	// RemoveAt takes an integer value which simply removes the element present at that specific index in the List. After removing the element, all the elements are moved to the left to fill the space and the indices of the objects are updated.
	RemoveAt(index int)
	// Get gets the value of the list at a specific index
	Get(index int) interface{}
}
