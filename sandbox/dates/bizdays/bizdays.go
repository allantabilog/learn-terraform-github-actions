package bizdays

import (
	"fmt"
	"time"
)

func BizDaysFrom(startDate time.Time, daysFrom int) time.Time {

	//if daysFrom == 0 {
	//	return startDate
	//}

	acc, nextDate := 0, startDate

	for acc < daysFrom {
		nextDate = nextDate.AddDate(0, 0, 1)
		if nextDate.Weekday() != time.Saturday && nextDate.Weekday() != time.Sunday {
			acc++
		}
	}
	return nextDate
}

func dateInfo(dateRequired time.Time) string {
	return fmt.Sprintf("%v is a %v", dateRequired, dateRequired.Weekday())
}
