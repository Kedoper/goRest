package helpers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Sex       int    `json:"sex"`
	VkID      int    `json:"vk_id"`
	//GoogleSecret string `json:"google_secret"`
	//Password     string `json:"password"`
	//VkToken      string `json:"vk_token"`
	//Avatar       string `json:"avatar"`
	//Level        string `json:"level"`
	//Active       int    `json:"active"`
	//LoginBanTo   int    `json:"login_ban_to"`
	//Bitmap       int    `json:"bitmap"`
	//RegisterDate int    `json:"register_date"`
}

func GetUsers() []*User {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3050)/panel")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT id,login,first_name,last_name,email,telephone,sex,vk_id FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(
			&user.ID,
			&user.Login,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Telephone,
			&user.Sex,
			&user.VkID,
			//&user.Password,
			//&user.GoogleSecret,
			//&user.VkToken,
			//&user.Avatar,
			//&user.Level,
			//&user.Active,
			//&user.LoginBanTo,
			//&user.Bitmap,
			//&user.RegisterDate,
		)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}
