package htmlParser

import (
	"os"
	"text/template"

	"github.com/google/uuid"
)

type htmlStruct struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
	return &htmlStruct{rootPath: rootPath}
}

func (a *htmlStruct) Create(templateName string, data interface{}) (string, error) {

	templateWriter, err := template.ParseFiles(templateName)
	if err != nil {
		return " ", err
	}
	fileName := a.rootPath + "/" + uuid.New().String() + ".html"
	if err != nil {
		return " ", err
	}
	fileWriter, err := os.Create(fileName)
	if err != nil {
		return " ", err
	}
	err = templateWriter.Execute(fileWriter, data)
	if err != nil {
		return " ", err
	}

	return fileName, nil
}
