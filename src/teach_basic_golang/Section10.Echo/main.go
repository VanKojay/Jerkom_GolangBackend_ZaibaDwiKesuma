package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    ` json:"id" form:"id" `
	Name     string ` json:"name" form:"name" `
	Email    string ` json:"email" form:"email" `
	Password string ` json:"password" form:"password" `
}

var users []User

// -------------------- controller --------------------
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// Dapatkan ID dari parameter URL
	id := c.Param("id")

	// Ubah ID menjadi tipe data int
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	// Cari user berdasarkan ID
	var foundUser *User
	for i, user := range users {
		if user.Id == userID {
			foundUser = &users[i]
			break
		}
	}

	if foundUser == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by ID",
		"user":     foundUser,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// Dapatkan ID dari parameter URL
	id := c.Param("id")

	// Ubah ID menjadi tipe data int
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	// Cari index user berdasarkan ID
	index := -1
	for i, user := range users {
		if user.Id == userID {
			index = i
			break
		}
	}

	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}

	// Hapus user dari slice users
	users = append(users[:index], users[index+1:]...)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// Dapatkan ID dari parameter URL
	id := c.Param("id")

	// Ubah ID menjadi tipe data int
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	// Cari user berdasarkan ID
	var foundUser *User
	for i, user := range users {
		if user.Id == userID {
			foundUser = &users[i]
			break
		}
	}

	if foundUser == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}

	// binding data baru
	updatedUser := new(User)
	c.Bind(updatedUser)

	// Update data user
	foundUser.Name = updatedUser.Name
	foundUser.Email = updatedUser.Email
	foundUser.Password = updatedUser.Password

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
		"user":     foundUser,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
//
//	func main() {
//		e := echo.New()
//		// routing with query parameter
//		e.GET("/users", GetUsersController)
//		e.POST("/users", CreateUserController)
//		// start the server, and log if it fails
//		e.Logger.Fatal(e.Start(" :8000 "))
//	}
func main() {
	e := echo.New()

	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)

	// Start the server
	err := e.Start("localhost:8000")
	if err != nil {
		// Handle error jika server tidak dapat dimulai
		fmt.Println("Error starting server:", err)
	}
}
