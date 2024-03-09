package model

type Fixture struct {
	Date     string `csv:"Date"`
	HomeTeam string `csv:"HomeTeam"`
	AwayTeam string `csv:"AwayTeam"`
	FTHG     int    `csv:"FTHG"`
	FTAG     int    `csv:"FTAG"`
	Referee  string `csv:"Referee"`
}
