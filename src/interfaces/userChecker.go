package interfaces

import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/domain"
)

// RealChecker empty struct implementing the domain.UserChecker interface
type RealChecker struct{}

// Check would do a real check
func (d RealChecker) Check(e domain.User, params interface{}) bool {
	log.Print("RealChecker Called ")
	return true
}
