package module

import (
	"agentx/core"
	"fmt"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

func init() {
	var Mib *{{.GoName}} = &{{.GoName}}{}
    
    {{range $v := .Oids }}
	core.RegisterOid( "{{.Oid}}", Mib.get{{.Name}},pdu.{{.Type}})
    {{end}}
	core.AddRootOid("{{.RootOid}}")
	fmt.Printf("register {{.Mib}}: {{.RootOid}} ...\n")

}
//{{.GoName}} code for MIB: {{.Mib}}
type {{.GoName}} struct {
}

{{range $v := .Oids }}
func (m *{{$.GoName}}) get{{.Name}}(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := get{{$.GoName}}_{{.Name}}(oid)
	if err!=nil {
		return ioid, pdu.{{.Type}}, value, nil
	}
	return ioid, pdu.{{.Type}}, value, nil
	//return core.NoSuchInstance(ioid)
}
{{end}}
