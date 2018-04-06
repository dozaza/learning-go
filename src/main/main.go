package main

import (
	"fmt"
	"unicode/utf8"
	"strconv"
	"container/list"
	"os"
	"bufio"
	"io"
	"runtime"
)

func main() {
	// 1.1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 1.2
	loop(0)

	// 1.3
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range a {
		fmt.Println(i)
	}

	// 2.1
	fizzBuzz()

	// 3.1
	printA()

	// 3.2
	str := "asSASA ddd dsjkdsjs dk"
	charArray := []byte(str)
	fmt.Printf("%d, %d\n", len(charArray), utf8.RuneCount(charArray))

	// 3.3
	newStr := []rune(str)
	copy(newStr[4:], []rune{'a', 'b', 'c'})
	fmt.Printf("%s\n", string(newStr))

	// 3.4
	reversePrint("foobar")

	// 4.1
	numbers := []float64{1.1, 1.2, 1.3}
	if len(numbers) > 0 {
		avg := 0.0
		for _, num := range numbers {
			avg += num
		}
		avg /= float64(len(numbers))
		fmt.Println(avg)
	}

	// 5.0
	nums_5 := []float64 {1.0, 1.1, 1.2}
	fmt.Println(avg(nums_5))

	// 6.0
	fmt.Println(sortInt(7, 2))

	// 7.1
	// 第9行 i 不存在

	// 8.1
	var s stack
	s.push(5)
	s.push(4)
	s.push(3)
	fmt.Println(s.pop())
	s.push(2)
	// 8.2
	fmt.Println(s)

	// 9.1
	printArguments(1, 2, 3)

	// 10.1
	fmt.Println(fibonacci(5))

	// 11.1
	fmt.Println(myMap(func(i int) int { return i * 2}, []int{1, 2, 3}))
	// 11.2
	fmt.Println(myMap2(func(s string) string { return s + "_new"}, []string{"a", "b"}))

	// 12.1
	fmt.Println(max([]int{3,1,2}))
	fmt.Println(min([]int{3,1,2}))

	// 13
	nums_13 := []int{3,1,2}
	bubbleSort(nums_13)
	fmt.Println(nums_13)

	// 14.1
	p := plusTwo()
	fmt.Printf("%v\n", p(2))
	// 14.2
	p2 := plusX(5)
	fmt.Printf("%v\n", p2(2))

	// 17.1
	// All types support operation +
	// 17.2
	// Not all types support operation +

	// 18
	fmt.Println(myMap3(func(i T) T { return i.(int)* 2}, []T{1, 2, 3}))
	fmt.Println(myMap3(func(i T) T { return i.(string) + "_new"}, []T{"a", "b"}))

	// 19.1
	// p1 is value, p2 is pointer
	// 19.2
	// first x is a pointer of what t is pointing, second x is a pointer of local variable t which is a copy

	// 20.1
	var l list.List
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Print(strconv.Itoa(i.Value.(int)) + ", ")
	}
	fmt.Println()
	// 20.2
	var l2 myList
	l2.pushBack(1)
	l2.pushBack(2)
	l2.pushBack(4)

	for i := l2.front(); i != nil; i = i.next {
		fmt.Print(strconv.Itoa(i.value.(int)) + ", ")
	}
	fmt.Println()

	// 21.1
	simpleCat("./src/main/test.txt")
	// 21.2
	cat("./src/main/test.txt", "n")

	// 22.1
	// k1: IntVector, k2: *IntVector, k3: *IntVector
	// 22.2
	// cause it will try to find *IntVector automatically

	// 23.1
	// type conversion check is in the runtime
	// Before conversion, we can check with if v, ok := i.(MyType); ok {}

	// 24.1
	// reflect works only on pointer, cause funcation argument is parsed by value

	// 25.2
	si := SliceInt{1,3,2}
	sf := SliceFloat64{1.1,3.3,2.2}
	fmt.Println(Max(si))
	fmt.Println(Max(sf))

	runtime.GOMAXPROCS(4)
	// 26
	c := make(chan int)
	done := make(chan bool)
	go display(c, done)
	for i := 0; i < 10; i++ {
		c <- i
	}
	close (c)
	<- done
	
}

func display(c chan int, done chan bool) {
	for {
		select {
		case i, more := <- c:
			if more {
				fmt.Print(strconv.Itoa(i) + " ")
			} else {
				fmt.Println("")
				done <- true
				break
			}
		}
	}
}

type SliceInt []int
type SliceFloat64 []float64

type BaseSlice interface {
	Compare(d1 interface{}, d2 interface{}) int
	Len() int
	Get(i int) interface{}
}

func Max(s BaseSlice) interface{} {
	if s.Len() == 0 {
		return nil
	}

	max := s.Get(0)
	for i := 1; i < s.Len(); i++ {
		if s.Compare(max, s.Get(i)) == -1 {
			max = s.Get(i)
		}
	}

	return max
}

func (s SliceInt) Compare(d1 interface{}, d2 interface{}) int {
	switch d1.(type) {
	case int:
		switch d2.(type) {
		case int:
			i1 := d1.(int)
			i2 := d2.(int)
			if i1 == i2 {
				return 0
			} else if i1 < i2 {
				return -1
			} else if i1 > i2 {
				return 1
			}
		}
	}
	panic("Wrong type")
}

func (s SliceInt) Len() int {
	return len(s)
}

func (s SliceInt) Get(i int) interface{} {
	return s[i]
}

func (s SliceFloat64) Compare(d1 interface{}, d2 interface{}) int {
	switch d1.(type) {
	case float64:
		switch d2.(type) {
		case float64:
			i1 := d1.(float64)
			i2 := d2.(float64)
			if i1 == i2 {
				return 0
			} else if i1 < i2 {
				return -1
			} else if i1 > i2 {
				return 1
			}
		}
	}
	panic("Wrong type")
}

func (s SliceFloat64) Len() int {
	return len(s)
}

func (s SliceFloat64) Get(i int) interface{} {
	return s[i]
}

func cat(fileName string, option string) {
	file, e := os.Open(fileName)
	if e != nil {
		fmt.Printf("unable to read file: %s", e.Error())
		return // exit the function on error
	}
	defer fmt.Println("")
	defer file.Close()

	reader := bufio.NewReader(file)
	num := 1
	for {
		str, e := reader.ReadString('\n')
		if option == "n" {
			fmt.Printf("%d: ", num)
		}
		fmt.Print(str)
		num++

		if e == io.EOF {
			break
		}
	}
}

func simpleCat(fileName string) {
	file, e := os.Open(fileName)
	if e != nil {
		fmt.Printf("unable to read file: %s", e.Error())
		return // exit the function on error
	}
	defer fmt.Println("")
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, e := reader.ReadString('\n')
		fmt.Print(str)
		if e == io.EOF {
			break
		}
	}
}

type myList struct {
	first *myElement
	last *myElement
	size int
}

func (l *myList) front() *myElement {
	return l.first
}

func (l *myList) back() *myElement {
	return l.last
}

func (l *myList) pushBack(v interface{}) {
	newElement := myElement{value: v}
	if l.first == nil {
		l.first = &newElement
		l.last = &newElement
	} else {
		l.last.next = &newElement
		l.last = &newElement
	}
}

type myElement struct {
	value interface{}
	prev *myElement
	next *myElement
}

type T interface {

}

func myMap3(f func(T) T, array []T) []T {
	tmp := make([]T, len(array))
	for i, d := range array {
		tmp[i] = f(d)
	}
	return tmp
}

func plusX(x int) func(int) int {
	return func(i int) int {
		return i + x
	}
}

func plusTwo() func(int) int {
	return func(i int) int {
		return i + 2
	}
}

func bubbleSort(nums []int) {
	for {
		swapped := false
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func min(nums []int) int {
	if len(nums) == 0 {
		panic("empty nums")
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}

	return res
}

func max(nums []int) int {
	if len(nums) == 0 {
		panic("empty nums")
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}

	return res
}

func myMap2(f func(string) string, array []string) []string {
	tmp := make([]string, len(array))
	for i, d := range array {
		tmp[i] = f(d)
	}
	return tmp
}

func myMap(f func(int) int, array []int) []int {
	tmp := make([]int, len(array))
	for i, d := range array {
		tmp[i] = f(d)
	}
	return tmp
}

func fibonacci(num int) (res []int) {
	switch num {
	case 0:
		// Do nothing
	case 1:
		res = make([]int, num)
		res[0] = 1
	default:
		res = make([]int, num)
		res[0] = 1
		res[1] = 1
		for i := 3; i <= num; i++ {
			res[i - 1] = res[i - 2] + res[i - 3]
		}
	}

	return res
}

func printArguments(nums ...int) {
	for _, num := range nums {
		fmt.Printf("%d, ", num)
	}
	fmt.Println("")
}

type stack struct {
	size int
	data [5]int
}

func (s stack) String() string {
	if s.size == 0 {
		return ""
	}

	str := ""
	for i, d := range s.data {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(d) + "] "
	}
	return str
}

func (s *stack) push(num int) {
	if s.size < 5 {
		s.data[s.size] = num
		s.size += 1
	}
}

func (s *stack) pop() int {
	if s.size == 0 {
		panic("empty stack")
	}

	s.size -= 1
	return s.data[s.size]
}

func sortInt(a int, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

func avg(nums []float64) float64 {
	if len(nums) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func loop(i int) {
Here:
	fmt.Println(i)
	i++
	if i < 10 {
		goto Here
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		is3Module := i % 3 == 0
		is5Module := i % 5 == 0
		if is3Module && is5Module {
			fmt.Print("fizzBuzz, ")
		} else if is3Module {
			fmt.Println("fizz, ")
		} else if is5Module {
			fmt.Println("buzz, ")
		}
	}
}

func printA() {
	for i := 1; i < 100; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Println("")
	}
}

func reversePrint(str string) {
	temp := []rune(str)
	for i, j := 0, len(temp) - 1; i < j; i, j = i + 1, j - 1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	fmt.Printf("%s\n", string(temp))
}