package user

import (
	"fmt"
	"net/http"
	db2 "web_magazin_v1/db"
	"web_magazin_v1/paigesrander"
)

func Register(w http.ResponseWriter, req *http.Request) {
	var data = make(map[string]interface{})

	if req.Method == http.MethodPost {
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")
		email := req.PostFormValue("email")

		data["Username"] = username // Передаем логин
		data["Password"] = password // Передаем пароль
		data["Email"] = email       // Передаем пароль

		fmt.Println(username, password)

		db := db2.DatabaseSQL()
		querry := db.QueryRow(`SELECT u.id FROM user u where u.login = ? OR u.email = ? LIMIT 1`, username, email)

		var id uint
		querry.Scan(&id)

		if id > 0 {
			data["Error"] = "Не уникальные данные пользователя или почты"
			paigesrander.Render(w, "template/register.html", data)
			return
		}

		hashedPassword, _ := hashPassword(password)

		fmt.Println(hashedPassword, username, email)

		_, err := db.Exec(`insert into user (login, password, email) values (?, ?, ?)`,
			username, hashedPassword, email)

		if err != nil {
			panic(err)
		}

		paigesrander.Render(w, "template/login.html", data)
	} else {
		paigesrander.Render(w, "template/register.html", data)
	}

}
