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
                "TimeWindow": "Daily",
                "GoalTime": "4h"
             },
             {
                "Type": "Unproductive",
                "TimeWindow": "Daily",
                "GoalTime": "1h15m"
             },
             {
                "Type": "Productive",
                "TimeWindow": "Weekly",
                "GoalTime": "20h"
             },
             {
                "Type": "Unproductive",
                "TimeWindow": "Weekly",
                "GoalTime": "6h15m"
             }
    ]
}

##Usage

	--$ rtgoals
    Daily Productive time: 3h36m23s (90.16% of goal)
    Daily Unproductive time: 50m50s (67.78% of goal)
    Weekly Productive time: 3h36m23s (18.03% of goal)
    Weekly Uproductive time: 50m50s (13.56% of goal)


