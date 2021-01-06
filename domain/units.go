package domain

import (
	"github.com/gocarina/gocsv"
	"os"
)

var(
	unitTypes []*UnitValue
)

func init(){
	unitFile, err := os.OpenFile("./resources/testEDE/testEDE_Units.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	unitTypes = []*UnitValue{}

	defer unitFile.Close()

	if _, err = unitFile.Seek(40, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	if err := gocsv.UnmarshalFile(unitFile, &unitTypes); err != nil { // Load clients from file
		panic(err)
	}

	//for _, unitType := range unitTypes {
	//	fmt.Printf("Code: %v, Object Type: %s\n", unitType.Code, unitType.UnitText)
	//}
}

type UnitValue struct {
	Code int `csv:"Code"`
	UnitText string `csv:"Unit Text"`
}
