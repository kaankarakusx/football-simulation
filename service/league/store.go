package league

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

func (s *Store) GetLeagueInfo() (types.League, error) {
	league := new(types.League)

	rows, err := s.db.Query("SELECT * FROM league")

	if err != nil {
		return types.League{}, err
	}

	for rows.Next() {
		league, err = scanRowsIntoLeague(rows)
		if err != nil {
			return types.League{}, err
		}
	}

	return *league, nil
}

func (s *Store) ClearFixtures() error {
	_, err := s.db.Exec("DELETE FROM matches")
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateLeague(league types.League) error {
	_, err := s.db.Exec(`UPDATE league SET name = $1, current_week = $2, total_weeks = $3, champion_team_name = $4 WHERE id = $5`,
		league.Name, league.CurrentWeek, league.TotalWeeks, league.ChampionTeamName, league.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetStandings() ([]types.Team, error) {

	rows, err := s.db.Query(`
		SELECT id, name, points, matches, wins, draws, losses, goals_for, goals_against, goals_difference, temporary_drop
		FROM teams
		ORDER BY points DESC, goals_difference DESC, goals_for DESC`)
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

func (s *Store) GetCurrentWeek() (int, error) {
	var currentWeek int
	err := s.db.QueryRow("SELECT current_week FROM league").Scan(&currentWeek)
	if err != nil {
		return 0, err
	}
	return currentWeek, nil
}

func (s *Store) GetMatchesForNextWeek() ([]types.Match, error) {
	currentWeek, err := s.GetCurrentWeek()
	if err != nil {
		return nil, err
	}

	var matches []types.Match
	rows, err := s.db.Query("SELECT id, week, team1_id, team2_id, team1_score, team2_score, played FROM matches WHERE played = FALSE AND week = $1", currentWeek)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		match, err := scanRowsIntoMatch(rows)
		if err != nil {
			return nil, err
		}
		matches = append(matches, *match)
	}

	return matches, nil
}

func (s *Store) GetMatchesByWeek(week int) ([]types.Match, error) {
	var matches []types.Match
	rows, err := s.db.Query("SELECT id, week, team1_id, team2_id, team1_score, team2_score, played FROM matches WHERE week = $1 AND played = TRUE", week)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		match, err := scanRowsIntoMatch(rows)
		if err != nil {
			return nil, err
		}
		matches = append(matches, *match)
	}

	return matches, nil
}

func (s *Store) GetMatchByID(id int) (*types.Match, error) {

	rows, err := s.db.Query("SELECT * FROM matches WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	match := new(types.Match)

	for rows.Next() {
		match, err = scanRowsIntoMatch(rows)

		if err != nil {
			return nil, err
		}

	}

	return match, nil
}

func (s *Store) UpdateMatch(match types.Match) error {
	_, err := s.db.Exec("UPDATE matches SET team1_score = $1, team2_score = $2, played = TRUE WHERE id = $3",
		match.Team1Score, match.Team2Score, match.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) SaveMatchResult(match types.Match) error {
	_, err := s.db.Exec("UPDATE matches SET team1_score = $1, team2_score = $2, played = TRUE WHERE id = $3",
		match.Team1Score, match.Team2Score, match.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) IncrementWeek() error {
	_, err := s.db.Exec("UPDATE league SET current_week = current_week + 1")
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAllMatches() ([]types.Match, error) {
	var matches []types.Match
	rows, err := s.db.Query("SELECT id, week, team1_id, team2_id, team1_score, team2_score, played FROM matches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		match, err := scanRowsIntoMatch(rows)
		if err != nil {
			return nil, err
		}
		matches = append(matches, *match)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func (s *Store) GetPredictions() ([]types.Prediction, error) {
	return nil, nil
}

func scanRowsIntoLeague(rows *sql.Rows) (*types.League, error) {
	league := new(types.League)
	var championTeamName sql.NullString
	err := rows.Scan(
		&league.ID,
		&league.Name,
		&league.CurrentWeek,
		&league.TotalWeeks,
		&championTeamName,
	)

	if err != nil {
		return nil, err
	}

	if championTeamName.Valid {
		league.ChampionTeamName = championTeamName.String
	} else {
		league.ChampionTeamName = ""
	}

	return league, nil
}

func scanRowsIntoTeam(rows *sql.Rows) (*types.Team, error) {
	team := new(types.Team)

	err := rows.Scan(
		&team.ID,
		&team.Name,
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

func scanRowsIntoMatch(rows *sql.Rows) (*types.Match, error) {
	match := new(types.Match)
	err := rows.Scan(
		&match.ID,
		&match.Week,
		&match.Team1ID,
		&match.Team2ID,
		&match.Team1Score,
		&match.Team2Score,
		&match.Played,
	)
	if err != nil {
		return nil, err
	}
	return match, nil
}
