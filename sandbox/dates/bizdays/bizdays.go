package bizdays

import (
	"fmt"
	"log"
	"time"
)

func BizDaysFrom(startDate time.Time, daysFrom int) time.Time {

	acc, nextDate := 0, startDate

	if daysFrom >= 0 {
		for acc < daysFrom {
			nextDate = nextDate.AddDate(0, 0, 1)
			if nextDate.Weekday() != time.Saturday && nextDate.Weekday() != time.Sunday {
				acc++
			}
		}
	} else {
		for acc > daysFrom {
			nextDate = nextDate.AddDate(0, 0, -1)
			if nextDate.Weekday() != time.Saturday && nextDate.Weekday() != time.Sunday {
				acc--
			}
		}
	}

	log.Printf("Calculated nextDate as : %v (%v)", nextDate, nextDate.Weekday())
	return nextDate
}

func dateInfo(dateRequired time.Time) string {
	return fmt.Sprintf("%v is a %v", dateRequired, dateRequired.Weekday())
}
