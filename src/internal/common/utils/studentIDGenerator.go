package utils

import (
	"fmt"
	"strconv"
	"time"
)

func StudentIDGenerator(universityCode int, semester string, lastSID int32) int32 {
	currentYear, _, _ := time.Now().Date()
	lastTwoDigitsOfYear := currentYear % 100

	var semesterCode int
	switch semester {
	case "Fall":
		semesterCode = 0
	case "Spring":
		semesterCode = 1
	case "Summer":
		semesterCode = 3
	default:
		semesterCode = 0
	}

	lastFourDigits := lastSID + 1

	studentNumber, _ := strconv.ParseInt(fmt.Sprintf("%d%02d%d%04d", universityCode, lastTwoDigitsOfYear, semesterCode, lastFourDigits), 10, 32)
	return int32(studentNumber)
}
