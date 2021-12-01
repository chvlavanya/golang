package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// try to get the value of the environmental variable PWD
	fmt.Println("The value of PWD appears to be:", os.Getenv("PWD"))
	
	fmt.Println("\n\n")

	for _, envValue := range os.Environ() {
		envpair := strings.SplitN(envValue, "=", 2)
		fmt.Println(envpair)    // both the env and the value
		fmt.Println(envpair[0]) // the env itself
		fmt.Println(envpair[1]) // the value of the env
	}

	// setting an environmental variable
	os.Setenv("ALTA", "4research")
	fmt.Println("The value of ALTA appears to be:", os.Getenv("ALTA"))
}
