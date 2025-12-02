package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

const (
	panicFuncName = "panic"
	mainPkgName   = "main"
	mainFuncName  = "main"
	logPkgName    = "log"
	fatalFuncName = "Fatal"
	osPkgName     = "os"
	exitFuncName  = "Exit"
)

var analyzer = &analysis.Analyzer{
	Name: "panic_check",
	Doc:  "checks the use of panic, log.Fatal and os.Exit",
	Run:  run,
}

func main() {
	singlechecker.Main(analyzer)
}

func run(pass *analysis.Pass) (interface{}, error) {
	isMainPackage := pass.Pkg.Name() == mainPkgName

	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if checkPanic(callExpr, pass) {
				return true
			}

			if checkLogFatal(callExpr, f, pass, isMainPackage) {
				return true
			}

			if checkOsExit(callExpr, f, pass, isMainPackage) {
				return true
			}

			return true
		})
	}

	return nil, nil
}

func checkPanic(callExpr *ast.CallExpr, pass *analysis.Pass) bool {
	ident, ok := callExpr.Fun.(*ast.Ident)
	if !ok {
		return false
	}

	if ident.Name == panicFuncName {
		pass.Reportf(callExpr.Pos(), "using panic is not allowed")
		return true
	}

	return false
}

func checkLogFatal(callExpr *ast.CallExpr, file *ast.File, pass *analysis.Pass, isMainPackage bool) bool {
	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	pkgIdent, ok := selExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	pkgName := pkgIdent.Name
	funcName := selExpr.Sel.Name

	if pkgName == logPkgName && funcName == fatalFuncName {
		if !isMainPackage || !isInMainFunction(callExpr, file, pass) {
			pass.Reportf(callExpr.Pos(), "calling log.Fatal outside the main function of package main is not allowed")
		}
		return true
	}

	return false
}

func checkOsExit(callExpr *ast.CallExpr, file *ast.File, pass *analysis.Pass, isMainPackage bool) bool {
	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	pkgIdent, ok := selExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	pkgName := pkgIdent.Name
	funcName := selExpr.Sel.Name

	if pkgName == osPkgName && funcName == exitFuncName {
		if !isMainPackage || !isInMainFunction(callExpr, file, pass) {
			pass.Reportf(callExpr.Pos(), "calling os.Exit outside the main function of package main is not allowed")
		}
		return true
	}

	return false
}

func isInMainFunction(callExpr *ast.CallExpr, file *ast.File, pass *analysis.Pass) bool {
	var mainFunc *ast.FuncDecl
	ast.Inspect(file, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			if fn.Name.Name == mainFuncName && fn.Recv == nil {
				mainFunc = fn
				return false
			}
		}
		return true
	})

	if mainFunc == nil {
		return false
	}

	callPos := callExpr.Pos()
	mainStart := mainFunc.Pos()
	mainEnd := mainFunc.End()

	return callPos >= mainStart && callPos <= mainEnd
}
