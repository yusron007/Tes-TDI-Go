package models

type Product struct {
	KodeProduk    string  `json:"kode_produk"`
	KodeEtalase   string  `json:"kode_etalase"`
	NamaProduk    string  `json:"nama_produk"`
	HargaProduk   float64 `json:"harga_produk"`
	StokProduk    int     `json:"stok_produk"`
	FotoProduk    string  `json:"foto_produk"`
	KondisiProduk string  `json:"kondisi_produk"`
	BeratProduk   int     `json:"berat_produk"`
	UkuranProduk  string  `json:"ukuran_produk"`
	Deskripsi     string  `json:"deskripsi"`
	VarianProduk  string  `json:"varian_produk"`
}
