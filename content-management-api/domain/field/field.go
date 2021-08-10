package field

type ID string

func (n ID) String() string {
	return string(n)
}

type FieldModel struct {
	Name     Name
	Type     Type
	Required Required
}

type Name string

func (n Name) String() string {
	return string(n)
}

type Required bool

type FieldModels []FieldModel
