package domain

type String string

func NewString() String {
	return String("string")
}

func (s String) String() string {
	return string(s)
}

func (s String) Ptr() *String {
	return &s
}
