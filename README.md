# Example of Clean Architecture in Golang

## Clean Architecture :
Dependencies go ***strictly*** from bottom to top (i.e. infra -> interfaces -> flows -> domain)
3rd party libs are forbidden above "Interfaces" level.

### 1 - Domain
*Pure business* : things still relevant if the project was operated by phone

### 2 - Flows
*Use cases* : what happens in a tech agnostic way. For example : "Save" (doesn't care if it's in a db or in RAM or whatever)
- **Declare** here the "tool Interfaces" necessary to operate and the Structs that will "use" them ("Interactors")
- These Structs "Interactors" will respect the Interface ("Interacter") containing all the use cases you want.
- Finally, **define** all the methods an Interactor must have to be an "Interacter".
*ie : do all you're business logic with the "tool Interfaces" (using also Domain stuff) in order "Interactor" respects the "Interacter" interface*

### 3 - Interfaces
Links with the outter world
- **Define** here the "tool Interfaces". That's where you write the tech specific code (use 3rd-party libs), declare new structs in order to map the Domain or Flow structs with another API (gorm, json...) (and in the opposite direction)
- **Declare** the structs (usually empty) that will hold an "Interacter" (whose methods can be called so the logic in flows is executed)

### 4 - Infra
Technical setup (boilerplate code) that will allow code written at Interfaces level to actually operate

### 5 - Mains
See the /mains/\*.go files.
Link everything together :
1. call if needed code in Infra

2. give the result to the toolsInterfaces you need

3. give these toolsInterfaces to an Interactor structs so they respect the interacterInterface

## Demo :
run main_test.go and read the logs
