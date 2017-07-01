package interfaces

import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/useCases"
)

// RealUserReadWriter implementing the UserReadWriter interface
type RealUserReadWriter struct {
	ucUser useCases.User
}

// RWUser : a tech/framework specific version of User
type RWUser struct {
	UsecaseUser useCases.User
	techStuff1  int
	techStuff2  Age
}

// ToUseCaseUser helper in order to return a useCases.User to useCases level
func (rURW RealUserReadWriter) ToUseCaseUser() useCases.User {
	log.Print("from RealUserRW : I use a special type of User but I can convert it back to ucUser")
	return rURW.ucUser
}

//GetDetails : allows to search by something the
func (rURW RealUserReadWriter) GetDetails(i interface{}) useCases.User {
	rwU := RWUser{}
	switch v := i.(type) {
	case int:
		rwU = rURW.getUserByID(v)
	case string:
		rwU = rURW.getUserByName(v)
	case Age:
		rwU = rURW.getUserByAge(v)
	default:
	}
	rURW.ucUser = rwU.UsecaseUser
	return rURW.ToUseCaseUser()
}

func (RealUserReadWriter) getUserByName(n string) RWUser {
	log.Print("realUserRW received a string : implement searchByName")
	rWu := RWUser{}
	rWu.UsecaseUser.Address = "address fetched somewhere"
	return rWu
}

func (RealUserReadWriter) getUserByID(id int) RWUser {
	log.Print("realUserRW received an int : implement searchByID")
	rWu := RWUser{}
	rWu.UsecaseUser.Address = "address fetched somewhere else"
	return rWu
}

func (RealUserReadWriter) getUserByAge(id Age) RWUser {
	log.Print("realUserRW received an Age ! -> a UseCase triggered a search by Age while it doesn't know this type !")
	rWu := RWUser{}
	rWu.UsecaseUser.Address = "address fetched somewhere else (again)"
	return rWu
}
