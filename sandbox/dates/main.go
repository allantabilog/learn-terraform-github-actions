package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	dateMain()
}

func dateMain() {
	bizDay := BizDaysFrom(time.Now(), 5)
	fmt.Println("5 biz days from today is : ", dateInfo(bizDay))
	bizDay = BizDaysFrom(time.Now(), 4)
	fmt.Println("4 biz days from today is : ", dateInfo(bizDay))
}

func BizDaysFrom_v1(startDate time.Time, days int) time.Time {
	// calculate N biz days from startDate, excluding weekends
	var nextDate time.Time
	for i := 1; i <= days; i++ {
		nextDate = startDate.AddDate(0, 0, i)
		log.Println(dateInfo(nextDate))
		if nextDate.Weekday() == time.Saturday {
			log.Println("Saturday: adjusting calculation")
			nextDate = nextDate.AddDate(0, 0, 2)
		}
		if nextDate.Weekday() == time.Sunday {
			log.Println("Sunday: adjusting calculation")
			nextDate = nextDate.AddDate(0, 0, 1)
		}
	}
	return nextDate
}

func BizDaysFrom(startDate time.Time, daysFrom int) time.Time {
	acc, nextDate := 0, startDate

	for acc < daysFrom {
		nextDate = nextDate.AddDate(0, 0, 1)
		//if nextDate.Weekday() == time.Saturday || nextDate.Weekday() == time.Sunday {
		//	log.Printf("skipping %v", nextDate.Weekday())
		//} else {
		//	acc++
		//}
		if nextDate.Weekday() != time.Saturday && nextDate.Weekday() != time.Sunday {
			acc++
		}
	}
	return nextDate
}

func dateInfo(dateRequired time.Time) string {
	return fmt.Sprintf("%v is a %v", dateRequired, dateRequired.Weekday())
}
