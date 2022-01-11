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
		tpm, err := template.ParseFiles("index.gohtml")
		handleError(err)
		tpm.Execute(w, nil)
	}

	if r.Method == "POST" {
		r.ParseMultipartForm(100)
		file, handler, err := r.FormFile("myFile")
		handleError(err)
		defer file.Close()
		fmt.Println(handler.Header)

		contenType := handler.Header["Content-Type"][0]
		fmt.Println(contenType)
		var tempFile *os.File
		if contenType == "image/jpeg" {
			tempFile, err = ioutil.TempFile("files/JPG", "*.jpg")
		} else if contenType == "application/pdf" {
			tempFile, err = ioutil.TempFile("files/PDFs", "*.pdf")
		} else if contenType == "image/png" {
			tempFile, err = ioutil.TempFile("files/png", "*.png")
		} else if contenType == "video/mp4" {
			tempFile, err = ioutil.TempFile("files/video", "*.mp4")
		}
		fmt.Println("err", err)
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		handleError(err)

		tempFile.Write(fileBytes)
		fmt.Fprintf(w, "File uploded")

	}

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
