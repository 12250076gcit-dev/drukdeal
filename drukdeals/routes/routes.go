package routes

import (
	"drukdeals/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", controller.Signup).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("POST")
	r.HandleFunc("/auth/check", controller.CheckAuth).Methods("GET")

	// Page routes for login/signup
	r.HandleFunc("/login", serveFile("./view/index.html")).Methods("GET")
	r.HandleFunc("/signup", serveFile("./view/signup.html")).Methods("GET")
	r.HandleFunc("/signup-page", serveFile("./view/signup.html")).Methods("GET")

	r.HandleFunc("/products", controller.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/admin/metrics", controller.AdminMetrics).Methods("GET")
	r.HandleFunc("/product", controller.AddProduct).Methods("POST")
	r.HandleFunc("/my-products", controller.GetMyProducts).Methods("GET")
	r.HandleFunc("/product/{id}", controller.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/categories", controller.GetCategories).Methods("GET")

	// Serve HTML pages
	r.HandleFunc("/", serveFile("./view/entry.html")).Methods("GET")
	r.HandleFunc("/entry", serveFile("./view/entry.html")).Methods("GET")
	r.HandleFunc("/entry.html", serveFile("./view/entry.html")).Methods("GET")
	r.HandleFunc("/home", serveFile("./view/home.html")).Methods("GET")
	r.HandleFunc("/add-item", serveFile("./view/add-item.html"))
	r.HandleFunc("/my-items", serveFile("./view/my-items.html"))
	r.HandleFunc("/admin", serveFile("./view/admin.html"))
	r.HandleFunc("/admin.html", serveFile("./view/admin.html"))
	r.HandleFunc("/home.html", serveFile("./view/home.html"))
	r.HandleFunc("/my-items.html", serveFile("./view/my-items.html"))
	r.HandleFunc("/add-item.html", serveFile("./view/add-item.html"))

	return r
}

func serveFile(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}
