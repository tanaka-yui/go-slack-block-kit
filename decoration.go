package slack

import "fmt"

func Bold(value string) string {
	return fmt.Sprintf("*%s*", value)
}
