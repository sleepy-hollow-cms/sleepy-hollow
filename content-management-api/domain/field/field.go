package field

type Field struct {
	Name      Name
	Type      Type
	Required  Required
}

type Name string


func (n Name) String() string {
	return string(n)
}

type Required bool

type Fields []Field
