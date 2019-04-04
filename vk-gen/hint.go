package main

import (
	cast "github.com/elliotchance/c2go/ast"
	"github.com/jinzhu/inflection"
)

func isMaybePlural(name string) bool {
	return inflection.Plural(name) == name
}

func analyzeHint(h *hint, src Source) {
	analyzeArrayAndSize := func(decls []*cast.FieldDecl, pid string) {
		/*Array Conditions:
		- It's pointer type.
		- It may be a plural name.
		- There must be one size of array, at least.
		- It's not explicitly excluded.
		*/

		/*Size of Array Conditions:
		- It's uint32_t or size_t type.
		- What next to it is an array.
		// - The name is singular form of array with suffix of "Size" or "Count".
		*/

		existArraySize := false
		for i := len(decls) - 1; i >= 0; i-- {
			decl := decls[i]
			id := pid + "." + decl.Name
			nid := ""
			if i < len(decls)-1 {
				nid = pid + "." + decls[i+1].Name
			}

			//array
			if _, ok := decl.ChildNodes[0].(*cast.PointerType); ok {
				switch decl.Name {
				case "pEnabledFeatures":
				case "pWaitDstStageMask":
					h.isArray[id] = true
				case "pData":
					h.isArray[id] = true
				case "pCode":
					h.isArray[id] = true
				default:
					if isMaybePlural(decl.Name) {
						h.isArray[id] = true
					}
				}
			}

			//size
			{
				if id == "VkShaderModuleCreateInfo.codeSize" {
					delete(h.isArray, id)
					existArraySize = true
					continue
				}
				switch n := decl.ChildNodes[0].(type) {
				case *cast.PointerType:
					if n.Type == "uint32_t *" || n.Type == "size_t *" {
						if nid != "" && h.isArray[nid] {
							delete(h.isArray, id)
							existArraySize = true
						}
					}
				case *cast.TypedefType:
					if n.Type == "uint32_t" || n.Type == "size_t" {
						if nid != "" && h.isArray[nid] {
							delete(h.isArray, id)
							existArraySize = true
							h.isArraySize[id] = true
						}
					}
				}
			}
		}

		//There must be one size of array, at least.
		if !existArraySize {
			for i := len(decls) - 1; i >= 0; i-- {
				id := pid + "." + decls[i].Name
				delete(h.isArray, id)
			}
		}
	}

	for _, node := range src {
		switch n := node.(type) {
		case *cast.TypedefDecl:
		case *cast.RecordDecl:
			{
				decls := []*cast.FieldDecl{}
				for _, child := range n.ChildNodes {
					decls = append(decls, child.(*cast.FieldDecl))
				}
				analyzeArrayAndSize(decls, n.Name)
			}
		case *cast.EnumDecl:
		case *cast.FunctionDecl:
			{
				decls := []*cast.FieldDecl{}
				argNames := []string{}
				for _, param := range n.ChildNodes[1:] {
					pvd, ok := param.(*cast.ParmVarDecl)
					if !ok {
						break
					}
					argNames = append(argNames, pvd.Name)
					decls = append(decls, &cast.FieldDecl{
						Name:       pvd.Name,
						ChildNodes: pvd.ChildNodes,
					})
				}
				analyzeArrayAndSize(decls, n.Name+"()")
				analyzeArrayAndSize(decls, "PFN_"+n.Name+"()")
				h.argNames["PFN_"+n.Name+"()"] = argNames
			}
		default:
			halt("Unkown node in source", node)
		}
	}
	//
	// for k, v := range m {
	// 	fmt.Printf("%s: %#v \n", k, v)
	// }
}
