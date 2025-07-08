package models

type Register struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type Tap struct {
	LoverId string `json:"loverID"`
}
