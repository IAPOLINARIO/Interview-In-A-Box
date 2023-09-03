package parser

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func ParseDate(dateStr string) (time.Time, error) {
	if len(dateStr) != 6 {
		return time.Time{}, errors.New("incorrect date format")
	}

	month, err := strconv.Atoi(dateStr[:2])
	if err != nil || month < 1 || month > 12 {
		return time.Time{}, errors.New("invalid month")
	}

	year, err := strconv.Atoi(dateStr[2:])
	if err != nil {
		return time.Time{}, errors.New("invalid year")
	}

	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC), nil
}

func GenerateMonthList(startDate, endDate string) ([]string, error) {
	start, err := ParseDate(startDate)
	if err != nil {
		return nil, err
	}

	end, err := ParseDate(endDate)
	if err != nil {
		return nil, err
	}

	if end.Before(start) {
		return nil, errors.New("endDate must be greater than or equal to startDate")
	}

	var monthList []string
	for !start.After(end) {
		monthList = append(monthList, GetMonthString(start))
		start = start.AddDate(0, 1, 0)
	}

	return monthList, nil
}

func GetMonthString(date time.Time) string {
	return strings.ToLower(date.Format("Jan_06"))
}
