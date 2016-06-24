package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Capture and parse flags from command line
	employeeSerial := flag.String("serial", "0", "employee serial number")
	filename := flag.String("filename", "0", "filename")
	flag.Parse()

	// Call function to fetch employee record
	readFile(*filename, *employeeSerial)
}

// Define Employee struct
type Employee struct {
	Name    string
	Serial  string
	Age     string
	Address string
}

// Function reads txt input and save it in a variable
func readFile(filename string, employeeSerial string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(data)

	//	Splits up input into slice of student record strings
	spltStr := strings.Split(str, "&")

	// Initialize map of student records
	mapOfRecords := make(map[string]Employee)

	//	Loop through chunks and place them in struct for each employee
	for i := 0; i < (len(spltStr)); i++ {
		spltSlice := strings.Split(spltStr[i], ",")
		studentStruct := Employee{
			Name:    spltSlice[0],
			Serial:  spltSlice[1],
			Age:     spltSlice[2],
			Address: spltSlice[3],
		}
		//	Place each struct into map where key is serial and map is record
		mapOfRecords[spltSlice[1]] = studentStruct
	}
	// Prints out record of serial passed as flag from command line
	fmt.Println("Name: " + mapOfRecords[employeeSerial].Name)
	fmt.Println("Serial: " + mapOfRecords[employeeSerial].Serial)
	fmt.Println("Age: " + mapOfRecords[employeeSerial].Age)
	fmt.Println("Address: " + mapOfRecords[employeeSerial].Address)
}

// Questions:

// Why one letter in all examples and not full word? Smaller file size?
// Why do you have to parse flags from command line?
