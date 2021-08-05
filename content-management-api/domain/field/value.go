package field

type TextValue struct {
	HasValue
	Value string
}

type MultipleTextValue struct {
	HasValue
	Value []string
}

type NumberValue uint64

type HasValue interface {
}

func FactoryValue(typeName Type, value interface{}) HasValue {
	switch typeName {
	case Text:
		return TextValue{
			Value: value.(string),
		}
	case MultipleText:
		ir := value.([]interface{})
		strings := make([]string, len(ir))
		for i, v := range ir {
			strings[i] = v.(string)
		}
		return MultipleTextValue{
			Value: strings,
		}
	default:
		return nil
	}
}
