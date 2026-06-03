package controller

import (
	"drukdeals/model"
	"drukdeals/utils"
	"net/http"
)

// AdminMetrics returns simple site metrics for the admin dashboard.
func AdminMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	users, err := model.CountUsers()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to count users")
		return
	}
	products, err := model.CountProducts()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to count products")
		return
	}
	sellers, err := model.CountSellers()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to count sellers")
		return
	}
	cats, err := model.CountCategories()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to count categories")
		return
	}
	revenue, err := model.SumProductPrices()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to sum prices")
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]interface{}{
		"total_users":      users,
		"total_products":   products,
		"total_sellers":    sellers,
		"total_categories": cats,
		"total_revenue":    revenue,
	})
}
