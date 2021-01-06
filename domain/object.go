package domain

type Object interface {
	GetName() string
	GetObjectId() string
	GetUnits() string
	GetRW() string
	GetBACType() string
	GetValueType() string
	GetDescription() string
	GetDefaultUnits() string
	GetFloatEncoding() string
}
