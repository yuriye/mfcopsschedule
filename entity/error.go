package entity

import "errors"

//ErrNotFound not found
//goland:noinspection ALL
var ErrNotFound = errors.New("Not found")

//ErrInvalidEntity invalid entity
//goland:noinspection GoErrorStringFormat
var ErrInvalidEntity = errors.New("Invalid entity")

//ErrCannotBeDeleted cannot be deleted
//goland:noinspection ALL
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")
