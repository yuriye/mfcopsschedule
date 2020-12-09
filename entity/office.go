package entity

type (
	//Office data
	Office struct {
		ID                ID
		Name              string
		DefaultEmployeers []*Employee
		WeekSchedule      OfficeWeekSchedule
	}
)

func NewOffice(name string, employees []*Employee) (*Office, error) {
	o := &Office{
		ID:                NewID(),
		Name:              name,
		DefaultEmployeers: employees,
	}

	err := o.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return o, nil
}

func (o *Office) Validate() error {
	if o.Name == "" {
		return ErrInvalidEntity
	}
	return nil
}

func (o Office) SetWeekSchedule(schedule OfficeWeekSchedule) {
	o.WeekSchedule = schedule
}

func (o Office) SetEmployees(employees []*Employee) {
	o.DefaultEmployeers = employees
}
