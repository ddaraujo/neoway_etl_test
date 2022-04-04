package main
import (
    "context"
    "fmt"
    //"gitlab.com/idoko/bucketeer/db"
    //"gitlab.com/idoko/bucketeer/handler"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)
func main() {
    addr := ":8888"
    listener, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalf("Error occurred: %s", err.Error())
    }
    dbUser, dbPassword, dbName :=
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB")

    database, err := db.Initialize(dbUser, dbPassword, dbName)
    if err != nil {
        log.Fatalf("Could not set up database (error): %v", err)
    }
    defer database.Conn.Close()

    httpHandler := handler.NewHandler(database)
    server := &http.Server{
        Handler: httpHandler,
    }
    go func() {
        server.Serve(listener)
    }()
    defer Stop(server)
    log.Printf("Started server on %s", addr)
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    log.Println(fmt.Sprint(<-ch))
    log.Println("Stopping API server.")
}
func Stop(server *http.Server) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Printf("Could not shut down server correctly: %v\n", err)
        os.Exit(1)
    }
}

/*package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "555407"
	dbname   = "postgres"
	MAX_UPLOAD_SIZE = 1024 * 1024 * 10
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Http Templates
var templates = template.Must(template.ParseFiles("index.html"))


// Display the named template
func displayTemplate(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

// File upload handler 
// GET: forbidden
// POST: upload file method
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit max file size to 10MB
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 10MB in size", http.StatusBadRequest)
		return
	}

	// Go to upload method
	upload(w, r)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Load file into imported_files table
func pgUploadData(importedFile string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)
	defer db.Close()

	file := fmt.Sprintf("/tmp/%s", filepath.Base(importedFile))
	os.Chmod(file, 0777)
	stats, err := os.Stat(file)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Permission File After: %s\n", stats.Mode())

	// Copy data to DB
	sql := fmt.Sprintf(`COPY imported_files(txt_dados_cliente) FROM '%s';`, file)

	// Execute statement
	sqlStatement, err := db.Exec(sql)
	if err != nil {
		fmt.Println(sqlStatement, err)
	}

	// Show how many rows was affected
	affected, err := sqlStatement.RowsAffected()
	checkErr(err)

	fmt.Println("Registros Inseridos: ", affected)
}

// File upload method
func upload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("File Upload Endpoint Hit")
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file
	tempFile, err := ioutil.TempFile("/tmp/", fmt.Sprintf("*-%s", handler.Filename))
	if err != nil {
		fmt.Println(err)
	}

    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }

    // write this byte array to our temporary file
    tempFile.Write(fileBytes)

	// return that we have successfully uploaded our file to server!
	fmt.Println("Successfully Uploaded File to Server")
	fmt.Fprintf(w, "Successfully Uploaded File to Server\n")

	// Save file data to postgres DB
	pgUploadData(fmt.Sprintf("%s", tempFile.Name()))
}

// Show landing page (Template: index.html)
func landingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show landing page")
	displayTemplate(w, "index", nil)
}

// Handle all requests
func setupRoutes() {
	fmt.Println("Start handling requests")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", landingPage)
	myRouter.HandleFunc("/upload", uploadHandler)
	log.Fatal(http.ListenAndServe(":8888", myRouter))
}

func main() {
	fmt.Println("Server started...")

	// Start handling requests
	fmt.Println("Setting up routes...")
	setupRoutes()
}*/