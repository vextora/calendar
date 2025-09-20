package scaffold

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateData struct {
	Entity      string
	EntityLower string
	Version     string
	ApiPrefix   string
	TargetDir   string
}

func Generate(version, entity string) error {
	data := TemplateData{
		Entity:      strings.Title(entity),
		EntityLower: strings.ToLower(entity),
		Version:     version,
		ApiPrefix:   "api",
	}

	targetDir := filepath.Join("internal", data.ApiPrefix, data.Version, data.EntityLower)
	data.TargetDir = targetDir

	domainDir := filepath.Join(targetDir, "domain")
	dtoDir := filepath.Join(targetDir, "dto")

	for _, dir := range []string{targetDir, domainDir, dtoDir} {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	files := map[string]string{
		"init.tpl":                      filepath.Join(targetDir, "init.go"),
		"handler.tpl":                   filepath.Join(targetDir, "handler.go"),
		"mapper.tpl":                    filepath.Join(targetDir, "mapper.go"),
		"repository_implementation.tpl": filepath.Join(targetDir, "repository_impl.go"),
		"service_implementation.tpl":    filepath.Join(targetDir, "service_impl.go"),
		"type_alias.tpl":                filepath.Join(targetDir, "type_alias.go"),
		"route.tpl":                     filepath.Join(targetDir, "route.go"),
		"model.tpl":                     filepath.Join(domainDir, "model.go"),
		"repository_interface.tpl":      filepath.Join(domainDir, "repository.go"),
		"service_interface.tpl":         filepath.Join(domainDir, "service.go"),
		"request.tpl":                   filepath.Join(dtoDir, "request.go"),
		"response.tpl":                  filepath.Join(dtoDir, "response.go"),
	}

	for tplName, outputPath := range files {
		fmt.Printf("Generating %v\n", outputPath)
		if err := renderTemplateToFile(tplName, outputPath, data); err != nil {
			return err
		}
	}

	if err := appendImportToRouter(data.ApiPrefix, data.Version, data.Entity); err != nil {
		return err
	}

	fmt.Printf("Scaffolding complete: %s/\n", data.TargetDir)
	return nil
}

func renderTemplateToFile(templateName, outputPath string, data TemplateData) error {
	tplContent, err := os.ReadFile(filepath.Join("pkg/scaffold/templates", templateName))
	if err != nil {
		return err
	}

	tpl, err := template.New(templateName).Parse(string(tplContent))
	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tpl.Execute(f, data)
}

func appendImportToRouter(api, version, entity string) error {
	apiVersion := strings.ToLower(api)
	entityLower := strings.ToLower(entity)
	versionLower := strings.ToLower(version)

	importPath := fmt.Sprintf("\t_ \"calendarapi/internal/%s/%s/%s\"", apiVersion, versionLower, entityLower)
	importFile := "routerloader/imports.go"

	contentBytes, err := ioutil.ReadFile(importFile)
	if err != nil {
		return err
	}
	content := string(contentBytes)

	if strings.Contains(content, importPath) {
		return nil
	}

	lines := strings.Split(content, "\n")
	var output []string
	inserted := false
	inImportBlock := false
	importInserted := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if trimmed == "import (" {
			inImportBlock = true
		}

		if inImportBlock && trimmed == ")" && !importInserted {
			output = append(output, importPath)
			importInserted = true
			inserted = true
		}

		output = append(output, line)
	}

	if !inserted {
		var finalOutput []string
		for _, line := range output {
			finalOutput = append(finalOutput, line)
			if strings.HasPrefix(strings.TrimSpace(line), "package ") {
				finalOutput = append(finalOutput, "\nimport (", importPath, ")")
			}
		}
		output = finalOutput
	}

	fmt.Printf("\nRoute %s registered to : %s\n\n", entityLower, importFile)

	return os.WriteFile(importFile, []byte(strings.Join(output, "\n")), 0644)
}
