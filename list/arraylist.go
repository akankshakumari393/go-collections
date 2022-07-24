package list

type arrayList struct {
	items []interface{}
}

func NewArrayList() List {
	return &arrayList{
		items: []interface{}{},
	}
}

func (al *arrayList) Size() int {
	return 0
}

func (al *arrayList) IsEmpty() bool {
	return true
}

func (al *arrayList) Clean() {

}

func (al *arrayList) Append(objects ...interface{}) {

}

func (al *arrayList) AddTo(index int, object interface{}) {

}

func (al *arrayList) Set(index int, object interface{}) {

}

func (al *arrayList) Remove(objects interface{}) {

}

func (al *arrayList) RemoveAt(index int) {

}

func (al *arrayList) Get(index int) interface{} {
	return 0
}
