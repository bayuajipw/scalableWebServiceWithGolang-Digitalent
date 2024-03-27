package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

var db *gorm.DB
var err error

type User struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	Age       uint   `gorm:"not null" json:"age"`
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Photo struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	Title     string `gorm:"size:255;not null;" json:"title"`
	Caption   string `gorm:"size:255" json:"caption"`
	PhotoURL  string `gorm:"size:255;not null;" json:"photo_url"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Comment struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	PhotoID   uint   `gorm:"not null" json:"photo_id"`
	Message   string `gorm:"size:255;not null;" json:"message"`
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type SocialMedia struct {
	ID             uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name           string `gorm:"size:255;not null;" json:"name"`
	SocialMediaURL string `gorm:"size:255;not null;" json:"social_media_url"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	CreatedAt      string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Initialize SQLite database
	db, err = gorm.Open("sqlite3", "mygram.db")
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &Photo{}, &Comment{}, &SocialMedia{})

	// Routes
	router.POST("/register", registerUser)
	router.POST("/login", loginUser)
	router.POST("/photos", createPhoto)
	router.GET("/photos", getAllPhotos)
	router.POST("/comments", createComment)
	router.GET("/comments/:photo_id", getCommentsByPhotoID)

	// Run server
	router.Run(":8080")
}

func registerUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Create user
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func loginUser(c *gin.Context) {
	var user User
	var foundUser User
	c.BindJSON(&user)

	// Find user by email
	db.Where("email = ?", user.Email).First(&foundUser)
	if foundUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": foundUser.Email,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func createPhoto(c *gin.Context) {
	var photo Photo
	c.BindJSON(&photo)

	// Create photo
	db.Create(&photo)
	c.JSON(http.StatusOK, gin.H{"message": "Photo created successfully"})
}

func getAllPhotos(c *gin.Context) {
	var photos []Photo
	db.Find(&photos)
	c.JSON(http.StatusOK, photos)
}

func createComment(c *gin.Context) {
	var comment Comment
	c.BindJSON(&comment)

	// Create comment
	db.Create(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}

func getCommentsByPhotoID(c *gin.Context) {
	photoID := c.Param("photo_id")
	var comments []Comment
	db.Where("photo_id = ?", photoID).Find(&comments)
	c.JSON(http.StatusOK, comments)
}
