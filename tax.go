package main

import (
	"fmt"
	"math"
	"strings"
)

type item struct {
	itemName  string
	itemType  string
	itemPrice float64
}

func (i item) checkoutAmount(state string) (float64, error) {

	if state != "NJ" {
		if state != "PA" {
			if state != "DE" {
				return 0.0, fmt.Errorf("invalid State Abbreviation. please enter a valid state abbreviation in all caps")
			}
		}
	}

	if i.itemType == "Wic Eligible food" {

		return math.Round(i.itemPrice*100) / 100, nil

	} else if i.itemType == "Clothing" {

		return math.Round(((i.itemPrice*clothingTaxCalculator(state, i))+i.itemPrice)*100) / 100, nil

	} else {

		return math.Round(((i.itemPrice*salesTaxCalculator(state))+i.itemPrice)*100) / 100, nil

	}

}

func salesTaxCalculator(state string) float64 {

	if state == "NJ" {

		return .066

	} else if state == "DE" {

		return 0.0

	} else if state == "PA" {

		return .06

	} else {

		return 0.0

	}

}

func clothingTaxCalculator(state string, i item) float64 {

	if state == "NJ" && (strings.Contains(i.itemName, "Fur") || strings.Contains(i.itemName, "fur")) {

		return .066

	} else {

		return 0.0
	}

}
