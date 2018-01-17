package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)


func main() {
	fmt.Println("It is currently:",time.Now())
	fmt.Println("A number from 1-100", rand.Intn(100))
	fmt.Println("The square root of 4 is",math.Sqrt(4))

}