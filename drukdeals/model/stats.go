package model

import (
	"drukdeals/db"
)

func CountUsers() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	return count, err
}

func CountProducts() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	return count, err
}

func CountSellers() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(DISTINCT user_id) FROM products").Scan(&count)
	return count, err
}

func CountCategories() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	return count, err
}

func SumProductPrices() (float64, error) {
	var sum float64
	err := db.DB.QueryRow("SELECT COALESCE(SUM(price),0) FROM products").Scan(&sum)
	return sum, err
}
