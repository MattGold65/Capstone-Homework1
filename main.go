package main

import (
	"fmt"
	"log"
)

func main() {
	//PA WICFOOD TEST
	paWicFood := item{"Bread Loaf", "Wic Eligible food", 5.95}
	paFastFood := item{"Big Mac", "Fast Food", 4.99}

	paWIC, err := paWicFood.checkoutAmount("PA")
	if err != nil {
		log.Fatal(err)
	}

	paNoWic, err := paFastFood.checkoutAmount("PA")
	if err != nil {
		log.Fatal(err)
	}

	//PA FurClothing TEST
	paFurClothing := item{"Fur Coat", "Clothing", 499.99}

	paNoFurClothing := item{"Leather Jacket", "Clothing", 399.99}

	paFUR, err := paFurClothing.checkoutAmount("PA")
	if err != nil {
		log.Fatal(err)
	}

	paNOFUR, err := paNoFurClothing.checkoutAmount("PA")
	if err != nil {
		log.Fatal(err)
	}

	//PA everything else test
	paCar := item{"Ford F150", "Vehicle", 45000}

	paWhip, err := paCar.checkoutAmount("PA")
	if err != nil {
		log.Fatal(err)
	}

	//NJ WICFOOD TEST
	njWicFood := item{"Apple", "Wic Eligible food", 1.50}
	njFastFood := item{"Fish Fillet", "Fast Food", 3.50}
	njWIC, err := njWicFood.checkoutAmount("NJ")
	if err != nil {
		log.Fatal(err)
	}

	njNOWIC, err := njFastFood.checkoutAmount("NJ")
	if err != nil {
		log.Fatal(err)
	}

	//NJ FurClothing TEST
	njFurClothing := item{"Fur Uggs", "Clothing", 199.99}
	njNoFurClothing := item{"Hat", "Clothing", 50}

	njFUR, err := njFurClothing.checkoutAmount("NJ")
	if err != nil {
		log.Fatal(err)
	}

	njNOFUR, err := njNoFurClothing.checkoutAmount("NJ")
	if err != nil {
		log.Fatal(err)
	}

	//NJ everything else test
	njCar := item{"Honda Civic", "Vehicle", 20000}

	njWHIP, err := njCar.checkoutAmount("NJ")
	if err != nil {
		log.Fatal(err)
	}

	//DE WICFOOD TEST
	deWicFood := item{"Water", "Wic Eligible food", 5.00}
	deFastFood := item{"McDouble", "Fast Food", 4.20}

	deWIC, err := deWicFood.checkoutAmount("DE")
	if err != nil {
		log.Fatal(err)
	}

	deNOWIC, err := deFastFood.checkoutAmount("DE")
	if err != nil {
		log.Fatal(err)
	}

	//DE FurClothing TEST
	deFurClothing := item{"Fur Hat", "Clothing", 120.50}
	deNoFurClothing := item{"Red Sox Shirt", "Clothing", 30}

	deFUR, err := deFurClothing.checkoutAmount("DE")
	if err != nil {
		log.Fatal(err)
	}

	deNOFUR, err := deNoFurClothing.checkoutAmount("DE")
	if err != nil {
		log.Fatal(err)
	}

	//DE everything else test
	deNoTaxLambo := item{"Lamborghini", "Vehicle", 420000}

	deLambo, err := deNoTaxLambo.checkoutAmount("DE")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PA FOOD TESTS:")
	fmt.Println(paWIC)
	fmt.Println(paNoWic)

	fmt.Println("PA Clothing TESTS:")
	fmt.Println(paFUR)
	fmt.Println(paNOFUR)

	fmt.Println("PA CAR TESTS:")
	fmt.Println(paWhip)

	fmt.Println("NJ FOOD TESTS:")
	fmt.Println(njWIC)
	fmt.Println(njNOWIC)

	fmt.Println("NJ Clothing TESTS:")
	fmt.Println(njFUR)
	fmt.Println(njNOFUR)

	fmt.Println("NJ CAR TESTS:")
	fmt.Println(njWHIP)

	fmt.Println("DE FOOD TESTS:")
	fmt.Println(deWIC)
	fmt.Println(deNOWIC)

	fmt.Println("DE Clothing TESTS:")
	fmt.Println(deFUR)
	fmt.Println(deNOFUR)

	fmt.Println("DE CAR TESTS:")
	fmt.Println(deLambo)

}
