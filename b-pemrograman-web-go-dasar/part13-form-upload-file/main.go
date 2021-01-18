package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan []byte)
		go func() {
			data := []byte("Welcome to my API")
			result <- data
		}()
		w.Write(<-result)
	})

	http.HandleFunc("/upload", routeFileUpload)
	http.HandleFunc("/upload-process", routeFileUploadProcess)
	http.HandleFunc("/insert", routeFileInsert)
	http.HandleFunc("/insert-process", routeFileInsertProcess)

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	logrus.Info(fmt.Sprintf("Run server at: %s", address))
	server.ListenAndServe()
}

func routeFileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl := template.Must(template.New("upload").ParseFiles("views/upload.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "No Content", http.StatusNoContent)
	}
}

func routeFileUploadProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	alias := r.FormValue("alias")
	uploadedFile, handler, err := r.FormFile("file") // uploadedFile= lokasi memori file upload; handler= identitas file upload
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd() // direktori aplikasi yang dijalankan
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	filename := handler.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename)) // filepath.Ext untuk melihat ekstensi file yang diupload
	}

	fileLocation := filepath.Join(dir, "files", filename) // membuat direktori file
	fileTarget, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(fileTarget, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Done"))
}

func routeFileInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.New("insert").ParseFiles("views/insert.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	}
}

func routeFileInsertProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	alias := r.FormValue("name")
	uploadedFile, handler, err := r.FormFile("upload_file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	filename := fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filedir := filepath.Join(dir, "files", filename)

	fileTarget, err := os.OpenFile(filedir, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer fileTarget.Close()

	if _, err := io.Copy(fileTarget, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Done"))
}
