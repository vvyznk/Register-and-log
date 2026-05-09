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
func space() {
	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println()
}
func createUsername() string {
	for {
		var username string
		fmt.Println("- Choose an avaliable username")
		fmt.Println()
		fmt.Println("- Rules ----------------------------")
		fmt.Println("| You username need be bigger at 4 characteres;")
		fmt.Println("| You username can't be bigger at 18 characteres;")
		fmt.Println("| You cant use a username already registred;")
		fmt.Println("| Do not use two words; the system will only pick up the first word;")
		fmt.Println("| Dont use special characters.")
		fmt.Println("- -----------------------------------------------------")
		fmt.Print(": ")
		fmt.Scanln(&username)

		username = strings.ToUpper(username)

		if LetterAndNumber(username) == false {
			fmt.Println()
			fmt.Println("-- Do not use special characters.")
			space()
			continue
		}

		if len(username) < 4 {
			fmt.Println()
			fmt.Println("-- Username too short, try another username bigger at 4 characters.")
			space()
			continue
		}

		if len(username) > 18 {
			fmt.Println()
			fmt.Println("-- Username too long, try another username bigger at 4 characters.")
			space()
			continue
		}

		if userExists(username) {
			fmt.Println()
			fmt.Println("-- Username already registred, Please try another username.")
			space()
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
		fmt.Println("| Your password need have 8 characters;")
		fmt.Println("| Do not use two words; the system will only pick up the first word.")
		fmt.Println("- -----------------------------------------------------")
		fmt.Print(": ")
		fmt.Scanln(&password)
		if len(password) < 8 {
			fmt.Println("-- Your password need have 8 characters.")
			space()
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
		fmt.Println("| Only use letters;")
		fmt.Println("| Do not use two words; the system will only pick up the first word.")
		fmt.Println("- -----------------------------------------------------")
		fmt.Print(": ")
		fmt.Scanln(&food)
		fmt.Print()

		if onlyLetters(food) == true {
			return food
		} else {
			fmt.Println("-- Please try again...")
			space()
		}
	}
}
func createNickname() string {
	for {
		var nickname string
		fmt.Println("- Write your childhood nickname.")
		fmt.Println()
		fmt.Println("- Rule ----------------------------")
		fmt.Println("| Only use letters;")
		fmt.Println("| Do not use two words; the system will only pick up the first word.")
		fmt.Println("- -----------------------------------------------------")
		fmt.Print(": ")
		fmt.Scanln(&nickname)

		if onlyLetters(nickname) == true {
			return nickname
		} else {
			fmt.Println("-- Please try again...")
			space()
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
	fmt.Println("- Write your age")
	fmt.Println()
	fmt.Println("- Rule ----------------------------")
	fmt.Println("| ONLY write a number.")
	fmt.Println("- -----------------------------------------------------")
	fmt.Print(": ")
	age := onlyNumber()

	return age
}
func register() {
	var newUser user

	newUser.username = createUsername()
	space()
	newUser.password = createPassword()
	space()
	newUser.food = favoriteFood()
	space()
	newUser.nickname = createNickname()
	space()
	newUser.Age = validAge()

	switch {
	case newUser.Age >= 18:
		users = append(users, newUser)
		fmt.Println("- Registred profile!")
		space()
		return
	default:
		fmt.Println("- Sorry, your age is not allowed in our system.")
		space()
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
	sess.logged = false

	fmt.Println("======= MENU =======")
	fmt.Println("1 - Register")
	fmt.Println("2 - Log")
	fmt.Println("3 - Reset password")
	fmt.Println("4 - Exit")
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
		changePassword()
	case "4":
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
		if !sess.logged {
			return
		}
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
			space()
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
			return
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

func changePassword() {
	var userType string
	if len(users) == 0 {
		fmt.Println("-- No registred users")
		fmt.Println()
		return
	}
	fmt.Println("====== Change Password ======")
	fmt.Println("- Write the username whose password you forgot.")
	fmt.Print(": ")
	fmt.Scanln(&userType)

	userType = strings.ToUpper(userType)

	found := false

	for i, u := range users {
		if u.username == userType {
			found = true
			fmt.Println()
			fmt.Println("- User found.")
			if validFood(i) == true && validNickname(i) == true {
				space()
				fmt.Println("Validation successful.")
				users[i].password = createPassword()
				fmt.Println("Password changed successfully")
				return
			}
			fmt.Println("-- Food or nickname wrong.")
			return
		}
	}
	if !found {
		fmt.Println("-- User does not exist.")
	}
}

func validFood(index int) bool {
	var validFood string

	fmt.Println("- Write this user's favorite food.")
	fmt.Print(":")
	fmt.Scanln(&validFood)

	if validFood == users[index].food {
		return true
	}
	return false
}
func validNickname(index int) bool {
	var validNick string

	fmt.Println("- Write this user's childHood nickname.")
	fmt.Print(":")
	fmt.Scanln(&validNick)

	if validNick == users[index].nickname {
		return true
	}
	return false
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
			sess.logged = false
			return
		case "2":
			fmt.Println("Ok, have fun with the system.")
			system()
		default:
			fmt.Println("Just choose one of these options.")
			continue
		}
	}
}
