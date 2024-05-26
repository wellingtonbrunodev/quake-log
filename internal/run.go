package internal

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	patterns "github.com/wellingtonbrunodev/quake-log/internal/constants"
	utils "github.com/wellingtonbrunodev/quake-log/pkg/utils"
)

var db Db

var currentGame int

var PATTERN_LINE, PATTERN_USER_INFO *regexp.Regexp

func initializeDb() {
	currentGame = 0
	db = Db{
		Matches: map[string]Match{},
	}
}

func initializeVars() {

	var err error

	PATTERN_LINE, err = regexp.Compile(patterns.REGEX_LINE)
	if err != nil {
		panic("error when compiling regex for line")
	}

	PATTERN_USER_INFO, err = regexp.Compile(patterns.REGEX_USER_INFO)
	if err != nil {
		panic("error when compiling regex for user info")
	}

	initializeDb()
}

func Run() {

	initializeVars()

	content := utils.ReadFile("./pkg/input_files/qgames.log")

	lines := strings.Split(content, "/n")
	for _, line := range lines {
		err := processLine(line)
		if err != nil {
			panic(err)
		}
	}
}

func processLine(line string) error {
	if line == "" {
		return nil
	}

	parsedGroups := PATTERN_LINE.FindStringSubmatch(line)
	switch parsedGroups[1] {
	case "InitGame":
		processInitGame()
	case "ShutdownGame":
		processShutdownGame()
	case "ClientUserinfoChanged":
		processClientUserinfoChanged(parsedGroups[2])
	default:
		return nil
	}
	return nil
}

func processInitGame() {
	currentGame = len(db.Matches) + 1
	id := fmt.Sprint("game_", currentGame)
	db.Matches[id] = Match{Id: id, Players: make(map[int]string), Kills: make(map[string]int)}
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
