package gen

type Function struct {
	Name     string
	Generics string
	Args     string
	Results  string
	Body     string
}

const fnTemplate = `
	func {{.Name}}{{.Generics}}({{.args}}){{.Results}}{
		{{.Body}}
	}
`
