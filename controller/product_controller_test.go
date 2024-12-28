package controllers

import (
	"bytes"
	"encoding/json"
	"inventory-management/database"
	"inventory-management/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/products", GetProducts)
	r.POST("/products", CreateProduct)
	return r
}

func TestGetProducts(t *testing.T) {

	router := SetupTestRouter()
	database.DB.Create(&models.Product{Name: "Test Product", Description: "Test Description", Price: 100.0, Category: "Test"})

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var products []models.Product
	err := json.Unmarshal(w.Body.Bytes(), &products)
	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, "Test Product", products[0].Name)
}

func TestCreateProduct(t *testing.T) {
	router := SetupTestRouter()
	newProduct := models.Product{Name: "New Product", Description: "New Description", Price: 200.0, Category: "New"}
	payload, _ := json.Marshal(newProduct)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var createdProduct models.Product
	err := json.Unmarshal(w.Body.Bytes(), &createdProduct)
	assert.NoError(t, err)
	assert.Equal(t, newProduct.Name, createdProduct.Name)
	assert.Equal(t, newProduct.Price, createdProduct.Price)
}
