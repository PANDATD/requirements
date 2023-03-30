package requirements

// Author: github.com/pandatd/
// License: MIT
// Description: This is a simple package for taking input from the user and all usefull funcations for daily use in Golang devlopment.
// Package : github.com/pandatd/requirements@latest.
// version : v1.1.2
// contact:	+91 7745014673
// mail : tejasdixit17@proton.me

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetStringInput() string {

	//Logic for Taking String Input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	//Error Handling Part

	if err != nil {
		log.Println("Oh Ohh Something Went WRONG ")
	}
	// Ruturning the input
	return input

}

func GetIntInput() int {

	//Logic for Taking Int Input
	var input int
	_, err := fmt.Scan(&input)

	//Error Handling Part

	if err != nil {
		log.Println("Oh Ohh Something Went WRONG ")
	}
	// Ruturning the input
	return input

}
