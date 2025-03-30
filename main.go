package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"web_magazin_v1/basket"
	"web_magazin_v1/midlware"
	"web_magazin_v1/render"
	"web_magazin_v1/user"
)

func server2() {
	// Создаём новый роутер
	r := mux.NewRouter()

	// Маршруты, доступные без аутентификации
	r.HandleFunc("/login/", user.Auth).Methods("GET", "POST")
	r.HandleFunc("/register/", user.Register).Methods("GET", "POST")
	r.HandleFunc("/logout/", user.Logout).Methods("GET", "POST")

	// Подмаршрутизатор для защищённых маршрутов
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(midlware.AuthMiddleware)

	// Защищённые маршруты
	protected.HandleFunc("/", render.TovarListPage).Methods("GET")
	protected.HandleFunc("/tovar/{Id:[0-9]+}/", render.TovarDetail).Methods("GET")
	protected.HandleFunc("/api/CheckTovarById/", basket.JsonChekTovarInbasket).Methods("GET")
	protected.HandleFunc("/api/AddToCart/", basket.JsonAddToBasket).Methods("POST")
	protected.HandleFunc("/basket/", basket.BasketListPage).Methods("GET")
	protected.HandleFunc("/api/GetCurrentPriceByIdTovar/", basket.JsonGetCurrentPriceByIdTovar).Methods("GET")
	protected.HandleFunc("/api/GetTotalSumByUserId/", basket.JsonGetTotalSumByUserId).Methods("GET")
	protected.HandleFunc("/api/UpdateToCart/", basket.JsonUpdateToCartById).Methods("POST")
	protected.HandleFunc("/api/DeleteToCart/", basket.JsonDeleteToCartById).Methods("POST")
	protected.HandleFunc("/api/ClearToCart/", basket.JsonClearToCartByUserId).Methods("POST")

	// Обслуживание статических файлов
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Запуск сервера
	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func main() {
	server2()

}
