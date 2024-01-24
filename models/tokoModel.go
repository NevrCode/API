package models

import (
	"time"
)

type Toko struct {
	ID_toko            string
	Nama_toko          string
	Alamat             string
	Desc               string
	Jam_operasional    time.Time
	Status_verif_Toko  int
	Gambar_toko        string
	Status_operasional int
	ID_user            string
	ID_kantin          string
	Date_delete        time.Time
}
