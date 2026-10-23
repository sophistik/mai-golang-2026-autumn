// Команда noprint — пример своего линтера на go/analysis:
// запрещает встроенные print/println (инструменты отладки и REPL);
// в прикладном коде — fmt.Print* или логгер.
//
// Запуск: go run ./04-testing/linter ./04-testing/linter/demo
// Как vet: go build -o noprint ./04-testing/linter
//
//	go vet -vettool=./noprint ./04-testing/...
package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var noprint = &analysis.Analyzer{
	Name: "noprint",
	Doc:  "запрещает print/println — используйте fmt.Print* или логгер",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			id, _ := call.Fun.(*ast.Ident)
			if id != nil && (id.Name == "print" || id.Name == "println") {
				pass.Reportf(call.Pos(),
					"print/println: используйте fmt.Print* или логгер")
			}
			return true
		})
	}
	return nil, nil
}

func main() {
	singlechecker.Main(noprint)
}
