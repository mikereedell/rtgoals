package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/mikereedell/rtgoals/config"
)

const (
	RT_API_BASE_URL = "https://www.rescuetime.com/anapi/data?key=%s&format=json&pv=member&rk=productivity&rb=%s&re=%s"
	TIME_FORMAT     = "2006-01-02"
	WEEKLY          = "Weekly"
	DAILY           = "Daily"
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

	hasDailyGoals := false
	hasWeeklyGoals := false

	for _, goal := range configuration.Goals {
		if goal.TimeWindow == DAILY {
			hasDailyGoals = true
		}
		if goal.TimeWindow == WEEKLY {
			hasWeeklyGoals = true
		}
	}
	if hasDailyGoals {
		computeGoals(configuration, DAILY)
	}

	if hasWeeklyGoals {
		computeGoals(configuration, WEEKLY)
	}
}

func computeGoals(configuration *config.Config, timeWindow string) {
	now := time.Now()
	var startTime string
	var endTime string

	if timeWindow == DAILY {
		startTime = now.Format(TIME_FORMAT)
		endTime = now.Add(24 * time.Hour).Format(TIME_FORMAT)
	} else {
		hoursToSubtract := 0 - (int(now.Weekday())-1)*24
		startTime = now.Add(time.Duration(hoursToSubtract) * time.Hour).Format(TIME_FORMAT)

		hoursToAdd := (6 - int(now.Weekday())) * 24
		endTime = now.Add(time.Duration(hoursToAdd) * time.Hour).Format(TIME_FORMAT)
	}

	data := getRescueTimeData(startTime, endTime, configuration.ApiKey)

	totalProductiveSeconds := 0
	totalUnproductiveSeconds := 0

	for _, row := range data.Rows {
		seconds := int(row[1].(float64))
		productivity := int(row[2].(float64))
		if productivity > 0 {
			totalProductiveSeconds += seconds
		}
		if productivity < 0 {
			totalUnproductiveSeconds += seconds
		}
	}

	goals := configuration.GoalsForTimeWindow(timeWindow)
	for _, goal := range goals {
		if goal.Type == "Productive" {
			printGoalProgress(timeWindow, totalProductiveSeconds, goal)
		} else {
			printGoalProgress(timeWindow, totalUnproductiveSeconds, goal)
		}
	}
}

func getRescueTimeData(startTime, endTime, apiKey string) *rtData {
	response, err := http.Get(fmt.Sprintf(RT_API_BASE_URL, apiKey, startTime, endTime))
	if err != nil {
		fmt.Printf("Error calling RescueTime API: %v\n", err)
	}

	decoder := json.NewDecoder(response.Body)
	rtData := &rtData{}

	if err = decoder.Decode(&rtData); err != nil {
		fmt.Printf("Unable to decode RescueTime JSON data: %v\n", err)
	}
	return rtData
}

func printGoalProgress(timeFrame string, seconds int, goal *config.Goal) {
	goalDuration, _ := time.ParseDuration(fmt.Sprintf("%s", goal.GoalTime))

	time, _ := time.ParseDuration(fmt.Sprintf("%ds", seconds))
	percentage := (float64(time) / float64(goalDuration)) * float64(100)
	fmt.Printf("%s %s time: %s (%4.2f%% of goal)\n", timeFrame, goal.Type, time.String(), percentage)
}
