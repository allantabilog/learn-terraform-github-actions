package main

import (
	"fmt"
	"log"
	"time"
)

type Workspace struct {
	WorkspaceID int
	Status      string
}

func main() {
	w1 := Workspace{WorkspaceID: 92900, Status: "IN_PREPARATION"}
	w2 := Workspace{WorkspaceID: 92900, Status: "ABANDONED"}
	fmt.Println(w1.IsOpen())
	fmt.Println(w2.IsOpen())
}

func (workspace Workspace) IsOpen() bool {
	return workspace.Status != "ABANDONED" && workspace.Status != "SETTLED"
}

func dateMain() {
	today := time.Now()
	settlementDate1 := today.AddDate(0, 0, 5)
	settlementDate2 := today.AddDate(0, 0, 7)
	settlementDate3 := today.AddDate(0, 0, 9)

	var datesToTest []time.Time

	datesToTest = append(datesToTest, settlementDate1)
	datesToTest = append(datesToTest, settlementDate2)
	datesToTest = append(datesToTest, settlementDate3)

	for _, dateToTest := range datesToTest {
		fmt.Printf("%s within 7 days (%s)? %v\n", dateToTest, today.AddDate(0, 0, 7), withinDays(dateToTest, 7))
	}

}

func withinDays(dateToTest time.Time, days int) bool {
	today := time.Now()
	todayPlusNDays := today.AddDate(0, 0, days)
	return (dateToTest.Equal(today) || dateToTest.After(today)) && (dateToTest.Equal(todayPlusNDays) || dateToTest.Before(todayPlusNDays))
}
func dateMain1() {
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
