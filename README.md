#RTGoals
Command line client for tracking RescueTime productivity goals, written in Go.

##Installation
Make sure your go environment is setup such that the directory where `go get` puts binaries is on your $PATH.

	go get github.com/mikereedell/rtgoals

##Setup
rtgoals uses configuration stored in ~/.rtgoals.  This file contains the information needed to connect to RescueTime and to calculate your goal progress,

To get rtgoals working:

1. Go to https://www.rescuetime.com/anapi/manage to create an API key for your RescueTime account.
2. Copy the below .rtgoals template to ~/.rtgoals.
3. Edit ~/.rtgoals, copying your API key into the "ApiKey" field.

	{
    "ApiKey": "$YOUR_RT_API_KEY",
    "Goals": [
             {
                "Type": "Productive",
                "TimeWindow": "day",
                "GoalTime": "4h"
             },
             {
                "Type": "Unproductive",
                "TimeWindow": "day",
                "GoalTime": "1h15m"
             },
             {
                "Type": "Productive",
                "TimeWindow": "week",
                "GoalTime": "20h"
             },
             {
                "Type": "Unproductive",
                "TimeWindow": "week",
                "GoalTime": "6h15m"
             }
    ]
}

##Usage

	--$ rtgoals
	Productive time: 2h21m26s (58.93% of goal)
	Unproductive time: 29m1s (38.69% of goal)


