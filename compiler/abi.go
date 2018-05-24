package compiler

import (
	"go/ast"
	"go/token"
	"strings"
)

type Functions struct {
	Functions []Function `json:"functions"`
}

type Function struct {
	Name       string      `json:"name"`
	Parameters []Paramster `json:"parameters"`
	Returntype string      `json:"returntype"`
}

type Paramster struct {
	PName string `json:"name"`
	PType string `json:"type"`
}

func MakeAbi(main *ast.FuncDecl) Functions {

	//make abi
	//step 1. make Main function
	funcs := Functions{}
	var fs []Function

	fmain := Function{Name: main.Name.Name}
	params := make([]Paramster, len(main.Type.Params.List))
	for i, p := range main.Type.Params.List {
		param := Paramster{PName: p.Names[0].Name}

		switch t := p.Type.(type) {
		case *ast.Ident:
			if t.Name == "string" {
				param.PType = "String"
			} else {
				param.PType = t.Name
			}

		case *ast.ArrayType:
			param.PType = "ByteArray"
		}

		params[i] = param
	}
	fmain.Parameters = params
	results := main.Type.Results.List
	if len(results) > 0 {
		switch t := results[0].Type.(type) {
		case *ast.Ident:
			if t.Name == "bool" {
				fmain.Returntype = "Boolean"
			}
		default:
			fmain.Returntype = "unknown"
		}
	}

	fs = append(fs, fmain)

	//step 2. Analyze operation
	for _, statement := range main.Body.List {
		fn := Function{}

		var method string
		var params []Paramster
		switch t := statement.(type) {
		case *ast.IfStmt:
			switch c := t.Cond.(type) {
			case *ast.BinaryExpr:
				//todo define the "operation"
				if c.Op == token.EQL && c.X.(*ast.Ident).Name == "operation" {
					method = strings.Replace(c.Y.(*ast.BasicLit).Value, "\"", "", -1)
				}
			}

			for _, st := range t.Body.List {
				switch stt := st.(type) {
				case *ast.AssignStmt:

					switch tae := stt.Rhs[0].(type) {
					case *ast.TypeAssertExpr:
						//todo defien the "args"
						switch tae.X.(type) {
						case *ast.IndexExpr:
							if tae.X.(*ast.IndexExpr).X.(*ast.Ident).Name == "args" {
								param := Paramster{}

								switch tae.Type.(type) {
								case *ast.Ident:
									param.PName = stt.Lhs[0].(*ast.Ident).Name
									param.PType = tae.Type.(*ast.Ident).Name
								case *ast.ArrayType:
									param.PName = stt.Lhs[0].(*ast.Ident).Name
									param.PType = "ByteArray"
								}
								params = append(params, param)
							}
						}
					}

				}

			}

		}
		fn.Name = method
		fn.Parameters = params
		fn.Returntype = fmain.Returntype
		if fn.Name != "" {
			fs = append(fs, fn)
		}
	}

	funcs.Functions = fs
	//end

	return funcs
}
