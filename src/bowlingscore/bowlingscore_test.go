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
		if currentFrame.frameNumnber != (i + 1) {
			t.Errorf("Frame has incorrect frameNumber: Expected %d, Actual %d", i+1, currentFrame.frameNumnber)
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
