package constants

// REGEX_LINE matches one line to be processed.
// This line should contain the log for the game.
// Lines that does not represent a game log is not matched by this regex.
// The first group represents the Title while the second group represents the content to be parsed.
const REGEX_LINE = "\\d+:\\d+\\s(\\w+):\\s?(.*)?"

// REGEX_USER_INFO matches the user info contained in a given line.
// The first group represents the userID while the second group represents the user name
const REGEX_USER_INFO = "(\\d+)\\sn\\\\(\\w+).*$"

// REGEX_KILL matches the following:
// Group 1 = the killerID;
// Group 2 = The Killed ID;
// Group 3 = The Death Cause
const REGEX_KILL = "^(\\d{1,})\\s(\\d{1,}).*by\\s(.*)$"

// WORLD_ID represents the constant id for the <world> register on a Kill line
const WORLD_ID = 1022
