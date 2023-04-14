package main

import (
	"log"
	"testing"
)

func TestCheckoutAmountDE(t *testing.T) {

	deWicFood := item{"Water", "Wic Eligible food", 5.00}
	deFastFood := item{"McDouble", "Fast Food", 4.20}
	deFurClothing := item{"Fur Hat", "Clothing", 120.50}
	deNoFurClothing := item{"Red Sox Shirt", "Clothing", 30}
	deNoTaxLambo := item{"Lamborghini", "Vehicle", 420000}

	deWic, err := deWicFood.checkoutAmount("DE")
	if err != nil {
		t.Fail()
		log.Fatal(err)
	}
	if deWic != 5.00 {
		t.Errorf("Expected 420000 but got %f", deWic)
	}

	deFood, err := deFastFood.checkoutAmount("DE")
	if err != nil {
		t.Fail()
		log.Fatal(err)
	}
	if deFood != 4.20 {

		t.Errorf("Expected 4.20 but got %f", deFood)
	}

	deFur, err := deFurClothing.checkoutAmount("DE")
	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	if deFur != 120.50 {
		t.Errorf("Expected 120.50 but got %f", deFur)
	}

	deNoFur, err := deNoFurClothing.checkoutAmount("DE")
	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	if deNoFur != 30.0 {
		t.Errorf("Expected 30.0 but got %f", deNoFur)
	}

	deLambo, err := deNoTaxLambo.checkoutAmount("DE")
	if err != nil {
		t.Fail()
		log.Fatal(err)
	}
	if deLambo != 420000 {
		t.Errorf("Expected 420000 but got %f", deLambo)
	}

}

func TestCheckoutAmountPA(t *testing.T) {

	//Not taxed -> return 5.95
	paWicFood := item{"Bread Loaf", "Wic Eligible food", 5.95}
	//Sales Taxed -> return 5.2894
	paFastFood := item{"Big Mac", "Fast Food", 4.99}
	//Not Taxed -> return 499.99
	paFurClothing := item{"Fur Coat", "Clothing", 499.99}
	//Not Taxed -> return 399.99
	paNoFurClothing := item{"Leather Jacket", "Clothing", 399.99}
	//Sales taxed -> return 47700
	paCar := item{"Ford F150", "Vehicle", 45000.00}

	paWic, err := paWicFood.checkoutAmount("PA")

	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	if paWic != 5.95 {
		t.Errorf("Expected 5.95 but got %f", paWic)
	}

	paNOWIC, err := paFastFood.checkoutAmount("PA")

	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	if paNOWIC != 5.29 {
		t.Errorf("Expected 5.29 but got %f", paNOWIC)
	}

	paFUR, err := paFurClothing.checkoutAmount("PA")

	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	if paFUR != 499.99 {
		t.Errorf("Expected 499.99 but got %f", paFUR)
	}

	paNOFUR, err := paNoFurClothing.checkoutAmount("PA")

	if err != nil {

		t.Fail()
		log.Fatal(err)
	}

	if paNOFUR != 399.99 {
		t.Errorf("Expected 399.99 but got %f", paNOFUR)
	}

	paWHIP, err := paCar.checkoutAmount("PA")

	if err != nil {

		t.Fail()
		log.Fatal(err)
	}

	if paWHIP != 47700.0 {
		t.Errorf("Expected 47700 but got %f", paWHIP)
	}

}

func TestCheckoutAmountNJ(t *testing.T) {

	//No tax -> return 1.50
	njWicFood := item{"Apple", "Wic Eligible food", 1.50}
	//Sales taxed -> return 3.73
	njFastFood := item{"Fish Fillet", "Fast Food", 3.50}
	//Sales taxed -> return 213.19
	njFurClothing := item{"Fur Uggs", "Clothing", 199.99}
	//No Tax -> return 50.0
	njNoFurClothing := item{"Hat", "Clothing", 50.00}
	//Sales Taxed -> return 21320
	njCar := item{"Honda Civic", "Vehicle", 20000}

	njWIC, err := njWicFood.checkoutAmount("NJ")

	if err != nil {

		t.Fail()
		log.Fatal(err)
	}

	if njWIC != 1.50 {
		t.Errorf("Expected 1.50 but got %f", njWIC)
	}

	njNOWIC, err := njFastFood.checkoutAmount("NJ")

	if err != nil {

		t.Fail()
		log.Fatal(err)
	}

	if njNOWIC != 3.73 {
		t.Errorf("Expected 3.73 but got %f", njNOWIC)
	}

	njFUR, err := njFurClothing.checkoutAmount("NJ")

	if err != nil {

		t.Fail()
		log.Fatal(err)
	}

	if njFUR != 213.19 {
		t.Errorf("Expected 213.19 but got %f", njFUR)
	}

	njNOFUR, err := njNoFurClothing.checkoutAmount("NJ")

	if err != nil {

		t.Fail()
		log.Fatal(err)
	}

	if njNOFUR != 50.00 {
		t.Errorf("Expected 50.0 but got %f", njNOFUR)
	}

	njWHIP, err := njCar.checkoutAmount("NJ")

	if err != nil {

		t.Fail()
		log.Fatal(err)
	}

	if njWHIP != 21320.0 {
		t.Errorf("Expected 21320.0 but got %f", njWHIP)
	}

}

func TestTestCheckoutAmountErrors(t *testing.T) {

	massAudi := item{"Audi S4", "Vehicle", 45000}

	noTaxonNJFUR := item{"Canada Goose Fur Jacket", "Clothing", 999.99}

	_, err := massAudi.checkoutAmount("MA")

	if err == nil {
		t.Errorf("MA is not a valid state abbriviation but an error was not returned")
	}

	canadaGoose, err := noTaxonNJFUR.checkoutAmount("NJ")

	if err != nil {

		t.Fail()
		log.Fatal(err)
	}

	if canadaGoose == 999.99 {
		t.Errorf("A FUR clothing item was not taxed in NJ")
	}

}
