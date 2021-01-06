package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("./resources/testEDE/testEDE_ede.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	defer csvfile.Close()

	// Parse the file
	r := csv.NewReader(csvfile)
	r.FieldsPerRecord = -1

	//r := csv.NewReader(bufio.NewReader(csvfile))
	fileData, _ := r.ReadAll()

	// Iterate through the records
	for _, value := range fileData {
		// Read each record from csv
		//record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(value) > 7 {
			//fmt.Printf("%v, Object Name: %v, Object Type: %v, Object Instance: %v, Default Value: %v, Min Present Value: %v, Max Present Value: %v, Commandable: %v, Supports COV: %v, High Limit: %v, Low Limit: %v, State Text Ref: %v, Unit Code: %v\n",
			//	e, value[2], value[3], value[4], value[5], value[6], value[7], value[8], value[9], value[10], value[11], value[12], value[13])

		}
	}
}
