package interfaces

import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/useCases"
)

// RealUserReadWriter implementing the UserReadWriter interface
type RealUserReadWriter struct{}

// RWUser : a tech/framework specific version of User
type RWUser struct {
	UsecaseUser useCases.User
	techStuff1  int
	techStuff2  Age
}

// ToUseCaseUser helper in order to return a useCases.User to useCases level
func (rwU RWUser) ToUseCaseUser() useCases.User {
	log.Print("from RealUserRW : I use a special type of User but I convert it back to ucUser")
	return rwU.UsecaseUser
}

//GetDetails : allows to search by something the
func (dRW RealUserReadWriter) GetDetails(i interface{}) useCases.User {
	rwU := RWUser{}
	switch v := i.(type) {
	case int:
		rwU = dRW.getUserByID(v)
	case string:
		rwU = dRW.getUserByName(v)
	case Age:
		rwU = dRW.getUserByAge(v)
	default:
	}
	return rwU.ToUseCaseUser()
}

func (RealUserReadWriter) getUserByName(n string) RWUser {
	log.Print("realUserRW received a string : implement searchByName")
	return RWUser{}
}

func (RealUserReadWriter) getUserByID(id int) RWUser {
	log.Print("realUserRW received an int : implement searchByID")
	return RWUser{}
}

func (RealUserReadWriter) getUserByAge(id Age) RWUser {
	log.Print("realUserRW received an Age ! -> a UseCase triggered a search by Age while it doesn't know this type !")
	return RWUser{}
}
