package module_test

import (
	"superman/module"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name             string
	ProtectLimit     uint
	RoofLength       uint
	ChickenPositions []uint
	Expected         uint
}

var cases = []TestCase{
	{
		Name:             "SupermanRescueChicken(5, 5, []uint{2, 5, 10, 12, 15})",
		ProtectLimit:     5,
		RoofLength:       5,
		ChickenPositions: []uint{2, 5, 10, 12, 15},
		Expected:         2,
	},
	{
		Name:             "SupermanRescueChicken(6, 10, []uint{1, 11, 30, 34, 35, 37})",
		ProtectLimit:     6,
		RoofLength:       10,
		ChickenPositions: []uint{1, 11, 30, 34, 35, 37},
		Expected:         4,
	},
	{
		Name:             "SupermanRescueChicken(5, 5, []uint{2, 5, 10, 12, 15, 16, 17, 18, 18, 20})",
		ProtectLimit:     5,
		RoofLength:       10,
		ChickenPositions: []uint{2, 5, 10, 12, 15, 16, 17, 18, 18, 20},
		Expected:         5,
	},
}

func TestSupermanRescueChicken(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			result := module.SupermanRescueChicken(test.ProtectLimit, test.RoofLength, test.ChickenPositions)
			assert.Equal(t, test.Expected, result)
		})
	}
}

func BenchmarkSupermanRescueChicken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		module.SupermanRescueChicken(5, 5, []uint{2, 5, 10, 12, 15, 16, 17, 18, 18, 20})
	}
}
