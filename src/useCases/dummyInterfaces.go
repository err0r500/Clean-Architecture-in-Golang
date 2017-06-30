package useCases

import "log"

// DummyUserReadWriter dummy struct for tests : implementing the UserReadWriter interface
// a real interface would be in interfaces package
type DummyUserReadWriter struct{}

//GetDetails : return a usecases.User found with by some criterias
func (dRW DummyUserReadWriter) GetDetails(i interface{}) User {
	switch v := i.(type) {
	case int:
		return dRW.getUserByID(v)
	case string:
		return dRW.getUserByName(v)
	default:
		return User{}
	}
}

func (DummyUserReadWriter) getUserByName(n string) User {
	log.Print("dummy userRW received an int : implement searchByID")
	return User{}
}

func (DummyUserReadWriter) getUserByID(id int) User {
	log.Print("dummy userRW received a string : implement searchByName")
	return User{}
}
