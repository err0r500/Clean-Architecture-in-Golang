package main

import (
	"log"
	"testing"

	"github.com/err0r500/cleanArchitectureGolang/src/domain"
	"github.com/err0r500/cleanArchitectureGolang/src/interfaces"
	"github.com/err0r500/cleanArchitectureGolang/src/useCases"
)

func TestFlow(t *testing.T) {
	log.Print("first implementation, call from Main.")
	log.Print("UserInteractor Setup : weakCheck is DummyChecker, strongCheck is DummyChecker, UserRW is RealUserRW")
	log.Print("===========")
	tI := useCases.NewUserInteractor(domain.DummyChecker{}, domain.DummyChecker{}, interfaces.RealUserReadWriter{})
	tI.UseCase1(nil)
	tI.UseCase2("userName")
	log.Print("===========\n\n\n")

	log.Print("Second implementation, call from Main.")
	log.Print("UserInteractor Setup : weakCheck is DummyChecker, strongCheck is RealChecker, UserRW is dummyUserRW")
	log.Print("===========")
	tI2 := useCases.NewUserInteractor(domain.DummyChecker{}, interfaces.RealChecker{}, useCases.DummyUserReadWriter{})
	tI2.UseCase1("hoho")
	tI2.UseCase2(12345)
	log.Print("===========\n\n\n")

	// use an input interface (could be a web router, a CLI etc.)
	// pass it the UserInteractor setup so it can call it
	log.Print("Third implementation : from main then to input 1.")
	log.Print("UserInteractor Setup : weakCheck is DummyChecker, strongCheck is DummyChecker, UserRW is RealUserRW")
	log.Print("===========")
	tIinput1 := interfaces.InputLayer1{}
	tIinput1.UserInteractor = useCases.NewUserInteractor(domain.DummyChecker{}, interfaces.RealChecker{}, interfaces.RealUserReadWriter{})
	tIinput1.CallUseCases()
	log.Print("===========\n\n\n")
}
