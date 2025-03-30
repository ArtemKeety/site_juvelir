package user

import (
	"database/sql"
	"web_magazin_v1/model"
)

func GetUserByID(db *sql.DB, userID uint) *model.UserModel {
	row := db.QueryRow("SELECT id, login, email FROM user WHERE id = ?", userID)

	var user model.UserModel
	err := row.Scan(&user.Id, &user.Login, &user.Email)
	if err != nil {
		return nil // Если не найден, возвращаем nil
	}
	return &user
}
