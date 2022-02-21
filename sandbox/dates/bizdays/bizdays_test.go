package bizdays

import (
	"fmt"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBizDaysFrom(t *testing.T) {
	assert := assert2.New(t)

	assert.True(true)

	startDate, _ := time.Parse(time.RFC822, "21 Feb 22 12:00 AEST")

	var testCases = []struct {
		description  string
		inputDate    time.Time
		daysFrom     int
		expectedDate time.Time
	}{
		{"1 biz day from today", startDate, 1, startDate.AddDate(0, 0, 1)},
		{"2 biz days from today", startDate, 2, startDate.AddDate(0, 0, 2)},
		{"3 biz days from today", startDate, 3, startDate.AddDate(0, 0, 3)},
		{"4 biz days from today", startDate, 4, startDate.AddDate(0, 0, 4)},
		{"5 biz days from today", startDate, 5, startDate.AddDate(0, 0, 7)},
	}
	for _, testCase := range testCases {
		assert.True(BizDaysFrom(testCase.inputDate, testCase.daysFrom).Equal(testCase.expectedDate), testCase.description)
	}
}

func TestBizDays_ArbitraryStartDates(t *testing.T) {
	assert := assert2.New(t)

	assert.True(true)

	testDate_Wed, _ := time.Parse(time.RFC822, "16 Feb 22 14:00 AEST")
	testDate_Sat, _ := time.Parse(time.RFC822, "19 Feb 22 14:00 AEST")
	testDate_Sun, _ := time.Parse(time.RFC822, "20 Feb 22 14:00 AEST")

	testCases := []struct {
		description  string
		startDate    time.Time
		daysFrom     int
		expectedDate time.Time
	}{
		{"Wed plus 0 biz day", testDate_Wed, 0, testDate_Wed},
		{"Wed plus 1 biz day", testDate_Wed, 1, testDate_Wed.AddDate(0, 0, 1)},
		{"Wed plus 2 biz days", testDate_Wed, 2, testDate_Wed.AddDate(0, 0, 2)},
		{"Wed plus 3 biz days", testDate_Wed, 3, testDate_Wed.AddDate(0, 0, 5)},
		{"Sat plus 2 biz days", testDate_Sat, 2, testDate_Sat.AddDate(0, 0, 3)},
		{"Sat plus 1 biz days", testDate_Sat, 1, testDate_Sat.AddDate(0, 0, 2)},
		{"Sun plus 1 biz days", testDate_Sun, 1, testDate_Sun.AddDate(0, 0, 1)},
		{"Sun plus 2 biz days", testDate_Sun, 2, testDate_Sun.AddDate(0, 0, 2)},
	}

	for _, testCase := range testCases {
		assert.True(BizDaysFrom(testCase.startDate, testCase.daysFrom).Equal(testCase.expectedDate), testCase.description)
	}
}

func TestBizDays_AcceptanceTests(t *testing.T) {
	assert := assert2.New(t)

	assert.True(true)

	testStartDate, _ := time.Parse(time.RFC822, "18 Feb 22 14:00 AEST")
	testStartDate2, _ := time.Parse(time.RFC822, "21 Feb 22 14:00 AEST")

	testCases := []struct {
		description       string
		startDate         time.Time
		daysFrom          int
		expectedDate      time.Time
		expectedDayOfWeek time.Weekday
	}{
		{"Fri plus 1 biz day", testStartDate, 1, testStartDate.AddDate(0, 0, 3), time.Monday},
		{"Fri plus 2 biz day", testStartDate, 2, testStartDate.AddDate(0, 0, 4), time.Tuesday},
		{"Fri plus 3 biz day", testStartDate, 3, testStartDate.AddDate(0, 0, 5), time.Wednesday},
		{"Fri plus 4 biz day", testStartDate, 4, testStartDate.AddDate(0, 0, 6), time.Thursday},
		{"Fri plus 4 biz day", testStartDate, 5, testStartDate.AddDate(0, 0, 7), time.Friday},
		{"Mon plus 5 biz day", testStartDate2, 5, testStartDate2.AddDate(0, 0, 7), time.Monday},
	}

	for _, testCase := range testCases {
		assert.True(BizDaysFrom(testCase.startDate, testCase.daysFrom).Equal(testCase.expectedDate), testCase.description)
		assert.True(BizDaysFrom(testCase.startDate, testCase.daysFrom).Weekday() == testCase.expectedDayOfWeek)
	}
}

func TestBizDays_EdgeCases(t *testing.T) {

	testStartDate, _ := time.Parse(time.RFC822, "18 Feb 22 14:00 AEST")

	testCases := []struct {
		description       string
		startDate         time.Time
		daysFrom          int
		expectedDate      time.Time
		expectedDayOfWeek time.Weekday
	}{
		{"Fri plus 0 biz day", testStartDate, 0, testStartDate, time.Friday},
		{"Negative biz days", testStartDate, -1, testStartDate.AddDate(0, 0, -1), time.Thursday},
	}

	for _, testCase := range testCases {
		assert2.True(t, BizDaysFrom(testCase.startDate, testCase.daysFrom).Equal(testCase.expectedDate), testCase.description)
	}

	fmt.Println(BizDaysFrom(time.Now(), -1))
}

func TestBizDays_NegativeParam(t *testing.T) {
	testStartDate, _ := time.Parse(time.RFC822, "21 Feb 22 14:00 AEST")

	testCases := []struct {
		description       string
		startDate         time.Time
		daysFrom          int
		expectedDate      time.Time
		expectedDayOfWeek time.Weekday
	}{
		{"Mon plus 1 biz day", testStartDate, 1, testStartDate.AddDate(0, 0, 1), time.Tuesday},
		{"Mon plus 5 biz days", testStartDate, 5, testStartDate.AddDate(0, 0, 7), time.Monday},
		{"Mon plus 0 biz day", testStartDate, 0, testStartDate, time.Monday},
		{"Mon minus 1 biz day", testStartDate, -1, testStartDate.AddDate(0, 0, -3), time.Friday},
		{"Mon minus 2 biz days", testStartDate, -2, testStartDate.AddDate(0, 0, -4), time.Thursday},
	}

	for _, testCase := range testCases {
		assert2.True(t, BizDaysFrom(testCase.startDate, testCase.daysFrom).Equal(testCase.expectedDate), testCase.description)
	}

}
