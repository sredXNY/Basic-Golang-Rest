package main

import ( 
     "github.com/gin-gonic/gin"
	  "github.com/jinzhu/gorm"
	  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Users struct {
    Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
    Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}

func PostUser(c *gin.Context) {
     db := InitDb()
	   defer db.Close()
	var user Users
    c.Bind(&user)
	
	if user.Firstname != "" && user.Lastname != "" {
        // INSERT INTO "users" (name) VALUES (user.Name);
        db.Create(&user)
        // Display error
        c.JSON(201, gin.H{"success": user})
    } else {
        // Display error
        c.JSON(422, gin.H{"error": "Fields are empty"})
    }
	 // curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
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

    // curl -i http://localhost:10000/api/v1/users
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

    if user.Id != 0 {
        // Display JSON result
        c.JSON(200, user)
    } else {
        // Display JSON error
        c.JSON(404, gin.H{"error": "User not found"})
    }

    // curl -i http://localhost:10000/api/v1/users/1
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

    if user.Firstname != "" && user.Lastname != "" {

        if user.Id != 0 {
            var newUser Users
            c.Bind(&newUser)

            result := Users{
                Id:        user.Id,
                Firstname: newUser.Firstname,
                Lastname:  newUser.Lastname,
            }

            // UPDATE users SET firstname='newUser.Firstname', lastname='newUser.Lastname' WHERE id = user.Id;
            db.Save(&result)
            // Display modified data in JSON message "success"
            c.JSON(200, gin.H{"success": result})
        } else {
            // Display JSON error
            c.JSON(404, gin.H{"error": "User not found"})
        }

    } else {
        // Display JSON error
        c.JSON(422, gin.H{"error": "Fields are empty"})
    }

    // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:10000/api/v1/users/1
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

    if user.Id != 0 {
        // DELETE FROM users WHERE id = user.Id
        db.Delete(&user)
        // Display JSON result
        c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
    } else {
        // Display JSON error
        c.JSON(404, gin.H{"error": "User not found"})
    }

    // curl -i -X DELETE http://localhost:1000/api/v1/users/1
}

func InitDb() *gorm.DB {
    
	db, err := gorm.Open("postgres", "host=xxx.xxx.xxx.xxx user=user dbname=test sslmode=disable password=ps")
  
    db.LogMode(true)
	
    // Error
    if err != nil {
        panic(err)
    }

    // Creating the table
    db.AutoMigrate(&Users{})

    return db
}

func main() {
    r := gin.Default()

    v1 := r.Group("api/v1")
    {
        v1.POST("/users", PostUser)
        v1.GET("/users", GetUsers)
        v1.GET("/users/:id", GetUser)
        v1.PUT("/users/:id", UpdateUser)
        v1.DELETE("/users/:id", DeleteUser)
    }

    r.Run(":10000")
}