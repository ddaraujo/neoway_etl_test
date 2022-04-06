package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/ddaraujo/neoway_etl_test/db"

	"github.com/gorilla/mux"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	dbInstance = db

	fmt.Println("Start handling requests")

	myRouter := mux.NewRouter().StrictSlash(true)
	// Landing page (index.html)
	myRouter.HandleFunc("/", landingPage).Methods("GET")

	// Upload file route (handler/files.go)
	myRouter.HandleFunc("/upload", upload).Methods("POST")

	// Valid Items Routes (handler/items.go)
	myRouter.HandleFunc("/items/valid", getAllValidItems).Methods("GET")                                // Get all valid items
	myRouter.HandleFunc("/items/valid/cpf/{cpf}", getAllItemsByCpf).Methods("GET")                      // Get all valid items by CPF
	myRouter.HandleFunc("/items/valid/lastSale/{cnpj}", getAllItemsLastSaleCnpj).Methods("GET")         // Gel all valid items by cnpj (last sale)
	myRouter.HandleFunc("/items/valid/frequentSale/{cnpj}", getAllItemsFrequentSaleCnpj).Methods("GET") // Get all valid items by cnpj (frequent sale)
	myRouter.HandleFunc("/items/valid/count", getValidRecordsCount)                                     // Get all valid records count
	myRouter.HandleFunc("/items/valid/delete", deleteValidItems)                                        // Delete all valid records

	// Invalid Items Routes (handler/items.go)
	myRouter.HandleFunc("/items/invalid", getAllInvalidItems).Methods("GET")
	myRouter.HandleFunc("/items/invalid/count", getInvalidRecordsCount) // Get all invalid (rejected) records count
	myRouter.HandleFunc("/items/invalid/delete", deleteInvalidItems)    // Delete all invalid records

	log.Fatal(http.ListenAndServe(":8888", myRouter))

	return myRouter
}

// Display HTML template
func displayTemplate(w http.ResponseWriter, page string, data interface{}) {
	//dir, _ := os.Getwd()
	//var templates = template.Must(template.ParseFiles(filepath.Join(dir, "handler", "index.html")))
	var templates = template.Must(template.ParseFiles("index.html"))
	templates.ExecuteTemplate(w, page+".html", data)
}

// Show landing page (Template: index.html)
func landingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show landing page")
	displayTemplate(w, "index", nil)
}

// Sanitize REST input data removing - . and /
func sanitize(data string) string {
	data = strings.Replace(data, ".", "", -1)
	data = strings.Replace(data, "-", "", -1)
	data = strings.Replace(data, "/", "", -1)
	return data
}
