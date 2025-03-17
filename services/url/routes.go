package url

import (
	"UrlShortenerGoLang/storage"
	"UrlShortenerGoLang/types"
	"UrlShortenerGoLang/utils"
	"net/http"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store storage.Storage
}

func NewHandler(store storage.Storage) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/urls", h.handleCreateUrl).Methods("POST")
	router.HandleFunc("/urls/{code}", h.handleGetUrl).Methods("GET")
	router.HandleFunc("/urls", h.handleGetAllUrls).Methods("GET")
}

func (h *Handler) handleCreateUrl(w http.ResponseWriter, r *http.Request) {
	var urlPayload types.CreateUrlPayload

	if err := utils.ParseJSON(r, &urlPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(urlPayload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if urlExists, _ := h.store.GetUrlByOriginalUrl(r.Context(), urlPayload.OriginalUrl); urlExists != nil {
		utils.WriteJSON(w, http.StatusOK, urlExists)
		return
	}

	url, err := h.store.CreateUrl(r.Context(), &types.Url{
		OriginalUrl: urlPayload.OriginalUrl,
		Code:        utils.GenerateUniqueCode(6),
		VisitCount:  0,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, url)
}

func (h *Handler) handleGetUrl(w http.ResponseWriter, r *http.Request) {
	code, ok := mux.Vars(r)["code"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing code"))
		return
	}

	url, err := h.store.GetUrlByCode(r.Context(), code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if url == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !utils.CheckURLExists(url.OriginalUrl) {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("URL does not exist"))
	}

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound) // Using 302 instead of 301 because 301 is cached by the browser so it won't update the visit count
}

func (h *Handler) handleGetAllUrls(w http.ResponseWriter, r *http.Request) {
	urls, err := h.store.GetAllUrls(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	utils.WriteJSON(w, http.StatusOK, urls)
}