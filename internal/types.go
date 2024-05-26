package internal

type Db struct {
	Matches map[string]Match
}

type Match struct {
	Id         string
	TotalKills uint
	Players    map[int]string
	Kills      map[string]int
}
