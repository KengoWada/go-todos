package auth

import (
	"net/http"

	"github.com/KengoWada/go-todos/internal/models"
	"github.com/KengoWada/go-todos/internal/store"
	"github.com/KengoWada/go-todos/internal/utils"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store store.Storage
}

func NewHandler(store store.Storage) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Post("/register", h.registerUser)

	return mux
}

type registerUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var payload registerUserPayload
	utils.ReadJSON(w, r, &payload)

	user := &models.User{Email: payload.Email, Password: payload.Password}
	userProfile := &models.UserProfile{Name: payload.Name}

	if err := h.store.Users.Create(r.Context(), user, userProfile); err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userProfile.User = user
	utils.WriteJSONResponse(w, http.StatusCreated, userProfile)
}
