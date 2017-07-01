package dummys

// DummyChecker respecting the UserAddressChecker interface for test
import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/domain"
)

// I think it's a good idea to keep the mocked interface at the level where they are declared
// They'll be actually implemented at the interfaces layer
type DummyChecker struct{}

// Check : checks if the User Address is valid
func (d DummyChecker) CheckAddress(e domain.User) bool {
	log.Print("dummyAddressCheck of : " + e.Address)
	return true
}
