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
	listHandler *ListHandler         = &ListHandler{}
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

//RegisterModules register module
func RegisterModules(session *agentx.Session) {
	fmt.Println("load...")
	session.Handler = listHandler

	if err := session.Register(127, value.MustParseOID("1.3.6.1.4.1.45995.3")); err != nil {
		log.Fatalf("%s", err)
	}

}

func RegisterCounter32(oid string, call GetOID) {
	item := listHandler.Add(oid)
	item.Type = pdu.VariableTypeCounter32
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
