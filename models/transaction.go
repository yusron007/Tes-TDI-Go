package models

type Transaction struct {
	IDTransaksi     int      `json:"id_transaksi"`
	JumlahProduk    int      `json:"jumlah_produk"`
	TotalHarga      *float64 `json:"total_harga"`
	StatusTransaksi bool     `json:"status_transaksi"`
}
