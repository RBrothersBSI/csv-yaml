package domain

type DeviceProfile struct {
	Name string `yaml:"name"`
	Manufacturer string `yaml:"manufacturer"`
	Model string `yaml:"model"`
	Labels []string `yaml:"labels"`
	Description string `yaml:"description"`
	DeviceResources []*DeviceResource `yaml:"deviceResources"`
	DeviceCommands []*DeviceCommand `yaml:"deviceCommands"`
	CoreCommands []*CoreCommand `yaml:"coreCommands"`
}

type DeviceResource struct {
	Name string `yaml:"name"`
	Description string `yaml:"description"`
	Attributes Attribute `yaml:"attributes"`
	Properties Property `yaml:"properties"`
}

type DeviceCommand struct{
	Name string
	Get []struct {
		DeviceResource string
	}
	Set []struct {
		DeviceResource string
	}
}

type CoreCommand struct {
	Name string
	Get Get
	Put Put `yaml:",omitempty"`
}

type Get struct {
	Path string
	Responses []Response
}

type Put struct {
	Path string `yaml:"path"`
	ParameterNames []string `yaml:"parameterNames"`
	Responses []Response `yaml:"responses"`
}

type Response struct {
	Code string
	Description string
	ExpectedValues []string
}

type Attribute struct {
	Type string `yaml:"type"`
	Instance string `yaml:"instance"`
	Property string `yaml:"property"`
	Index string `yaml:"index"`
}

type Property struct {
	Value Value `yaml:"value"`
	Units Unit `yaml:"units"`
}

type Value struct {
	Type string `yaml:"type"`
	ReadWrite string `yaml:"readWrite"`
	FloatEncoding string `yaml:"floatEncoding,omitempty"`
}

type Unit struct {
	Type string `yaml:"type"`
	ReadWrite string `yaml:"readWrite"`
	DefaultValue string `yaml:"defaultValue"`
}

func NewProfile(name string, manufacturer string, model string, labels []string, description string, dr []*DeviceResource, cc []*CoreCommand) (DeviceProfile, error){
	dp := DeviceProfile{
		Name:            name,
		Manufacturer:    manufacturer,
		Model:           model,
		Labels:          labels,
		Description:    description,
		DeviceResources: dr,
		DeviceCommands:  nil,
		CoreCommands:    cc,
	}
	return dp, nil
}

func ObjToDeviceResource(obj Object) (DeviceResource, error){
	dr := DeviceResource{
		Name:        obj.GetName(),
		Description: obj.GetDescription(),
		Attributes:  Attribute{
				Type: obj.GetBACType(),
				Instance: obj.GetObjectId(),
				Property: "presentValue",
				Index: "none",
		},
		Properties:  Property{
				Value: Value{
					Type:          obj.GetValueType(),
					ReadWrite:     "RW",
					FloatEncoding: obj.GetFloatEncoding(),
			},
			Units: Unit{
				Type:         "String",
				ReadWrite:    "R",
				DefaultValue: obj.GetDefaultUnits(),
			},
		},
	}

	return dr, nil
}


func ObjToCoreCommand(obj Object) (CoreCommand, error) {
	getRs := constructGetResponses(obj.GetName())
	get := Get{
		Path:      "/api/v1/device/{deviceId}/" + obj.GetName(),
		Responses: getRs,
	}
	getPs := constructPutResponses()
	put := Put{
		Path:      "/api/v1/device/{deviceId}/" + obj.GetName(),
		ParameterNames: []string{ obj.GetName() },
		Responses: getPs,
	}
	var dc CoreCommand
	if obj.GetRW() == "RW" {
		dc = CoreCommand{
			Name: obj.GetName(),
			Get:  get,
			Put:  put,
		}
	} else {
		dc = CoreCommand{
			Name: obj.GetName(),
			Get:  get,
		}
	}
	return dc, nil
}

func constructGetResponses(name string) ([]Response) {
	res := []Response{
		{
			Code: "200",
			Description: "Success",
			ExpectedValues: []string{name},
		},
		{
			Code: "400",
			Description: "BadRequest",
			ExpectedValues: []string{},
		},
		{
			Code: "404",
			Description: "Not Found",
			ExpectedValues: []string{},
		},
		{
			Code: "500",
			Description: "FailedTransaction",
			ExpectedValues: []string{},
		},
	}
	return res
}

func constructPutResponses() ([]Response) {
	res := []Response{
		{
			Code: "200",
			Description: "Success",
			ExpectedValues: []string{},
		},
		{
			Code: "400",
			Description: "BadRequest",
			ExpectedValues: []string{},
		},
		{
			Code: "404",
			Description: "Not Found",
			ExpectedValues: []string{},
		},
		{
			Code: "500",
			Description: "FailedTransaction",
			ExpectedValues: []string{},
		},
	}
	return res
}
