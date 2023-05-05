package health

import (
	"net/http"

	"github.com/kikils/vercel-api-sample/internal/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler.Health(w, r)
}
