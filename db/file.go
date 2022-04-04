package db

import (
    "database/sql"
    "gitlab.com/idoko/bucketeer/models"
	"fmt"
	"path/filepath"
	"os"
	"log"
	//"database/sql"
)


/*func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
    item := models.Item{}
    query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
    err := db.Conn.QueryRow(query, itemData.Name, itemData.Description, itemId).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return item, ErrNoMatch
        }
        return item, err
    }
    return item, nil
}*/


// Load file into imported_files table
func (db Database) PgUploadData(importedFile string) error {
	/*psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)
	defer db.Close()*/

	file := fmt.Sprintf("/tmp/%s", filepath.Base(importedFile))
	os.Chmod(file, 0777)
	//stats, err := os.Stat(file)

	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Permission File After: %s\n", stats.Mode())

	// Copy data to DB
	query := `COPY imported_files(txt_dados_cliente) FROM $1;`
	_, err := db.Conn.Exec(query, importedFile)
    switch err {
    case sql.ErrNoRows:
        return ErrNoMatch
    default:
        return err
    }
	//sql := fmt.Sprintf(`COPY imported_files(txt_dados_cliente) FROM '%s';`, file)

	// Execute statement
	//sqlStatement, err := db.Exec(sql)
	//if err != nil {
	//	fmt.Println(sqlStatement, err)
	//}

	// Show how many rows was affected
	//affected, err := sqlStatement.RowsAffected()

	//fmt.Println("Inserted Records: ", affected)
}