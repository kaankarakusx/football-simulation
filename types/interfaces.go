package types

type LeagueStore interface {
	GetLeagueInfo() (League, error)
	ClearFixtures() error
	GetStandings() ([]Team, error)
	GetMatchesByWeek(week int) ([]Match, error)
	GetMatchByID(id int) (*Match, error)
	UpdateMatch(match Match) error
	GetCurrentWeek() (int, error)
	GetMatchesForNextWeek() ([]Match, error)
	SaveMatchResult(match Match) error
	IncrementWeek() error
	UpdateLeague(league League) error
	GetAllMatches() ([]Match, error)
	GetPredictions() ([]Prediction, error)
}

type LeagueService interface {
	StartLeague() error
	NextWeek() ([]Match, *Team, error)
	PlayAll() ([]MatchResult, *Team, error)
	GetWeekResults() ([]MatchResult, error)
	UpdateMatch(match Match) error
	GetMatchesByWeek(id int) ([]MatchResult, error)
	GetAllMatches() ([]MatchResult, error)
	RestartLeague() error
	GetStandings() ([]Team, error)
	GetPredictions() ([]Prediction, error)
}

type MatchStore interface {
	GetAllMatches() ([]Match, error)
	UpdateMatch(match Match) error
	GetMatchByID(id int) (*Match, error)
}

type MatchService interface {
	GetAllMatches() ([]MatchResult, error)
	UpdateMatch(match Match) error
}

type Teamstore interface {
	GetTeams() ([]Team, error)
	GetTeamByID(id int) (*Team, error)
	GetTeamByName(name string) (*Team, error)
	UpdateTeam(Team) error
	ResetTeams() error
}

type TeamService interface {
	GetTeams() ([]Team, error)
	GetTeamByID(id int) (*Team, error)
	GetTeamByName(name string) (*Team, error)
	UpdateTeamStatsReverse(match Match) error
	UpdateTeamStats(team1, team2 Team, team1Score, team2Score int, isUpdate bool) error
	UpdateTeam(Team) error
	ResetTeams() error
}

type SimulationStore interface {
	SaveFixture(matches []Match) error
}

type SimulationService interface {
	GenerateFixture([]Team) error
	PlayMatch(team1, team2 Team) (int, int)
	CalculateChampionshipOdds(teams []Team, matches []Match) ([]Prediction, error)
}
