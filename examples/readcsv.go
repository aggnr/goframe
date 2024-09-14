// This example demonstrates how to read a CSV string into a DataFrame.
//go:build ignoreme
// +build ignoreme

package main

import (
	"github.com/aggnr/goframe"
)

func main() {
	csvString := `Name,Age,Salary,IsMarried
	John,30,50000.50,true
	Jane,25,60000.75,false`

	type Person struct {
		Name      string
		Age       int
		Salary    float64
		IsMarried bool
	}

	df, _ := cow.ReadCSVFromString(csvString, &Person{})

	defer df.Close()
	df.Display()
}
