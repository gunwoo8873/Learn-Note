package error

import (
	"errors"
	"fmt"
)

func Error() {
	err1 := errors.New("Test Error")

	fmt.Println(err1.Error())
}
