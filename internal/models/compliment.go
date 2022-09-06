package models

type Compliment struct {
	ID            int
	Text          string
	ProperWeather string `json:"proper_weather"`
	IsGreeting    bool   `json:"is_greeting"`
}
