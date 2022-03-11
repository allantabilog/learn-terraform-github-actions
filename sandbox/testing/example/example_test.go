package example

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	testCases := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zurich", "13:05"},
		{"12:31", "America/New_York", "07:34"},
		{"08:08", "Australia/Sydney", "18:12"},
	}
	for _, tc := range testCases {

		t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatalf("Could not load location %q", tc.loc)
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("In(%s, %s) = %s; want %s", tc.gmt, tc.loc, got, tc.want)
			}
		})
	}

}
