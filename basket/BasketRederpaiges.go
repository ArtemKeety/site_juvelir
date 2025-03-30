package basket

import (
	"net/http"
	db2 "web_magazin_v1/db"
	"web_magazin_v1/midlware"
	"web_magazin_v1/paigesrander"
)

func BasketListPage(w http.ResponseWriter, r *http.Request) {
	database := db2.DatabaseSQL()

	session, _ := midlware.GetSession(w, r)
	userID, ok := session.Values["user_id"].(uint)
	if !ok {
		http.Error(w, "Не удалось найти user_id в сессии", http.StatusUnauthorized)
		return
	}

	data, _ := GetBasketList(database, userID)

	paigesrander.Render(w, "template/basket.html", data)
}
