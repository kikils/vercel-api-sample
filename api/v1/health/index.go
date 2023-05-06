package health

import (
	"net/http"

	"github.com/kikils/vercel-api-sample/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fn := handler.MiddlewareHandler(handler.Health)
	fn(w, r)
}
