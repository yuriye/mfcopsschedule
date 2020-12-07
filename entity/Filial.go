package entity

type (
	Filial struct {
		ID        ID
		Name      string
		Employees []*Employee
		Offices   []*Office
	}
)

func NewFilial(name string, employees []*Employee, offices []*Office) (*Filial, error) {
	f := &Filial{
		ID:        NewID(),
		Name:      name,
		Employees: employees,
		Offices:   offices,
	}
	err := f.Validate()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f Filial) Validate() error {
	if f.Name == "" {
		return ErrInvalidEntity
	}
	return nil
}
