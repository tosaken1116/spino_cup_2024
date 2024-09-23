package main

import "google.golang.org/protobuf/compiler/protogen"

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, protoFile := range gen.Files {
			if !protoFile.Generate {
				continue
			}
			generateFile(gen, protoFile)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + ".go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("package ", file.GoPackageName)

	for _, message := range file.Messages {
		g.P("type ", message.GoIdent.GoName, " struct {")
		for _, field := range message.Fields {
			g.P(field.GoName, " ", field.Desc.Kind())
		}
		g.P("}")
	}
}
