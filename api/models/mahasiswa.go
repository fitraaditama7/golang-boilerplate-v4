package models

type Mahasiswa struct {
	ID    int    `json:"id"`
	Nim   string `validate:"required" json:"nim"`
	Nama  string `validate:"required" json:"nama"`
	Kelas string `validate:"required" json:"kelas"`
}
