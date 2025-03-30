package render

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"web_magazin_v1/db"
	"web_magazin_v1/midlware"
	"web_magazin_v1/model"
	"web_magazin_v1/paigesrander"
	"web_magazin_v1/user"
)

func TovarListPage(w http.ResponseWriter, r *http.Request) {
	db := db.DatabaseSQL()
	tovarList := model.GetTovarList(db) // Получаем список товаров

	// Получаем пользователя из сессии
	session, _ := midlware.GetSession(w, r)
	userID, ok := session.Values["user_id"]

	var usr *model.UserModel
	if ok && userID != nil {
		usr = user.GetUserByID(db, userID.(uint)) // Функция, которая получает пользователя по ID
	}

	// Создаём объект с данными для шаблона
	data := model.PageData{
		User:   usr,
		Tovars: tovarList,
	}

	paigesrander.Render(w, "template/index.html", data) // Передаём данные в ша
}

func TovarDetail(w http.ResponseWriter, r *http.Request) {
	db2 := db.DatabaseSQL()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["Id"])
	tovar, _ := model.GetTovarByID(db2, uint(id))

	session, _ := midlware.GetSession(w, r)
	userID, ok := session.Values["user_id"]

	var usr *model.UserModel
	if ok && userID != nil {
		usr = user.GetUserByID(db2, userID.(uint))
	}

	data := model.PageDataForOneProduct{
		User:  usr,
		Tovar: tovar,
	}

	paigesrander.Render(w, "template/tovar_detail.html", data)
}
