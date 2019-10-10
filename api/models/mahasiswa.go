package models

type Mahasiswa struct {
	ID    int    `json:"id"`
	Nim   string `validate:"required" json:"nim" bind:"required"`
	Nama  string `validate:"required" json:"nama" bind:"required"`
	Kelas string `validate:"required" json:"kelas" bind:"required"`
}
