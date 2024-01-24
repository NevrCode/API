package models

import "time"

type Varian struct {
	Id_varian      string
	Nama_varian    string
	Harga_tambahan float32
	Id_menu        string
	Date_delete    time.Time
}
