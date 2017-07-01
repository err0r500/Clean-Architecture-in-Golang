package useCases

import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/domain"
)

// User : a more complete user with all app-specific details
// it should contains all the fields that may be sent to a presentationnal layer (http response for example)
// a simple Format() would then be implemented to filter out or alter the properties name/values for the actual response
type User struct {
	domain.User
	ID   int
	name string
}

// NewUseCasesUser : constructor to be sure everything is well initialized
func NewUseCasesUser(dU domain.User, id int, name string) User {
	return User{dU, id, name}
}

// UserInteractor : a struct with all the interfaces needed in order to execute all the useCases
// the job of an Interactor is to gather all the interfaces needed by the use cases
type UserInteractor struct {
	WeakCheck      domain.UserAddressChecker
	StrongCheck    domain.UserAddressChecker
	UserReadWriter UserReadWriter
}

//NewUserInteractor : constructor to be sure everything is well initialized
func NewUserInteractor(wC, sC domain.UserAddressChecker, uRW UserReadWriter) UserInteractor {
	return UserInteractor{wC, sC, uRW}
}

//UserReadWriter interface that will typîcally be defined at "interfaces" level
// NB : a dummy version for test is in interfaces/dummys
type UserReadWriter interface {
	ToUseCaseUser() User
	GetDetails(interface{}) User
}

//
// the useCases !
//

//UseCase1 : a usecase representing a complete "scenario"
func (t UserInteractor) UseCase1(address string) {
	log.Print("--- start of UC 1")

	fU := NewUseCasesUser(
		domain.NewUser(address, t.WeakCheck),
		112345678,
		"userName")
	fU.User.CheckAddress(fU.User)

	log.Print("--- end of UC 1")
}

//UseCase2 : another use case
func (t UserInteractor) UseCase2(params interface{}) {
	log.Print("--- start of UC 2")

	// GetDetails receives params potentially unkown at "useCases" level
	// and returns a usecases.User according to that
	fU := t.UserReadWriter.GetDetails(params)
	fU.User.UserAddressChecker = t.StrongCheck
	fU.User.CheckAddress(fU.User)

	log.Print("--- end of UC 2")
}
