package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tsaqiffatih/ecommerce-bakcend/controllers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello This Is Ecommerce Backend"))
	})
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")

	return router
}
