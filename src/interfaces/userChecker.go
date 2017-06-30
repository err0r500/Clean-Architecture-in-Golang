package interfaces

import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/domain"
)

// RealChecker empty struct implementing the domain.UserAddressChecker interface
type RealChecker struct{}

// CheckAddress checks the Address of a user
func (d RealChecker) CheckAddress(e domain.User, params interface{}) bool {
	log.Print("RealChecker Called ")
	return true
}
