# Usage

1. Add an import statement pointing to `github.com/mikerapa/bowlingscores`.
2. Prepare a slice of strings representing each roll. See the valid input section of this document for more information. 
3. Pass the values into the GetGameStatsFromRolls function. 
4. Check for errors returned from the GetGameStatsFromRolls function. See the section titled Possible Errors for additional information. 
5. Use the Score and GameState values returned. 

```go
package main

import "fmt"

import bowling "github.com/mikerapa/bowlingscores"

func main() {
	fmt.Println("Application running")
	rolls := []string{"X", "2", "/", "8", "/", "7", "1", "X", "8", "/", "9", "/", "8", "-", "X", "8", "/", "5"}
	result, err := bowling.GetGameStatsFromRolls(rolls)
	if err == nil {
		fmt.Println("Score", result.Score)
		fmt.Println("GameState:", result.GameState)
		fmt.Println("guess that worked", result)
	}
}

```

# Valid Input
The slice representing the roll data should contain only characters listed in the table below. Each item in the slice should have just one character. 

Character|Notes
---| --- 
_ | No pins
-| No pins
0| No Pins
1| 1 Pin
2| 2 Pins
3| 3 Pins 
4| 4 Pins 
5| 5 Pins 
6| 6 Pins
7| 7 Pins
8| 8 Pins
9| 9 Pins
/| Spare
\\| Spare
X| Strike


# Possible Errors
If the supplied roll data is not valid, the GetGameStatsFromRolls function will return a error with the following message. `Error: Game stats could not be calculated because the roll data is invalid.` If invalid input is provided, the module will not attempt to calculate the score and the GameState will remain in the NotStarted status. 

# Tests
There are tests included in this packages which cover the following:
1. Validation of roll data
2. Normalization of the roll data
3. Calculation of the game state
4. Calculation of the scores 

Currently the tests cover 100% of the code base. Run the tests with one of the following commands. 

Run tests with code coverage stats

`go test -cover`

Run tests in verbose mode

`go test -v`