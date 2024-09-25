package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func main() {
	opts := &protogen.Options{}
	opts.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if strings.Contains(*f.Proto.Name, "resource") {

				generateDomainModel(gen, f)
				continue

			}
			if strings.Contains(*f.Proto.Name, "rpc") {
				generateSchema(gen, f)
				continue
			}
			if !strings.Contains(*f.Proto.Name, "rpc") && !strings.Contains(*f.Proto.Name, "resource") {

				generateApiClient(gen, gen.Files)
			}
		}
		return nil
	})
}

func generateSchema(gen *protogen.Plugin, file *protogen.File) {
	// "domain/[package]/schema.ts" にファイルを生成
	packageName := strings.Replace(strings.ToLower(strings.Replace(string(*file.Proto.Package), "api.", "", -1)), ".rpc", "", -1)
	outputPath := filepath.Join("domain", packageName, "schema.ts")
	if _, err := os.Stat(outputPath); err == nil {
		// ファイルが既に存在する場合、生成をスキップ
		return
	}
	g := gen.NewGeneratedFile(outputPath, "")

	// import 文を先に書き出すためのマップ
	imports := map[string]bool{}
	usedTypes := map[string]bool{}

	g.P("// Code generated by protoc-gen-ts. DO NOT EDIT.")
	g.P(fmt.Sprintf("// source: %s", file.Desc.Path()))

	// 各メッセージのフィールドを処理し、メッセージ型を使っている場合に import を追加
	for _, message := range file.Messages {
		for _, field := range message.Fields {
			if field.Message != nil {
				// 他のメッセージ型の場合、import を追加
				importPath := generateImportPath(file, field.Message)
				imports[importPath] = true
				usedTypes[field.Message.GoIdent.GoName] = true
			}
		}
	}

	// import 文をファイルの先頭に書き出し
	for importPath := range imports {
		g.P(fmt.Sprintf("import { %s } from '%s';", getImportedTypes(usedTypes), importPath))
	}

	// スキーマの生成
	for _, message := range file.Messages {
		g.P(fmt.Sprintf("export type %s = {", message.GoIdent.GoName))
		for _, field := range message.Fields {
			if field.Message != nil {
				// メッセージ型のフィールドの場合、型名を出力
				tsType := field.Message.GoIdent.GoName
				arrayContext := func() string {
					if field.Desc.IsList() {
						return tsType + "[]"
					}
					return tsType
				}()

				g.P(fmt.Sprintf("  %s: %s;", field.Desc.Name(), arrayContext))
			} else {
				// プリミティブ型の場合
				tsType := goToTSType(field.Desc.Kind().String())
				arrayContext := func() string {
					if field.Desc.IsList() {
						return tsType + "[]"
					}
					return tsType
				}()

				g.P(fmt.Sprintf("  %s: %s;", field.Desc.Name(), arrayContext))
			}
		}
		g.P("};")
	}
}

// import パスの生成
func generateImportPath(file *protogen.File, message *protogen.Message) string {
	// 他の proto ファイルに含まれる message 型の場合の正しいパスを生成
	sourceFile := message.Location.SourceFile
	packageName := strings.Replace(strings.Split(sourceFile, "/")[strings.Count(sourceFile, "/")], ".proto", "", -1)
	return fmt.Sprintf("../../domain/%s/model", strings.ToLower(packageName))
}

// インポートする型名を抽出
func getImportedTypes(usedTypes map[string]bool) string {
	types := ""
	for t := range usedTypes {
		types += t + ", "
	}
	return strings.TrimRight(types, ", ")
}

// Goの型をTypeScriptの型に変換する関数
func goToTSType(goType string) string {
	switch goType {
	case "string":
		return "string"
	case "int32", "int64", "uint32", "uint64":
		return "number"
	case "bool":
		return "boolean"
	default:
		return "any"
	}
}
func generateDomainModel(gen *protogen.Plugin, file *protogen.File) {
	// "domain/[package]/model.ts" にファイルを生成
	packageName := strings.Replace(strings.ToLower(strings.Replace(string(*file.Proto.Package), "api.", "", -1)), ".resources", "", -1)
	outputPath := filepath.Join("domain", packageName, "model.ts")
	if _, err := os.Stat(outputPath); err == nil {
		// ファイルが既に存在する場合、生成をスキップ
		return
	}
	g := gen.NewGeneratedFile(outputPath, "")

	// import 文を先に書き出すためのマップ
	imports := map[string]bool{}
	usedTypes := map[string]bool{}

	g.P("// Code generated by protoc-gen-ts. DO NOT EDIT.")
	g.P(fmt.Sprintf("// source: %s", file.Desc.Path()))

	// 各メッセージのフィールドを処理し、メッセージ型を使っている場合に import を追加
	for _, message := range file.Messages {
		for _, field := range message.Fields {
			if field.Message != nil {
				// 他のメッセージ型の場合、import を追加
				importPath := generateImportPath(file, field.Message)
				imports[importPath] = true
				usedTypes[field.Message.GoIdent.GoName] = true
			}
		}
	}

	// import 文をファイルの先頭に書き出し
	for importPath := range imports {
		g.P(fmt.Sprintf("import { %s } from '%s';", getImportedTypes(usedTypes), importPath))
	}

	// 各メッセージの型定義を生成
	for _, message := range file.Messages {
		g.P(fmt.Sprintf("export type %s = {", message.GoIdent.GoName))
		for _, field := range message.Fields {
			if field.Message != nil {
				// メッセージ型のフィールドの場合、型名を出力
				tsType := field.Message.GoIdent.GoName
				arrayContext := func() string {
					if field.Desc.IsList() {
						return tsType + "[]"
					}
					return tsType
				}()

				g.P(fmt.Sprintf("  %s: %s;", field.Desc.Name(), arrayContext))
			} else {
				// プリミティブ型の場合
				tsType := goToTSType(field.Desc.Kind().String())
				arrayContext := func() string {
					if field.Desc.IsList() {
						return tsType + "[]"
					}
					return tsType
				}()

				g.P(fmt.Sprintf("  %s: %s;", field.Desc.Name(), arrayContext))
			}
		}
		g.P("}")
	}
}

func generateApiClient(gen *protogen.Plugin, files []*protogen.File) {
	// "apiclient/index.ts" にファイルを生成
	outputPath := filepath.Join("index.ts")
	if _, err := os.Stat(outputPath); err == nil {
		// ファイルが既に存在する場合、生成をスキップ
		return
	}
	g := gen.NewGeneratedFile(outputPath, "")

	// import 文を先に生成
	g.P("// Code generated by protoc-gen-ts. DO NOT EDIT.")
	for _, file := range files {
		packageName := strings.ToLower(string(*file.Proto.Package))
		if !strings.Contains(file.GoImportPath.String(), "tosaken1116") {
			continue
		}
		if strings.Contains(packageName, "rpc") || strings.Contains(packageName, "resource") {
			continue // rpc や resource を含むパッケージをスキップ
		}
		packageName = strings.Replace(packageName, "api.", "", -1)
		g.P(fmt.Sprintf("import type * as %sSchema from './domain/%s/schema';", ToUpperCamelCase(packageName), packageName))
	}
	g.P("")

	// apiClient 関数の生成
	g.P("export const apiClient = (baseUrl: string) => ({")
	for _, file := range files {
		if !strings.Contains(file.GoImportPath.String(), "tosaken1116") {
			continue
		}
		packageName := strings.ToLower(string(*file.Proto.Package))
		if strings.Contains(packageName, "rpc") || strings.Contains(packageName, "resource") {
			continue // rpc や resource を含むパッケージをスキップ
		}
		packageName = strings.Replace(packageName, "api.", "", -1)
		g.P(fmt.Sprintf("  %s: {", packageName))
		for _, service := range file.Services {
			for _, method := range service.Methods {
				// HTTP メソッドと URL パスを取得
				httpRule := getHttpRule(method)
				httpMethod := getHttpMethod(httpRule)
				urlPath := applyPathParamsToURL(getUrlPath(httpRule))

				// メソッド名やリクエスト/レスポンス型の生成
				reqType := fmt.Sprintf("%sSchema.%s", ToUpperCamelCase(packageName), method.Input.GoIdent.GoName)
				respType := fmt.Sprintf("%sSchema.%s", ToUpperCamelCase(packageName), method.Output.GoIdent.GoName)
				methodName := strings.ToLower(string(method.GoName[0])) + method.GoName[1:]

				// 各APIメソッドの生成
				g.P(fmt.Sprintf("		%s: async (req: %s): Promise<%s> => {", methodName, reqType, respType))
				g.P(fmt.Sprintf("			const res = await fetch(`${baseUrl}%s`, {", urlPath))
				g.P(fmt.Sprintf("				method: '%s',", httpMethod))
				g.P("				headers: { 'Content-Type': 'application/json' },")
				if httpMethod != "GET" {
					g.P("				body: JSON.stringify(req)")
				}
				g.P("			});")
				g.P("			if (!res.ok) {")
				g.P("				throw new Error('Network response was not ok');")
				g.P("			}")
				g.P("			return await res.json();")
				g.P("		},")
			}
		}
		g.P("  },")
	}
	g.P("});")
	g.P("export type ApiClient = ReturnType<typeof apiClient>")
}

// HTTP ルールを取得するための関数
func getHttpRule(method *protogen.Method) *annotations.HttpRule {
	if proto.HasExtension(method.Desc.Options(), annotations.E_Http) {
		ext := proto.GetExtension(method.Desc.Options(), annotations.E_Http)
		if httpRule, ok := ext.(*annotations.HttpRule); ok {
			return httpRule
		}
	}
	return nil
}

// HTTP メソッドを取得するための関数
func getHttpMethod(httpRule *annotations.HttpRule) string {
	if httpRule == nil {
		return "POST"
	}
	switch httpRule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return "GET"
	case *annotations.HttpRule_Post:
		return "POST"
	case *annotations.HttpRule_Put:
		return "PUT"
	case *annotations.HttpRule_Delete:
		return "DELETE"
	default:
		return "POST"
	}
}

// URL パスを取得するための関数
func getUrlPath(httpRule *annotations.HttpRule) string {
	if httpRule == nil {
		return "/defaultPath"
	}
	switch pattern := httpRule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return pattern.Get
	case *annotations.HttpRule_Post:
		return pattern.Post
	case *annotations.HttpRule_Put:
		return pattern.Put
	case *annotations.HttpRule_Delete:
		return pattern.Delete
	default:
		return "/defaultPath"
	}
}

func applyPathParamsToURL(path string) string {
	if !strings.Contains(path, "{") {
		return path
	}

	properties := strings.Split(path, "/")
	for i, prop := range properties {
		if strings.Contains(prop, "{") {
			properties[i] = "${req." + strings.Trim(prop, "{}") + "}"
		}
	}
	return strings.Join(properties, "/")
}
func ToUpperCamelCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
