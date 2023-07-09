package errFormat

import "fmt"

func FormatError(location, function string, err any) error {
	return fmt.Errorf("%s.%s: %v", location, function, err)
}
