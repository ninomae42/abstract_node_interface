package domain

type ID string

func NewID() ID {
	return ID("id")
}

func (i ID) String() string {
	return string(i)
}
