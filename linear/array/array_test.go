package array

import (
	"fmt"
	"testing"
	"unsafe"
)

type testStruct struct {
	next *testStruct
	str string
	c []int
}

func TestString(t *testing.T) {
	//a := 5
	//b := &a
	//c := &a
	//d := 5
	//e := 10
	//fmt.Printf("%p, %p, %p, %p, %p, %p, %p", &a, b, c, &b, &c, &d, &e)

	//t1 := &testStruct{
	//	next: nil,
	//	str: "111",
	//	c: []int{1},
	//}
	//t2 := &testStruct{
	//	next: t1,
	//	str: "222",
	//	c: []int{2},
	//}
	//t3 := t2
	//t4 := &testStruct{
	//	next: t1,
	//	str: "222",
	//	c: []int{2},
	//}
	//fmt.Println(t2 == t3, t4 == t3, t2.next == t1)

	var str1 string
	str2 := new(string)
	fmt.Printf("%p, %p, %p, %p, [%v], [%v], [%v]" , &str1, str2, &str2, &(*str2), str1, str2, *str2)
}

func TestArrayList(t *testing.T) {

	pri := []interface{}{1,2,3,4,5}
	list := Init(10)
	list.Append(pri...)
	list.RemoveFirst()
	list.Insert(0, 100)
	list.Insert(0,1000)
	list.Insert(0, 10000)
	list.Append(9)
	list.Append(99)
	list.Append(88)
	list.Append(888)
	list.Append(777)
	list.Append(666)
	list.Remove(99)
	list.RemoveAt(3)
	list.Set(5, 55555)

	fmt.Println(list)

	new := list.SubArray([]int{1,3,5,7}...)
	fmt.Println(new)
}


func TestGoArray(t *testing.T) {

	a := [5]string{"aaa", "bbb", "a", "ddddddddd", "eeeeee"}
	a[0] = "hello arraywqeweeqeqew"
	fmt.Printf("%p, %p, %p, %p", &a, &a[0], &a[1], &a[2])
	fmt.Println()
	//s := []int{1,2,3}
	//f := func(a int, b string) {
	//	fmt.Println(a, b)
	//}
	//c := make(chan string, 1)
	mm := map[string]interface{}{
		"111": 111,
		"222": "222",
		"333": 333.33,
		"444": []int{4,4,4,4},
	}
	size := unsafe.Sizeof(mm)
	fmt.Println(size)
}


func TestGoSlice(t *testing.T) {

	a := []string{"aaa", "bbb", "a", "ddddddddd", "eeeeee"}
	a[0] = "hello arraywqeweeqeqew"
	fmt.Printf("%p, %p, %p, %p\n", &a, &a[0], &a[1], &a[2])
	fmt.Println(cap(a), len(a))
	// slice 扩容后，底层数组会拷贝
	a = append(a, "111", "222", "eee", "ffff", "24242")
	fmt.Printf("%p, %p, %p, %p\n", &a, &a[0], &a[1], &a[2])
}
