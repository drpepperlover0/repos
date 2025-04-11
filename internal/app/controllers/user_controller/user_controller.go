package user_controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/drpepperlover0/internal/app/service"
	"github.com/drpepperlover0/internal/app/types"
	"github.com/drpepperlover0/internal/models"
	"gorm.io/gorm"
)

type UserControl struct {
	service types.Service
}

func New(service types.Service) *UserControl {
	return &UserControl{
		service: service,
	}
}

func (h *UserControl) Create(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	json.NewDecoder(r.Body).Decode(&user)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := h.service.Create(ctx, user); err != nil {
		fmt.Fprintf(w, "Create user error: %v", err)
		return
	}

	fmt.Fprintf(w, "Username in DB: %s\nID: %d", user.Username, user.ID)
}

func (h *UserControl) GetAll(w http.ResponseWriter, r *http.Request) {
	var users []*models.User

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	found, err := h.service.GetAll(ctx)
	if err != nil {
		errorHandle(w, err)
		return
	}
	users = append(users, found...)

	json.NewEncoder(w).Encode(users)
}

func (h *UserControl) Get(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	user, err := h.service.Get(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorHandle(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserControl) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := h.service.Delete(ctx, id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorHandle(w, err)
		return
	}

	fmt.Fprintf(w, "Deleted ID: %d\n", id)
}

func getID(path string) (int, error) {
	q := strings.Split(path, "/")[2]
	id, err := strconv.Atoi(q)
	if err != nil {
		return 0, errors.New("ID must to be an integer")
	}

	return id, nil
}

func errorHandle(wr io.Writer, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprint(wr, "not a single record was found")
	} else {
		fmt.Fprintf(wr, "users finding error: %v", err)
	}
}
