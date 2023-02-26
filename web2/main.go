package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	errorCheck(err)
	err = parsedTemplate.Execute(w, nil)
	errorCheck(err)

}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, request *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func addHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello internet\n")
	sum := getSum(5, 4)
	output := fmt.Sprintf("5 + 4 = %d\n", sum)
	write(writer, output)
}

func divideHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello internet\n")
	result, err := getQuotient(5, 0)
	if err != nil {
		output := fmt.Sprintf("Sorry, %s", err)
		write(writer, output)
	} else {
		output := fmt.Sprintf("5 / 4 = %.2f\n", result)
		write(writer, output)
	}

}

func getSum(x, y int) int {
	return x + y
}

func getQuotient(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	return x / y, nil
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getsum", addHandler)
	http.HandleFunc("/getdivide", divideHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	errorCheck(err)
	fmt.Println("App started")

}
