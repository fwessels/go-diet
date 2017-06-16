package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"
	"unicode"
)

type NameReturnValuesASTVisitor struct {
	scanCode  bool
	info      *types.Info
	funcName  string
	funcEnd   token.Pos
	namedVars map[string]string
	eligible  map[string]bool
}

func nameReturnValues(file *ast.File) *ast.File {

	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	scanner := &NameReturnValuesASTVisitor{info: &info, eligible: make(map[string]bool), scanCode: true}
	ast.Walk(scanner, file)

	if len(scanner.eligible) == 0 {
		fmt.Println("No eligible functions.")
	} else {
		fmt.Println("Eligible functions:")
		for k, _ := range scanner.eligible {
			fmt.Printf("    - %s()\n", k)
		}
	}

	fixer := &NameReturnValuesASTVisitor{info: &info, eligible: scanner.eligible, scanCode: false}
	ast.Walk(fixer, file)

	return file
}

func getShortIdent(ident string) (short string) {

	// Convert consecutive capitals into lower case, i.e. "APIErrorCode" --> ApiErrorCode
	capsStreak, indexStreak := "", 0
	for i, c := range ident {
		if i == 0 {
			continue
		}
		if unicode.IsUpper(c) {
			if len(capsStreak) == 0 {
				indexStreak = i
			}
			capsStreak += string(c)
		} else if unicode.IsLower(c) {
			if len(capsStreak) > 1 {
				ident = ident[:indexStreak] + strings.ToLower(capsStreak[:len(capsStreak)-1]) + capsStreak[len(capsStreak)-1:] + ident[i:]
			}
			capsStreak = ""
		}
	}

	// Create abbreviation but concatenating capital letters, starting with first char
	for i, c := range ident {
		if i == 0 {
			short += string(unicode.ToLower(c))
		} else if unicode.IsUpper(c) {
			short += string(unicode.ToLower(c))
		}
	}
	return
}

func checkHasNamedReturnValues(list []*ast.Field) bool {

	for _, l := range list {
		if len(l.Names) != 0 {
			return true
		}
	}
	return false
}

func (v *NameReturnValuesASTVisitor) Visit(node ast.Node) ast.Visitor {
	if node != nil {
		if v.funcEnd > 0 && node.Pos() > v.funcEnd {
			v.funcName, v.funcEnd = "", 0
		}
		if !v.scanCode && v.funcName != "" && !v.eligible[v.funcName] {
			return v
		}
		//fmt.Printf("%d: %s", nr, reflect.TypeOf(node).String())
		switch n := node.(type) {
		case *ast.FuncDecl:

			if n.Body != nil {
				v.funcEnd = n.Body.Rbrace
				v.funcName = n.Name.Name
				v.namedVars = make(map[string]string)
			} else {
				v.funcName, v.funcEnd = "", 0
				v.namedVars = make(map[string]string)
			}

			if !v.scanCode && !v.eligible[v.funcName] {
				return v
			}

			if n.Type.Results != nil {
				if !checkHasNamedReturnValues(n.Type.Results.List) {

					for _, l := range n.Type.Results.List {
						ident := ""
						switch i := l.Type.(type) {
						case *ast.Ident:
							ident = i.Name
						}
						shortIdent := getShortIdent(ident)
						v.namedVars[ident] = shortIdent

						if len(l.Names) == 0 {
							if !v.scanCode {
								l.Names = append(l.Names, &ast.Ident{Name: shortIdent})
							}
						}
					}

				}
			}

		case *ast.ReturnStmt:
			if len(n.Results) > 0 {
				switch t := n.Results[0].(type) {
				case *ast.CompositeLit:
					switch typ := t.Type.(type) {
					case *ast.Ident:
						short, found := v.namedVars[typ.Name]
						if found && t.Lbrace+1 == t.Rbrace {
							if v.scanCode {
								// Add this function as eligible
								v.eligible[v.funcName] = true
							} else {
								// Modify code to use name of return value
								n.Results[0] = &ast.Ident{Name: short}
							}
						}
					}
				}
			}

		case ast.Expr:
			//t := v.info.TypeOf(node.(ast.Expr))
			//if t != nil {
			//	fmt.Printf(" : %s", t.String())
			//}
		}
	}
	return v
}
