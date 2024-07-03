package main

type StatusErr struct {
	Status  int
	Message string
}

// implement error interface
func (se StatusErr) Error() string {
	return se.Message
}

func ValidateSomethingAndFails(data string) (bool, error) {
	if len(data) == 0 {
		return false, StatusErr{
			Status:  1,
			Message: "big mistake",
		}
	}
	return true, nil
}
