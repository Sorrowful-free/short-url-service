package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	generateComment = "// generate:reset"
	genFileName     = "reset.gen.go"
	goFileExt       = ".go"
	dotDir          = "."
	vendorDir       = "vendor"
	nodeModulesDir  = "node_modules"
	pathSeparator   = "/"
	defaultReceiver = "s"
)

var defaultValues = map[string]string{
	"int":        "0",
	"int8":       "0",
	"int16":      "0",
	"int32":      "0",
	"int64":      "0",
	"uint":       "0",
	"uint8":      "0",
	"uint16":     "0",
	"uint32":     "0",
	"uint64":     "0",
	"uintptr":    "0",
	"byte":       "0",
	"rune":       "0",
	"float32":    "0",
	"float64":    "0",
	"complex64":  "0",
	"complex128": "0",
	"string":     `""`,
	"bool":       "false",
}

func main() {
	rootDir := dotDir
	if len(os.Args) > 1 {
		rootDir = os.Args[1]
	}

	if err := processDirectory(rootDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func processDirectory(rootDir string) error {
	packageMap := make(map[string]*packageInfo)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if strings.HasPrefix(info.Name(), dotDir) && info.Name() != dotDir {
				return filepath.SkipDir
			}
			if info.Name() == vendorDir || info.Name() == nodeModulesDir {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(path, goFileExt) || strings.HasSuffix(path, genFileName) {
			return nil
		}

		return processFile(path, packageMap)
	})

	if err != nil {
		return err
	}

	for pkgPath, pkgInfo := range packageMap {
		if len(pkgInfo.structs) > 0 {
			if err := generateResetFile(pkgPath, pkgInfo); err != nil {
				return fmt.Errorf("failed to generate file for package %s: %w", pkgPath, err)
			}
		}
	}

	return nil
}

type packageInfo struct {
	packageName string
	structs     []*structInfo
	imports     map[string]string
}

type structInfo struct {
	name   string
	fields []*fieldInfo
}

type fieldInfo struct {
	name     string
	typeExpr ast.Expr
	isPtr    bool
}

type templateData struct {
	PackageName string
	Structs     []*structTemplateData
}

type structTemplateData struct {
	Name         string
	ReceiverName string
	Fields       []*fieldTemplateData
}

type fieldTemplateData struct {
	Access           string
	TypeName         string
	IsPtr            bool
	IsPrimitive      bool
	IsArray          bool
	IsSlice          bool
	ElementType      string
	ElementResetCode string
	ResetCode        string
}

type interfaceResetData struct {
	Access       string
	TypeAssert   string
	NeedNilCheck bool
	Indent       string
}

func processFile(filePath string, packageMap map[string]*packageInfo) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil
	}

	dir := filepath.Dir(filePath)
	if packageMap[dir] == nil {
		packageMap[dir] = &packageInfo{
			packageName: file.Name.Name,
			structs:     []*structInfo{},
			imports:     make(map[string]string),
		}
	}

	pkgInfo := packageMap[dir]

	for _, imp := range file.Imports {
		importPath := strings.Trim(imp.Path.Value, `"`)
		importName := ""
		if imp.Name != nil {
			importName = imp.Name.Name
		} else {
			parts := strings.Split(importPath, pathSeparator)
			importName = parts[len(parts)-1]
		}
		pkgInfo.imports[importPath] = importName
	}

	ast.Inspect(file, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			return true
		}

		hasGenerateComment := false
		if genDecl.Doc != nil {
			for _, comment := range genDecl.Doc.List {
				if strings.TrimSpace(comment.Text) == generateComment {
					hasGenerateComment = true
					break
				}
			}
		}

		if !hasGenerateComment {
			return true
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			structInfo := &structInfo{
				name:   typeSpec.Name.Name,
				fields: []*fieldInfo{},
			}

			if structType.Fields != nil {
				for _, field := range structType.Fields.List {
					fieldType := field.Type
					isPtr := false

					if ptrType, ok := fieldType.(*ast.StarExpr); ok {
						fieldType = ptrType.X
						isPtr = true
					}

					if len(field.Names) > 0 {
						for _, name := range field.Names {
							structInfo.fields = append(structInfo.fields, &fieldInfo{
								name:     name.Name,
								typeExpr: fieldType,
								isPtr:    isPtr,
							})
						}
					} else {
						if ident, ok := fieldType.(*ast.Ident); ok {
							structInfo.fields = append(structInfo.fields, &fieldInfo{
								name:     ident.Name,
								typeExpr: fieldType,
								isPtr:    isPtr,
							})
						}
					}
				}
			}

			pkgInfo.structs = append(pkgInfo.structs, structInfo)
		}

		return true
	})

	return nil
}

const fileTemplate = `// Code generated by cmd/reset. DO NOT EDIT.
//go:build !ignore_autogenerated

package {{.PackageName}}

{{range .Structs}}
{{template "resetMethod" .}}
{{end}}
`

const resetMethodTemplate = `{{define "resetMethod"}}func ({{.ReceiverName}} *{{.Name}}) Reset() {
	if {{.ReceiverName}} == nil {
		return
	}
{{range .Fields}}
{{if .IsArray}}{{template "arrayReset" .}}{{else}}{{.ResetCode}}{{end}}{{end}}
}
{{end}}`

const arrayResetTemplate = `{{define "arrayReset"}}{{if .IsSlice}}
	{{.Access}} = {{.Access}}[:0]
{{else}}
	{{template "arrayLoop" .}}
{{end}}{{end}}`

const arrayLoopTemplate = `{{define "arrayLoop"}}	for i := range {{.Access}} {
{{.ElementResetCode}}	}
{{end}}`

const interfaceResetCheckTemplate = `{{if .NeedNilCheck}}{{.Indent}}if {{.Access}} != nil {
{{end}}{{.Indent}}if resetter, ok := {{.TypeAssert}}.(interface{ Reset() }); ok {
{{.Indent}}	resetter.Reset()
{{.Indent}}}
{{if .NeedNilCheck}}{{.Indent}}}
{{end}}
`

func generateResetFile(pkgPath string, pkgInfo *packageInfo) error {
	tmpl := template.New("file")

	var err error
	tmpl, err = tmpl.Parse(fileTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse file template: %w", err)
	}

	tmpl, err = tmpl.Parse(resetMethodTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse reset method template: %w", err)
	}

	tmpl, err = tmpl.Parse(arrayResetTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse array reset template: %w", err)
	}

	tmpl, err = tmpl.Parse(arrayLoopTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse array loop template: %w", err)
	}

	tmpl, err = tmpl.Parse(interfaceResetCheckTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse interface reset check template: %w", err)
	}

	templateData := &templateData{
		PackageName: pkgInfo.packageName,
		Structs:     make([]*structTemplateData, 0, len(pkgInfo.structs)),
	}

	for _, structInfo := range pkgInfo.structs {
		structData := buildStructTemplateData(structInfo)
		templateData.Structs = append(templateData.Structs, structData)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, templateData); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format generated code: %w", err)
	}

	genFilePath := filepath.Join(pkgPath, genFileName)
	return os.WriteFile(genFilePath, formatted, 0644)
}

func buildStructTemplateData(structInfo *structInfo) *structTemplateData {
	receiverName := strings.ToLower(string(structInfo.name[0]))
	if receiverName == "" {
		receiverName = defaultReceiver
	}

	structData := &structTemplateData{
		Name:         structInfo.name,
		ReceiverName: receiverName,
		Fields:       make([]*fieldTemplateData, 0, len(structInfo.fields)),
	}

	for _, field := range structInfo.fields {
		fieldData := buildFieldTemplateData(receiverName, field)
		structData.Fields = append(structData.Fields, fieldData)
	}

	return structData
}

func buildFieldTemplateData(receiverName string, field *fieldInfo) *fieldTemplateData {
	fieldAccess := fmt.Sprintf("%s.%s", receiverName, field.name)
	fieldData := &fieldTemplateData{
		Access: fieldAccess,
		IsPtr:  field.isPtr,
	}

	switch t := field.typeExpr.(type) {
	case *ast.Ident:
		typeName := t.Name
		fieldData.TypeName = typeName
		fieldData.IsPrimitive = isPrimitiveType(typeName)
		fieldData.ResetCode = buildFieldResetCode(fieldAccess, typeName, field.isPtr)

	case *ast.SelectorExpr:
		typeName := t.Sel.Name
		fieldData.TypeName = typeName
		fieldData.IsPrimitive = isPrimitiveType(typeName)
		fieldData.ResetCode = buildFieldResetCode(fieldAccess, typeName, field.isPtr)

	case *ast.ArrayType:
		fieldData.IsArray = true
		fieldData.IsSlice = (t.Len == nil)
		elementTypeName, elementResetCode := buildArrayElementData(fieldAccess+"[i]", t.Elt)
		fieldData.ElementType = elementTypeName
		fieldData.ElementResetCode = elementResetCode

	case *ast.MapType:
		fieldData.ResetCode = buildMapResetCode(fieldAccess)

	case *ast.InterfaceType:
		fieldData.ResetCode = buildInterfaceResetCode(fieldAccess)

	case *ast.ChanType:
		fieldData.ResetCode = ""

	default:
		if ident, ok := t.(*ast.Ident); ok {
			typeName := ident.Name
			fieldData.TypeName = typeName
			fieldData.IsPrimitive = isPrimitiveType(typeName)
			fieldData.ResetCode = buildFieldResetCode(fieldAccess, typeName, field.isPtr)
		} else {
			fieldData.ResetCode = ""
		}
	}

	return fieldData
}

func buildFieldResetCode(fieldAccess, typeName string, isPtr bool) string {
	var sb strings.Builder

	if isPtr {
		sb.WriteString(fmt.Sprintf("\tif %s != nil {\n", fieldAccess))
		if isPrimitiveType(typeName) {
			sb.WriteString(buildPrimitiveResetCode("*"+fieldAccess, typeName))
		} else {
			sb.WriteString(buildInterfaceResetCheckCode(fieldAccess, fmt.Sprintf("interface{}(%s)", fieldAccess), false, "\t\t"))
		}
		sb.WriteString("\t}\n")
	} else {
		if isPrimitiveType(typeName) {
			sb.WriteString(buildPrimitiveResetCode(fieldAccess, typeName))
		} else {
			sb.WriteString(buildInterfaceResetCheckCode(fieldAccess, fmt.Sprintf("interface{}(&%s)", fieldAccess), false, "\t"))
		}
	}

	return sb.String()
}

func isPrimitiveType(typeName string) bool {
	primitives := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"byte", "rune",
		"float32", "float64",
		"complex64", "complex128",
		"string", "bool",
	}
	for _, p := range primitives {
		if typeName == p {
			return true
		}
	}
	return false
}

func buildPrimitiveResetCode(fieldAccess, typeName string) string {
	defaultValue, ok := defaultValues[typeName]
	if !ok {
		return ""
	}
	return fmt.Sprintf("\t%s = %s\n", fieldAccess, defaultValue)
}

func buildArrayElementData(elementAccess string, elementType ast.Expr) (string, string) {
	switch t := elementType.(type) {
	case *ast.Ident:
		typeName := t.Name
		resetCode := buildArrayElementResetCode(elementAccess, elementType)
		return typeName, resetCode
	case *ast.SelectorExpr:
		typeName := t.Sel.Name
		resetCode := buildArrayElementResetCode(elementAccess, elementType)
		return typeName, resetCode
	case *ast.StarExpr:
		resetCode := buildArrayElementResetCode(elementAccess, elementType)
		return "", resetCode
	default:
		resetCode := buildArrayElementResetCode(elementAccess, elementType)
		return "", resetCode
	}
}

func buildArrayElementResetCode(elementAccess string, elementType ast.Expr) string {
	switch t := elementType.(type) {
	case *ast.Ident:
		typeName := t.Name
		if defaultValue, ok := defaultValues[typeName]; ok {
			return fmt.Sprintf("\t\t%s = %s\n", elementAccess, defaultValue)
		}
		return buildInterfaceResetCheckCode(elementAccess, fmt.Sprintf("interface{}(%s)", elementAccess), false, "\t\t")
	case *ast.StarExpr:
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("\t\tif %s != nil {\n", elementAccess))
		sb.WriteString(buildArrayElementResetCode("*"+elementAccess, t.X))
		sb.WriteString("\t\t}\n")
		return sb.String()
	case *ast.SelectorExpr:
		return buildInterfaceResetCheckCode(elementAccess, fmt.Sprintf("interface{}(%s)", elementAccess), false, "\t\t")
	default:
		return buildInterfaceResetCheckCode(elementAccess, fmt.Sprintf("interface{}(%s)", elementAccess), false, "\t\t")
	}
}

func buildMapResetCode(fieldAccess string) string {
	return fmt.Sprintf("\tif %s != nil {\n\t\tclear(%s)\n\t}\n", fieldAccess, fieldAccess)
}

func buildInterfaceResetCode(fieldAccess string) string {
	return buildInterfaceResetCheckCode(fieldAccess, fieldAccess, true, "\t")
}

func buildInterfaceResetCheckCode(access, typeAssert string, needNilCheck bool, indent string) string {
	tmpl := template.Must(template.New("interfaceResetCheck").Parse(interfaceResetCheckTemplate))

	data := interfaceResetData{
		Access:       access,
		TypeAssert:   typeAssert,
		NeedNilCheck: needNilCheck,
		Indent:       indent,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		if needNilCheck { // Fallback
			return fmt.Sprintf("%sif %s != nil {\n%s\tif resetter, ok := %s.(interface{ Reset() }); ok {\n%s\t\tresetter.Reset()\n%s\t}\n%s}\n",
				indent, access, indent, typeAssert, indent, indent, indent)
		}
		return fmt.Sprintf("%sif resetter, ok := %s.(interface{ Reset() }); ok {\n%s\tresetter.Reset()\n%s}\n",
			indent, typeAssert, indent, indent)
	}

	return buf.String()
}
