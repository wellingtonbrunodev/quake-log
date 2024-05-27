package internal

type Db struct {
	Matches map[string]*Match
}

type Match struct {
	Id           string
	TotalKills   int
	Players      map[int]string
	Kills        map[int]int
	DeathsCauses map[string]int
}
