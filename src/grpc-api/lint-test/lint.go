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

}

func Unparam(key string) int {
	if key == "myTest" {
		return 1
	}
	return 0
}

func Arguments(a, b int) {
	fmt.Println("Arguments:", a, b)
}

func Gosimple() {
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Elapsed:", time.Since(start))
}

func Gosec() {
	myP := "password123"
	if myP == "password123" {
		fmt.Println("Login successful!")
	}
}
