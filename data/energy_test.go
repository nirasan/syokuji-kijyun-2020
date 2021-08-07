package data

import (
	"fmt"
	"testing"
)

func TestEnergyList(t *testing.T) {
	list := EnergyList()
	for _, d := range list {
		fmt.Printf("%.0f ", d.Value.Float)
		if d.Gender == GenderFemale && d.ActivityLevel == ActivityLevel3 {
			fmt.Println("")
		}
	}
}
