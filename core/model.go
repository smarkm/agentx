package core

//VarBind variable binding
type VarBind struct {
	Name   string
	Oid    string
	Type   string
	GoType string
}

//Module MIB module
type Module struct {
	GoName  string
	Mib     string
	RootOid string
	Oids    map[string]VarBind
}
