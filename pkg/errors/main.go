package errors

import "fmt"

func NewEf(e error, s string) error {
	fmt.Println(s)
	return e
}
