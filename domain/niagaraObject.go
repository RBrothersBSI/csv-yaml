package domain

import (
	"siteconnect/distechdp/utils"
	"strings"
)

type NiagaraObject struct {
	Path string `csv:"Path"`
	Name string `csv:"Name"`
	Type string `csv:"Type"`
	Out string `csv:"Out"`
	Enable string `csv:"Enable"`
	ObjectID string `csv:"Object ID"`
	PropertyID string `csv:"Property ID"`
	Index string `csv:"Index"`
	TuningPolicy string `csv:"Tuning Policy Name"`
	DataType string `csv:"Data Type"`
	Read string `csv:"Read"`
	Write string `csv:"Write"`
	DeviceFacets string `csv:"Device Facets"`
	Facets string `csv:"Facets"`
	Conversion string `csv:"Conversion"`
	ReadValue string `csv:"Read Value"`
	WriteValue string `csv:"Write Value"`
	FaultCause string `csv:"Fault Cause"`
}

func (n NiagaraObject) GetFloatEncoding() string {
	str := n.GetValueType()
	if strings.Contains(str, "Float"){
		return "eNotation"
	}
	return ""
}

func (n NiagaraObject) GetDefaultUnits() string {
	str := strings.Split(n.Facets, ";")
	runes := []rune(str[0])
	switchOn := string(runes[0:4])
	switch switchOn {
	case "unit":
		return string(runes[8:])
	case "true":
		return "TRUE/FALSE"
	}
	return ""
}

func (n NiagaraObject) GetName() string {
	s := utils.SpaceStringsBuilder(n.Name)
	return s
}

func (n NiagaraObject) GetObjectId() string {
	str := strings.Split(n.ObjectID, ":")
	return str[1]
}

func (n NiagaraObject) GetUnits() string {
	return n.DeviceFacets
}

func (n NiagaraObject) GetRW() string {
	if n.Write == "OK"{
		return "RW"
	} else {
		return "R"
	}
}

func (n NiagaraObject) GetBACType() string {
	str := strings.Split(n.ObjectID, ":")
	return str[0]
}

func (n NiagaraObject) GetValueType() string {
	if strings.Contains(n.Type, "Boolean"){
		return "Bool"
	} else if strings.Contains(n.Type, "Numeric") {
		return "Float32"
	} else {
		return "String"
	}
}

func (n NiagaraObject) GetDescription() string {
	return n.Path
}
