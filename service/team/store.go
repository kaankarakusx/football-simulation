package team

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

func (s *Store) GetTeams() ([]types.Team, error) {

	rows, err := s.db.Query("SELECT * FROM teams")

	if err != nil {
		return nil, err
	}

	teams := make([]types.Team, 0)

	for rows.Next() {
		team, err := scanRowsIntoTeam(rows)

		if err != nil {
			return nil, err
		}

		teams = append(teams, *team)
	}

	return teams, nil
}

func (s *Store) GetTeamByID(id int) (*types.Team, error) {
	rows, err := s.db.Query("SELECT * FROM teams WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	team := new(types.Team)

	for rows.Next() {
		team, err = scanRowsIntoTeam(rows)

		if err != nil {
			return nil, err
		}
	}

	return team, nil
}
func (s *Store) GetTeamByName(name string) (*types.Team, error) {
	rows, err := s.db.Query("SELECT * FROM teams WHERE name = $1", name)

	if err != nil {
		return nil, err
	}

	team := new(types.Team)

	for rows.Next() {
		team, err = scanRowsIntoTeam(rows)

		if err != nil {
			return nil, err
		}
	}

	return team, nil
}

func (s *Store) UpdateTeam(team types.Team) error {
	_, err := s.db.Exec(`UPDATE teams
	SET name = $1, strength = $2, points = $3, matches = $4, wins = $5, draws = $6, losses = $7, goals_for = $8, goals_against = $9, goals_difference = $10, temporary_drop = $11
	WHERE id = $12`, team.Name, team.Strength, team.Points, team.Matches, team.Wins, team.Draws, team.Losses, team.GoalsFor, team.GoalsAgainst, team.GoalsDifference, team.TemporaryDrop, team.ID)

	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ResetTeams() error {
	_, err := s.db.Exec("UPDATE teams SET points = 0, matches = 0, wins = 0, draws = 0, losses = 0, goals_for = 0, goals_against = 0, goals_difference = 0, temporary_drop = 0")
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoTeam(rows *sql.Rows) (*types.Team, error) {
	team := new(types.Team)

	err := rows.Scan(
		&team.ID,
		&team.Name,
		&team.Strength,
		&team.Points,
		&team.Matches,
		&team.Wins,
		&team.Draws,
		&team.Losses,
		&team.GoalsFor,
		&team.GoalsAgainst,
		&team.GoalsDifference,
		&team.TemporaryDrop,
	)

	if err != nil {
		return nil, err
	}

	return team, nil
}
