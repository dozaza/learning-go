package main

import (
	"fmt"
	"unicode/utf8"
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