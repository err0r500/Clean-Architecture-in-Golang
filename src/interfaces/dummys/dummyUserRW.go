package dummys

import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/useCases"
)

// DummyUserReadWriter dummy struct for tests : implementing the UserReadWriter interface
// a real interface would be in interfaces package
type DummyUserReadWriter struct {
	ucUser useCases.User
}

// DummyUser : an enhanced User specific to this interface
type DummyUser struct {
	UsecaseUser useCases.User
	otherStuff  int
}

//ToUseCaseUser : just to be sure, the interfaceLayer will be able to return this type
func (dURW DummyUserReadWriter) ToUseCaseUser() useCases.User {
	return dURW.ucUser
}

//GetDetails : return a usecases.User found with by some criterias
func (dURW DummyUserReadWriter) GetDetails(i interface{}) useCases.User {
	dU := DummyUser{}
	switch v := i.(type) {
	case int:
		dU = dURW.getUserByID(v)
	case string:
		dU = dURW.getUserByName(v)
	default:
		dU = DummyUser{}
	}
	dURW.ucUser = dU.UsecaseUser
	return dURW.ToUseCaseUser()
}

func (dURW DummyUserReadWriter) getUserByName(n string) DummyUser {
	log.Print("dummy userRW received an int : implement searchByName")
	dU := DummyUser{}
	dU.UsecaseUser.Address = "dummyAddressFoundbyName"
	return dU
}

func (dURW DummyUserReadWriter) getUserByID(id int) DummyUser {
	log.Print("dummy userRW received a string : implement searchByID")
	dU := DummyUser{}
	dU.UsecaseUser.Address = "dummyAddressFoundbyID"
	return dU
}
