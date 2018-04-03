package main

import (
	"fmt"
	"unicode/utf8"
	"strconv"
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