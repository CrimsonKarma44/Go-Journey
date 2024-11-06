package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"to-do/database"
	"to-do/models"
	"to-do/utils"
)

func main() {
	db, err := database.GormPostgresqlInit()
	if err != nil {
		panic(err)
	}
	//err = utils.InsertingUsers(db, &models.User{Email: "vincentprincewill44@yahoo.com", UserName: "Karma", Password: "password"})
	users := utils.RetrievingAllUser(db)
	//fmt.Println(users)
	for _, user := range users {
		fmt.Println(user.Email)
	}
	//if err != nil {
	//	panic(err)
	//}
	//result := db.Preload("Lists").Find(&dummy)
	//err = utils.DeletingUser(db, "vincentprincewill44@gmail.com")
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("Deleted successfully")
	//}
	//fmt.Println(utils.RetrievingAllUser(db))
	//var user_ []models.User
	//for _, user := range users {
	//	fmt.Println(user.UserID)
	//	db.Find(&user_, "id = ?", user.UserID)
	//	for _, u := range user_ {
	//		fmt.Println(u.UserName)
	//	}

	db, err = database.GormInit()
	if err != nil {
		panic(err)
	}
	err = database.Migration(db)
	if err != nil {
		panic(err)
	}

	var username, password, email, condition string

	for {
		fmt.Println("Are you a new user?(y/n):")
		fmt.Scanln(&condition)
		if condition == "y" {
			username = createUsername(username, db)
			email = createEmail(email, db)
			password = createPassword(password)

			createUser(username, password, email, db)
			if stop := start(db); stop == "break" {
				break
			}

		} else if condition == "n" {
			if stop := start(db); stop == "break" {
				break
			}
		} else {
			fmt.Println("Invalid input")
		}
	}

	//	db.Find(&user_, "id = ?", user.UserID)
	//result := db.Preload("Lists").Find(&dummy)
}

func start(db *gorm.DB) string {
	var condition, username, password string
	fmt.Println("Do you want to log in?(y/n):")
	fmt.Scanln(&condition)
	if condition == "y" {
		fmt.Println("Username:")
		fmt.Scanln(&username)
		fmt.Println("Password:")
		fmt.Scanln(&password)
		err := login(username, password, db)
		if err != nil {
			fmt.Println(err)
		}
		return start(db)
	} else {
		fmt.Println("Bye!..")
		return "break"
	}
}

func createUsername(username string, db *gorm.DB) string {
	for {
		fmt.Println("Creating user...:")
		fmt.Println("Username:")
		fmt.Scanln(&username)
		result := db.First(&models.User{}, "user_name = ?", username)
		if result.RowsAffected == 1 {
			fmt.Println("User already exists.")
		} else {
			break
		}
	}
	return username
}

func createEmail(email string, db *gorm.DB) string {
	for {
		fmt.Println("Email:")
		fmt.Scanln(&email)
		err := utils.EmailValidator(email)
		if err != nil {
			fmt.Println("Invalid email address")
		}
		result := db.First(&models.User{}, "email = ?", email)
		if result.RowsAffected == 1 || !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("Email already exists.")
		} else {
			break
		}
	}
	return email
}
func createPassword(password string) string {
	for {
		fmt.Println("Password:")
		fmt.Scanln(&password)
		err := utils.PasswordValidator(password)
		if err != nil {
			fmt.Println("Invalid password")
		} else {
			password, err = utils.HashPassword(password)
			if err != nil {
				fmt.Println("Error hashing password")
			} else {
				break
			}
		}
	}
	return password
}

func createUser(userName string, password string, email string, db *gorm.DB) {
	newUser := models.User{UserName: userName, Password: password, Email: email}
	err := utils.InsertingUsers(db, newUser)
	if err != nil {
		log.Println("error creating user")
		panic(err)
	}
}
func login(userName string, password string, db *gorm.DB) error {
	var user models.User
	result := db.First(&user, "user_name = ?", userName)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if err := utils.ComparePasswords(user.Password, password); err == nil {
			fmt.Println("Logged in")
			return nil
		} else {
			fmt.Println("Invalid password...")
		}
	} else {
		fmt.Println("User not found")
	}
	return result.Error
}
