# A Clean Architecture in Golang

## Clean Architecture :
> The purpose of this architecture is to be as flexible as possible in order to develop as quickly as possible and to maintain this speed during the whole developpement lifespan, whatever its size and the changes that will have to be done, for whatever reason.

How ? Dependencies always go in the opposite direction of Abstraction

In other words : Dependencies go ***strictly*** from the bottom up : **Infra** (server Setup and stuff like that) **depends on** **Interfaces** (external libs, Db queries, output formatting) **depends on** **Usecases** (pure logic) **depends on** **Domain** (pure Business & eternal rules))

References to 3rd party libs & frameworks are then forbidden in layers above "Interfaces" : no `import "github.com/..."` things in your useCases or Domain!

### 1 - Domain
#### Purpose
***Pure business & eternal rules*** : think of the things that will be still relevant if the project was operated in a completely different context (you're building a website backend, what will stay intact if it was replaced by a call center ?).

The product owner must able to tell you what to put in here :)

#### Example
- I've got users
- they have a name
- this name can't be a set of numbers
	- I need a way to update the name and stay assured it always respects this rule

*SO... WHAT DO I DO ?*
- declare the User struct with a name property
- about the UpdateUserName() method, you've got 2 options :
	- it's easy -> define it directly in Domain
	- you need a lib *[personal remark : Really ? Sounds pretty bad ! ]* -> declare the interface containing the method in Domain, you'll define how it actually works in Interfaces  

#### Tests
- test the consistency of your rules.
  - *Example :* you defined the min & max for something (say, a price). Check that the min is not superior to the max (remember that the product owner may play with this part of your code :) ).
- test the methods defined here :
  - *Example :* if you defined the UpdateUserName() method here, check that it actually does its job.

### 2 - Use Cases
***Pure Logic*** : This layer knows an outter world exists but absolutely doesn't care about how it looks like. It just does its job without caring about how it gets done.

#### Examples
##### #1
USER_INTERACTOR.*CREATE_AN_ACCOUNT (name, password)* {
- create a new useCase.User
- use ***UserNameUpdater*** to **update(name)** useCase.User.Name with the name received
- if it's fine, set also the useCase.User.Password with the password received
- use ***UserReadWriter*** to **save(user)** useCase.User

}

---
##### #2
USER_INTERACTOR.*DELETE_AN_ACCOUNT (name, password)* {
- use ***UserReadWriterLive*** to **get(name)** the useCase.User it finds with this name
- use ***PasswordChecker*** to check if the password received **isTheSameAs(password)** the password of the useCase.User returned by ***UserReadWriterLive***
- if it's the same, use ***UserReadWriterLive*** to **delete(user)** the User
- use ***UserReadWriterHistory*** to **createDeleted(user)** the User

}

*SO... WHAT DO I DO ?*

***A few naming conventions first :***
- CAPS -> the INTERACTOR
- CAPS, italic with parenthesis -> a *USE_CASE()*
- bold camel-cased -> a ***ToolInterface***
- bold with parenthesis -> a **toolMethod()**

***How it works :***
- a USE_CASE() is a method of an INTERACTOR
- a **toolMethod()** is a method in a  ***ToolInterfaces***
> The INTERACTOR will "hold" all the ***ToolInterfaces*** in order that, when any of its USE_CASE() is called, it will find all the **toolMethods()** it needs.

***So... ?***

What do you know ? The USE_CASE() ! Let's take the second one, it's more complicated ! (in fact I even overcomplicated it just to show how this architecture is flexible) :)

OK, you've got a name & password (the params of the USE_CASE(), we'll see later how you get them) and you would like to delete the user account if the passwords matches. We'll declare the set of methods we'll need at the moment we encounter them and group them in interfaces

- If the accound has to be deleted, it means it relies somewhere. Where ? You don't care ! ... but whatever this place is, you'll have to interact with it.
  - Say, we need **get()**, let's put it in something (an interface) called ***UserReadWriterLive***.
- ... But wait ! There's some chance that in the outter world, the data don't look quite the same as ours :
  - Just to be sure, we'll ask for a **ToUseCaseUser()** method. We won't call it, ever, but it's mandatory for anything wanting to talk with us !
- Let's move on ! We now need to check if passwords match. For the sake of the example, we won't check that directly in the use case. It may not be related to Users only :
  - Let's ask for an **isTheSameAs()** method and put it somewhere else : ***PasswordChecker***
- We'll also have to delete the user :
  - let's ask for a **delete()**, we'll mostly likely have to interact with the same place as where we got the user from, so let's put **delete()** in ***UserReadWriterLive*** too.
- Finally, we group all these ***ToolInterfaces*** in a single structure : the USER_INTERACTOR.
  - We'll wire everything up when we construct it (most likely in the main file) but this way, when any USER_INTERACTOR.*USE_CASE()* is called, it will find an implementation of all the **toolMethods()** it needs !


- ... some time passes and the product owner says it would be cool to keep track of these deleted users but well separated from the "live" ones. No problem, we simply add a line to our use case and
  - ask for a **createDeleted()** method but we'll use another interface, ***UserReadWriterHistory***. This way, we don't care where and how it's actually saved : another table in the same DB, another DB, another type of DB, a file system... we simple don't care, our use case is finished !

#### Tests
- Test the logic of your useCases, try to detect edge cases
  - *Example :* vary the inputs and verify the execution flows through your code the way it's supposed to.
- Test the error detection of the results coming from the Interfaces layer (especially those not throwing any error ! )
  - *Example :* say your **get()** method returns an empty User with no error, would it be fine for you ? You don't know this code, don't trust it !
- Test your inconsistency detection : your use case will receive its params from the Interfaces layer too, do you check that it doesn't make your code act a funny way ?
    - *Example :*  You've got a SEND_ORDER(customerID, cartID) use case, do you check that the customerID retrieved with your (perfectly working) **getCustomerDetails(cartID)** method is the same as the one your use case received as paramater ?

### 3 - Interfaces
Implementation of the links with the outter world
- **Define** here the "tool Interfaces". That's where you write the tech specific code (use 3rd-party libs), declare new structs in order to map the Domain or UseCase structs with another API (gorm, json...) (and in the opposite direction)
- **Declare** the structs that will hold an INTERACTOR (whose methods can be called so the logic in usecases is executed)

### 4 - Infra
Technical setup (boilerplate code) that will allow code written at Interfaces level to actually operate
(not implemented in this minimal example, think about CORS setup for your http router and this kind of stuff)

### 5 - Main
See the main_test.go file.
That's where everything is linked together :
- call if needed code in Infra (not shown here)

- give the result to the toolsInterfaces you need

- give these toolsInterfaces to an Interactor struct that will have everything in hand in order to execute the use cases


## Demo :
run main_test.go and read the logs
