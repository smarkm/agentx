package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/sleepinggenius2/gosmi"
	"github.com/sleepinggenius2/gosmi/types"
)

//LoadAndGenerateCode RT
func LoadAndGenerateCode(path, mib string, dryRun bool) {
	gosmi.Init()
	gosmi.AppendPath(path)
	_, err := gosmi.LoadModule(mib)
	if err != nil {
		fmt.Printf("loading mib: %s, error: %s\n", mib, err.Error())
		return
	}

	loadedModules := gosmi.GetLoadedModules()
	fmt.Printf("loading mib: %s, success\n", mib)
	for _, m := range loadedModules {
		if m.Name == mib {
			//Find root OID
			nodes := m.GetNodes(types.NodeAny)
			rootOid := nodes[0].Oid.String()
			fmt.Printf("Root OID: %s\n", rootOid)

			//Find Node Scalar
			nodes = m.GetNodes(types.NodeScalar)
			nodes = append(nodes, m.GetNodes(types.NodeColumn)...)
			fileName := strings.ReplaceAll(m.Name, "-", "_")
			tm := Module{Mib: m.Name, GoName: fileName, RootOid: rootOid, Oids: make(map[string]VarBind)}
			for _, n := range nodes {
				oid := n.Oid.String()
				vType, goType := CoverGosmiTyp2PduType(n.Type)
				if n.Kind == types.NodeScalar {
					oid = oid + ".0"
				}
				fmt.Println(n.Name, n.Oid, n.Type, vType, goType)
				tm.Oids[n.Name] = VarBind{Name: n.Name, Oid: oid, Type: vType.String(), GoType: goType}
			}

			//sort.Sort(tm.Oids)
			GenerateFiles(tm, dryRun)

		}
	}
}

//GenerateFiles generate go code base on loading mib
func GenerateFiles(m Module, dryRun bool) {
	generateFile(m, "tpl/m.tpl", false, dryRun)

	impTpl := "tpl/m_impl.tpl"
	targetImplFile := fmt.Sprintf("module/%s_Impl.go", m.GoName)
	if dryRun {
		fmt.Println("-----------")
		generateFile(m, impTpl, true, dryRun)
		return
	}
	if _, err := os.Stat(targetImplFile); os.IsNotExist(err) {
		generateFile(m, impTpl, true, dryRun)
		return
	}

	fmt.Printf("generate file: %s already exist(skip)\n", targetImplFile)
}

func generateFile(m Module, tplPath string, impl bool, dryRun bool) {
	f, err := ioutil.ReadFile(tplPath)
	if err != nil {
		log.Print(err)
		return
	}

	// Parse requires a string
	t, err := template.New(m.GoName).Parse(string(f))
	if err != nil {
		log.Print(err)
		return
	}

	Impl := ""
	if impl {
		Impl = "_Impl"
	}
	targetFileName := "module/" + m.GoName + Impl + ".go"
	if dryRun {
		fmt.Printf("generate file: %s\n", targetFileName)
		err = t.Execute(os.Stdout, m)
		return

	}

	o, err := os.Create(targetFileName)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}
	defer o.Close()
	err = t.Execute(o, m)
	fmt.Printf("generate file: %s\n", targetFileName)

	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
}
