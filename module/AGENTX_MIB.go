package module

import (
	"agentx/core"
	"fmt"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

func init() {
	var Mib *AGENTX_MIB = &AGENTX_MIB{}
    
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.1.1.0", Mib.getagentxDefaultTimeout,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.1.2.0", Mib.getagentxMasterAgentXVer,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.2.1.0", Mib.getagentxConnTableLastChange,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.2.2.1.1", Mib.getagentxConnIndex,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.2.2.1.2", Mib.getagentxConnOpenTime,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.2.2.1.3", Mib.getagentxConnTransportDomain,pdu.VariableTypeObjectIdentifier)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.2.2.1.4", Mib.getagentxConnTransportAddress,pdu.VariableTypeOctetString)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.3.1.0", Mib.getagentxSessionTableLastChange,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.3.2.1.1", Mib.getagentxSessionIndex,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.3.2.1.2", Mib.getagentxSessionObjectID,pdu.VariableTypeObjectIdentifier)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.3.2.1.3", Mib.getagentxSessionDescr,pdu.VariableTypeOctetString)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.3.2.1.4", Mib.getagentxSessionAdminStatus,pdu.VariableTypeInteger)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.3.2.1.5", Mib.getagentxSessionOpenTime,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.3.2.1.6", Mib.getagentxSessionAgentXVer,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.3.2.1.7", Mib.getagentxSessionTimeout,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.1.0", Mib.getagentxRegistrationTableLastChange,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.2.1.1", Mib.getagentxRegIndex,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.2.1.2", Mib.getagentxRegContext,pdu.VariableTypeOctetString)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.2.1.3", Mib.getagentxRegStart,pdu.VariableTypeObjectIdentifier)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.2.1.4", Mib.getagentxRegRangeSubId,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.2.1.5", Mib.getagentxRegUpperBound,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.2.1.6", Mib.getagentxRegPriority,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.2.1.7", Mib.getagentxRegTimeout,pdu.VariableTypeCounter32)
    
	core.RegisterOid( "1.3.6.1.2.1.74.1.4.2.1.8", Mib.getagentxRegInstance,pdu.VariableTypeInteger)
    
	core.AddRootOid("1.3.6.1.2.1.74")
	fmt.Printf("register AGENTX-MIB: 1.3.6.1.2.1.74 ...\n")

}
//AGENTX_MIB code for MIB: AGENTX-MIB
type AGENTX_MIB struct {
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

func (m *AGENTX_MIB) getagentxConnTableLastChange(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxConnTableLastChange(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxConnIndex(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxConnIndex(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxConnOpenTime(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxConnOpenTime(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxConnTransportDomain(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxConnTransportDomain(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeObjectIdentifier, value, nil
	}
	return ioid, pdu.VariableTypeObjectIdentifier, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxConnTransportAddress(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxConnTransportAddress(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeOctetString, value, nil
	}
	return ioid, pdu.VariableTypeOctetString, value, nil
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

func (m *AGENTX_MIB) getagentxSessionIndex(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxSessionIndex(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxSessionObjectID(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxSessionObjectID(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeObjectIdentifier, value, nil
	}
	return ioid, pdu.VariableTypeObjectIdentifier, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxSessionDescr(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxSessionDescr(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeOctetString, value, nil
	}
	return ioid, pdu.VariableTypeOctetString, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxSessionAdminStatus(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxSessionAdminStatus(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeInteger, value, nil
	}
	return ioid, pdu.VariableTypeInteger, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxSessionOpenTime(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxSessionOpenTime(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxSessionAgentXVer(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxSessionAgentXVer(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxSessionTimeout(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxSessionTimeout(oid)
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

func (m *AGENTX_MIB) getagentxRegIndex(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxRegIndex(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxRegContext(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxRegContext(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeOctetString, value, nil
	}
	return ioid, pdu.VariableTypeOctetString, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxRegStart(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxRegStart(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeObjectIdentifier, value, nil
	}
	return ioid, pdu.VariableTypeObjectIdentifier, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxRegRangeSubId(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxRegRangeSubId(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxRegUpperBound(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxRegUpperBound(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxRegPriority(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxRegPriority(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxRegTimeout(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxRegTimeout(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeCounter32, value, nil
	}
	return ioid, pdu.VariableTypeCounter32, value, nil
	//return core.NoSuchInstance(ioid)
}

func (m *AGENTX_MIB) getagentxRegInstance(oid string) (value.OID, pdu.VariableType, interface{}, error) {
	ioid, _ := value.ParseOID(oid)
	value,err := getAGENTX_MIB_agentxRegInstance(oid)
	if err!=nil {
		return ioid, pdu.VariableTypeInteger, value, nil
	}
	return ioid, pdu.VariableTypeInteger, value, nil
	//return core.NoSuchInstance(ioid)
}

