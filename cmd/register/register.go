package register

import (
	"net/http"
	"pig/internal/emailValidation"
	"pig/internal/hash"

	"pig/cmd/dbConn"

	"github.com/gin-gonic/gin"
)

type UserForRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserForResult struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func RegisterTheUser(c *gin.Context) {

	var user UserForRegister

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid JSON provided."})
		return
	}

	if user.Username == "" && user.Password == "" && user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username, password and email are required."})
		return
	}

	if user.Username == "" && user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and email are required."})
		return
	}

	if user.Password == "" && user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password and email are required."})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required."})
		return
	}

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required."})
		return
	}

	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required."})
		return
	}

	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be 8 characters long."})
		return
	}

	if e := user.Email; !emailValidation.IsEmailValid(e) {
		c.JSON(http.StatusBadRequest, gin.H{"error": user.Email + "is not a valid email."})
		return
	}

	var isUsernameExists int
	var isEmailExists int
	db := dbConn.DbConn()
	if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM PIGusers WHERE username = (?));", user.Username).Scan(&isUsernameExists); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error11"})
		return
	}

	if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM PIGusers WHERE email = (?));", user.Email).Scan(&isEmailExists); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error12"})
		return
	}

	if isUsernameExists == 0 && isEmailExists == 0 {
		insData, err := db.Prepare("INSERT INTO PIGusers (username, password, email) VALUES (?,?,?);")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error2"})
			return
		}

		hashedPassword, _ := hash.Password(user.Password)
		insData.Exec(user.Username, hashedPassword, user.Email)

		var userResult UserForResult
		if err := db.QueryRow("SELECT userID, username FROM PIGusers WHERE userID = LAST_INSERT_ID();").Scan(&userResult.Id, &userResult.Username); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error3"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": userResult.Id, "username": userResult.Username})
		return

	} else if isUsernameExists != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is already taken."})
		return
	} else if isEmailExists != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already taken."})
		return
	}

}
