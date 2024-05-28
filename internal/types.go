package internal

// Db Represents a database containing the matches info
type Db struct {
	Matches map[string]*Match
}

// Match represents a game and its info
type Match struct {
	Id           string
	TotalKills   int
	Players      map[int]string
	Kills        map[int]int
	DeathsCauses map[string]int
}
