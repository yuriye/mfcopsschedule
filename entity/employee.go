package entity

type (
	//Employee data
	Employee struct {
		ID         ID
		FirstName  string
		Patronymic string
		LastName   string
	}
)

func NewEmployee(firstName, patronymic, lastName string) (*Employee, error) {
	e := &Employee{
		ID:         NewID(),
		FirstName:  firstName,
		Patronymic: patronymic,
		LastName:   lastName,
	}
	err := e.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return e, nil
}

//Validate validate data
func (e *Employee) Validate() error {
	if e.FirstName == "" || e.LastName == "" {
		return ErrInvalidEntity
	}

	return nil
}
