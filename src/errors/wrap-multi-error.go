package main

import (
	"errors"
	"fmt"
)

func joinError() error {
	var errs []error
	errs = append(errs, errors.New("field FirstName cannot be empty"))
	errs = append(errs, errors.New("save Corinthians please"))
	errs = append(errs, errors.New("field Email cannot be empty"))
	return errors.Join(errs...)
}

func wError() error {
	err1 := errors.New("Corinthians please stop")
	err2 := errors.New("AAAAAAA")
	err3 := errors.New("third error")
	return fmt.Errorf("first: %w, second: %w, third: %w", err1, err2, err3)
}

func main() {
	err := joinError()
	if err != nil {
		fmt.Println(err)
	}

	err = wError()
	if err != nil {
		fmt.Println(err)
	}
}
