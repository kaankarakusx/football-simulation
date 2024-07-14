package simulation

import (
	"database/sql"
	"football-simulation/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveFixture(matches []types.Match) error {

	for _, match := range matches {
		_, err := s.db.Exec(`INSERT INTO matches (week, team1_id, team2_id, team1_score, team2_score, played) VALUES ($1, $2, $3, $4, $5, $6)`,
			match.Week, match.Team1ID, match.Team2ID, match.Team1Score, match.Team2Score, match.Played)
		if err != nil {
			return err
		}
	}

	return nil
}
