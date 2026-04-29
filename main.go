package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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

func createUsername() string {
	for {
		var username string
		fmt.Println("- Choose an avaliable username")
		fmt.Println()
		fmt.Println("- Rules ----------------------------")
		fmt.Println("| You username need be bigger at 4 characteres.")
		fmt.Println("| You cant use a username already registred.")
		fmt.Println("| Do not use two words; the system will only pick up the first word.")
		fmt.Println("| Dont use special characters.")
		fmt.Println("- -----------------------------------------------------")
		fmt.Print(": ")
		fmt.Scanln(&username)

		username = strings.ToUpper(username)

		if LetterAndNumber(username) == false {
			fmt.Println()
			fmt.Println("-- Do not use special characters.")
			continue
		}

		if len(username) < 4 {
			fmt.Println()
			fmt.Println("-- Username too short, try another username bigger at 4 characters.")
			continue
		}

		if userExists(username) {
			fmt.Println()
			fmt.Println("-- Username already registred, Please try another username.")
			continue
		}
		return username
	}
}

func createPassword() string {
	for {
		var password string
		fmt.Println("- Choose your password")
		fmt.Println()
		fmt.Println("- Rules ----------------------------")
		fmt.Println("| Your password need have 8 characters.")
		fmt.Println("| Do not use two words; the system will only pick up the first word.")
		fmt.Println("- -----------------------------------------------------")
		fmt.Print(": ")
		fmt.Scanln(&password)
		if len(password) < 8 {
			fmt.Println("-- Your password need have 8 characters.")
			continue
		}
		return password
	}
}
func favoriteFood() string {
	for {
		var food string
		fmt.Println("- Write your favorite food.")
		fmt.Println()
		fmt.Println("- Rule ----------------------------")
		fmt.Println("| Do not use two words; the system will only pick up the first word.")
		fmt.Println("- -----------------------------------------------------")
		fmt.Print(": ")
		fmt.Scanln(&food)

		if onlyLetters(food) == true {
			return food
		} else {
			fmt.Println("-- Please try again...")
			fmt.Println()
		}
	}
}
func createNickname() string {
	for {
		var nickname string
		fmt.Println("- Write your childhood nickname.")
		fmt.Println()
		fmt.Println("- Rule ----------------------------")
		fmt.Println("| Do not use two words; the system will only pick up the first word.")
		fmt.Println("- -----------------------------------------------------")
		fmt.Print(": ")
		fmt.Scanln(&nickname)

		if onlyLetters(nickname) == true {
			return nickname
		} else {
			fmt.Println("-- Please try again...")
			fmt.Println()
		}
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

func onlyLetters(text string) bool {
	if text == "" {
		return false
	}
	for _, r := range text {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
func LetterAndNumber(text string) bool {
	if text == "" {
		return false
	}
	for _, r := range text {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
func validAge() int {
	for {
		fmt.Println("- Write your age")
		fmt.Println()
		fmt.Println("- Rule ----------------------------")
		fmt.Println("| ONLY write a number.")
		fmt.Println("- -----------------------------------------------------")
		fmt.Print(": ")
		age := onlyNumber()

		return age
	}
}
func register() {
	var newUser user

	newUser.username = createUsername()
	newUser.password = createPassword()
	newUser.food = favoriteFood()
	newUser.nickname = createNickname()

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
	choice := ""
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
		choice := ""
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
	for {
		choice := ""
		fmt.Println()
		fmt.Println("====== information ======")
		fmt.Print("Current user: ")
		fmt.Println(users[sess.currentIndex].username)
		fmt.Println("Password: ****")
		fmt.Print("Favority food: ")
		fmt.Println(users[sess.currentIndex].food)
		fmt.Print("Childhood nickname: ")
		fmt.Println(users[sess.currentIndex].nickname)
		fmt.Print("Age: ")
		fmt.Println(users[sess.currentIndex].Age)
		fmt.Println("=========================")
		fmt.Println("- Do you want change any information?")
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

}

func changeInfo() {
	choice := ""
	for {
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
			fmt.Println("Just choose one of these options.")
		}
	}

}

func changeUser() {
	fmt.Println("====== Changer User ======")
	fmt.Println("- Your current user is", users[sess.currentIndex].username, "!")
	fmt.Println()

	users[sess.currentIndex].username = createUsername()
	fmt.Println("User changed successfully")
	return
}

func changePassword() {
	fmt.Println("====== Changer password ======")
	fmt.Println("- We can't acess your account password :/")
	fmt.Println()

	users[sess.currentIndex].password = createPassword()
	fmt.Println("Password changed successfully")
	return
}

func changeFood() {
	fmt.Println("====== Changer Favorite Food ======")
	fmt.Println("- Your current favorite food is", users[sess.currentIndex].food, "!")
	fmt.Println()

	users[sess.currentIndex].food = favoriteFood()
	fmt.Println("Favorite food changed successfully")
	return
}

func changeNickname() {
	fmt.Println("====== Changer Childhood Nickname ======")
	fmt.Println("- Your current childhood nickname is", users[sess.currentIndex].nickname, "!")
	fmt.Println()

	users[sess.currentIndex].nickname = createNickname()
	fmt.Println("Childhood Nickname changed successfully")
	return
}

func deleteUser() {
	for {
		choice := ""

		fmt.Println("====== Delete USER ======")
		fmt.Println("- Are you sure about that?")
		fmt.Println("1 - Yes, delete the user:", users[sess.currentIndex].username)
		fmt.Println("2 - No.")
		fmt.Println("=========================")
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Println("Deleting user...")
			users = append(users[:sess.currentIndex], users[sess.currentIndex+1:]...)
			fmt.Println("User deleted. Return to main menu")
			fmt.Println(users)
			main()
		case "2":
			fmt.Println("Ok, have fun with the system.")
			system()
		default:
			fmt.Println("Just choose one of these options.")
			continue
		}
	}
}
