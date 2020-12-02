package core_test

import (
	"agentx/core"
	"testing"
)

func TestLoadAndGenerateCode(t *testing.T) {
	m := core.Module{}
	m.Mib = "TestMib"
	m.Oids = make(map[string]core.VarBind)
	oid := core.VarBind{Name: "test1", Oid: "1.2.3.4", Type: "Type"}
	m.Oids[oid.Name] = oid

	core.GenerateFiles(m, true)
}
