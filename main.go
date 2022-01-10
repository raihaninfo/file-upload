package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":8080", r)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		tpm, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}
		tpm.Execute(w, nil)
	}

	if r.Method == "POST" {
		r.ParseMultipartForm(100)
		file, handler, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Println(handler.Header)

		// tempFile, err := ioutil.TempFile("files", "upload-*.png")
		// if err != nil {
		// 	panic(err)
		// }
		// defer tempFile.Close()

		contenType := handler.Header["Content-Type"][0]
		fmt.Println(contenType)
		var tempFile *os.File
		if contenType == "image/jpeg" {
			tempFile, err = ioutil.TempFile("files/JPG", "*.jpg")
		} else if contenType == "application/pdf" {
			tempFile, err = ioutil.TempFile("files/PDFs", "*.pdf")
		} else if contenType == "image/png" {
			tempFile, err = ioutil.TempFile("files/png", "*.png")
		}
		fmt.Println("err", err)
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		tempFile.Write(fileBytes)
		fmt.Fprintf(w, "File uploded")

	}

}
