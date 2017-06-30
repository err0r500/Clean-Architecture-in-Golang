package interfaces

import "github.com/err0r500/cleanArchitectureGolang/src/domain"

//User a common User struct
type User struct {
	domain.User
	name string
}

//Age is a custom struct, flows can call the getByDetails
type Age int
