package main

import (
	"fmt"
	"os"
	"strings"
)

type user struct {
	username string
	password string
	food     string
	nickname string
	Age      int
}
type login struct {
	username string
	password string
}
type session struct {
	currentIndex int
	logged       bool
}

var choice string
var users []user
var sess session

func main() {
	for {
		menu()
	}
}
func userExists(username string) bool {
	for _, u := range users {
		if u.username == username {
			return true
		}
	}
	return false
}
func createPassword() string {
	for {
		var password string
		fmt.Println("- Choose your password")
		fmt.Print(": ")
		fmt.Scanln(&password)
		if len(password) < 8 {
			fmt.Println("-- Your password need have 8 caracteres.")
			continue
		}
		return password
	}
}

func createUsername() string {
	for {
		var username string
		fmt.Println("- Choose an avaliable username")
		fmt.Print(": ")
		fmt.Scanln(&username)

		username = strings.ToUpper(username)

		if len(username) < 4 {
			fmt.Println()
			fmt.Println("-- Username too short, try another username bigger at 4 caracteres.")
			continue
		}

		if userExists(username) {
			fmt.Println()
			fmt.Println("-- Username already registred, Please try another username")
			continue
		}
		return username
	}
}
func onlyNumber() int {
	var number int
	for {
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println()
			fmt.Println("-- Just type numbers, please try again.")
			fmt.Print(": ")

		} else {
			return number
		}
	}
}
func validAge() int {
	for {
		fmt.Println("- Write your age")
		fmt.Print(": ")
		age := onlyNumber()

		return age
	}
}
func register() {
	var newUser user

	newUser.username = createUsername()
	newUser.password = createPassword()

	fmt.Println("- Write your favorite food")
	fmt.Print(": ")
	fmt.Scanln(&newUser.food)
	fmt.Println("- Write your childhood nickname")
	fmt.Print(": ")
	fmt.Scanln(&newUser.nickname)

	newUser.Age = validAge()

	switch {
	case newUser.Age >= 18:
		users = append(users, newUser)
		fmt.Println("- Registred profile!")
		fmt.Println()
		return
	default:
		fmt.Println("- Sorry, your age is not allowed in our system.")
		fmt.Println()
		return
	}
}

func loginUser() {
	var validation login
	if len(users) == 0 {
		fmt.Println("- No registred users")
		fmt.Println()
		return
	}

	fmt.Println("- Choose an user registered in the system")
	fmt.Print(": ")
	fmt.Scanln(&validation.username)
	fmt.Println("- Enter your password")
	fmt.Print(": ")
	fmt.Scanln(&validation.password)
	validation.username = strings.ToUpper(validation.username)

	for i, u := range users {
		if u.username == validation.username && u.password == validation.password {
			sess.currentIndex = i
			sess.logged = true
			fmt.Println()
			fmt.Println("- Welcome,", validation.username, "!")
			fmt.Println()
			system()
			return
		}
	}
	fmt.Println()
	fmt.Println("-- Wrong user or password.")
	fmt.Println()
}

func menu() {
	var choice string
	fmt.Println("======= MENU =======")
	fmt.Println("1 - Register")
	fmt.Println("2 - Log")
	fmt.Println("3 - Exit")
	fmt.Println("=====================")
	fmt.Print("Choose: ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		fmt.Println("===================")
		register()
	case "2":
		fmt.Println("===================")
		loginUser()
	case "3":
		fmt.Println("===================")
		fmt.Println("- Leaving the system...")
		os.Exit(0)
	default:
		fmt.Println()
		fmt.Println("-- That was not a valid option. ")
		fmt.Println()
		return
	}
}

func system() {
	for {
		fmt.Println("======== SYSTEM ========")
		fmt.Println("1 - Information about this user")
		fmt.Println("2 - Back to menu")
		fmt.Println("3 - Exit")
		fmt.Println("4 - Delete this user")
		fmt.Println("========================")
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			info()
		case "2":
			return
		case "3":
			fmt.Println()
			fmt.Println("- Leaving the system...")
			os.Exit(0)
		case "4":
			deleteUser()
		default:
			fmt.Println()
			fmt.Println("- Just choose one of these options. ")
		}
	}
}

func info() {

	fmt.Println()
	fmt.Println("====== information ======")
	fmt.Print("Current user: ")
	fmt.Println(users[sess.currentIndex].username)
	fmt.Print("Favority food: ")
	fmt.Println(users[sess.currentIndex].food)
	fmt.Print("Childhood nickname: ")
	fmt.Println(users[sess.currentIndex].nickname)
	fmt.Print("Age: ")
	fmt.Println(users[sess.currentIndex].Age)
	fmt.Println("=========================")
	fmt.Println("- Do you want change any information?")
	fmt.Println()
	fmt.Println("1 - Yes")
	fmt.Println("2 - Not, bring me back to the system")
	fmt.Println("=========================")
	fmt.Print("Choose: ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		changeInfo()
	case "2":
		system()
	default:
		fmt.Println()
		fmt.Println("Just choose one of these options.")
	}

}

func changeInfo() {
	fmt.Println()
	fmt.Println("====== Changer Information ======")
	fmt.Println("Which one of these information do you wanna change? ")
	fmt.Println("1 - Username ")
	fmt.Println("2 - Password ")
	fmt.Println("3 - Favority food ")
	fmt.Println("4 - Childhood nickname")
	fmt.Println("5 - Exit")
	fmt.Println("=================================")
	fmt.Print("Choose: ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		changeUser()
	case "2":
		changePassword()
	case "3":
		changeFood()
	case "4":
		changeNickname()
	case "5":
		return
	default:
		fmt.Println("Just choose one of these options. ")
	}
}

func changeUser() {
	fmt.Println("====== Changer User ======")
	users[sess.currentIndex].username = createUsername()
	fmt.Println("User changed successfully")
	return
}

func changePassword() {
	fmt.Println("====== Changer User ======")
	users[sess.currentIndex].password = createPassword()
	fmt.Println("Password changed successfully")
	return
}

func changeFood() {

}

func changeNickname() {

}

func deleteUser() {

}

// monstra regras ao criar usuario "sem espaço" e etc. e ao criar senha tambem, e coloca uma validação e regra ao coloca nome de infacia e comida facvorita, ou seja, sem espaços
