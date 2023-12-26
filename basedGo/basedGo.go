package basedGo

import "errors"
import "fmt"

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty name !")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
