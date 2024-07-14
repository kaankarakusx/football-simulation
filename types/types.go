package types

type League struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	CurrentWeek      int    `json:"current_week"`
	TotalWeeks       int    `json:"total_weeks"`
	ChampionTeamName string `json:"champion_team_name,omitempty"`
}

type Team struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Strength        int    `json:"strength"`
	Points          int    `json:"points"`
	Matches         int    `json:"matches"`
	Wins            int    `json:"wins"`
	Draws           int    `json:"draws"`
	Losses          int    `json:"losses"`
	GoalsFor        int    `json:"goals_for"`
	GoalsAgainst    int    `json:"goals_against"`
	GoalsDifference int    `json:"goals_difference"`
	TemporaryDrop   int    `json:"temporary_drop,omitempty"`
}

type Match struct {
	ID         int  `json:"id"`
	Week       int  `json:"week"`
	Team1ID    int  `json:"team1_id"`
	Team2ID    int  `json:"team2_id"`
	Team1Score int  `json:"team1_score"`
	Team2Score int  `json:"team2_score"`
	Played     bool `json:"played"`
}

type MatchResult struct {
	ID         int    `json:"id"`
	Week       int    `json:"week"`
	Team1Name  string `json:"team1_name"`
	Team2Name  string `json:"team2_name"`
	Team1Score int    `json:"team1_score"`
	Team2Score int    `json:"team2_score"`
	Played     bool   `json:"played"`
}

type Prediction struct {
	TeamID           int     `json:"team_id"`
	TeamName         string  `json:"team_name"`
	ChampionshipOdds float64 `json:"championship_odds"`
}

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

type UpdateMatchRequest struct {
	Team1Score int `json:"team1_score"`
	Team2Score int `json:"team2_score"`
}
