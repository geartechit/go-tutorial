package timeutil

import "time"

func MustParseTime(layout string) time.Time {
	t, err := time.Parse(time.RFC3339, layout)
	if err != nil {
		panic(err)
	}
	return t
}
