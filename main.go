package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" // Menggunakan middleware Echo yang benar
	_ "github.com/lib/pq"

	"E-commerce/controllers" // Sesuaikan path dengan struktur direktori Anda
)

var db *sql.DB

func main() {
	// Konfigurasi koneksi database PostgreSQL
	dbConnStr := "user=postgres dbname=TDI sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dbConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Inisialisasi Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Inisialisasi controller untuk masing-masing entitas
	controllers.InitProductController(db)
	controllers.InitCartController(db)
	controllers.InitTransactionController(db)

	// Routes
	// Routes untuk Produk
	// Contoh konversi fungsi CreateProduct menjadi HandlerFunc
	e.POST("/products", func(c echo.Context) error {
		return controllers.CreateProduct(c)
	})

	// Lakukan hal yang sama untuk rute lainnya
	e.GET("/products", func(c echo.Context) error {
		return controllers.GetProducts(c)
	})

	e.GET("/products/:kode_produk", func(c echo.Context) error {
		return controllers.GetProduct(c)
	})

	e.PUT("/products/:kode_produk", func(c echo.Context) error {
		return controllers.UpdateProduct(c)
	})

	e.DELETE("/products/:kode_produk", func(c echo.Context) error {
		return controllers.DeleteProduct(c)
	})

	// Routes untuk Keranjang Belanja
	e.POST("/carts", func(c echo.Context) error {
		return controllers.AddToCart(c)
	})

	e.GET("/carts/:id_transaksi", func(c echo.Context) error {
		return controllers.GetCart(c)
	})

	e.DELETE("/carts/:id_transaksi/:kode_produk", func(c echo.Context) error {
		return controllers.RemoveFromCart(c)
	})

	// Routes untuk Transaksi
	e.POST("/checkout", func(c echo.Context) error {
		return controllers.Checkout(c)
	})

	e.GET("/transactions", func(c echo.Context) error {
		return controllers.GetTransactions(c)
	})

	e.GET("/transactions/:id_transaksi", func(c echo.Context) error {
		return controllers.GetTransaction(c)
	})

	// Start server
	fmt.Println("Server started at :8080")
	e.Start(":8080")
}
