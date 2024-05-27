package internal

import (
	"encoding/json"
	"testing"
)

func TestProcessInitGame(t *testing.T) {

	initializeDb()
	processInitGame()

	const expected = 1
	const expectedId = "game_1"

	if len(db.Matches) != expected {
		t.Fatalf(`DB size should be = %d, but is %d`, expected, len(db.Matches))
	}

	if db.Matches[expectedId].Id != expectedId {
		t.Fatalf(`DB ID should be = %s, but is %s`, expectedId, db.Matches[expectedId].Id)
	}
}

func TestProcessShutdownGame(t *testing.T) {
	initializeDb()
	processInitGame()
	if currentGame != 1 {
		t.Fatalf(`DB size should be = 1, but is %d`, currentGame)
	}

	processInitGame()
	processInitGame()

	if currentGame != 3 {
		t.Fatalf(`DB size should be = 3, but is %d`, currentGame)
	}

	processShutdownGame()

	if currentGame != 2 {
		t.Fatalf(`DB size should be = 2, but is %d`, currentGame)
	}
}

func TestProcessClientUserinfoChanged(t *testing.T) {
	initializeVars()
	processInitGame()
	var expectedId = 2
	var expectedName = "Isgalamido"
	var expectedResult = map[int]string{expectedId: expectedName}

	processClientUserinfoChanged(" 20:34 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\xian/default\\hmodel\\xian/default\\g_redteam\\\\g_blueteam\\\\c1\\4\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0")

	expectedJson, _ := json.Marshal(expectedResult)
	returnedJson, _ := json.Marshal(db.Matches["game_1"].Players)
	if string(expectedJson) != string(returnedJson) {
		t.Fatalf(`expected result should be %s, but it was %s`, expectedJson, returnedJson)
	}

}

func TestProcessKill(t *testing.T) {
	initializeVars()
	processInitGame()

	lines := [...]string{"1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
		"1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
		"1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
		"1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
		"2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH",
		"2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH",
		"2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH",
		"3 2 6: Isgalamido killed Mocinha by MOD_ROCKET"}

	var expectedResult = Match{
		Id:           "game_1",
		TotalKills:   8,
		Kills:        map[int]int{2: -3, 3: 1},
		DeathsCauses: map[string]int{"MOD_ROCKET_SPLASH": 3, "MOD_ROCKET": 1},
	}

	for _, line := range lines {
		err := processKill(line)
		if err != nil {
			panic(err)
		}
	}

	if expectedResult.TotalKills != db.Matches["game_1"].TotalKills {
		t.Fatalf(`expected TotalKills result should be %d, but it was %d`, expectedResult.TotalKills, db.Matches["game_1"].TotalKills)
	}

	expectedKillJson, _ := json.Marshal(expectedResult.Kills)
	returnedKillJson, _ := json.Marshal(db.Matches["game_1"].Kills)

	if string(expectedKillJson) != string(returnedKillJson) {
		t.Fatalf(`expected kills result should be %s, but it was %s`, expectedKillJson, returnedKillJson)
	}

	expectedDeathsCauseJson, _ := json.Marshal(expectedResult.DeathsCauses)
	returnedDeathsCauseJson, _ := json.Marshal(db.Matches["game_1"].DeathsCauses)

	if string(expectedDeathsCauseJson) != string(returnedDeathsCauseJson) {
		t.Fatalf(`expected kills result should be %s, but it was %s`, expectedDeathsCauseJson, returnedDeathsCauseJson)
	}

}
