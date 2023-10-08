package main

import (
	"strings"
	"testing"
	"time"
)

func BenchmarkFyne(b *testing.B) {
	for i := 0; i < 1000; i++ {
		w := &Week{
			Monday:    true,
			Tuesday:   true,
			Wednesday: false,
			Thursday:  false,
			Friday:    false,
			Saturday:  false,
			Sunday:    false,
		}
		countOfDays := w.GetCountOfDays()
		now := time.Now()
		var (
			finishDate           = time.Date(now.Year()+1, now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			stringBuilder        = strings.Builder{}
			weekDayInYear uint64 = 52
		)
		stringBuilder.Grow(int(weekDayInYear * countOfDays))
		stringBuilder.WriteString(getDateByWeekdays(w, &now))
		for now.UnixNano() < finishDate.UnixNano() {
			now = now.AddDate(0, 0, 1)
			stringBuilder.WriteString(getDateByWeekdays(w, &now))
		}
		stringBuilder.Reset()
	}
}
