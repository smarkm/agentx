package module

import (
	"agentx/core"
	"fmt"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

func init() {
	var Mib *AGENTX_MIB = &AGENTX_MIB{}
    
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.2.1.0", Mib.getagentxConnTableLastChange,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.1.1.0", Mib.getagentxDefaultTimeout,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.1.2.0", Mib.getagentxMasterAgentXVer,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.1.0", Mib.getagentxRegistrationTableLastChange,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.3.1.0", Mib.getagentxSessionTableLastChange,pdu.VariableTypeCounter32)
    
	core.AddRootOid("1.3.6.1.2.1.74")
	fmt.Printf("register AGENTX-MIB: 1.3.6.1.2.1.74 ...\n")

}
//AGENTX_MIB code for MIB: AGENTX-MIB
type AGENTX_MIB struct {
}


func (m *AGENTX_MIB) getagentxConnTableLastChange(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxConnTableLastChange(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxDefaultTimeout(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxDefaultTimeout(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxMasterAgentXVer(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxMasterAgentXVer(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxRegistrationTableLastChange(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxRegistrationTableLastChange(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxSessionTableLastChange(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxSessionTableLastChange(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

