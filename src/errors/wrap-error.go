package main

import (
	"errors"
	"fmt"
)

type StatusErrWrap struct {
	Status  int
	Message string
	Err     error
}

// implement error interface
func (se StatusErrWrap) Error() string {
	return se.Message
}

func (se StatusErrWrap) Unwrap() error {
	return se.Err
}

func login() (string, error) {
	return "", errors.New("fake error")
}

func ValidateAndFails(data string) (string, error) {
	token, err := login()
	if err != nil {
		return "", StatusErrWrap{
			Status:  1,
			Message: "login error",
			Err:     err,
		}
	}
	return token, nil
}

func main() {
	_, err := ValidateAndFails("")
	fmt.Println(err)
	wrapErr := errors.Unwrap(err)
	fmt.Println(wrapErr)

	wrappingErr := fmt.Errorf("error in main: %w", err)
	fmt.Println(wrappingErr)
	wrapErr = errors.Unwrap(wrappingErr)
	fmt.Println(wrapErr)
}
