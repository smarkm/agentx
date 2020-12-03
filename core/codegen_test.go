package core_test

import (
	"agentx/core"
	"fmt"
	"sort"
	"testing"
)

func TestLoadAndGenerateCode(t *testing.T) {
	m := core.Module{}
	m.Mib = "TestMib"
	m.Oids = make([]core.VarBind, 0)
	oid := core.VarBind{Name: "test1", Oid: "1.2.3.4", Type: "Type"}
	m.Oids = append(m.Oids, oid)

	core.GenerateFiles(m, true)
}

func TestVarBindSort(t *testing.T) {
	vbs := []core.VarBind{
		{Oid: "1"},
		{Oid: "0"},
	}

	sort.Sort(core.SortByOID(vbs))
	fmt.Println(vbs)

}
