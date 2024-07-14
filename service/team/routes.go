package team

import (
	"football-simulation/types"
	"football-simulation/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service types.TeamService
}

func NewHandler(service types.TeamService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/teams", h.handleGetTeams).Methods("GET")

}

func (h *Handler) handleGetTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.service.GetTeams()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, teams)
}
