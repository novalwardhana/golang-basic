package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan []byte)
		go func() {
			result <- []byte("Welcome to my API service.....")
		}()
		w.Write(<-result)
	})

	http.HandleFunc("/upload-file", routeUploadFile)
	http.HandleFunc("/upload-file-process", routeUploadFileProcess)
	http.HandleFunc("/insert-file", routeInsertFile)
	http.HandleFunc("/insert-file-process", routeInsertFileProcess)

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	log.Printf("Run server at: %s\n", address)
	server.ListenAndServe()
}

func routeUploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl := template.Must(template.New("upload-file").ParseFiles("views/upload-file.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func routeUploadFileProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(dir, "files", part.FileName())
		createFiles, err := os.Create(fileLocation)
		if createFiles != nil {
			defer createFiles.Close()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(createFiles, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte("Success upload all files"))
}

func routeInsertFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl := template.Must(template.New("insert-file").ParseFiles("views/insert-file.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeInsertFileProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	files, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		file, err := files.NextPart()
		if err == io.EOF {
			break
		}

		filedir := filepath.Join(dir, "files", file.FileName())
		createFile, err := os.Create(filedir)
		if createFile != nil {
			defer createFile.Close()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(createFile, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte("File success fully uploaded"))
}
