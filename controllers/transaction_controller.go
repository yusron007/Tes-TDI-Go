package controllers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"E-commerce/models"
)

func InitTransactionController(database *sql.DB) {
	db = database
}

// Handler untuk proses checkout
func Checkout(c echo.Context) error {
	checkoutData := new(models.Transaction)
	if err := c.Bind(checkoutData); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Hitung total harga produk yang akan dibeli oleh pelanggan
	var totalHarga float64
	err := db.QueryRow("SELECT SUM(P.HARGA_PRODUK * TRANSAKSI.JUMLAH_PRODUK) FROM TRANSAKSI JOIN PRODUK P ON P.KODE_PRODUK = P.KODE_PRODUK WHERE TRANSAKSI.ID_TRANSAKSI = $1", checkoutData.IDTransaksi).Scan(&totalHarga)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Masukkan transaksi ke dalam tabel TRANSAKSI
	_, err = db.Exec("INSERT INTO TRANSAKSI (ID_TRANSAKSI, JUMLAH_PRODUK, TOTAL_HARGA, STATUS_TRANSAKSI) VALUES ($1, $2, $3, $4)",
		checkoutData.IDTransaksi, checkoutData.JumlahProduk, totalHarga, checkoutData.StatusTransaksi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// // Hapus produk dari keranjang belanja setelah berhasil checkout
	// _, err = db.Exec("DELETE FROM TRANSAKSI WHERE ID_TRANSAKSI = $1", checkoutData.IDTransaksi)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	return c.JSON(http.StatusCreated, checkoutData)
}

// Handler untuk mendapatkan transaksi pelanggan
func GetTransactions(c echo.Context) error {
	rows, err := db.Query("SELECT ID_TRANSAKSI, JUMLAH_PRODUK, TOTAL_HARGA, STATUS_TRANSAKSI FROM TRANSAKSI")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	transactions := []models.Transaction{}
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.IDTransaksi, &transaction.JumlahProduk, &transaction.TotalHarga, &transaction.StatusTransaksi)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		transactions = append(transactions, transaction)
	}
	return c.JSON(http.StatusOK, transactions)
}

// Handler untuk mendapatkan detail transaksi berdasarkan ID Transaksi
func GetTransaction(c echo.Context) error {
	idTransaksi := c.Param("id_transaksi")
	transaction := new(models.Transaction)
	err := db.QueryRow("SELECT ID_TRANSAKSI, JUMLAH_PRODUK, TOTAL_HARGA, STATUS_TRANSAKSI FROM TRANSAKSI WHERE ID_TRANSAKSI = $1", idTransaksi).Scan(&transaction.IDTransaksi, &transaction.JumlahProduk, &transaction.TotalHarga, &transaction.StatusTransaksi)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Transaksi tidak ditemukan")
	}
	return c.JSON(http.StatusOK, transaction)
}
