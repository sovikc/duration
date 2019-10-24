package duration

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	separator             = "/"
	lowerDayLimit         = 1
	upperDayLimit         = 31
	lowerMonthLimit       = 1
	upperMonthLimit       = 12
	lowerYearLimit        = 1901
	upperYearLimit        = 2999
	validInputDayLength   = 2
	validInputMonthLength = 2
	validInputYearLength  = 4
	leapFebruaryDays      = 29
	startDateLabel        = "Start"
	endDateLabel          = "End"
)

var (
	incorrectDateFormatErrMessage = "%s Date format is incorrect"
	incorrectDateRangeErrMessage  = "%s Date is not within acceptable range"
	incorrectDateErrMessage       = "%s Date is not a valid date"
)

// Gregorian calendar consists of 12 months
// February - 28 days in a common year and 29 days in leap years
var (
	daysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

func isValidFormat(date, label string) (bool, error) {
	parts := getParts(date)

	if !hasThreeValidDateParts(parts) {
		return false, fmt.Errorf(getErrorMessage(incorrectDateFormatErrMessage, label))
	}

	dd := parts[0]
	mm := parts[1]
	yyyy := parts[2]

	if !hasValidIndvidualPartLengths(dd, mm, yyyy) {
		return false, fmt.Errorf(getErrorMessage(incorrectDateFormatErrMessage, label))
	}

	if !isValidDay(dd) {
		return false, fmt.Errorf(getErrorMessage(incorrectDateErrMessage, label))
	}

	if !isValidMonth(mm) {
		return false, fmt.Errorf(getErrorMessage(incorrectDateErrMessage, label))
	}

	if !isYearWithinRange(yyyy) {
		return false, fmt.Errorf(getErrorMessage(incorrectDateRangeErrMessage, label))
	}

	return true, nil

}

func getErrorMessage(message, label string) string {
	return fmt.Sprintf(message, label)
}

func getDatePart(date, label string) (*datePart, error) {
	parts := getParts(date)
	dd := parts[0]
	mm := parts[1]
	yyyy := parts[2]

	day, err := getNumericDatePart(dd)
	if err != nil {
		return nil, fmt.Errorf(getErrorMessage(incorrectDateFormatErrMessage, label))
	}

	month, err := getNumericDatePart(mm)
	if err != nil {
		return nil, fmt.Errorf(getErrorMessage(incorrectDateFormatErrMessage, label))
	}

	year, err := getNumericDatePart(yyyy)
	if err != nil {
		return nil, fmt.Errorf(getErrorMessage(incorrectDateFormatErrMessage, label))
	}

	if !isValidCombination(day, month, year) {
		return nil, fmt.Errorf(getErrorMessage(incorrectDateErrMessage, label))
	}

	return &datePart{
		dd:   day,
		mm:   month,
		yyyy: year,
	}, nil
}

func getParts(date string) []string {
	return strings.Split(date, separator)
}

func hasThreeValidDateParts(parts []string) bool {
	return len(parts) == 3
}

func hasValidIndvidualPartLengths(dd, mm, yyyy string) bool {
	return hasValidLength(dd, validInputDayLength) &&
		hasValidLength(mm, validInputMonthLength) &&
		hasValidLength(yyyy, validInputYearLength)
}

func hasValidLength(part string, length int) bool {
	return len(part) == length
}

func isYearWithinRange(yyyy string) bool {
	return isValidDatePart(yyyy, lowerYearLimit, upperYearLimit)
}

func isValidDay(dd string) bool {
	return isValidDatePart(dd, lowerDayLimit, upperDayLimit)
}

func isValidMonth(mm string) bool {
	return isValidDatePart(mm, lowerMonthLimit, upperMonthLimit)
}

func isValidDatePart(part string, lowerLimit, upperLimit int) bool {
	number, err := getNumericDatePart(part)
	if err != nil {
		log.Println(err)
		return false
	}

	return number >= lowerLimit && number <= upperLimit
}

func getNumericDatePart(part string) (int, error) {
	return strconv.Atoi(part)
}

func isValidCombination(dd, mm, yyyy int) bool {
	switch mm {
	case 2:
		return isValidFebruaryDay(dd, mm, yyyy)

	default:
		validNumberOfDays := daysInMonth[mm-1]
		return validNumberOfDays >= dd
	}

}

func isValidFebruaryDay(dd, mm, yyyy int) bool {
	if mm != 2 {
		return false
	}

	if dd > leapFebruaryDays {
		return false
	}

	switch dd {
	case leapFebruaryDays:
		if !isYearDivisible(yyyy, 4) {
			return false
		}

		if isYearDivisible(yyyy, 100) &&
			!isYearDivisible(yyyy, 400) {
			return false
		}

		return true

	default:
		return true
	}
}

func isYearDivisible(yyyy, divisor int) bool {
	return yyyy%divisor == 0
}
