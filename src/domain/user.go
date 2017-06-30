package domain

import "log"

// User at domain level, he only has an address and a way to check it
type User struct {
	Address string
	UserChecker
}

type UserChecker interface {
	Check(User, interface{}) bool
}

// DummyChecker respecting the UserChecker interface for test
// I think it's a good idea to keep the mocked interface at the level where they are declared
// They'll be actually implemented at the interfaces layer
type DummyChecker struct{}

// Check : checks if the User Address is valid
func (d DummyChecker) Check(e User, params interface{}) bool {
	log.Print("dummyCheck called !")
	return true
}

// NewUser constructor
func NewUser(a string, c UserChecker) User {
	return User{a, c}
}
