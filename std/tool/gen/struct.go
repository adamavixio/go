package gen

type Struct struct {
	Name string
}

const StructTemplate = `type {{.Name}}{{.Generics}} struct {{.Fields}}`
