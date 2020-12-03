package core

//VarBind variable binding
type VarBind struct {
	Name   string
	Oid    string
	Type   string
	GoType string
}
type SortByOID []VarBind // Len is the number of elements in the collection.

//Module MIB module
type Module struct {
	GoName  string
	Mib     string
	RootOid string
	Oids    []VarBind
}

func (vbs SortByOID) Len() int {
	return len(vbs)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (vbs SortByOID) Less(i int, j int) bool {
	return vbs[i].Oid < vbs[j].Oid
}

// Swap swaps the elements with indexes i and j.
func (vbs SortByOID) Swap(i int, j int) {
	vbs[i], vbs[j] = vbs[j], vbs[i]
}
