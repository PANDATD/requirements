package main 

import (
	"os"
	"bufio"
	"log"
)

func GetStringInput() string {

	//Logic for Taking String Input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	//Error Handling Part 

	if err != nil{
		log.Println("Oh Ohh Something Went WRONG ")
	}
	// Ruturning the input 
	return input

}

