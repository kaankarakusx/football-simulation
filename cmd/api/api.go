package api

import (
	"database/sql"
	"football-simulation/service/league"
	"football-simulation/service/simulation"
	"football-simulation/service/team"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	//Store
	teamStore := team.NewStore(s.db)
	leagueStore := league.NewStore(s.db)
	simulationStore := simulation.NewStore(s.db)

	//Service
	teamService := team.NewService(teamStore)
	simulationService := simulation.NewService(simulationStore)
	leagueService := league.NewService(leagueStore, simulationService, teamService)

	//Handler
	teamHandler := team.NewHandler(teamService)
	leagueHandler := league.NewHandler(leagueService)

	leagueHandler.RegisterRoutes(subRouter)
	teamHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)

	if err := http.ListenAndServe(s.addr, router); err != nil {
		log.Fatalf("could not start server: %v\n", err)
		return err
	}
	return nil
}
