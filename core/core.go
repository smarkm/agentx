package core

import (
	"fmt"
	"log"

	"github.com/posteo/go-agentx"
	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
	"github.com/sleepinggenius2/gosmi/models"
	"github.com/sleepinggenius2/gosmi/types"
)

// Core
var (
	OidGets     map[string]GetOID    = make(map[string]GetOID)
	Types       map[interface{}]Type = make(map[interface{}]Type)
	listHandler                      = &ListHandler{}
	RootOids    map[string]struct{}  = make(map[string]struct{})
)

func init() {
	Types[types.BaseTypeUnknown] = Type{PDU: pdu.VariableTypeNull, Go: ""}
	Types[types.BaseTypeInteger32] = Type{PDU: pdu.VariableTypeCounter32, Go: "int32"}
	Types[types.BaseTypeOctetString] = Type{PDU: pdu.VariableTypeOctetString, Go: "string"}
	Types[types.BaseTypeObjectIdentifier] = Type{PDU: pdu.VariableTypeObjectIdentifier, Go: "string"}
	Types[types.BaseTypeUnsigned32] = Type{PDU: pdu.VariableTypeCounter32, Go: "uint32"}
	Types[types.BaseTypeInteger64] = Type{PDU: pdu.VariableTypeCounter64, Go: "int64"}
	Types[types.BaseTypeUnsigned64] = Type{PDU: pdu.VariableTypeCounter64, Go: "uint64"}
	Types[types.BaseTypeFloat32] = Type{PDU: pdu.VariableTypeGauge32, Go: "float32"}
	Types[types.BaseTypeFloat64] = Type{PDU: pdu.VariableTypeCounter64, Go: "float64"}
	Types[types.BaseTypeFloat128] = Type{PDU: pdu.VariableTypeCounter64, Go: "float128"}
	Types[types.BaseTypeEnum] = Type{PDU: pdu.VariableTypeInteger, Go: "int32"}
	Types[types.BaseTypeBits] = Type{PDU: pdu.VariableTypeGauge32, Go: "int"}
	Types[types.BaseTypePointer] = Type{PDU: pdu.VariableTypeObjectIdentifier, Go: "interface{}"}
	Types["name"] = Type{PDU: pdu.VariableTypeCounter32, Go: ""}

}

//Type type wapper
type Type struct {
	PDU pdu.VariableType
	Go  string
}

//GetOID implement for serve snmp get
type GetOID func(oid string) (value.OID, pdu.VariableType, interface{}, error)

func NoSuchInstance(oid value.OID) (value.OID, pdu.VariableType, interface{}, error) {
	return nil, pdu.VariableTypeNoSuchObject, nil, nil
}

//AddRootOid RT
func AddRootOid(rootOid string) {
	RootOids[rootOid] = struct{}{}
}

//RegisterModules register module
func RegisterModules(session *agentx.Session) {
	fmt.Println("load...")
	session.Handler = listHandler
	fmt.Print(RootOids)
	for oid := range RootOids {
		fmt.Printf("Register Root Oid: %s\n", oid)
		if err := session.Register(127, value.MustParseOID(oid)); err != nil {
			log.Fatalf("%s", err)
		}
	}

}

//RegisterOid RT
func RegisterOid(oid string, call GetOID, pduType pdu.VariableType) {
	item := listHandler.Add(oid)
	item.Type = pduType
	OidGets[oid] = call
}

//CoverGosmiTyp2PduType RT
func CoverGosmiTyp2PduType(t *models.Type) (vType pdu.VariableType, goType string) {
	vType, goType = pdu.VariableTypeNull, "interface{}"
	registerType, ok := Types[t.BaseType]
	if ok {
		vType, goType = registerType.PDU, registerType.Go
	}
	return
}
