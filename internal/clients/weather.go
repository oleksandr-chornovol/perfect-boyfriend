package clients

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func GetCurrentWeather() (string, error) {
	res, err := http.Get(os.Getenv("WEATHER_API_URL"))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		Current struct {
			WeatherDescriptions []string `json:"weather_descriptions"`
		}
		Error struct {
			Info string `json:"info"`
		}
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result.Current.WeatherDescriptions) > 0 {
		weather := result.Current.WeatherDescriptions[0]
		log.Println(weather)

		return weather, nil
	} else {
		return "", errors.New(result.Error.Info)
	}
}
