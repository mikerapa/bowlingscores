package bowlingscore

import (
	"errors"
	"fmt"
	"strings"
)

const (
	validRollString = "_0123456789/X"
)

func isValidRollString(rollString string) (valid bool) {
	rollString = strings.Trim(rollString, " ")
	valid = strings.Contains(validRollString, rollString) && len(rollString) == 1
	return
}

func getCleanedRollsData(rolls []string) (rollData []string, err error) {
	for i := range rolls {
		if isValidRollString(rolls[i]) {
			rollData = append(rollData, strings.Trim(rolls[i], " "))
		} else {
			// found an invalid entry, stop everything
			errorMessage := fmt.Sprintf("Error importing roll data. Roll %d is invalid.", i)
			return nil, errors.New(errorMessage)
		}
	}
	return rollData, err
}
