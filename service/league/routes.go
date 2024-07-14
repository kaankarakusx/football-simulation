package league

import (
	"fmt"
	"football-simulation/types"
	"football-simulation/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service types.LeagueService
}

func NewHandler(service types.LeagueService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/league/nextweek", h.handleNextWeek).Methods("POST")
	router.HandleFunc("/league/playall", h.handlePlayAll).Methods("POST")
	router.HandleFunc("/league/restart", h.handleRestartLeague).Methods("POST")
	router.HandleFunc("/league/standings", h.handleGetStandings).Methods("GET")
	router.HandleFunc("/league/weekresults", h.handleGetWeekResults).Methods("GET")
	router.HandleFunc("/league/matches", h.handleGetAllMatches).Methods("GET")
	router.HandleFunc("/league/matches/{week}", h.handleGetMatchesByWeek).Methods("GET")
	router.HandleFunc("/league/match/{id}", h.handleUpdateMatch).Methods("PUT")
	router.HandleFunc("/league/predictions", h.handleGetPredictions).Methods("GET")
}

func (h *Handler) handleNextWeek(w http.ResponseWriter, r *http.Request) {
	playedMatches, champion, err := h.service.NextWeek()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := map[string]interface{}{
		"status":        "success",
		"message":       "Next week played successfully",
		"playedMatches": playedMatches,
		"champion":      nil,
	}

	if champion != nil {
		response["champion"] = champion
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handlePlayAll(w http.ResponseWriter, r *http.Request) {
	playedMatches, champion, err := h.service.PlayAll()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := map[string]interface{}{
		"status":        "success",
		"message":       "Next week played successfully",
		"playedMatches": playedMatches,
		"champion":      champion,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handleGetMatchesByWeek(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	week, err := strconv.Atoi(vars["week"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	matches, err := h.service.GetMatchesByWeek(week)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, matches)
}
func (h *Handler) handleGetWeekResults(w http.ResponseWriter, r *http.Request) {
	weekResults, err := h.service.GetWeekResults()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, weekResults)
}

func (h *Handler) handleGetAllMatches(w http.ResponseWriter, r *http.Request) {
	matches, err := h.service.GetAllMatches()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, matches)
}

func (h *Handler) handleUpdateMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, nil)
		return
	}

	var req types.UpdateMatchRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println(req)
	match := types.Match{
		ID:         id,
		Team1Score: req.Team1Score,
		Team2Score: req.Team2Score,
	}

	if err := h.service.UpdateMatch(match); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, nil)
}

func (h *Handler) handleRestartLeague(w http.ResponseWriter, r *http.Request) {
	err := h.service.RestartLeague()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "")

}

func (h *Handler) handleGetStandings(w http.ResponseWriter, r *http.Request) {
	teams, err := h.service.GetStandings()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, teams)

}

func (h *Handler) handleGetPredictions(w http.ResponseWriter, r *http.Request) {
	predictions, err := h.service.GetPredictions()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, predictions)
}
