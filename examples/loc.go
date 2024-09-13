// Description: This example demonstrates how to retrieve rows from a DataFrame using the Loc method.
//go:build ignoreme
// +build ignoreme

package main

import (
	"fmt"
	"github.com/aggnr/goframe"
	"log"
)

func main() {

	type Person struct {
		Name      string
		Age       int
		Salary    float64
		IsMarried bool
	}

	people := []Person{
		{"John", 30, 50000.50, true},
		{"Jane", 25, 60000.75, false},
	}

	df, err := goframe.NewDataFrame(people)
	if err != nil {
		log.Fatalf("Error creating DataFrame: %v", err)
	}
	defer df.Close()

	rows, err := df.Loc(2)
	if err != nil {
		log.Fatalf("Error retrieving rows: %v", err)
	}

	fmt.Println("Retrieved rows for index 2:")
	for _, row := range rows {
		fmt.Println(row)
	}
}
