package linttest

import "fmt"

func GoCyclo() {
	age := 10
	if age > 1 {
		fmt.Println("GoCyclo test")
	}
	if age > 2 {
		fmt.Println("GoCyclo test")
	}
	if age > 3 {
		fmt.Println("GoCyclo test")
	}
	if age > 4 {
		fmt.Println("GoCyclo test")
	}
	if age > 5 {
		fmt.Println("GoCyclo test")
	}
	if age > 6 {
		fmt.Println("GoCyclo test")
	}
	if age > 7 {
		fmt.Println("GoCyclo test")
	}
	if age > 8 {
		fmt.Println("GoCyclo test")
	}
	if age > 9 {
		fmt.Println("GoCyclo test")
	}
	if age > 10 {
		fmt.Println("GoCyclo test")
	}
	if age > 11 {
		fmt.Println("GoCyclo test")
	}

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println("Even:", i)
		} else {
			fmt.Println("Odd:", i)
		}
	}

}

func Funlen(a int, b int, c int, d int, e int, f int) {
	fmt.Println("Funlen test")
}

func Unparam(a int) {
	fmt.Println("Funlen test")
}
