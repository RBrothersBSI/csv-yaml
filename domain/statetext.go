package domain

import (
	"encoding/csv"
	"os"
)

var(
	stateTexts []*StateValue
)

func init(){
	stateFile, err := os.OpenFile("./resources/ECBVAVN/ECBVAVN_StateText.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	stateTexts = []*StateValue{}

	defer stateFile.Close()
	if _, err = stateFile.Seek(169, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	r := csv.NewReader(stateFile)
	r.FieldsPerRecord = -1
	str, _ := r.ReadAll()
	for _, v := range str {
		sv := StateValue{}
		for i,v1 := range v {
			switch i {
			case 0:
				sv.RefNumber = v1
			case 1:
				sv.State1 = v1
			case 2:
				sv.State2 = v1
			case 3:
				sv.State3 = v1
			case 4:
				sv.State4 = v1
			case 5:
				sv.State5 = v1
			case 6:
				sv.State6 = v1
			case 7:
				sv.State7 = v1
			case 8:
				sv.State8 = v1
			case 9:
				sv.State9 = v1
			case 10:
				sv.State10 = v1
			case 11:
				sv.State12 = v1
			case 12:
				sv.State12 = v1
			}
			stateTexts = append(stateTexts, &sv)
		}
	}
	//for _, stateText := range stateTexts {
	//	fmt.Printf("StateText: %v", stateText)
	//}
}

type StateValue struct {
	RefNumber string `csv:"Reference Number"`
	State1 string `csv:"Text 1 or Inactive-Text"`
	State2 string `csv:"Text 2 or Active-Text"`
	State3 string `csv:"Text 3,omitempty"`
	State4 string `csv:"Text 4,omitempty"`
	State5 string `csv:"Text 5,omitempty"`
	State6 string `csv:"Text 6,omitempty"`
	State7 string `csv:"Text 7,omitempty"`
	State8 string `csv:"Text 8,omitempty"`
	State9 string `csv:"Text 9,omitempty"`
	State10 string `csv:"Text 10,omitempty"`
	State11 string `csv:"Text 11,omitempty"`
	State12 string `csv:"Text 12,omitempty"`
	State13 string `csv:"Text 13,omitempty"`
	State14 string `csv:"Text 14,omitempty"`
}
