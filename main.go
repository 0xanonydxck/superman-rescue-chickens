package main

import (
	"fmt"
	"superman/module"
)

func main() {
	result := module.SupermanRescueChicken(5, 5, []uint{2, 5, 10, 12, 15})
	fmt.Println(result)
}
