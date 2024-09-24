package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

type HandlerInfo struct {
	HandlerName   string
	InterfaceName string
}

type MethodInfo struct {
	HTTPMethod  string
	Path        string
	HandlerName string
	MethodName  string
}

var (
	handlerInfos []HandlerInfo
	methodInfos  []MethodInfo
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		collectHandlerInfos(gen)

		generateFile(handlerInfos, methodInfos)

		return nil
	})
}

func generateFile(handlerInfos []HandlerInfo, methodInfos []MethodInfo) {
	filename := "register_routes.go"
	filePath := filepath.Join("../backend/internal/router", filename)

	// ディレクトリが存在しない場合は作成
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	// ファイルを作成
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()

	// ファイルに書き込む内容を組み立てる
	var content strings.Builder

	content.WriteString("package router\n\n")
	content.WriteString("import (\n")
	content.WriteString(`	"github.com/labstack/echo/v4"` + "\n")
	content.WriteString(`	"github.com/tosaken1116/spino_cup_2024/backend/handler"` + "\n")
	content.WriteString(")\n\n")
	content.WriteString("func registerRoutes(\n")
	content.WriteString("	e *echo.Echo,\n")

	for _, handlerInfo := range handlerInfos {
		handlerArg := fmt.Sprintf("	%s handler.%s,\n", handlerInfo.HandlerName, handlerInfo.InterfaceName)
		content.WriteString(handlerArg)
	}

	content.WriteString(") {\n")

	for _, methodInfo := range methodInfos {
		route := fmt.Sprintf(`	e.%s("%s", %s.%s)`, methodInfo.HTTPMethod, methodInfo.Path, methodInfo.HandlerName, methodInfo.MethodName)
		content.WriteString(route + "\n")
	}

	content.WriteString("}\n")

	// ファイルに書き込む
	_, err = f.WriteString(content.String())
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func collectHandlerInfos(gen *protogen.Plugin) {
	for _, protoFile := range gen.Files {
		if !protoFile.Generate {
			continue
		}

		for _, service := range protoFile.Services {
			serviceName := service.GoName
			handlerName := ToLowerCamelCase(strings.TrimSuffix(serviceName, "Service")) + "Handler"
			interfaceName := ToUpperCamelCase(handlerName)
			handlerInfos = append(handlerInfos, HandlerInfo{
				HandlerName:   handlerName,
				InterfaceName: interfaceName,
			})

			for _, method := range service.Methods {
				httpRule := getHttpRule(method)
				httpMethod := getHttpMethod(httpRule)
				urlPath := convertPathParamsToEchoFormat(getUrlPath(httpRule))
				methodInfos = append(methodInfos, MethodInfo{
					HTTPMethod:  httpMethod,
					Path:        urlPath,
					HandlerName: handlerName,
					MethodName:  method.GoName,
				})
			}
		}
	}
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

func convertPathParamsToEchoFormat(path string) string {
	if !strings.Contains(path, "{") {
		return path
	}

	properties := strings.Split(path, "/")
	for i, prop := range properties {
		if strings.HasPrefix(prop, "{") && strings.HasSuffix(prop, "}") {
			paramName := strings.Trim(prop, "{}")
			properties[i] = ":" + paramName
		}
	}
	return strings.Join(properties, "/")
}

func ToLowerCamelCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func ToUpperCamelCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
