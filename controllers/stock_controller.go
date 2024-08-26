package controllers

import (
	"encoding/json"
	"inventory-api/config"
	"inventory-api/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		logrus.WithError(err).Error("Failed to decode request")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&stock).Error; err != nil {
		logrus.WithError(err).Error("Failed to create stock")
		http.Error(w, "Failed to create stock", http.StatusInternalServerError)
		return
	}

	logrus.WithField("request", stock).Info("Stock created")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

func ListStock(w http.ResponseWriter, r *http.Request) {
	var stocks []models.Stock
	if err := config.DB.Find(&stocks).Error; err != nil {
		logrus.WithError(err).Error("Failed to list stocks")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.Info("Stocks listed")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var stock models.Stock

	if err := config.DB.First(&stock, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logrus.WithError(err).Error("Stock not found GetStock")
			http.Error(w, "Stock not found", http.StatusNotFound)
		} else {
			logrus.WithError(err).Error("Failed to get stock")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	logrus.WithField("id", id).Info("Stock retrieved")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var stock models.Stock

	if err := config.DB.First(&stock, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logrus.WithError(err).Error("Stock not found UpdateStock")
			http.Error(w, "Stock not found", http.StatusNotFound)
		} else {
			logrus.WithError(err).Error("Failed to get stock for update")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var updatedStock models.Stock
	if err := json.NewDecoder(r.Body).Decode(&updatedStock); err != nil {
		logrus.WithError(err).Error("Failed to decode request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update fields
	stock.ItemName = updatedStock.ItemName
	stock.Quantity = updatedStock.Quantity
	stock.SerialNumber = updatedStock.SerialNumber
	stock.ItemImage = updatedStock.ItemImage
	stock.CreatedBy = updatedStock.CreatedBy
	stock.UpdatedBy = updatedStock.UpdatedBy
	stock.UpdatedAt = time.Now()

	if err := config.DB.Save(&stock).Error; err != nil {
		logrus.WithError(err).Error("Failed to update stock")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.WithField("id", id).Info("Stock updated")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result := config.DB.Delete(&models.Stock{}, id)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("Failed to delete stock")
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		logrus.Error("Stock not found DeleteStock")
		http.Error(w, "Stock not found", http.StatusNotFound)
		return
	}

	logrus.WithField("id", id).Info("Stock deleted")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Stock deleted successfully"})
}
