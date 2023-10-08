package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/goodsign/monday"
)

const (
	format = "02.01.2006"
)

var (
	mapWeekday = map[time.Weekday]string{
		time.Monday:    "Понедельник",
		time.Tuesday:   "Вторник",
		time.Wednesday: "Среда",
		time.Thursday:  "Четверг",
		time.Friday:    "Пятница",
		time.Saturday:  "Суббота",
		time.Sunday:    "Воскресенье",
	}
)

type Week struct {
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool
	Sunday    bool
}

func (w *Week) GetCountOfDays() uint64 {
	var (
		count uint64 = 0
	)
	if w.Monday {
		count++
	}
	if w.Tuesday {
		count++
	}
	if w.Wednesday {
		count++
	}
	if w.Thursday {
		count++
	}
	if w.Friday {
		count++
	}
	if w.Saturday {
		count++
	}
	if w.Sunday {
		count++
	}
	return count
}

func (w *Week) GetWeekdays() []time.Weekday {
	var (
		weekdays []time.Weekday = make([]time.Weekday, 0, 1)
	)
	if w.Monday {
		weekdays = append(weekdays, time.Monday)
	}
	if w.Tuesday {
		weekdays = append(weekdays, time.Tuesday)
	}
	if w.Wednesday {
		weekdays = append(weekdays, time.Wednesday)
	}
	if w.Thursday {
		weekdays = append(weekdays, time.Thursday)
	}
	if w.Friday {
		weekdays = append(weekdays, time.Friday)
	}
	if w.Saturday {
		weekdays = append(weekdays, time.Saturday)
	}
	if w.Sunday {
		weekdays = append(weekdays, time.Sunday)
	}
	return weekdays
}

func (w *Week) Contains(weekday *time.Weekday) bool {
	switch *weekday {
	case time.Monday:
		return w.Monday
	case time.Tuesday:
		return w.Tuesday
	case time.Wednesday:
		return w.Wednesday
	case time.Thursday:
		return w.Thursday
	case time.Friday:
		return w.Friday
	case time.Saturday:
		return w.Saturday
	case time.Sunday:
		return w.Sunday
	default:
		return false
	}
}

func main() {
	startProgramm()
}

func startProgramm() {
	var (
		week *Week = new(Week)
	)
	myApp := app.New()
	window := myApp.NewWindow("Получение дней недели по датам")
	window.Resize(fyne.Size{
		Width:  700,
		Height: 700,
	})
	display := widget.NewEntry()
	scroll := container.NewVScroll(display)
	scroll.SetMinSize(fyne.NewSize(500, 500))

	currentDate := widget.NewEntry()
	errorLabel := widget.NewLabel("")
	buttonMonday := widget.NewCheck("Понедельник", func(value bool) {
		week.Monday = value
	})
	buttonTuesday := widget.NewCheck("Вторник", func(value bool) {
		week.Tuesday = value
	})
	buttonWednesday := widget.NewCheck("Среда", func(value bool) {
		week.Wednesday = value
	})
	buttonThursday := widget.NewCheck("Четверг", func(value bool) {
		week.Thursday = value
	})
	buttonFriday := widget.NewCheck("Пятница", func(value bool) {
		week.Friday = value
	})
	buttonSaturday := widget.NewCheck("Суббота", func(value bool) {
		week.Saturday = value
	})
	buttonSunday := widget.NewCheck("Воскресенье", func(value bool) {
		week.Sunday = value
	})
	buttonStart := widget.NewButton("Получить даты", func() {
		display.Refresh()
		errorLabel.Refresh()
		countOfDays := week.GetCountOfDays()
		if countOfDays != 0 {
			if err := setWeekdays(currentDate, display, errorLabel, week, countOfDays); err != nil {
				errorLabel.SetText(err.Error())
			}
		}
	})

	window.SetContent(container.NewVBox(
		errorLabel,
		currentDate,
		buttonMonday,
		buttonTuesday,
		buttonWednesday,
		buttonThursday,
		buttonFriday,
		buttonSaturday,
		buttonSunday,
		buttonStart,
		scroll,
	))

	window.ShowAndRun()
}

func setWeekdays(currentDate *widget.Entry, display *widget.Entry, errorLabel *widget.Label, week *Week, countOfDays uint64) error {
	now, err := time.Parse(format, currentDate.Text)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Неверный формат данных. Должен быть %s", format))
	}
	var (
		finishDate           = time.Date(now.Year()+1, now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		stringBuilder        = strings.Builder{}
		weekDayInYear uint64 = 52
	)
	stringBuilder.Grow(int(weekDayInYear * countOfDays))
	stringBuilder.WriteString(getDateByWeekdays(week, &now))
	for now.UnixNano() < finishDate.UnixNano() {
		now = now.AddDate(0, 0, 1)
		stringBuilder.WriteString(getDateByWeekdays(week, &now))
	}
	display.SetText(stringBuilder.String())
	stringBuilder.Reset()
	return nil
}

func getDateByWeekdays(week *Week, now *time.Time) string {
	weekday := now.Weekday()
	if week.Contains(&weekday) {
		return fmt.Sprintln(mapWeekday[now.Weekday()] + " : " + monday.Format(*now, format, monday.LocaleRuRU))
	}
	return ""
}
