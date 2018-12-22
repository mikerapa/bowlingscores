package bowlingscore

import (
	"errors"
	"fmt"
	"strings"
)

const (
	validRollString = "_0123456789/X"
)

// remove extra spaces and make alpha chars upper case
func nomalizeRollString(rollString string) string {
	return strings.ToUpper(strings.Trim(rollString, " "))
}

func isValidRollString(rollString string) (valid bool) {
	rollString = nomalizeRollString(rollString)
	valid = strings.Contains(validRollString, rollString) && len(rollString) == 1
	return
}

func getCleanedRollsData(rolls []string) (rollData []string, err error) {
	for i := range rolls {
		if isValidRollString(rolls[i]) {
			rollData = append(rollData, nomalizeRollString(rolls[i]))
		} else {
			// found an invalid entry, stop everything
			errorMessage := fmt.Sprintf("Error importing roll data. Roll %d is invalid.", i)
			return nil, errors.New(errorMessage)
		}
	}
	return rollData, err
}
