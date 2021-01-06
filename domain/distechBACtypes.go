package domain

import (
	"github.com/gocarina/gocsv"
	"os"
)

var(
	bacTypes []*DBacType
)

func init(){
	bacFile, err := os.OpenFile("./resources/testEDE/testEDE_ObjTypes.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	bacTypes = []*DBacType{}

	defer bacFile.Close()


	if _, err = bacFile.Seek(35, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	if err := gocsv.UnmarshalFile(bacFile, &bacTypes); err != nil { // Load clients from file
		panic(err)
	}

	//for _, bactype := range bacTypes {
	//	fmt.Printf("Code: %v, Object Type: %s\n", bactype.Code, bactype.ObjectType)
	//}
}

type DBacType struct {
	Code int `csv:"Code"`
	ObjectType string `csv:"Object Type"`
}

