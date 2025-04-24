package linttest

import (
	"fmt"
	"time"
)

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

func Unparam(key string) int {
	str := "myTest"

	if str == "myTest" {
		return 1
	}
	return 0
}

func Arguments(a, b, c, d, e int) {
	fmt.Println("Arguments:")
}

func Gosimple() {
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Elapsed:", time.Now().Sub(start))
}

func Gosec() {
	password := "password123" // G101: Hardcoded credentials
	if password == "password123" {
		fmt.Println("Login successful!")
	}
}
