package common

import "regexp"

const EMAIL_REGEX string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
const NUMERIC_REGEX string = "[0-9]+"
const NAME_REGEX string = "[A-Za-z ]+"

func ValidateString(regex string, value string) bool {
	ret, _ := regexp.MatchString(regex, value)

	return ret
}
