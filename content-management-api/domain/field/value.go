package field

type TextValue struct {
	Value string
}

func (t TextValue) FieldValue() interface{} {
	return t.Value
}

type MultipleTextValue struct {
	Value []string
}

func (m MultipleTextValue) FieldValue() interface{} {
	return m.Value
}

type NumberValue uint64

type HasValue interface {
	FieldValue() interface{}
}

type HasName interface {
	FieldName() string
}

func FactoryValue(typeName Type, value interface{}) HasValue {
	switch typeName {
	case Text:
		return TextValue{
			Value: value.(string),
		}
	case MultipleText:
		switch value.(type) {
		case []string:
			return MultipleTextValue{
				Value: value.([]string),
			}
		default:
			ir := value.([]interface{})
			strings := make([]string, len(ir))
			for i, v := range ir {
				strings[i] = v.(string)
			}
			return MultipleTextValue{
				Value: strings,
			}
		}
	default:
		return nil
	}
}
