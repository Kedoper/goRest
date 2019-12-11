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

func GetUsers() ([]*User, bool) {
	users := make([]*User, 0)

	db, err := sql.Open("mysql", "root:123@tcp(localhost:3050)/panel")
	if err != nil {
		log.Print("Error to connect to database")
		log.Print(err)
	}
	rows, err := db.Query("SELECT id,login,first_name,last_name,email,telephone,sex,vk_id FROM users")
	if err != nil {
		log.Print("Error query")
		log.Print(err)
	}

	if err != nil {
		return users, true
	}

	defer rows.Close()

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
		)
		if err != nil {
			log.Print("Error scan row")
			log.Print(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Print("zx")
		log.Print(err)
	}
	return users, false
}

func GetUserById(id int64) ([]*User, bool) {
	user := make([]*User, 0)

	db, err := sql.Open("mysql", "root:123@tcp(localhost:3050)/panel")
	if err != nil {
		log.Print("Error to connect to database")
		log.Print(err)
	}
	rows, err := db.Query("SELECT id,login,first_name,last_name,email,telephone,sex,vk_id FROM user WHERE id = ?", id)
	if err != nil {
		log.Print("Error query")
		log.Print(err)
	}

	if err != nil {
		return user, true
	}

	defer rows.Close()

	for rows.Next() {
		user_ := new(User)
		err := rows.Scan(
			&user_.ID,
			&user_.Login,
			&user_.FirstName,
			&user_.LastName,
			&user_.Email,
			&user_.Telephone,
			&user_.Sex,
			&user_.VkID,
		)
		if err != nil {
			log.Print("Error scan row")
			log.Print(err)
		}
		user = append(user, user_)
	}
	if err = rows.Err(); err != nil {
		log.Print("zx")
		log.Print(err)
	}
	return user, false
}
