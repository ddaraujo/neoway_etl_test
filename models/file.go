package models

import (
    "fmt"
    "net/http"
)
type File struct {
	FileName string `json:"filename"`
}

func (i *File) Bind(r *http.Request) error {
    if i.FileName == "" {
        return fmt.Errorf("Filename is a mandatory field")
    }
    return nil
}
func (*File) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}