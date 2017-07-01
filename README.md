# Example of Clean Architecture in Golang

## Clean Architecture :
Dependencies always go in the opposite direction of Abstraction

So, ***strictly*** from bottom to top (i.e. Infra -> Interfaces -> Usecases -> Domain)
3rd party libs are forbidden above "Interfaces" level.

### 1 - Domain
#### Purpose
**Pure business** : think of the things that will be still relevant if the project was operated in a completely different context (you're building a website backend, what will stay intact if it was replaced by a call center ?)

#### Example
- I've got users
- they have a name
- this name can't be a set of numbers
	- I need a way to update the name and stay assured it always respects this rule

*TODO :*
- declare the User struct with a name property
- about the UpdateUserName() method, you've got 2 options :
	1. it's easy -> define it directly in Domain
	2. you need a lib *[personal remark : Really ? Sounds pretty bad ! ]* -> declare the interface containing the method in Domain, you'll define how it actually works in Interfaces  

#### Tests
- test the consistency of your rules.
  - *for example :* you defined the min & max for something (say, a price) check that the min value you set is not superior to the max
- test the methods defined here :
  - *for example :* if you defined the UpdateUserName() method here, check that it actually does its job.

### 2 - Use Cases
*Use cases* : what happens in a tech agnostic way. For example : "Save" (doesn't care if it's in a db or in RAM or whatever)
- **Declare** here the "tool Interfaces" necessary to operate and the Structs that will "use" them ("Interactors")
- These Structs "Interactors" will respect the Interface ("Interacter") containing all the use cases you want.
- Finally, **define** all the methods an Interactor must have to be an "Interacter".
*ie : do all you're business logic with the "tool Interfaces" (using also Domain stuff) in order "Interactor" respects the "Interacter" interface*

### 3 - Interfaces
Links with the outter world
- **Define** here the "tool Interfaces". That's where you write the tech specific code (use 3rd-party libs), declare new structs in order to map the Domain or Flow structs with another API (gorm, json...) (and in the opposite direction)
- **Declare** the structs (usually empty) that will hold an "Interacter" (whose methods can be called so the logic in usecases is executed)

### 4 - Infra
Technical setup (boilerplate code) that will allow code written at Interfaces level to actually operate
(not implemented in this minimal example, think about CORS setup for your http router and this kind of stuff)

### 5 - Main
See the main_test.go file.
That's where everything is linked together :
1. call if needed code in Infra

2. give the result to the toolsInterfaces you need

3. give these toolsInterfaces to an Interactor struct that will have everything in hand in order to execute the use cases

4. (just so you can figure out everything : there's usually an input layer that will receive request from the outside and use the Interactor in order to trigger its methods )



## Demo :
run main_test.go and read the logs
