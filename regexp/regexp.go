package regexp

import (
	"fmt"
	"regexp"
)

const email = `^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`

func CheckEmail(s string) bool {
	return regexp.MustCompile(email).MatchString(s)
}

const password = `^(?=.*[A-Z])(?=.*[.?/-])[a-zA-Z0-9.?/-]{5,24}$`

func CheckPassword(s string) {
	fmt.Println(regexp.MustCompile("[a-z]").MatchString(s))
	fmt.Println(regexp.MustCompile("[A-Z]").MatchString(s))
	fmt.Println(regexp.MustCompile("[0-9]").MatchString(s))
	fmt.Println(regexp.MustCompile(`^[0-9a-zA-Z\-^$*.@]+$`).MatchString(s))
	fmt.Println(regexp.MustCompile(`[\-^$*.@]`).MatchString(s))
}
