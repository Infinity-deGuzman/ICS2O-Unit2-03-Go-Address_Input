// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program registers your address

package main

import "fmt"

func main() {
	// This function prints the address
	var streetNumber int
	var streetName string

	// input
	fmt.Println("This program gets a user's address.")
	fmt.Println()
	fmt.Print("Enter your street number: ")
	fmt.Scanln(&streetNumber)
	fmt.Print("Enter your street name: ")
	fmt.Scanln(&streetName)

	// output
	fmt.Println("Your address is:", streetNumber, " ", streetName, ".")
	fmt.Println("\nDone.")
}
