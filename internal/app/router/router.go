package router

import (
	"net/http"

	"github.com/drpepperlover0/internal/app/controllers/user_controller"
	_ "github.com/drpepperlover0/internal/app/service"
	"github.com/drpepperlover0/internal/app/types"
)

func InitRoutes(s types.Service) *http.ServeMux {
	r := http.NewServeMux()
	h := user_controller.New(s)

	r.HandleFunc("POST /user", h.Create)
	r.HandleFunc("GET /user", h.GetAll)
	r.HandleFunc("GET /user/{id}", h.Get)
	r.HandleFunc("DELETE /user/{id}", h.Delete)

	return r
}
