package models

import (
    "fmt"
    "net/http"
)
type Item struct {
	Cpf string `json:"cpf"`
	Private string `json:"private"`
	Incompleto string `json:"incompleto"`
	DataUltimaCompra string `json:"data_ultima_compra"`
	ValorTicketMedio string `json:"valor_ticket_medio"`
	ValorTicketUltimaCompra string `json:"valor_ticket_ultima_compra"`
	CnpjLojaMaisFrequente string `json:"cpnj_loja_mais_frequente"`
	CnpjLojaUltimaCompra string `json:"cnpj_loja_utima_compra"`
}
type ItemList struct {
    Items []Item `json:"items"`
}
func (i *Item) Bind(r *http.Request) error {
    if i.Cpf == "" {
        return fmt.Errorf("CPF é um campo obrigatório")
    }
    return nil
}
func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}