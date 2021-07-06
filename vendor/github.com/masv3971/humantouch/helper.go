package humantouch

import (
	"math/rand"
	"time"
)

func randomGender(gender string) string {
	if gender == "" {
		switch r := rand.Intn(2); r {
		case 0:
			return GenderFemale
		case 1:
			return GenderMale
		}
	}
	return gender
}

func month(month time.Month) int {
	switch m := month; m {
	case time.Month(1):
		return 1
	case time.Month(2):
		return 2
	case time.Month(3):
		return 3
	case time.Month(4):
		return 4
	case time.Month(5):
		return 5
	case time.Month(6):
		return 6
	case time.Month(7):
		return 7
	case time.Month(8):
		return 8
	case time.Month(9):
		return 9
	case time.Month(10):
		return 10
	case time.Month(11):
		return 11
	case time.Month(12):
		return 12
	}
	return 0
}
