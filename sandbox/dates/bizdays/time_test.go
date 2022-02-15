package bizdays

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetBusinessDayFrom(t *testing.T) {
	assert := assert2.New(t)

	var tests = []struct {
		inputDate    time.Time
		daysFrom     int
		expectedDate time.Time
	}{
		{time.Now(), 0, time.Now()},
		//{time.Now(), 1, time.Now()},
	}
	for _, test := range tests {
		assert.True(DatesEqual(GetBusinessDayFrom(test.inputDate, test.daysFrom), test.expectedDate))
	}
}

func TestDateEquality(t *testing.T) {
	assert := assert2.New(t)

	today := time.Now()
	assert.True(DatesEqual(today, today))

	day1FromToday := time.Now().Add(time.Hour * 24)
	assert.True(DatesEqual(day1FromToday, day1FromToday))
	assert.True(DatesEqual(time.Now().Add(time.Hour*24), day1FromToday))
}

func TestParseDateString(t *testing.T) {
	assert := assert2.New(t)

	//monday, _ := time.Parse(time.RFC822, "14 Feb 22 00:01 AEST")
	//tuesday, _ := time.Parse(time.RFC822, "15 Feb 22 00:01 AEST")
	//wednesday, _ := time.Parse(time.RFC822, "16 Feb 22 00:01 AEST")
	//thursday, _ := time.Parse(time.RFC822, "17 Feb 22 00:01 AEST")
	//friday, _ := time.Parse(time.RFC822, "18 Feb 22 00:01 AEST")

	assert.True(true)
}
