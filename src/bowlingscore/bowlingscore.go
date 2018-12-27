package bowlingscore

import (
	"errors"
	"fmt"
	"strconv"
)

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
	frameNumber int
	frameState  string
	frameScore  int
	currentBall int
	rolls       [3]string
	bonusCount  int
}

// Game : Game score and state data
type Game struct {
	gameState    string
	score        int
	currentFrame int
	frames       [10]Frame
}

func (g *Game) addRoll(newRoll string) {
	g.frames[g.currentFrame-1].addRoll(newRoll)
}

func (f *Frame) addRoll(newRoll string) {
	f.rolls[f.currentBall-1] = newRoll

	if (f.frameNumber < 10 && (newRoll == "/" || newRoll == "X" || f.currentBall == 2)) || (f.frameNumber == 10) && (newRoll == "/" || newRoll == "X" || f.currentBall == 3) {
		f.frameState = frameComplete
		f.bonusCount = f.calculateFrameBonusCount()
	} else {
		f.frameState = frameInProgress
		f.currentBall++
	}
	f.frameScore = f.calculateFrameScore()
}

func (f *Frame) calculateFrameScore() (frameScore int) {
	skipPrevious := false
	for i := 2; i >= 0; i-- {
		if skipPrevious {
			skipPrevious = false
			continue
		}

		switch f.rolls[i] {
		case "X":
			frameScore += 10
		case "/":
			frameScore += 10
			skipPrevious = true
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			rollValue, _ := strconv.Atoi(f.rolls[i])
			frameScore += rollValue
		}
	}
	return
}

// TODO need to add a test for this
func (f *Frame) calculateFrameBonusCount() (bonusCount int) {
	if f.frameNumber < 10 {
		for i := 1; i >= 0; i-- {
			switch f.rolls[i] {
			case "x":
				return 2
			case "/":
				return 1
			}
		}
	}
	return
}

func getEmptyFrames() (frames [10]Frame) {
	// set up each empty frame
	for i := range frames {
		frames[i].frameNumber = i + 1
		frames[i].frameState = frameUnstarted
		frames[i].frameScore = 0
		frames[i].bonusCount = 0
	}
	return
}

func getEmptyGame() (game Game) {
	game.frames = getEmptyFrames()
	game.score = 0
	game.gameState = gameNotStarted
	game.currentFrame = 1
	game.frames[0].currentBall = 1
	return
}

// GetGameStatsFromRolls : get the score of a game based on roll data
func GetGameStatsFromRolls(rolls []string) (game Game, gameError error) {
	// set up the game state
	game = getEmptyGame()

	rollData, err := getCleanedRollsData(rolls)

	if err != nil {
		// roll data is invalid. Cannot continue
		gameErrorMessage := fmt.Sprintf("Error: Game stats could not be calculated because the roll data is invalid. %v", err)
		gameError = errors.New(gameErrorMessage)
		return
	}

	for _, currentRoll := range rollData {
		game.gameState = gameInProgress
		// add roll to the current frame
		game.addRoll(currentRoll)
		game.updateGameScore()

		// move to the next frame
		game.advanceGameState()
	}

	return
}

func (g *Game) getCurrentFrame() Frame {
	return g.frames[g.currentFrame-1]
}

// TODO need a test for this
func (g *Game) advanceGameState() {
	// move to the next frame
	if g.getCurrentFrame().frameState == frameComplete && g.currentFrame < 10 {
		// advance to the next frame
		g.currentFrame++
		g.frames[g.currentFrame-1].currentBall = 1

	} else if g.getCurrentFrame().frameState == frameComplete && g.currentFrame == 10 {
		// end the game
		g.gameState = gameCompleted
	}
}

func (g *Game) updateGameScore() {
	// apply bonuses
	latestFrameScore := g.frames[g.currentFrame-1].frameScore
	for i := g.currentFrame - 2; i > 0; i-- {
		if g.frames[i].bonusCount > 0 {
			// apply the bonus
			g.frames[i].frameScore += latestFrameScore
			g.frames[i].bonusCount--
		} else {
			// there are no remaining bonuses, stop calculating
			break
		}
	}
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
