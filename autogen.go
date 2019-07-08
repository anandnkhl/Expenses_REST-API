package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type MetaData struct{
	DbObj string
	DbType string
	Operation string
}

func main(){

	var dbType string
	var dbObj string
	var operation string

	flag.StringVar(&dbType, "dbtype", "", "enter type of database")
	flag.StringVar(&operation, "op", "", "Enter the CRUD operation")
	flag.Parse()

	if dbType == "" || operation == "" {
		log.Fatal("dbtype/op cannot be nil")
	}

	dbObj = strings.ToLower(dbType)[:1]

	metadata := MetaData{
		DbObj: dbObj,
		DbType: dbType,
		Operation: operation,
	}
	file,err := os.Create("./"+metadata.Operation+"Handler.go")
	if err != nil {
		log.Fatal(err)
	}

	templFile, err := ioutil.ReadFile("../templates/HandlerTempl.gotpl")
	tpl,err := template.New("template").Parse(string(templFile))
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(file, metadata)
	if err != nil {
		log.Fatal(err)
	}
}