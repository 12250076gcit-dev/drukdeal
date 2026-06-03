package main

import (
	"drukdeals/db"
	"drukdeals/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println(" Database connected")

	router := routes.InitializeRoutes()

	// Serve static files
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./view/css/"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./view/js/"))))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./view/images/"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./view/"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println(" Server starting on port", port)
	http.ListenAndServe(":"+port, router)
}
