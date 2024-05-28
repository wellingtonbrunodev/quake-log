package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	constants "github.com/wellingtonbrunodev/quake-log/internal/constants"
	utils "github.com/wellingtonbrunodev/quake-log/pkg/utils"
)

var db Db

var currentGame int

var PATTERN_LINE, PATTERN_USER_INFO, PATTERN_KILL *regexp.Regexp

func initializeDb() {
	currentGame = 0
	db = Db{
		Matches: map[string]*Match{},
	}
}

func initializeVars() {

	var err error

	PATTERN_LINE, err = regexp.Compile(constants.REGEX_LINE)
	if err != nil {
		panic("error when compiling regex for line")
	}

	PATTERN_USER_INFO, err = regexp.Compile(constants.REGEX_USER_INFO)
	if err != nil {
		panic("error when compiling regex for user info")
	}

	PATTERN_KILL, err = regexp.Compile(constants.REGEX_KILL)
	if err != nil {
		panic("error when compiling regex for kill")
	}

	initializeDb()
}

func Run() {

	initializeVars()

	lines := utils.ReadFile("./pkg/input_files/qgames.log")

	for _, line := range lines {
		err := processLine(line)
		if err != nil {
			panic(err)
		}
	}

	report, err := generateReport()
	if err != nil {
		panic(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, []byte(report), "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	reportString := out.String()
	fmt.Println(reportString)

	utils.WriteFile(reportString, "./pkg/output_files/output.json")
}

func generateReport() (string, error) {
	json, err := json.Marshal(db.Matches)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func processLine(line string) error {
	if line == "" {
		return nil
	}

	parsedGroups := PATTERN_LINE.FindStringSubmatch(line)
	if len(parsedGroups) > 1 {
		switch parsedGroups[1] {
		case "InitGame":
			processInitGame()
		case "ShutdownGame":
			processShutdownGame()
		case "ClientUserinfoChanged":
			err := processClientUserinfoChanged(parsedGroups[2])
			if err != nil {
				return err
			}
		case "Kill":
			err := processKill(parsedGroups[2])
			if err != nil {
				return err
			}
		default:
			return nil
		}
	}

	return nil
}

func processInitGame() {
	currentGame = len(db.Matches) + 1
	id := fmt.Sprint("game_", currentGame)
	db.Matches[id] = &Match{Id: id,
		Players:      make(map[int]string),
		Kills:        make(map[int]int),
		DeathsCauses: make(map[string]int),
		TotalKills:   0,
	}
}

func processShutdownGame() {
	currentGame--
}

func processClientUserinfoChanged(info string) error {
	parsedGroups := PATTERN_USER_INFO.FindStringSubmatch(info)

	if len(parsedGroups) < 2 {
		return errors.New("Failed while parsing user info: " + info)
	}

	gameId := fmt.Sprint("game_", currentGame)
	userId, err := strconv.Atoi(parsedGroups[1])
	if err != nil {
		return err
	}
	userName := parsedGroups[2]

	db.Matches[gameId].Players[userId] = userName

	return nil
}

func processKill(info string) error {
	parsedGroups := PATTERN_KILL.FindStringSubmatch(info)

	if len(parsedGroups) < 4 {
		return errors.New("Failed while parsing kill info: " + info)
	}

	gameId := fmt.Sprint("game_", currentGame)

	var killerId, killedId int
	var err error

	killerId, err = strconv.Atoi(parsedGroups[1])
	if err != nil {
		return err
	}

	killedId, err = strconv.Atoi(parsedGroups[2])
	if err != nil {
		return err
	}

	deathMode := parsedGroups[3]

	if killerId != constants.WORLD_ID {

		if killerId != killedId {
			userKills := db.Matches[gameId].Kills[killerId]
			userKills++
			db.Matches[gameId].Kills[killerId] = userKills
		}

		deathsCauses := db.Matches[gameId].DeathsCauses[deathMode]
		deathsCauses++
		db.Matches[gameId].DeathsCauses[deathMode] = deathsCauses

	} else {
		userKills := db.Matches[gameId].Kills[killedId]
		userKills--
		db.Matches[gameId].Kills[killedId] = userKills
	}

	totalKills := db.Matches[gameId].TotalKills
	totalKills++
	db.Matches[gameId].TotalKills = totalKills

	return nil
}
