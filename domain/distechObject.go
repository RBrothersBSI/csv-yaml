package domain

import (
	"siteconnect/distechdp/utils"
	"strconv"
	"strings"
)

type DistechObject struct {
	Keyname string `csv:"-"`
	DeviceInstance string `csv:"-"`
	ObjectName string `csv:"object-name"`
	ObjectType string `csv:"object-type"`
	ObjectInstance string `csv:"object-instance"`
	Description string `csv:"description"`
	DefaultValue string `csv:"present-value-default"`
	MinValue string `csv:"min-present-value"`
	MaxValue string `csv:"max-present-value"`
	Commandable string `csv:"commandable"`
	SupportsCOV string `csv:"supports-COV"`
	HighLimit string `csv:"hi-limit"`
	LowLimit string `csv:"low-limit"`
	StateTextReference string `csv:"state-text-reference"`
	UnitCode string `csv:"unit-code"`
	VendorAddress string `csv:"-"`
	ProgramObject string `csv:"-"`
}

func (d DistechObject) GetFloatEncoding() string {
	str := d.GetValueType()
	if strings.Contains(str, "Float") {
		return "eNotation"
	}
	return ""
}

func (d DistechObject) GetDefaultUnits() string {
	if d.StateTextReference != "" {
		num, _ := strconv.Atoi(d.StateTextReference)
		return strings.Join([]string{stateTexts[num-1].State1, stateTexts[num-1].State2}, "/")
	}
	if d.UnitCode != "" {
		num, _ := strconv.Atoi(d.UnitCode)
		if strings.Contains(d.ObjectName, "Temp"){
			return "degrees-fahrenheit"
		}
		if strings.Contains(d.ObjectName, "Humidity"){
			return "percent"
		}
		return unitTypes[num].UnitText
	}
	return ""
}

func (d DistechObject) GetName() string {
	return utils.SpaceStringsBuilder(d.ObjectName)
}

func (d DistechObject) GetObjectId() string {
	return d.ObjectInstance
}

func (d DistechObject) GetUnits() string {
	return d.StateTextReference
}

func (d DistechObject) GetRW() string {
	if d.Commandable == "Y"{
		return "RW"
	} else {
		return "R"
	}
}

func (d DistechObject) GetBACType() string {
	return d.ObjectType
}

func (d DistechObject) GetValueType() string {
	vtype := d.ObjectType
	if vtype == "0" || vtype == "1" || vtype == "2" {
		return "Float32"
	} else if vtype == "3" || vtype == "4" || vtype == "5" {
		return "Bool"
	} else {
		return "String"
	}
}

func (d DistechObject) GetDescription() string {
	return d.Description
}