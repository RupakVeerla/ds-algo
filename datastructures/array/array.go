package array

type New struct {
	length int
	data   map[int]interface{}
}

func (a *New) Get(index int) interface{} {
	return a.data[index]
}

func (a *New) Append(item interface{}) int {
	if a.data == nil {
		a.data = make(map[int]interface{})
	}
	a.data[a.length] = item
	a.length++
	return a.length
}

func (a *New) Delete() (val interface{}) {
	val = a.data[a.length-1]
	delete(a.data, a.length-1)
	a.length--
	return
}

func (a *New) DeleteAtIndex(index int) (val interface{}) {
	val = a.data[index]
	a.shiftValues(index)
	return
}

func (a *New) shiftValues(index int) {
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	delete(a.data, a.length-1)
	a.length--
}

func (a *New) Length() int {
	return a.length
}
