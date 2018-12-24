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

// func TestGetGameFromRolls(t *testing.T) {
// 	var sampleRolls [21]string
// 	sampleRolls[0] = "X"
// 	frames := GetGameFromRolls(sampleRolls)
// 	if frames[0].frameState != frameComplete {
// 		t.Fail()
// 	}

// }

// func TestGetGameFromRolls(t *testing.T) {
// 	type GameRollTestData struct {
// 		rollData                     [21]string
// 		expectedFramesCompletedCount int
// 	}

// 	var gameTestData = []GameRollTestData{
// 		GameRollTestData{
// 			rollData:                     [21]string{"X", "1"},
// 			expectedFramesCompletedCount: 1,
// 		},
// 		GameRollTestData{
// 			rollData:                     [21]string{"X", "1", "/"},
// 			expectedFramesCompletedCount: 2,
// 		},
// 		GameRollTestData{
// 			rollData:                     [21]string{"X", "1", "0"},
// 			expectedFramesCompletedCount: 2,
// 		},
// 	}

// 	for _, currentGameTestData := range gameTestData {
// 		testName := fmt.Sprintf("Test the number of closed frames:")
// 		t.Run(testName, func(t *testing.T) {
// 			frames := GetGameFromRolls(currentGameTestData.rollData)
// 			completedFrameCount := 0
// 			for h, currentFrame := range frames {
// 				if currentFrame.frameState != frameComplete {
// 					completedFrameCount = h - 1
// 					break
// 				}
// 			}
// 			if completedFrameCount != currentGameTestData.expectedFramesCompletedCount {
// 				t.Fail()
// 			}

// 		})
// 	}

// }

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO do something with the error state
			gotGameState, gotScore, gotCurrentFrame, _, err := GetGameStatsFromRolls(tt.rolls)
			if gotGameState != tt.wantGameState {
				t.Errorf("GetGameStatsFromRolls() gotGameState = %v, want %v", gotGameState, tt.wantGameState)
			}
			if gotScore != tt.wantScore {
				t.Errorf("GetGameStatsFromRolls() gotScore = %v, want %v", gotScore, tt.wantScore)
			}
			if gotCurrentFrame != tt.wantCurrentFrame {
				t.Errorf("GetGameStatsFromRolls() gotCurrentFrame = %v, want %v", gotCurrentFrame, tt.wantCurrentFrame)
			}
			if tt.wantError != (err != nil) {
				t.Errorf("GetGameStatsFromRolls() error gotError=%v, wantError=%v", err != nil, tt.wantError)
			}
		})
	}
}
