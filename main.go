package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"superman/module"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

func inputProtectLimitAndRoofLength() (uint, uint) {
	answer, err := prompt.New().
		Ask("Enter protectLimit and roofLength (ex. 5 5):").
		Input("5 5", input.WithValidateFunc(func(s string) error {
			isEmpty := strings.TrimSpace(s) == ""
			if isEmpty {
				return fmt.Errorf("please enter a value")
			}

			parts := strings.Split(s, " ")
			if len(parts) != 2 {
				return fmt.Errorf("please enter 2 values")
			}

			return nil
		}))

	if err != nil {
		log.Panic(err)
	}

	parts := strings.Split(answer, " ")
	protectLimit, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		log.Panic(err)
	}

	roofLength, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		log.Panic(err)
	}

	return uint(protectLimit), uint(roofLength)
}

func inputChickenPositions() []uint {
	answer, err := prompt.New().
		Ask("Enter chicken positions ex(2 5 10 12 15):").
		Input("2 5 10 12 15", input.WithValidateFunc(func(s string) error {
			isEmpty := strings.TrimSpace(s) == ""
			if isEmpty {
				return fmt.Errorf("please enter a value")
			}

			parts := strings.Split(s, " ")
			if len(parts) == 0 {
				return fmt.Errorf("please enter more than 1 value")
			}

			return nil
		}))

	if err != nil {
		log.Panic(err)
	}

	parts := strings.Split(answer, " ")
	result := make([]uint, len(parts))
	for i := 0; i < len(parts); i++ {
		chickenPosition, err := strconv.ParseUint(parts[i], 10, 64)
		if err != nil {
			log.Panic(err)
		}

		result[i] = uint(chickenPosition)
	}

	return result
}

func main() {
	protectLimit, roofLength := inputProtectLimitAndRoofLength()
	chickenPositions := inputChickenPositions()
	result := module.SupermanRescueChicken(protectLimit, roofLength, chickenPositions)
	fmt.Printf("Maximum chickens protected by Superman: %v\n", result)
}
