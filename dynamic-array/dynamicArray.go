package dynamic_array

import (
	"errors"
	"fmt"
)

var defaultCapacity = 10

type DynamicArray struct {
	size        int
	capacity    int
	elementData []interface{}
}

func (array *DynamicArray) Get(index int) (interface{}, error) {
	err := array.CheckRangeFromIndex(index)
	if err != nil {
		return nil, err
	}
	return array.elementData[index], nil
}

func (array *DynamicArray) Put(index int, element interface{}) error {
	err := array.CheckRangeFromIndex(index)
	if err != nil {
		return err
	}
	array.elementData[index] = element
	return nil
}

func (array *DynamicArray) Add(element interface{}) {
	if array.size == array.capacity {
		//resize
		array.Resize()
	}
	array.elementData[array.size] = element
	array.size++
}

func (array *DynamicArray) Remove(index int) error {
	err := array.CheckRangeFromIndex(index)
	if err != nil {
		return err
	}
	copy(array.elementData[index:], array.elementData[index+1:array.size])
	array.elementData[array.size-1] = nil
	array.size--
	return nil
}

func (array *DynamicArray) GetData() []interface{} {
	return array.elementData[:array.size]
}

func (array *DynamicArray) IsEmpty() bool {
	return array.size == 0
}

func (array *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= array.size || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

func (array *DynamicArray) Resize() {
	if array.capacity == 0 {
		array.capacity = defaultCapacity
	} else {
		array.capacity = array.capacity << 1
	}
	newArray := make([]interface{}, array.capacity)
	copy(newArray, array.elementData)
	array.elementData = newArray
}

func main() {
	numbers := DynamicArray{}
	fmt.Println(numbers.IsEmpty())

	numbers.Add(10)
	numbers.Add(20)
	numbers.Add(30)
	numbers.Add(40)
	numbers.Add(50)

	fmt.Println(numbers.IsEmpty())

	fmt.Println(numbers.GetData())

	numbers.Remove(1)

	fmt.Println(numbers.GetData())

	numberFound, _ := numbers.Get(1)
	fmt.Println(numberFound)

	numbers.Put(0, 100)
	fmt.Println(numbers.GetData())
}
