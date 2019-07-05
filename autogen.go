package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type MetaData struct{
	DbObj string
	DbType string
	Operation string
}

func main(){
	metadata := MetaData{
		DbObj: "mongo",
		DbType: "MongoDB",
		Operation: "default",
	}
	file,err := os.Create("./handlers/"+metadata.Operation+"Handler.go")
	if err != nil {
		log.Fatal(err)
	}

	templFile, err := ioutil.ReadFile("./templates/HandlerTempl.gotpl")
	tpl,err := template.New("template").Parse(string(templFile))
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(file, metadata)
	if err != nil {
		log.Fatal(err)
	}
}