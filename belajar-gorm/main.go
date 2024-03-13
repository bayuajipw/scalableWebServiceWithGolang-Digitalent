package main

import (
	"belajar-gorm/database"
	"belajar-gorm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func createUser(email string) {
	db := database.GetDB()

	user := models.User{
		Email: email,
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		fmt.Println("Failed to create new user, err:", err)
		return
	}

	fmt.Println("Success to create new user:", user)
}

func getUserByID(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.Debug().First(&user, "id=?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}

		fmt.Println("Failed to find user data, err:", err)
		return
	}

	fmt.Printf("User data: %+v\n", user)
}

func updateUserByID(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id=?", id).Updates(models.User{Email: email}).Error
	if err != nil {
		fmt.Println("Failed to update user data, err:", err)
		return
	}
	fmt.Printf("Success to update user data: %+v\n", user)
}

func createProduct(userID uint, name, brand string) {
	db := database.GetDB()

	product := models.Product{
		UserID: userID,
		Name:   name,
		Brand:  brand,
	}

	err := db.Create(&product).Error
	if err != nil {
		fmt.Println("Failed to create product, err:", err)
		return
	}

	fmt.Printf("Product added: %+v\n", product)
}

func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}

	err := db.Preload("Products").Find(&users).Error
	if err != nil {
		fmt.Println("Failed to find all users with products, err:", err)
		return
	}

	fmt.Printf("Users with products: %+v", users)
}

func deleteProductByID(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Where("id=?", id).Delete(&product).Error
	if err != nil {
		fmt.Println("Failed to delete product, err:", err)
		return
	}

	fmt.Printf("Product with id %d has been deleted", id)
}

func main() {
	database.StartDB()
	// createUser("ilham@god.com")
	getUserByID(1)
	// updateUserByID(1, "ilham@theking.com")
	// createProduct(1, "Redmi Note", "Xiaomi")
	// getUsersWithProducts()
	deleteProductByID(1)
}
