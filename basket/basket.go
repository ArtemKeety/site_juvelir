package basket

import (
	"database/sql"
	"errors"
	"fmt"
	"web_magazin_v1/model"
)

func AddTovarInBasket(database *sql.DB, tovarid, userid, count uint) error {
	if !CheckTovarinBasket(database, int(tovarid), userid) {
		querry := `INSERT INTO basket(user_id, tovar_id, counts) VALUES (?, ?, ?)`
		_, err := database.Exec(querry, userid, tovarid, count)
		if err != nil {
			return err
		} else {
			return nil
		}

	}
	return fmt.Errorf("Товар с ID %d уже есть в корзине", tovarid)
}

func CheckTovarinBasket(database *sql.DB, tovarid int, userid uint) bool {
	querry := database.QueryRow(`SELECT * FROM basket b where  b.tovar_id = ? and b.user_id = ?`, tovarid, userid)

	var basket model.Basket
	err := querry.Scan(&basket.Id, &basket.UserId, &basket.TovarId, &basket.Counts)
	if err != nil {
		return false
	}
	return true
}

func GetBasketList(database *sql.DB, userid uint) ([]model.Basket, error) {
	var list_basket []model.Basket

	rows, err := database.Query(`
		SELECT b.id, b.tovar_id, b.counts, 
		       t.id, t.name, t.price, t.counts, p.filepath AS filepath
		FROM basket b
		JOIN Tovar t ON b.tovar_id = t.id
		LEFT JOIN (
		    SELECT idTovar, MIN(id) AS min_id 
		    FROM Photo 
		    GROUP BY idTovar
		) AS first_photos ON t.id = first_photos.idTovar
		LEFT JOIN Photo p ON first_photos.min_id = p.id
		WHERE b.user_id = ?`, userid)

	if err != nil {
		return list_basket, err
	}

	defer rows.Close() // Закрываем результат после выхода из функции

	// Читаем данные из rows
	for rows.Next() {
		var basket model.Basket
		err := rows.Scan(
			&basket.Id, &basket.TovarId, &basket.Counts,
			&basket.Tovar.Id, &basket.Tovar.Name, &basket.Tovar.Price, &basket.Tovar.Count, &basket.Tovar.PhotoPath)

		if err != nil {
			return list_basket, err // Если ошибка, сразу выходим
		}

		list_basket = append(list_basket, basket)
	}

	// Проверяем, была ли ошибка во время перебора
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return list_basket, nil
}

func GetTotalPriceByIdTovar(db *sql.DB, tovarid, count int) (int, error) {
	var Product model.Product
	err := db.QueryRow(`SELECT t.price, t.counts From Tovar t Where t.id = ?`, tovarid).Scan(&Product.Price, &Product.Count)
	if err != nil {
		return 0, err
	}
	if count > Product.Count {
		return 0, errors.New("Больше чем количество в магазине")
	}

	return count * int(Product.Price), nil
}

func GetTotalPriceByIdUser(db *sql.DB, userId uint) (int, error) {
	totalSum := 0
	err := db.QueryRow(`SELECT sum(b.counts * t.price) as Totalsum
							from basket b 
							Join Tovar t ON t.id = b.tovar_id
							where b.user_id = ?`, userId).Scan(&totalSum)

	if err != nil {
		return totalSum, err
	}

	return totalSum, nil
}

func CheckBasketById(database *sql.DB, basketId int) bool {
	var id int
	err := database.QueryRow(`SELECT b.id FROM basket b WHERE b.id = ?`, basketId).Scan(&id)
	if err != nil || err == sql.ErrNoRows {
		return false
	}
	return true
}

func UpdateTovarInBasketById(database *sql.DB, basketId, count int, userId uint) error {
	if !CheckBasketById(database, basketId) {
		return errors.New("Товара нет в корзине")
	}

	querry := ` UPDATE basket
				SET counts = ?
				WHERE user_id = ? AND id = ?`

	_, err := database.Exec(querry, count, userId, basketId)
	if err != nil {
		return err
	}
	return nil

}

func DeleteBasketItemById(database *sql.DB, basketId int) error {
	if !CheckBasketById(database, basketId) {
		return errors.New("Товара нет в корзине")
	}

	querry := `DELETE FROM basket WHERE id = ?`
	_, err := database.Exec(querry, basketId)
	if err != nil {
		return err
	}
	return nil

}

func ClearCartByUserId(database *sql.DB, userid uint) error {
	querry := `DELETE FROM basket WHERE user_id = ?`
	_, err := database.Exec(querry, userid)
	if err != nil {
		return err
	}
	return nil
}
