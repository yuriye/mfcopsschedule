package entity

import "time"

type (
	WorkingDayProperties struct {
		Start    time.Time
		End      time.Time
		Duration time.Duration
	}

	TypeOfDay int

	ScheduleDay struct {
		Date       time.Time
		Type       TypeOfDay
		Properties WorkingDayProperties
	}

	MonthSchedule struct {
		Year  int
		Month int
		Days  [31]ScheduleDay
	}

	WorkedDay struct {
		Date       time.Time
		Type       TypeOfDay
		Properties WorkingDayProperties
	}

	MonthWorkedHistory struct {
		Year  int
		Month int
		Days  [31]WorkedDay
	}

	OfficeDay struct {
		Open  bool
		Start time.Time
		End   time.Time
	}

	OfficeWeekSchedule [7]OfficeDay
)

const (
	Working TypeOfDay = iota
	Rest
	Vacation
	Ill
)

func (m MonthSchedule) NumberOfDays() int {
	return 32 - time.Date(m.Year, time.Month(m.Month), 32, 0, 0, 0, 0, time.UTC).Day()
}
