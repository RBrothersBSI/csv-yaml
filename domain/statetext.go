package domain

import (
	"github.com/gocarina/gocsv"
	"os"
)

var(
	stateTexts []*StateValue
)

func init(){
	stateFile, err := os.OpenFile("./resources/testEDE/testEDE_StateText.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	stateTexts = []*StateValue{}

	defer stateFile.Close()

	if _, err = stateFile.Seek(24, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	if err := gocsv.UnmarshalFile(stateFile, &stateTexts); err != nil { // Load clients from file
		panic(err)
	}

	//for _, stateText := range stateTexts {
	//	fmt.Printf("RefNumber: %v, State1: %s, State2: %s, State3: %s\n", stateText.RefNumber, stateText.State1, stateText.State2, stateText.State3)
	//}
}

type StateValue struct {
	RefNumber string `csv:"Reference Number"`
	State1 string `csv:"Text 1 or Inactive-Text"`
	State2 string `csv:"Text 2 or Active-Text"`
	State3 string `csv:"Text 3"`
}
