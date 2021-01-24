package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type m map[string]interface{}

func main() {
	http.HandleFunc("/", routeIndex)
	http.HandleFunc("/list", routeList)
	http.HandleFunc("/list-get-data", routeListGetData)
	http.HandleFunc("/list-download", routeListDownload)

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	log.Printf("Start service at: %s", address)
	server.ListenAndServe()
}

func routeIndex(w http.ResponseWriter, r *http.Request) {
	result := make(chan []byte)
	go func() {
		result <- []byte("*---Welcome to my API---*")
	}()
	w.Write(<-result)
}

func routeList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl := template.Must(template.New("list").ParseFiles("views/list.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeListGetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	content := []m{}
	basePath, _ := os.Getwd()
	fileDir := filepath.Join(basePath, "files")

	err := filepath.Walk(fileDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		content = append(content, m{"filename": info.Name(), "path": path})
		return nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contentJSON, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(contentJSON)
}

func routeListDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := r.FormValue("path")
	file, err := os.Open(path)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", file.Name())
	w.Header().Set("Content-Disposition", contentDisposition)
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
