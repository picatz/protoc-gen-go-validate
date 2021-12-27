package example

import (
	fmt "fmt"
	regexp "regexp"
	strings "strings"
	unicode "unicode"
)

// Validate applies configured validation rule options from the protobuf.
func (x *Request) Validate() error {
	if len(x.GetFirstName()) <= 0 {
		return fmt.Errorf("invalid length for FirstName, must be greater than 0")
	}
	if len(x.GetFirstName()) >= 256 {
		return fmt.Errorf("invalid length for FirstName, must be less than 256")
	}
	if len(x.GetFirstName()) < 1 {
		return fmt.Errorf("invalid length for FirstName, must be at least 1")
	}
	if len(x.GetFirstName()) > 255 {
		return fmt.Errorf("invalid length for FirstName, cannot be more than 255")
	}
	if strings.Contains(x.GetFirstName(), ".") {
		return fmt.Errorf("invalid value for FirstName, must not contain \".\"")
	}
	if strings.Contains(x.GetFirstName(), " ") {
		return fmt.Errorf("invalid value for FirstName, cannot have spaces")
	}
	for _, c := range x.GetFirstName() {
		if c > unicode.MaxASCII {
			return fmt.Errorf("invalid value for FirstName, can only contain ASCII characters")
		}
	}
	if len(x.GetLastName()) <= 0 {
		return fmt.Errorf("invalid length for LastName, must be greater than 0")
	}
	if len(x.GetLastName()) >= 256 {
		return fmt.Errorf("invalid length for LastName, must be less than 256")
	}
	if len(x.GetLastName()) < 1 {
		return fmt.Errorf("invalid length for LastName, must be at least 1")
	}
	if len(x.GetLastName()) > 255 {
		return fmt.Errorf("invalid length for LastName, cannot be more than 255")
	}
	if strings.Contains(x.GetLastName(), ".") {
		return fmt.Errorf("invalid value for LastName, must not contain \".\"")
	}
	if strings.Contains(x.GetLastName(), " ") {
		return fmt.Errorf("invalid value for LastName, cannot have spaces")
	}
	for _, c := range x.GetLastName() {
		if c > unicode.MaxASCII {
			return fmt.Errorf("invalid value for LastName, can only contain ASCII characters")
		}
	}
	if len(x.GetEmail()) <= 0 {
		return fmt.Errorf("invalid length for Email, must be greater than 0")
	}
	if len(x.GetEmail()) >= 256 {
		return fmt.Errorf("invalid length for Email, must be less than 256")
	}
	if len(x.GetEmail()) < 1 {
		return fmt.Errorf("invalid length for Email, must be at least 1")
	}
	if len(x.GetEmail()) > 255 {
		return fmt.Errorf("invalid length for Email, cannot be more than 255")
	}
	if !strings.HasSuffix(x.GetEmail(), "@gmail.com") {
		return fmt.Errorf("invalid value for Email, must have suffix \"@gmail.com\"")
	}
	for _, c := range x.GetEmail() {
		if c > unicode.MaxASCII {
			return fmt.Errorf("invalid value for Email, can only contain ASCII characters")
		}
	}
	if len(x.GetNickname()) < 1 {
		return fmt.Errorf("invalid length for Nickname, must be at least 1")
	}
	if len(x.GetNickname()) > 255 {
		return fmt.Errorf("invalid length for Nickname, cannot be more than 255")
	}
	matchNickname, err := regexp.Match("^.", []byte(x.GetNickname()))
	if err != nil {
		return fmt.Errorf("failed to validate: %w", err)
	}
	if !matchNickname {
		return fmt.Errorf("invalid value for Nickname, did not match \"^.\"")
	}
	notMatchNickname, err := regexp.Match("[[:^alpha:]]", []byte(x.GetNickname()))
	if err != nil {
		return fmt.Errorf("failed to validate: %w", err)
	}
	if notMatchNickname {
		return fmt.Errorf("invalid value for Nickname, can not match \"[[:^alpha:]]\"")
	}
	if len(x.GetTeam()) == 0 {
		return fmt.Errorf("invalid value for Team, cannot be empty")
	}
	if len(x.GetTeam()) < 1 {
		return fmt.Errorf("invalid length for Team, must be at least 1")
	}
	if len(x.GetTeam()) > 255 {
		return fmt.Errorf("invalid length for Team, cannot be more than 255")
	}
	for _, c := range x.GetTeam() {
		if c > unicode.MaxASCII {
			return fmt.Errorf("invalid value for Team, can only contain ASCII characters")
		}
	}
	if x.GetPoints() < 1 {
		return fmt.Errorf("invalid value for Points, must be greater than or equal to 1")
	}
	if x.GetPoints() > 1000 {
		return fmt.Errorf("invalid value for Points, must be less than or equal to 1000")
	}
	if x.GetExtraPoints() < 0 {
		return fmt.Errorf("invalid value for ExtraPoints, must be greater than or equal to 0")
	}
	if x.GetSomething() <= 0.000000 {
		return fmt.Errorf("invalid value for Something, must be greater than 0.000000")
	}
	if x.GetKey() == nil {
		return fmt.Errorf("invalid value for Key, is required")
	}
	if x.GetKey() != nil {
		if len(x.GetKey()) != 2048 {
			return fmt.Errorf("invalid value for Key, must equal 2048")
		}
		if len(x.GetKey()) == 0 {
			return fmt.Errorf("invalid value for Key, must use non-empty value")
		}
	}
	for _, v := range x.GetFriends() {
		if len(v) == 0 {
			return fmt.Errorf("invalid value for Friends, cannot be empty")
		}
	}
	if x.GetSomethingElse() <= 0.000000 {
		return fmt.Errorf("invalid value for SomethingElse, must be greater than 0.000000")
	}
	return nil // is valid
}

// Validate applies configured validation rule options from the protobuf.
func (x *Request2) Validate() error {
	return fmt.Errorf("no validation options configured") // has no validations
}
