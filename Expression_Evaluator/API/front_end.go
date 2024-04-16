package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func getHtmlFile(writer http.ResponseWriter, req *http.Request) {
	template, errTemplate := template.ParseFiles("templates/index.tpl")

	if errTemplate != nil {
		fmt.Fprintln(writer, errTemplate.Error())
		return
	}

	var data string = ""

	template.Execute(writer, data)
}
