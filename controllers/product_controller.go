package controllers

import (
	"database/sql"
	"net/http"
	//"strconv"

	"github.com/labstack/echo/v4"

	"E-commerce/models"
)

var db *sql.DB

func InitProductController(database *sql.DB) {
	db = database
}

// Handler untuk membuat produk baru
func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Insert produk ke database
	_, err := db.Exec("INSERT INTO PRODUK (KODE_PRODUK, KODE_ETALASE, NAMA_PRODUK, HARGA_PRODUK, STOK_PRODUK, FOTO_PRODUK, KONDISI_PRODUK, BERAT_PRODUK, UKURAN_PRODUK, DESKRIPSI_PRODUK, VARIAN_PRODUK) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		product.KodeProduk, product.KodeEtalase, product.NamaProduk, product.HargaProduk, product.StokProduk, product.FotoProduk, product.KondisiProduk, product.BeratProduk, product.UkuranProduk, product.Deskripsi, product.VarianProduk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, product)
}

// Handler untuk mendapatkan daftar produk
func GetProducts(c echo.Context) error {
	rows, err := db.Query("SELECT * FROM PRODUK")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.KodeProduk, &product.KodeEtalase, &product.NamaProduk, &product.HargaProduk, &product.StokProduk, &product.FotoProduk, &product.KondisiProduk, &product.BeratProduk, &product.UkuranProduk, &product.Deskripsi, &product.VarianProduk)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		products = append(products, product)
	}
	return c.JSON(http.StatusOK, products)
}

// Handler untuk mendapatkan detail produk berdasarkan Kode Produk
func GetProduct(c echo.Context) error {
	kodeProduk := c.Param("kode_produk")
	product := new(models.Product)
	err := db.QueryRow("SELECT * FROM PRODUK WHERE KODE_PRODUK = $1", kodeProduk).Scan(&product.KodeProduk, &product.KodeEtalase, &product.NamaProduk, &product.HargaProduk, &product.StokProduk, &product.FotoProduk, &product.KondisiProduk, &product.BeratProduk, &product.UkuranProduk, &product.Deskripsi, &product.VarianProduk)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Produk tidak ditemukan")
	}
	return c.JSON(http.StatusOK, product)
}

// Handler untuk memperbarui produk berdasarkan Kode Produk
func UpdateProduct(c echo.Context) error {
	kodeProduk := c.Param("kode_produk")
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Update produk di database
	_, err := db.Exec("UPDATE PRODUK SET NAMA_PRODUK=$1, HARGA_PRODUK=$2, STOK_PRODUK=$3, FOTO_PRODUK=$4, KONDISI_PRODUK=$5, BERAT_PRODUK=$6, UKURAN_PRODUK=$7, DESKRIPSI_PRODUK=$8, VARIAN_PRODUK=$9 WHERE KODE_PRODUK=$10",
		product.NamaProduk, product.HargaProduk, product.StokProduk, product.FotoProduk, product.KondisiProduk, product.BeratProduk, product.UkuranProduk, product.Deskripsi, product.VarianProduk, kodeProduk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

// Handler untuk menghapus produk berdasarkan Kode Produk
func DeleteProduct(c echo.Context) error {
	kodeProduk := c.Param("kode_produk")

	// Hapus produk dari database
	_, err := db.Exec("DELETE FROM PRODUK WHERE KODE_PRODUK=$1", kodeProduk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
