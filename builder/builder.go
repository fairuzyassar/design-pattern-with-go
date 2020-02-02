package main

import "fmt"

type user struct {
	firstName string
	lastName string
	email string
}

type userBuilder struct {
	firstName string
	lastName string
	email string
}

func(ub *userBuilder) setFirstName(firstName string) *userBuilder {
	ub.firstName = firstName
	return ub
}

func(ub *userBuilder) setLastName(lastName string) *userBuilder {
	ub.lastName = lastName
	return ub
}

func(ub *userBuilder) setEmail(email string) *userBuilder {
	ub.email = email
	return ub
}

func(ub *userBuilder) build() *user {
	return &user{
		firstName: ub.firstName,
		lastName: ub.lastName,
		email: ub.email,
	}
}

func main(){
	builder := &userBuilder{}
	user1 := builder.setFirstName("andi").setLastName("pavlo")
	user2 := builder.setFirstName("pavlo").setEmail("pavlo@cmu.edu") 

	fmt.Println(user1.firstName)
	fmt.Println(user2.email)

}
