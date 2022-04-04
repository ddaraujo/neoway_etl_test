package db

import (
	"database/sql"

	"github.com/ddaraujo/neoway_etl_test/models"
)

// Get all valid items
func (db Database) GetAllValidItems() (*models.ItemList, error) {
	list := &models.ItemList{}
	query := `SELECT * FROM customer_data ORDER BY cpf DESC`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Cpf, &item.Private, &item.Incompleto, &item.DataUltimaCompra, &item.ValorTicketMedio, &item.CnpjLojaMaisFrequente, &item.DataUltimaCompra, &item.CnpjLojaUltimaCompra)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

// Get all invalid(rejected) items
func (db Database) GetAllInvalidItems() (*models.ItemList, error) {
	list := &models.ItemList{}
	query := `SELECT * FROM customer_data_rejected ORDER BY cpf DESC;`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Cpf, &item.Private, &item.Incompleto, &item.DataUltimaCompra, &item.ValorTicketMedio, &item.CnpjLojaMaisFrequente, &item.DataUltimaCompra, &item.CnpjLojaUltimaCompra)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

// Get all valid items by CPF
func (db Database) GetAllItemsByCpf(itemCpf string) (*models.ItemList, error) {
	list := &models.ItemList{}
	query := `SELECT * FROM customer_data WHERE cpf = $1;`

	rows, err := db.Conn.Query(query, itemCpf)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Cpf, &item.Private, &item.Incompleto, &item.DataUltimaCompra, &item.ValorTicketMedio, &item.ValorTicketUltimaCompra, &item.CnpjLojaMaisFrequente, &item.CnpjLojaUltimaCompra)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

// Get all valid items by Last Sale CNPJ
func (db Database) GetAllItemsLastSale(lastSaleCnpj string) (*models.ItemList, error) {
	list := &models.ItemList{}
	query := `SELECT * FROM customer_data WHERE cnpj_loja_ultima_compra = $1;`

	rows, err := db.Conn.Query(query, lastSaleCnpj)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Cpf, &item.Private, &item.Incompleto, &item.DataUltimaCompra, &item.ValorTicketMedio, &item.ValorTicketUltimaCompra, &item.CnpjLojaMaisFrequente, &item.CnpjLojaUltimaCompra)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

// Get all valid items by frequent sale CNPJ
func (db Database) GetAllItemsFrequentSale(frequentSaleCnpj string) (*models.ItemList, error) {
	list := &models.ItemList{}
	query := `SELECT * FROM customer_data WHERE cpf = $1;`

	rows, err := db.Conn.Query(query, frequentSaleCnpj)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Cpf, &item.Private, &item.Incompleto, &item.DataUltimaCompra, &item.ValorTicketMedio, &item.ValorTicketUltimaCompra, &item.CnpjLojaMaisFrequente, &item.CnpjLojaUltimaCompra)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

// Get valid records count
func (db Database) GetValidRecordsCount() (models.Record, error) {
	record := models.Record{}

	query := `SELECT count(*) FROM customer_data;`
	row := db.Conn.QueryRow(query)
	switch err := row.Scan(&record.TotalRecords); err {
	case sql.ErrNoRows:
		return record, ErrNoMatch
	default:
		return record, err
	}
}

// Get invalid records count
func (db Database) GetInvalidRecordsCount() (models.Record, error) {
	record := models.Record{}

	query := `SELECT count(*) FROM customer_data_rejected;`
	row := db.Conn.QueryRow(query)
	switch err := row.Scan(&record.TotalRecords); err {
	case sql.ErrNoRows:
		return record, ErrNoMatch
	default:
		return record, err
	}
}

// Delete valid records table
func (db Database) DeleteValidItems(itemId int) error {
	query := `DELETE FROM customer_data;`
	_, err := db.Conn.Exec(query, itemId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

// Delete invalid records table
func (db Database) DeleteInvalidItems(itemId int) error {
	query := `DELETE FROM customer_data_rejected;`
	_, err := db.Conn.Exec(query, itemId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
