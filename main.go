package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
}

func getCity(reader *bufio.Reader) string {
	// Bold yellow prompt
	fmt.Print("\033[1;33mEnter city name: \033[0m")
	city, _ := reader.ReadString('\n')
	return strings.TrimSpace(city)
}

func getWeather(city, API_KEY string) (*WeatherResponse, error) {
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", encodedCity, API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}
	return &weather, nil
}

func kelvintoCelcius(temp float64) float64 {
	reslut := temp - 273.15
	roundedResult := math.Round(reslut)
	return roundedResult
}

func getWeatherEmoji(condition string) string {
	switch strings.ToLower(condition) {
	case "clear sky":
		return "☀️"
	case "few clouds", "scattered clouds", "broken clouds":
		return "⛅"
	case "shower rain", "rain":
		return "🌧️"
	case "thunderstorm":
		return "⛈️"
	case "snow":
		return "❄️"
	case "mist":
		return "🌫️"
	default:
		return "🌈"
	}
}

func displayWeather(weather *WeatherResponse) {
	tempC := kelvintoCelcius(weather.Main.Temp)
	emoji := getWeatherEmoji(weather.Weather[0].Description)
	fmt.Println("\033[1;34m---------------------------------\033[0m")
	fmt.Printf("\033[1;32mCity: %s\033[0m\n", weather.Name)
	fmt.Printf("\033[1;32mTemperature: %.2f°C\033[0m\n", tempC)
	fmt.Printf("\033[1;32mCondition: %s %s\033[0m\n", weather.Weather[0].Description, emoji)
	fmt.Println("\033[1;34m---------------------------------\033[0m")
}

func main() {
	API_KEY := os.Getenv("WEATHER_API_KEY")
	if API_KEY == "" {
		fmt.Println("Error: WEATHER_API_KEY environment variable not set")
		os.Exit(1)
	}
	fmt.Println("\033[1;36m") // Start cyan color
	fmt.Println(`
 ██████   █████  ██████  ██ ███████     ██     ██ ███████  █████  ████████ ██   ██ ███████ ██████         ██████ ██      ██ 
██    ██ ██   ██ ██   ██ ██ ██          ██     ██ ██      ██   ██    ██    ██   ██ ██      ██   ██       ██      ██      ██ 
██    ██ ███████ ██   ██ ██ ███████     ██  █  ██ █████   ███████    ██    ███████ █████   ██████  █████ ██      ██      ██ 
██ ▄▄ ██ ██   ██ ██   ██ ██      ██     ██ ███ ██ ██      ██   ██    ██    ██   ██ ██      ██   ██       ██      ██      ██ 
 ██████  ██   ██ ██████  ██ ███████      ███ ███  ███████ ██   ██    ██    ██   ██ ███████ ██   ██        ██████ ███████ ██ 
    ▀▀                                                                                                                                                                                                                                                
    `)
	fmt.Println("\033[0m")
	fmt.Println("\033[1;36mWelcome to the Weather CLI!\033[0m")
	fmt.Println("Type a city name to get its current weather.")
	fmt.Println("Enter 'exit' to quit anytime.")

	reader := bufio.NewReader(os.Stdin)
	for {
		city := getCity(reader)

		if city == "exit" {
			fmt.Println("\033[1;31mGoodbye!\033[0m")
			os.Exit(0)
		}
		if city == "" {
			fmt.Println("\033[1;31mPlease enter a city name.\033[0m")
			continue
		}
		weather, err := getWeather(city, API_KEY)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		displayWeather(weather)

		fmt.Print("\033[1;33mCheck another city? [y/n]: \033[0m")
		retry, _ := reader.ReadString('\n')
		if strings.TrimSpace(retry) == "n" {
			fmt.Println("\033[1;31mGoodbye!\033[0m")
			os.Exit(0)
		}
	}

}
