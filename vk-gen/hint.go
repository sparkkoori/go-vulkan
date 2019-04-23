package main

import (
	"strings"

	cast "github.com/elliotchance/c2go/ast"
	"github.com/jinzhu/inflection"
)

func analyzeHint(h *hint, src Source) {
	markArray := func(id string, bl bool) {
		if bl {
			h.isArray[id] = true
		} else {
			delete(h.isArray, id)
		}
	}

	markArraySize := func(id string, bl bool) {
		if bl {
			h.isArraySize[id] = true
		} else {
			delete(h.isArraySize, id)
		}
	}

	analyzeArrayAndSize := func(decls []*cast.FieldDecl, pid string) {
		/*Size of Array Conditions:
		- It's uint32_t or size_t type.
		- The name contains suffix of "Size" or "Count".
		- It's follow by Array.
		*/

		/*Array Conditions:
		- It's next to an size param or array param.
		- It's pointer type.
		- The name is not singular form.
		*/

		nextToSizeOrArray := false
		for i := 0; i < len(decls); i++ {
			name := decls[i].Name
			id := pid + "." + name

			isArray := func(decl *cast.FieldDecl, nextToSizeOrArray bool) bool {
				name := decl.Name

				//Cond 1
				{
					if !nextToSizeOrArray {
						return false
					}
				}

				//Cond 2
				{
					var ok bool
					switch n := decl.ChildNodes[0].(type) {
					case *cast.PointerType:
						_ = n
						ok = true
					}
					if !ok {
						return false
					}
				}

				//Cond 3
				if inflection.Plural(name) != name {
					return false
				}

				return true
			}

			isArraySize := func(decl, ndecl *cast.FieldDecl) (bool, bool) {
				name := decl.Name
				var in bool = false
				//Cond 1
				{
					var ok bool
					switch n := decl.ChildNodes[0].(type) {
					case *cast.PointerType:
						if n.Type == "uint32_t *" || n.Type == "size_t *" {
							ok = true
							in = false
						}
					case *cast.TypedefType:
						if n.Type == "uint32_t" || n.Type == "size_t" {
							ok = true
							in = true
						}
					}
					if !ok {
						return false, in
					}
				}
				//Cond 2
				{
					check := func(suffix, name string) bool {
						if name != strings.ToLower(suffix) && !strings.HasSuffix(name, suffix) {
							return false
						}
						// name = strings.TrimSuffix(name, suffix)
						// if inflection.Singular(name) != name {
						// 	// println(inflection.Singular(name), "!=", name)
						// 	if !strings.Contains(name, "Data") && !strings.Contains(name, "data") {
						// 		return false
						// 	}
						// }

						return true
					}

					if !check("Size", name) && !check("Count", name) {
						return false, false
					}
				}

				//Cond 3
				{
					if ndecl == nil {
						return false, false
					}

					// println("isArray", decl.Name, ndecl.Name, isArray(ndecl, true))
					if !isArray(ndecl, true) {
						return false, false
					}
				}

				return true, in
			}

			decl := decls[i]
			var ndecl *cast.FieldDecl = nil
			if i < len(decls)-1 {
				ndecl = decls[i+1]
			}

			if ok, in := isArraySize(decl, ndecl); ok {
				if in {
					markArraySize(id, true)
				}
				nextToSizeOrArray = true
				continue
			}

			if isArray(decl, nextToSizeOrArray) {
				markArray(id, true)
				continue
			}

			{
				nextToSizeOrArray = false
				continue
			}
		}

		//Pass 2: explicty
		for i := 0; i < len(decls); i++ {
			name := decls[i].Name
			id := pid + "." + decls[i].Name

			switch name {
			case "pPipelines":
				markArray(id, true)
			case "pCommandBuffers":
				markArray(id, true)
			case "pData":
				markArray(id, true)
			case "pCode":
				markArray(id, true)
			}

			switch id {
			case "VkDeviceCreateInfo.pEnabledFeatures":
				markArray(id, false)
			case "VkSubmitInfo.pWaitDstStageMask":
				markArray(id, true)

			case "PFN_vkCreateSharedSwapchainsKHR().swapchainCount":
				markArraySize(id, false)
			case "PFN_vkCreateSharedSwapchainsKHR().pSwapchains":
				markArray(id, true)
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
