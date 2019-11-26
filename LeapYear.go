package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	years := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &years[i])
	}

	for i := range years {
		if years[i]%400 == 0 {
			fmt.Printf("%d is a leap year\n", years[i])
		} else if years[i]%100 == 0 {
			fmt.Printf("%d is not a leap year\n", years[i])
		} else if years[i]%4 == 0 {
			fmt.Printf("%d is a leap year\n", years[i])
		} else {
			fmt.Printf("%d is not a leap year\n", years[i])
		}
	}
}
