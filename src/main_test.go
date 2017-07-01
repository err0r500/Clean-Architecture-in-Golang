package main

import (
	"log"
	"testing"

	"github.com/err0r500/cleanArchitectureGolang/src/interfaces"
	"github.com/err0r500/cleanArchitectureGolang/src/interfaces/dummys"
	"github.com/err0r500/cleanArchitectureGolang/src/useCases"
)

func TestUsesCases(t *testing.T) {
	log.Print("first implementation, called from Main.")
	log.Print("UserInteractor Setup -> weakCheck : DummyChecker, strongCheck : RealChecker, UserRW : DummyUserRW")
	tI := useCases.NewUserInteractor(dummys.DummyChecker{}, interfaces.RealChecker{}, dummys.DummyUserReadWriter{})
	tI.UseCase1("1st Try Address")
	tI.UseCase2(12345)
	log.Print("===========\n\n\n")

	log.Print("Second implementation, called from Main.")
	log.Print("UserInteractor Setup -> weakCheck : DummyChecker, strongCheck : RealChecker, UserRW : realUserRW")
	tI2 := useCases.NewUserInteractor(dummys.DummyChecker{}, interfaces.RealChecker{}, interfaces.RealUserReadWriter{})
	tI2.UseCase1("2nd try Address")
	tI2.UseCase2("userName")
	log.Print("===========\n\n\n")

	// use an input interface (could be a web router, a CLI etc.)
	// pass it the UserInteractor setup so it can call it
	log.Print("Third implementation : from main then to input 1.")
	log.Print("UserInteractor Setup -> weakCheck : DummyChecker, strongCheck : RealChecker, UserRW : RealUserRW")
	tIinput1 := interfaces.InputLayer1{}
	tIinput1.UserInteractor = useCases.NewUserInteractor(dummys.DummyChecker{}, interfaces.RealChecker{}, interfaces.RealUserReadWriter{})
	tIinput1.CallUseCases()
	log.Print("===========")
}
