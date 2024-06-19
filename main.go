package main

import (
	"fmt"
	"superman/module"
)

func main() {
	result := module.SupermanRescueChicken(5, 5, []uint{2, 5, 10, 12, 15, 16, 17, 18, 18, 20})
	fmt.Println(result)
}
