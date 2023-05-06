package handler

import (
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "ok")
	return nil
}
