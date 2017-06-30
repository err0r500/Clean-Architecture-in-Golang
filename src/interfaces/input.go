package interfaces

import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/useCases"
)

//InputLayer1 : could be a web router or whatever
type InputLayer1 struct {
	useCases.UserInteractor
}

//CallUseCases : just here to trigger the useCases
func (i InputLayer1) CallUseCases() {
	// usually some logic would happen here to trigger a useCase or another
	// (ie : if it was an http router, it would be a different useCase for each route)
	log.Print("from the inputLayer : I received a raw request and will format it in order to call UseCase 1")
	i.UseCase1("aze")
	log.Print("from the inputLayer : I received the response from UC 1, I can format it to another format (think json for example)")

	//here, the custom type "Age" defined only at the interface level is passed to the useCase layer,
	//this layer will simply carry it without knowing what it is
	log.Print("from the inputLayer : I received another raw request so I will format it in order to call UseCase 2")
	i.UseCase2(Age(12))
	log.Print("from the inputLayer : I received the response from UC 2, I can format it too")
}
