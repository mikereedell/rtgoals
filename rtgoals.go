package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/mikereedell/rtgoals/config"
)

type rtData struct {
	Notes      string          `json:"notes"`
	RowHeaders []string        `json:"row_headers"`
	Rows       [][]interface{} `json:"rows"`
}

func readConfig() *config.Config {
	homedir := os.Getenv("HOME")
	configFilePath := fmt.Sprintf("%s/.rtgoals", homedir)
	configuration, err := config.NewConfig(configFilePath)
	if err != nil {
		fmt.Printf("Error reading configuration file %s: %v\n", configFilePath, err)
		os.Exit(1)
	}
	return configuration
}

func main() {
	configuration := readConfig()

	now := time.Now()
	timeFormat := "2006-01-02"
	startTime := now.Format(timeFormat)
	endTime := now.Add(24 * time.Hour).Format(timeFormat)

	baseURL := "https://www.rescuetime.com/anapi/data?key=%s&format=json&pv=member&rk=productivity&rb=%s&re=%s"
	response, err := http.Get(fmt.Sprintf(baseURL, configuration.ApiKey, startTime, endTime))
	if err != nil {
		fmt.Printf("Error calling RescueTime API: %v\n", err)
	}

	decoder := json.NewDecoder(response.Body)
	rtData := &rtData{}

	if err = decoder.Decode(&rtData); err != nil {
		fmt.Printf("Unable to decode RescueTime JSON data: %v\n", err)
	}

	totalProductiveSeconds := 0
	totalUnproductiveSeconds := 0

	for _, row := range rtData.Rows {
		seconds := int(row[1].(float64))
		productivity := int(row[2].(float64))
		if productivity > 0 {
			totalProductiveSeconds += seconds
		}
		if productivity < 0 {
			totalUnproductiveSeconds += seconds
		}
	}

	printProgress(totalProductiveSeconds, configuration.Goals[0])
	printProgress(totalUnproductiveSeconds, configuration.Goals[1])
}

func printProgress(seconds int, goal config.Goal) {
	goalDuration, _ := time.ParseDuration(fmt.Sprintf("%s", goal.GoalTime))

	time, _ := time.ParseDuration(fmt.Sprintf("%ds", seconds))
	percentage := (float64(time) / float64(goalDuration)) * float64(100)
	fmt.Printf("%s time: %s (%4.2f%% of goal)\n", goal.Type, time.String(), percentage)
}
