package handler

import (
    "context"
    "fmt"
    "net/http"
    "strconv"
	"io/ioutil"
    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "gitlab.com/idoko/bucketeer/db"
    "gitlab.com/idoko/bucketeer/models"
)

//var itemIDKey = "itemID"
func files(router chi.Router) {
	router.Put("/upload", FileUpload)                     // REDIRECIONA PARA O UPLOAD DO ARQUIVO
}

// File upload method
func FileUpload(w http.ResponseWriter, r *http.Request) {

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
	PgUploadData(fmt.Sprintf("%s", tempFile.Name()))
}