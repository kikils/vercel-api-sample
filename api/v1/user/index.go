package user

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/kikils/vercel-api-sample/handler"
	_ "github.com/lib/pq"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	conn := os.Getenv("POSTGRES_CONN")
	db, err := sqlx.Open("postgres", conn)
	defer db.Close()
	if err != nil {
		http.Error(w, fmt.Sprintf("user.Handler: %v", err), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		http.Error(w, fmt.Sprintf("user.Handler: %v", err), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	uHandler := handler.NewUserHandler(db)

	switch r.Method {
	case http.MethodPost:
		h := handler.MiddlewareHandler(uHandler.CreateUser)
		h(w, r)
	case http.MethodGet:
		h := handler.MiddlewareHandler(uHandler.SearchUser)
		h(w, r)
	case http.MethodPatch:
		h := handler.MiddlewareHandler(uHandler.UpdateUser)
		h(w, r)
	case http.MethodDelete:
		h := handler.MiddlewareHandler(uHandler.DeleteUser)
		h(w, r)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
