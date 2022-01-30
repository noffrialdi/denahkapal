package main

import (
	"fmt"
	"log"
)

type PerahuMtr struct {
	Nama             string
	Kapasitasmanusia int
}

func TmbhPerahuMotor(m *PerahuMtr) {
	fmt.Printf("Tambah perahu di insert, nama perahu %s dan kapasitas manusia %d", m.Nama, m.Kapasitasmanusia)
}

func (p PerahuMtr) TmbhPerahuMotor() {
	fmt.Printf("Name: '%s', kapasitas manusia: %t\n", p.Nama, p.Kapasitasmanusia)
}

type PerahuPesiar struct {
	PerahuMtr
	Jumlahlantai       int
	Kapasitaskendaraan int
}

func (ps PerahuPesiar) TmbhPerahuMotor() {
	fmt.Printf("Perahu %s, Jumlah lantai %d, jumlah kendaraan %d", ps.Nama, ps.Jumlahlantai, ps.Kapasitaskendaraan)
}

type PerahuLayar struct {
	PerahuPesiar
	Jumlahlayar   int
	Jmlhpendayung int
}

func perahuLayer(layar int, dayung int) *PerahuLayar {
	temp := &PerahuLayar{
		PerahuPesiar{
			PerahuMtr{"Perahu Motor", 4},
			4,
			6,
		},
		layar,
		dayung,
	}
	return temp
}

func main() {
	perahuMotor := &PerahuMtr{
		"Perahu Motor",
		20,
	}
	TmbhPerahuMotor(perahuMotor)
	lyr2 := perahuLayer(8, 6)
	log.Println("layar 2 ", lyr2)
}
