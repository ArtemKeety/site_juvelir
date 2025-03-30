package model

import (
	"database/sql"
)

func GetTovarByID(db *sql.DB, id uint) (TovarDetail, error) {
	var tovar TovarDetail

	// Запрос для получения деталей товара
	queryTovar := `SELECT id, name, price, counts, idCategory FROM tovar WHERE id = ?`
	err := db.QueryRow(queryTovar, id).Scan(&tovar.Id, &tovar.Name, &tovar.Price, &tovar.Counts, &tovar.IdCategory)
	if err != nil {
		return tovar, err
	}

	// Запрос для получения фотографий товара
	queryPhotos := `SELECT id, filepath, idTovar FROM photo WHERE idTovar = ?`
	rows, err := db.Query(queryPhotos, id)
	if err != nil {
		return tovar, err
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		var photo Photo
		if err := rows.Scan(&photo.Id, &photo.Filepath, &photo.TovarId); err != nil {
			return tovar, err
		}
		photos = append(photos, photo)
	}
	tovar.Photos = photos

	return tovar, nil
}

func GetTovarList(db *sql.DB) []Product {
	var products []Product

	str := `SELECT t.id, t.name, t.price, p.filepath 
			FROM Tovar t 
			JOIN (
				SELECT idTovar, MIN(id) AS min_id 
				FROM Photo 
				GROUP BY idTovar
				) 
			AS first_photos ON t.id = first_photos.idTovar 
			JOIN Photo p ON first_photos.min_id = p.id `

	rows, err := db.Query(str)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var product Product
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.PhotoPath)
		products = append(products, product)
	}

	return products
}
