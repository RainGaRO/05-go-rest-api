package user

import (
	"net/http"
	"restapi/internal/handlers"
	"restapi/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{} //for Register

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

// logger and service (логгирование - что сейчас приложение делает)
type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler { //return interface
	return &handler{
		logger: logger,
	} //create struct
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserById)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write(([]byte("this is list of users")))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write(([]byte("this is list of users")))
}
func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write(([]byte("this is list of users by id")))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write(([]byte("this is list of users")))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write(([]byte("this is list of users")))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write(([]byte("this is list of users")))
}
