package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Load file into imported_files table
func (db Database) PgUploadData(importedFile string) error {
	file := fmt.Sprintf("/tmp/%s", filepath.Base(importedFile))
	os.Chmod(file, 0777)
	stats, err := os.Stat(file)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Permission File After: %s\n", stats.Mode())

	// Copy data to DB
	query := fmt.Sprintf("COPY imported_files(txt_dados_cliente) FROM '%s';", importedFile)
	_, dbErr := db.Conn.Exec(query)

	switch dbErr {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return dbErr
	}
}
