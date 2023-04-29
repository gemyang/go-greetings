package hello

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// GO: Hello is an exported func as H is capitalized
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	// GO: := declare and init a var
	message := fmt.Sprintf("Hi, %v.", name)

	// GO: declare, then init
	var message2 string
	message2 = "Welcome!"
	// GO: %v is called fmt verb
	return fmt.Sprintf(randomFormat(), message, message2), nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	// range return two values
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice | An Array of message formats.
	formats := []string{
		"Hi, %v. Welcome! %v",
		"Great to see you, %v! %v",
		"Hail, %v! Well met %v!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
