package module


{{range $v := .Oids }}
func get{{$.GoName}}_{{.Name}}(oid string) (v {{.GoType}},e error) {
	//implement here
	return 
}
{{end}}
