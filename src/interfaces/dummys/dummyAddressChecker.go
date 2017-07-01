package dummys

// DummyChecker respecting the UserAddressChecker interface for test
import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/domain"
)

// DummyChecker is just here to hold a "version" of the UserAddressChecker interface
// declare in the domain layer
type DummyChecker struct{}

// CheckAddress : checks if the User Address is valid
func (d DummyChecker) CheckAddress(e domain.User) bool {
	log.Print("dummyAddressCheck of : " + e.Address)
	return true
}
