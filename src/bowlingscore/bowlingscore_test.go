package bowlingscore

import (
	"testing"
)

const (
	IncorrectValueInt = "%s has incorrect value: Expected %d, Actual %d"
)

func TestGetEmptyFrames(t *testing.T) {
	emptyFrames := getEmptyFrames()

	// check the correct number of frames exist
	if len(emptyFrames) != 10 {
		t.Fail()
	}

	// Make sure all frames are set up correctly
	for i, currentFrame := range emptyFrames {
		if currentFrame.frameState != frameUnstarted {
			t.Errorf("Frame %d is not in UnStarted state", i+1)
		}

		// check the frame number
		if currentFrame.frameNumber != (i + 1) {
			t.Errorf("Frame has incorrect frameNumber: Expected %d, Actual %d", i+1, currentFrame.frameNumber)
		}

		// the framescore should be 0
		if currentFrame.frameScore != 0 {
			t.Errorf(IncorrectValueInt, "frameScore", 0, currentFrame.frameScore)
		}
	}

}

func TestGetGameStatsFromRolls(t *testing.T) {
	tests := []struct {
		name             string
		rolls            []string
		wantGameState    string
		wantScore        int
		wantCurrentFrame int
		wantError        bool
	}{
		{name: "no rolls", rolls: []string{}, wantGameState: gameNotStarted, wantScore: 0, wantCurrentFrame: 1, wantError: false},
		{name: "invalid roll", rolls: []string{""}, wantGameState: gameNotStarted, wantScore: 0, wantCurrentFrame: 1, wantError: true},
		{name: "invalid roll2", rolls: []string{"A"}, wantGameState: gameNotStarted, wantScore: 0, wantCurrentFrame: 1, wantError: true},
		{name: "1 Strike", rolls: []string{"X"}, wantGameState: gameInProgress, wantScore: 10, wantCurrentFrame: 2, wantError: false},
		{name: "12 Strikes", rolls: []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X"}, wantGameState: gameCompleted, wantScore: 300, wantCurrentFrame: 10, wantError: false},
		{name: "163", rolls: []string{"X", "2", "/", "8", "/", "7", "1", "X", "8", "/", "9", "/", "8", "-", "X", "8", "/", "5"}, wantGameState: gameCompleted, wantScore: 163, wantCurrentFrame: 10, wantError: false},
		{name: "163 with unmarked spare", rolls: []string{"X", "2", "8", "8", "/", "7", "1", "X", "8", "/", "9", "/", "8", "-", "X", "8", "/", "5"}, wantGameState: gameCompleted, wantScore: 163, wantCurrentFrame: 10, wantError: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotGame, err := GetGameStatsFromRolls(tt.rolls)
			if gotGame.gameState != tt.wantGameState {
				t.Errorf("GetGameStatsFromRolls() gotGameState = %v, want %v", gotGame.gameState, tt.wantGameState)
			}
			if gotGame.score != tt.wantScore {
				t.Errorf("GetGameStatsFromRolls() gotScore = %v, want %v", gotGame.score, tt.wantScore)
			}
			if gotGame.currentFrame != tt.wantCurrentFrame {
				t.Errorf("GetGameStatsFromRolls() gotCurrentFrame = %v, want %v", gotGame.currentFrame, tt.wantCurrentFrame)
			}
			if tt.wantError != (err != nil) {
				t.Errorf("GetGameStatsFromRolls() error gotError=%v, wantError=%v", err != nil, tt.wantError)
			}
		})
	}
}

func Test_getEmptyGame(t *testing.T) {
	gotGame := getEmptyGame()

	// make sure the frames exist
	if len(gotGame.frames) != 10 {
		t.Errorf("Test_getEmptyGame: There should be 10 frames")
	}

	// current game
	if gotGame.currentFrame != 1 {
		t.Errorf("Test_getEmptyGame: currentGame not set correctly. Should be 1, got %d", gotGame.currentFrame)
	}
	//score
	if gotGame.score != 0 {
		t.Errorf("Test_getEmptyGame: Score should be 0. Got %d", gotGame.score)
	}

	if gotGame.gameState != gameNotStarted {
		t.Errorf("Test_getEmptyGame: GameState not set to GameNotStarted. Got %v", gotGame.gameState)
	}
}

func TestFrame_calculateFrameScore(t *testing.T) {
	tests := []struct {
		name           string
		rolls          []string
		wantFrameScore int
	}{
		{name: "no rolls", rolls: []string{}, wantFrameScore: 0},
		{name: "srike", rolls: []string{"X"}, wantFrameScore: 10},
		{name: "four", rolls: []string{"4"}, wantFrameScore: 4},
		{name: "four and two", rolls: []string{"4", "2"}, wantFrameScore: 6},
		{name: "four and six", rolls: []string{"4", "6"}, wantFrameScore: 10},
		{name: "four and spare", rolls: []string{"4", "/"}, wantFrameScore: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newGame := getEmptyGame()
			for _, newRoll := range tt.rolls {
				newGame.addRoll(newRoll)
			}
			if gotFrameScore := newGame.frames[newGame.currentFrame-1].calculateFrameScore(); gotFrameScore != tt.wantFrameScore {
				t.Errorf("Frame.calculateFrameScore() = %v, want %v", gotFrameScore, tt.wantFrameScore)
			}
		})
	}
}
