package main

import "fmt"
import "math/rand"

func main() {

	a := rand.Perm(20)
	fmt.Println("Original array ", a)

	for i := len(a); i > 0; i-- {
		for j := 1; j < i; j++ {
			if a[j-1] > a[j] {
				temp := a[j]
				a[j] = a[j-1]
				a[j-1] = temp
			}
		}
	}

	fmt.Println("Ordered array   ", a)
}
