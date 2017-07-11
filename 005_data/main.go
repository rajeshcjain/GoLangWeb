package main

import (
	"text/template"
	"os"
)

var tpl *template.Template


func init(){

	tpl = template.Must(template.ParseFiles("tpl.html"))
}
/*
So in tpl.html to put the current data,we represent it as "{{.}}".So whatever is the current value of
the data it will be populated there.If we range it through different numbers..then on each iteration
it will be changed.

The important point is that we can pass only single piece of data;So we can pass
int,string,char or ant structure or map or slice here.This is a restriction of it.
*/
func main(){
	_ = tpl.ExecuteTemplate(os.Stdout,"tpl.html"," meditate")
}
