package main

import (
	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"siteconnect/distechdp/domain"
)

func main() {
	/////////////////////////
	//Distech Logic
	objectsFile, err := os.OpenFile("./resources/testEDE/testEDE_ede.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer objectsFile.Close()

	if _, err := objectsFile.Seek(165, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	objects := []*domain.DistechObject{}

	if err := gocsv.UnmarshalFile(objectsFile, &objects); err != nil { // Load clients from file
		panic(err)
	}

	dr := []*domain.DeviceResource{}
	cc := []*domain.CoreCommand{}

	for _, object := range objects {
		data, _ := domain.ObjToDeviceResource(object)
		dPuts, _ := domain.ObjToCoreCommand(object)
		dr = append(dr, &data)
		cc = append(cc, &dPuts)
	}
	dp, _ := domain.NewProfile(
		"Distech S100 Chiller Plant",
		"Distech",
		"S100",
		[]string{"distech", "CH", "LL"},
		"Distech chiller plant programmed for 2 chiller lead/lag/failover",
		dr,
		cc,
	)
	d, err := yaml.Marshal(dp)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile("./distech.yml", d, 0644)
	if err != nil {
		panic(err)
	}


	///////////////////////
	//Niagara Logic
	niagaraFile, err := os.OpenFile("./resources/niagara/niagara.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer niagaraFile.Close()

	niagaraObjs := []*domain.NiagaraObject{}

	if err := gocsv.UnmarshalFile(niagaraFile, &niagaraObjs); err != nil { // Load clients from file
		panic(err)
	}

	drn := []*domain.DeviceResource{}
	ccn := []*domain.CoreCommand{}

	for _, obj := range niagaraObjs {
		nData, _ := domain.ObjToDeviceResource(obj)
		nPuts, _ := domain.ObjToCoreCommand(obj)
		drn = append(drn, &nData)
		ccn = append(ccn, &nPuts)
	}

	dpN, _ := domain.NewProfile(
		"Niagara",
		"Niagara Test",
		"llll10000",
		[]string{"really", "who", "cares"},
		"This is a test",
		 drn,
		 ccn,
	)

	dN, err := yaml.Marshal(dpN)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile("./niagara.yml", dN, 0644)
	if err != nil {
		panic(err)
	}

}
