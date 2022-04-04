package db

import (
    "database/sql"
    "gitlab.com/idoko/bucketeer/models"
)

// Get all valid items
func (db Database) GetAllValidItems() (*models.ItemList, error) {
    list := &models.ItemList{}
    rows, err := db.Conn.Query("SELECT * FROM customer_data ORDER BY cpf DESC")
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

// Get all rejected items
func (db Database) GetAllRejectedItems() (*models.ItemList, error) {
    list := &models.ItemList{}
    rows, err := db.Conn.Query("SELECT * FROM customer_data_rejected ORDER BY cpf DESC")
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

// Get all items by CPF
func (db Database) GetAllItemsByCpf(itemCpf int) (models.Item, error) {
    item := models.Item{}
    query := `SELECT * FROM customer_data WHERE cpf = $1;`
    row := db.Conn.QueryRow(query, itemCpf)
    switch rows.Scan(&item.Cpf, &item.Private, &item.Incompleto, &item.DataUltimaCompra, &item.ValorTicketMedio, &item.CnpjLojaMaisFrequente, &item.DataUltimaCompra, &item.CnpjLojaUltimaCompra); err {
    case sql.ErrNoRows:
        return item, ErrNoMatch
    default:
        return item, err
    }
}

// Get all most frequent items by CNPJ
func (db Database) GetAllItemsByCnpjFrequent(itemCnpj int) (models.Item, error) {
    item := models.Item{}
    query := `SELECT * FROM customer_data WHERE cnpj_loja_mais_frequente = $1;`
    row := db.Conn.QueryRow(query, itemCnpj)
    switch rows.Scan(&item.Cpf, &item.Private, &item.Incompleto, &item.DataUltimaCompra, &item.ValorTicketMedio, &item.CnpjLojaMaisFrequente, &item.DataUltimaCompra, &item.CnpjLojaUltimaCompra); err {
    case sql.ErrNoRows:
        return item, ErrNoMatch
    default:
        return item, err
    }
}



// Get all last store items by CNPJ
func (db Database) GetAllItemsByCnpjLast(itemCnpj int) (models.Item, error) {
    item := models.Item{}
    query := `SELECT * FROM customer_data WHERE cnpj_loja_ultima_compra = $1;`
    row := db.Conn.QueryRow(query, itemCnpj)
    switch rows.Scan(&item.Cpf, &item.Private, &item.Incompleto, &item.DataUltimaCompra, &item.ValorTicketMedio, &item.CnpjLojaMaisFrequente, &item.DataUltimaCompra, &item.CnpjLojaUltimaCompra); err {
    case sql.ErrNoRows:
        return item, ErrNoMatch
    default:
        return item, err
    }
}
