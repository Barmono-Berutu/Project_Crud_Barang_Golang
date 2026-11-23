package Models

type Barang struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
}
