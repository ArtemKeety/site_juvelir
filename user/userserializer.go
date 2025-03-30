package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	db2 "web_magazin_v1/db"
	"web_magazin_v1/midlware"
	"web_magazin_v1/model"
	"web_magazin_v1/paigesrander"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Auth(w http.ResponseWriter, req *http.Request) {
	var data = make(map[string]interface{})

	if req.Method == http.MethodPost {
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")

		data["Username"] = username // Передаем логин
		data["Password"] = password // Передаем пароль

		fmt.Println(username, password)

		db := db2.DatabaseSQL()

		row := db.QueryRow(`SELECT id, login, email, password FROM user WHERE login = ?`, username)

		var User model.UserModel
		err := row.Scan(&User.Id, &User.Login, &User.Email, &User.Password)

		if !checkPasswordHash(password, User.Password) {
			data["Error"] = "Неверное имя пользователя или пароль"
			paigesrander.Render(w, "template/login.html", data)
			return
		}

		session, err := midlware.GetSession(w, req)
		if err != nil {
			data["Error"] = "Ошибка при создании сессии"
			paigesrander.Render(w, "template/login.html", data)
			return
		}

		session.Values["user_id"] = User.Id
		session.Save(req, w)

		http.Redirect(w, req, "/", http.StatusSeeOther)
	} else {
		paigesrander.Render(w, "template/login.html", data)
	}

}

func Logout(w http.ResponseWriter, req *http.Request) {
	session, err := midlware.GetSession(w, req)
	if err != nil {
		http.Error(w, "Ошибка при получении сессии", http.StatusInternalServerError)
		return
	}

	// Удаляем данные сессии
	delete(session.Values, "user_id")
	session.Options.MaxAge = -1 // Сразу удаляет cookie
	session.Save(req, w)

	// Перенаправляем на страницу входа
	http.Redirect(w, req, "/login/", http.StatusSeeOther)
}
