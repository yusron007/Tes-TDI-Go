package models

type CartItem struct {
	KodeProduk   string `json:"kode_produk"`
	JumlahProduk int    `json:"jumlah_produk"`
}
