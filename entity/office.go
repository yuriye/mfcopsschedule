package entity

type (
	//Office data
	Office struct {
		ID                ID
		Name              string
		DefaultEmployeers []*Employee
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