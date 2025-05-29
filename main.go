package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const version = "v0.0.0"

func main() {
	flag.Parse()
	if flag.Arg(0) == "version" {
		fmt.Println(version)
		os.Exit(0)
	}

	fileSet := token.NewFileSet()
	err := filepath.WalkDir(".", func(path string, entry os.DirEntry, err error) error {
		if err != nil || entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		file, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
		if err != nil {
			return nil
		}
		typedConstsToEnumerate := map[string][]string{}
		ast.Inspect(file, func(node ast.Node) bool {
			decl, ok := node.(*ast.GenDecl)
			if !ok || decl.Tok != token.TYPE {
				return true
			}
			spec, ok := decl.Specs[0].(*ast.TypeSpec)
			if ok && spec.Name.IsExported() && hasEnumerateAnnotation(spec.Comment) {
				typedConstsToEnumerate[spec.Name.Name] = []string{}
			}

			return false
		})

		if len(typedConstsToEnumerate) == 0 {
			return nil
		}

		ast.Inspect(file, func(node ast.Node) bool {
			decl, ok := node.(*ast.GenDecl)
			if !ok || decl.Tok != token.CONST {
				return true
			}
			for _, spec := range decl.Specs {
				value, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				ident, ok := value.Type.(*ast.Ident)
				if !ok {
					continue
				}

				typ := ident.Name
				if _, ok := typedConstsToEnumerate[typ]; ok {
					for _, name := range value.Names {
						typedConstsToEnumerate[typ] = append(typedConstsToEnumerate[typ], name.Name)
					}
				}
			}

			return false
		})
		pkg := file.Name.Name
		dir := filepath.Dir(path)
		enumeratedFileName := fmt.Sprintf(`enumerated_%s`, filepath.Base(path))
		outputPath := filepath.Join(dir, enumeratedFileName)
		writeOutput(pkg, typedConstsToEnumerate, outputPath)

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func hasEnumerateAnnotation(commentGroup *ast.CommentGroup) bool {
	if commentGroup == nil {
		return false
	}
	for _, comment := range commentGroup.List {
		if strings.Contains(comment.Text, "@enumerate") {
			return true
		}
	}
	return false
}

func writeOutput(pkg string, typedConstsToEnumerate map[string][]string, outputPath string) {
	enumeratedBlocks := make([]string, len(typedConstsToEnumerate))
	for typ, constNames := range typedConstsToEnumerate {
		lines := make([]string, len(constNames))
		for i, constName := range constNames {
			lines[i] = "\t" + constName + ","
		}
		enumeratedBlocks = append(enumeratedBlocks, fmt.Sprintf(`
			var Enumerated%ss = []%s{
				%s
			}

			func Is%s(a any) bool {
				v, ok := a.(%s)
				return ok && slices.Contains(Enumerated%ss, v)
			}`, typ, typ, strings.Join(lines, "\n"), typ, typ, typ))
	}

	code := fmt.Sprintf(`
			package %s

			import "slices"

			%s`, pkg, strings.Join(enumeratedBlocks, "\n"))

	formattedCode, err := format.Source([]byte(code))
	if err != nil {
		fmt.Println("format error:", err)
		return
	}

	os.WriteFile(outputPath, formattedCode, 0o644)
}
