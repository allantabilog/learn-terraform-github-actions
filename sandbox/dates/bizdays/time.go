package bizdays

import (
	"fmt"
	"time"
)

func TimeMain() {
	fmt.Printf("1 business day from today: %v\n", GetBusinessDayFrom(time.Now(), 1))
	fmt.Printf("2 business days from today: %v\n", GetBusinessDayFrom(time.Now(), 2))
	fmt.Printf("3 business days from today: %v\n", GetBusinessDayFrom(time.Now(), 3))
	fmt.Printf("4 business days from today: %v\n", GetBusinessDayFrom(time.Now(), 4))
	fmt.Printf("5  business days from today: %v\n", GetBusinessDayFrom(time.Now(), 5))
}

func ElapsedTime(start time.Time) time.Duration {
	fmt.Println("sleeping...")
	time.Sleep(10 * time.Second)
	t := time.Now()
	elapsed := t.Sub(start)
	return elapsed
}

func GetBusinessDayFromMethod1(startDate time.Time) time.Time {
	return startDate.Add(time.Hour * 24)
}

func GetBusinessDayFromMethod2(startDate time.Time) time.Time {
	return startDate.AddDate(0, 0, 1)
}

func GetBusinessDayFrom(startDate time.Time, days int) time.Time {
	offset := time.Duration(days)
	return startDate.Add(time.Hour * 24 * offset)

}

func ParseDateString(dateStr string) time.Time {
	t, _ := time.Parse(time.RFC822, dateStr)
	fmt.Println(t)
	return t
}

func DatesEqual(date1 time.Time, date2 time.Time) bool {
	return date1.Year() == date2.Year() && date1.Month() == date2.Month() && date1.Day() == date2.Day()
}
