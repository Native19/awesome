package go3

import (
	"errors"
	"fmt"
)

type MyErr1 struct{}
type MyErr2 struct{}

func Er() {
	var newErr error
	err1, err2 := MyErr1{}, MyErr2{}

	newErr = fmt.Errorf("первая ошибка: %w, %w", err1, err2)
	err3 := fmt.Errorf("вторая ошибка: %w", newErr)
	fmt.Println(err3)

	if errors.Is(err3, MyErr1{}) {
		fmt.Printf("yes")
	} else {
		fmt.Printf("nope")
	}
}

func (MyErr1) Error() string {
	return "MyErr1"
}

func (MyErr2) Error() string {
	return "MyErr2"
}
