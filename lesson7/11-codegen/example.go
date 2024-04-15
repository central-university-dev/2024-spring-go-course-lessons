package main

// Example the simple struct
type Example struct {
	UserName string
}

func (e *Example) String() string {
	return e.UserName
}
