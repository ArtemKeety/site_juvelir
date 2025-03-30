package midlware

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("секретный-ключ"))

func GetSession(w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return nil, err
	}
	return session, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := GetSession(w, r)
		if err != nil {
			http.Error(w, "Ошибка при получении сессии", http.StatusInternalServerError)
			return
		}

		userID, ok := session.Values["user_id"]
		if !ok || userID == nil {
			// Если пользователя нет в сессии, редиректим на страницу входа
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
			return
		}

		// Передаём управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}
