package basedGo

import "errors"
import "fmt"
import "math/rand"

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty name !")
	}

	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	fmt.Println(messages)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
