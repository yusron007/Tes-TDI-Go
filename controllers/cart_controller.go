package controllers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"E-commerce/models"

)

func InitCartController(database *sql.DB) {
	db = database
}

// Handler untuk menambahkan produk ke keranjang belanja
func AddToCart(c echo.Context) error {
	cartItem := new(models.CartItem)
	if err := c.Bind(cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validasi input
	if cartItem.KodeProduk == "" || cartItem.JumlahProduk <= 0 {
		return c.JSON(http.StatusBadRequest, "Invalid input data")
	}

	// Gunakan prepared statement untuk insert atau update
	stmt := "INSERT INTO PESANAN (KODE_PRODUK, JUMLAH_PRODUK) VALUES (?, ?) ON DUPLICATE KEY UPDATE JUMLAH_PRODUK = JUMLAH_PRODUK + VALUES(JUMLAH_PRODUK)"
	_, err := db.Exec(stmt, cartItem.KodeProduk, cartItem.JumlahProduk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, cartItem)
}

// Handler untuk mendapatkan keranjang belanja pelanggan
func GetCart(c echo.Context) error {
	rows, err := db.Query("SELECT KODE_PRODUK, JUMLAH_PRODUK FROM PESANAN")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	cartItems := []models.CartItem{}
	for rows.Next() {
		var cartItem models.CartItem
		err := rows.Scan(&cartItem.KodeProduk, &cartItem.JumlahProduk)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		cartItems = append(cartItems, cartItem)
	}
	return c.JSON(http.StatusOK, cartItems)
}

// Handler untuk menghapus produk dari keranjang belanja
func RemoveFromCart(c echo.Context) error {
	kodeProduk := c.Param("kode_produk")

	_, err := db.Exec("DELETE FROM PESANAN WHERE KODE_PRODUK = ?", kodeProduk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
