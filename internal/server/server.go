package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)


type PageData struct {
	Title string
	Component string
  Version string
}

func getJsonResponseBytes(detail string, code int) ([]byte, error) {
	statusCode := fmt.Sprintf("%d", code)
	response := map[string]string{"status": statusCode, "detail": detail}
	jsonResponse, _ := json.Marshal(response)
	return jsonResponse, nil
}

func invalidMethod(w http.ResponseWriter) {
	detail := "Invalid method"
	statusCode := http.StatusMethodNotAllowed

	http.Error(w, detail, http.StatusMethodNotAllowed)
	msg, _ := getJsonResponseBytes(detail, statusCode)
	w.Write(msg)

	return
}

func faviconHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		invalidMethod(w)
	}

	favicon, err := os.ReadFile("assets/favicon.png")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(favicon)
}

func logoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		invalidMethod(w)
	}

	logo, err := os.ReadFile("assets/logo.png")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(logo)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		invalidMethod(w)
	}

  data := PageData{
    Title: "imager",
    Component: "backend:v2",
    Version: "1.0",
  }
  
  tmpl, err := template.ParseFiles("web/index.html") 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func Run(address string) {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/logo", logoHandler)
	http.HandleFunc("/favicon", faviconHandler)
	http.ListenAndServe(address, nil)
}
