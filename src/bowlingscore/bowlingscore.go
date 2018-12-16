package bowlingscore

const (
	frameComplete   = "Completed"
	frameUnstarted  = "Unstarted"
	frameInProgress = "InProgress"
)

const (
	gameCompleted  = "Completed"
	gameInProgress = "InProgress"
	gameNotStarted = "NotStarted"
)

// Frame : represtentation of a frame
type Frame struct {
	frameNumnber int
	frameState   string
	frameScore   int
}

func getEmptyFrames() (frames [10]Frame) {
	// set up each empty frame
	for i := range frames {
		frames[i].frameNumnber = i + 1
		frames[i].frameState = frameUnstarted
		frames[i].frameScore = 0
	}
	return
}

// GetGameStatsFromRolls : get the score of a game based on roll data
func GetGameStatsFromRolls(rolls []string) (gameState string, score int, currentFrame int, frames [10]Frame) {

	gameState = gameNotStarted
	score = 0
	currentFrame = 1

	return
}

// GetGameFromRolls : Get the frames from a list of rolls
func GetGameFromRolls(rolls [21]string) [10]Frame {
	var frames [10]Frame
	currentFrame := 0
	for currentRoll := 0; currentRoll < 21; currentRoll++ {
		if rolls[currentRoll] == "X" || rolls[currentRoll] == "/" {
			frames[currentFrame].frameState = frameComplete
			currentFrame++
		}
	}
	return (frames)
}
