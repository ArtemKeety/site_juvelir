package basket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	db2 "web_magazin_v1/db"
	"web_magazin_v1/midlware"
)

func JsonChekTovarInbasket(w http.ResponseWriter, r *http.Request) {
	database := db2.DatabaseSQL()
	tovarIDStr := r.URL.Query().Get("tovar_id")
	tovarID, _ := strconv.Atoi(tovarIDStr)

	session, _ := midlware.GetSession(w, r)
	userID, _ := session.Values["user_id"]

	data := make(map[string]bool)

	data["status"] = CheckTovarinBasket(database, tovarID, userID.(uint))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func JsonAddToBasket(w http.ResponseWriter, r *http.Request) {
	// Читаем JSON тело запроса
	var request struct {
		Id string `json:"id"`
	}

	// Декодируем JSON
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	Id, _ := strconv.Atoi(request.Id)
	// Проверяем, что ID передан
	if Id == 0 {
		http.Error(w, "ID не может быть пустым", http.StatusBadRequest)
		return
	}

	// Получаем подключение к базе данных
	database := db2.DatabaseSQL()

	// Получаем сессию пользователя
	session, _ := midlware.GetSession(w, r)
	userID, ok := session.Values["user_id"].(uint)
	if !ok {
		http.Error(w, "Не удалось найти user_id в сессии", http.StatusUnauthorized)
		return
	}

	// Структура ответа
	response := make(map[string]string)

	// Попытка добавить товар в корзину
	err = AddTovarInBasket(database, uint(Id), userID, 1)
	if err != nil {
		// Логируем ошибку в консоль для отладки
		fmt.Println("Ошибка при добавлении товара в корзину:", err)
		// Возвращаем ошибку пользователю
		response["status"] = "error"
		response["message"] = "Ошибка при добавлении товара: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Ответ на успешное добавление
	response["status"] = "success"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func JsonGetCurrentPriceByIdTovar(w http.ResponseWriter, r *http.Request) {
	database := db2.DatabaseSQL()
	tovarIDStr := r.URL.Query().Get("tovar_id")
	tovarID, _ := strconv.Atoi(tovarIDStr)
	countStr := r.URL.Query().Get("count")
	count, _ := strconv.Atoi(countStr)

	data := make(map[string]string)
	CurrPrice, err := GetTotalPriceByIdTovar(database, tovarID, count)

	if err != nil {
		data["error"] = err.Error()
	}
	data["CurrentPrice"] = strconv.Itoa(CurrPrice)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func JsonGetTotalSumByUserId(w http.ResponseWriter, r *http.Request) {
	database := db2.DatabaseSQL()
	session, _ := midlware.GetSession(w, r)
	userID, _ := session.Values["user_id"]

	data := make(map[string]string)

	totalSum, err := GetTotalPriceByIdUser(database, userID.(uint))
	if err != nil {
		data["error"] = err.Error()
	}
	data["totalsum"] = strconv.Itoa(totalSum)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func JsonUpdateToCartById(w http.ResponseWriter, r *http.Request) {

	var request struct {
		BasketId string `json:"basketId"`
		Quantity string `json:"quantity"`
	}

	// Декодируем JSON
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	Id, _ := strconv.Atoi(request.BasketId)
	Count, _ := strconv.Atoi(request.Quantity)
	// Проверяем, что ID передан
	if Id == 0 || Count == 0 {
		http.Error(w, "ID или Количество не равно не может быть пустым", http.StatusBadRequest)
		return
	}

	database := db2.DatabaseSQL()
	session, _ := midlware.GetSession(w, r)
	userID, _ := session.Values["user_id"]
	data := make(map[string]string)

	err = UpdateTovarInBasketById(database, Id, Count, userID.(uint))
	if err != nil {
		data["error"] = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func JsonDeleteToCartById(w http.ResponseWriter, r *http.Request) {
	var request struct {
		BasketId string `json:"basketId"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	Id, _ := strconv.Atoi(request.BasketId)
	if Id == 0 {
		http.Error(w, "ID не равно не может быть пустым", http.StatusBadRequest)
		return
	}

	database := db2.DatabaseSQL()

	data := make(map[string]string)
	err = DeleteBasketItemById(database, Id)
	if err != nil {
		data["error"] = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func JsonClearToCartByUserId(w http.ResponseWriter, r *http.Request) {
	session, _ := midlware.GetSession(w, r)
	userID, _ := session.Values["user_id"]
	database := db2.DatabaseSQL()
	data := make(map[string]string)
	err := ClearCartByUserId(database, userID.(uint))
	if err != nil {
		data["error"] = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
