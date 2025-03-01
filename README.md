# Weather CLI

A sleek, colorful command-line tool to fetch and display current weather data for any city using the OpenWeatherMap API.

## Features
- Fetch real-time weather data by city name.
- Displays temperature (in Celsius), weather conditions, and city name.
- Stylish ASCII art title and colorful terminal output.
- Supports multi-word city names (e.g., "New York").
- Interactive: keeps running until you choose to exit.

## Prerequisites
- [Go](https://golang.org/dl/) (version 1.16 or later) installed.
- An [OpenWeatherMap API key](https://openweathermap.org/api) (free tier available).
- Git installed for cloning the repo.

## Installation
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/MRQ67/my-weather-cli.git
   cd my-weather-cli
   ```


## Usage
1. **Set Your API Key**: Store your OpenWeatherMap API key as an environment variable:
    - Windows (Powershell):
    ```powershell
    $env:WEATHER_API_KEY = "your-api-key-here"
    ```

    - Mac OS/Linux:
    ```bash
    export WEATHER_API_KEY="your-api-key-here"
    ```


2. **Run the Program**:
    ```bash
    go run main.go
    ```

3. **Follow the prompts**:
   - Enter a city name (e.g., "London" or "New York").
   - Type `exit` to quit.

### Example Output
```


 ██████   █████  ██████  ██ ███████     ██     ██ ███████  █████  ████████ ██   ██ ███████ ██████         ██████ ██      ██ 
██    ██ ██   ██ ██   ██ ██ ██          ██     ██ ██      ██   ██    ██    ██   ██ ██      ██   ██       ██      ██      ██ 
██    ██ ███████ ██   ██ ██ ███████     ██  █  ██ █████   ███████    ██    ███████ █████   ██████  █████ ██      ██      ██ 
██ ▄▄ ██ ██   ██ ██   ██ ██      ██     ██ ███ ██ ██      ██   ██    ██    ██   ██ ██      ██   ██       ██      ██      ██ 
 ██████  ██   ██ ██████  ██ ███████      ███ ███  ███████ ██   ██    ██    ██   ██ ███████ ██   ██        ██████ ███████ ██ 
    ▀▀                                                                                                                                            

    

Welcome to the Weather CLI!
Type a city name to get its current weather.
Enter 'exit' to quit anytime.
Enter city name: Adama
---------------------------------
City: Adama
Temperature: 21.00°C
Condition: clear sky ☀️
---------------------------------
Check another city? [y/n]: 
```
## Configuration
- **API Key**: Store your key in the `WEATHER_API_KEY` environment variable. Do not hardcode it in the source code.
- **Customization**: Edit `main.go` to change colors, emojis, or add features like Fahrenheit support.

## Contributing
1. Fork the repo.
2. Create a branch: `git checkout -b feature/your-feature`.
3. Commit changes: `git commit -m "Add your feature"`.
4. Push to GitHub: `git push origin feature/your-feature`.
5. Open a Pull Request.

## License
This project is licensed under the MIT License—see [LICENSE](LICENSE) for details.

## Acknowledgments
- Built with [Go](https://golang.org/).
- Weather data provided by [OpenWeatherMap](https://openweathermap.org/).
- Inspired by a love for clean CLI tools!


    
