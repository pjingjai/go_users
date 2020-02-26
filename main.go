package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	ID         int    `gorm:"AUTO_INCREMENT" form:"ID" json:"id"`
	First_name string `gorm:"not null" form:"First_name" json:"first_name"`
	Last_name  string `gorm:"not null" form:"Last_name" json:"last_name"`
	Email      string `gorm:"not null" form:"Email" json:"email"`
	Gender     string `gorm:"not null" form:"Gender" json:"gender"`
	Age        int    `gorm:"not null" form:"Age" json:"age"`
}

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./user.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	}

	return db
}

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Add("Access-Control-Allow-Headers", "POST, GET, PUT, DELETE")
// 		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Option, Authorization")
// 		c.Next()
// 	}
// }

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	v1 := r.Group("api/v1")
	{
		v1.GET("/users", GetUsers)

		v1.GET("/user/:id", GetUser)
		v1.GET("/first_name/:first_name", GetUserByFirstName)
		v1.GET("/last_name/:last_name", GetUserByLastName)
		v1.GET("/email/:email", GetUserByEmail)

		v1.GET("/gender/:gender", GetUsersByGender)
		v1.GET("/age/:age", GetUsersByAge)
		v1.GET("/ages/:age1/:age2", GetUsersByAges)

		v1.POST("/user", PostUser)
		v1.PUT("/user/:id", UpdateUser)
		v1.DELETE("/user/:id", DeleteUser)
	}

	r.Run(":9000")
}

func PostUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user Users
	var u Users
	c.Bind(&user)
	
	if user.First_name != "" && user.Last_name != "" && user.Email != "" && user.Gender != "" && user.Age != 0 {
		// Fixing AUTO_INCREMENT bug
		db.Select("id").Last(&u)
		user.ID = u.ID + 1
		// INSERT INTO "users" (name) VALUES (user.Name);
		db.Create(&user)

		// Display error
		c.JSON(201, gin.H{"success": user})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func GetUsers(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var users []Users
	// SELECT * FROM users
	db.Find(&users)

	// Display JSON result
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// Display JSON result
		c.JSON(200, user)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetUserByFirstName(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	first_name := c.Params.ByName("first_name")
	var user Users
	// SELECT * FROM users WHERE first_name = ?;
	db.First(&user, "first_name = ?", first_name)

	if user.First_name != "" {
		// Display JSON result
		c.JSON(200, user)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetUserByLastName(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	last_name := c.Params.ByName("last_name")
	var user Users
	// SELECT * FROM users WHERE last_name = ?;
	db.First(&user, "last_name = ?", last_name)

	if user.Last_name != "" {
		// Display JSON result
		c.JSON(200, user)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetUserByEmail(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	email := c.Params.ByName("email")
	var user Users
	// SELECT * FROM users WHERE email = ?;
	db.First(&user, "email = ?", email)

	if user.Email != "" {
		// Display JSON result
		c.JSON(200, user)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetUsersByGender(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	gender := c.Params.ByName("gender")
	var users []Users
	// SELECT * FROM users WHERE gender = ?;
	db.Where("gender = ?", gender).Find(&users)

	if gender != "" {
		// Display JSON result
		c.JSON(200, users)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetUsersByAge(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	age := c.Params.ByName("age")
	var users []Users
	// SELECT * FROM users WHERE age = ?;
	db.Where("age = ?", age).Find(&users)

	if age != "" {
		// Display JSON result
		c.JSON(200, users)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetUsersByAges(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	age1 := c.Params.ByName("age1")
	age2 := c.Params.ByName("age2")
	var users []Users
	// SELECT * FROM users WHERE age BETWEEN ? AND ?;
	db.Where("age BETWEEN ? AND ?", age1, age2).Find(&users)

	if age1 != "" && age2 != "" {
		// Display JSON result
		c.JSON(200, users)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func UpdateUser(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		var newUser Users
		c.Bind(&newUser)

		// Only If all fields are empty
		if newUser.First_name != "" || newUser.Last_name != "" || newUser.Email != "" || newUser.Gender != "" || newUser.Age != 0 {

			newFirst_name := newUser.First_name
			if newUser.First_name == "" {
				newFirst_name = user.First_name
			}

			newLast_name := newUser.Last_name
			if newUser.Last_name == "" {
				newLast_name = user.Last_name
			}

			newEmail := newUser.Email
			if newUser.Email == "" {
				newEmail = user.Email
			}

			newGender := newUser.Gender
			if newUser.Gender == "" {
				newGender = user.Gender
			}

			newAge := newUser.Age
			if newUser.Age == 0 {
				newAge = user.Age
			}

			result := Users{
				ID:         user.ID,
				First_name: newFirst_name,
				Last_name:  newLast_name,
				Email:      newEmail,
				Gender:     newGender,
				Age:        newAge,
			}

			// UPDATE users SET First_name='newUser.First_name', Last_name='newUser.Last_name' WHERE id = user.ID;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(400, gin.H{"error": "Fields are empty"})
		}
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func DeleteUser(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// DELETE FROM users WHERE id = user.ID
		db.Delete(&user)
		// Display JSON result
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
