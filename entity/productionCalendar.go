package entity

type (
	ProductionCalendar struct {
		Year   int
		Months [12]*[]int
	}
)
