package models

import "time"

type Etalase struct {
	Id_etalase   string
	Nama_etalase string
	Id_toko      string
	Data_delete  time.Time
}
