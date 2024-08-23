package gen

type GenericType string

const (
	Any        GenericType = "any"
	Slice      GenericType = "slice"
	Map        GenericType = "map"
	Seq        GenericType = "seq"
	Comparable GenericType = "comparable"
)

type Generic struct {
	Label  string
	Type   GenericType
	Amount int
}

func NewGeneric(label string, generticType GenericType, amount int) Generic {
	return Generic{
		Label:  label,
		Type:   generticType,
		Amount: amount,
	}
}
