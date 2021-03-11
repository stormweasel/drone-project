package login

import (
	"net/http"
	"pig/cmd/dbConn"
	"pig/internal/hash"
	"pig/internal/jwt"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginFunction(c *gin.Context) {
	var userFromWeb User

	if err := c.ShouldBindJSON(&userFromWeb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	var isUsernameExists int

	db := dbConn.DbConn()
	if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM PIGusers WHERE username = (?));", userFromWeb.Username).Scan(&isUsernameExists); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error11"})
		return
	}

	if isUsernameExists == 0 {
		c.JSON(http.StatusUnauthorized, "Username or password is incorrect.")
		return

	} else {

		var userFromDB User
		if err := db.QueryRow("SELECT userID, username, password FROM PIGusers WHERE username = (?);", userFromWeb.Username).Scan(&userFromDB.ID, &userFromDB.Username, &userFromDB.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error3"})
			return
		}

		if userFromWeb.Username == "" && userFromDB.Password == userFromWeb.Password {
			c.JSON(http.StatusUnauthorized, "Username is required.")
			return
		}
		if userFromDB.Username == userFromWeb.Username && userFromWeb.Password == "" {
			c.JSON(http.StatusUnauthorized, "Password is required.")
			return
		}
		if userFromWeb.Username == "" || userFromWeb.Password == "" {
			c.JSON(http.StatusUnauthorized, "All fields are required")
			return
		}
		if userFromDB.Username == userFromWeb.Username && !hash.Match(userFromWeb.Password, userFromDB.Password) {
			c.JSON(http.StatusUnauthorized, "Username or password is incorrect.")
			return
		}
		token, err := jwt.CreateToken(userFromDB.ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "token": token})

	}
}
