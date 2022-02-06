package main

import "fmt"

type hitungMuatan interface {
	perahuMotor() muatanPerahu
	perahuPesiar() muatanPerahu
	perahuLayar() muatanPerahu
}

type inputKategory struct {
	jmlahMuatan int
}

type muatanPerahu struct {
	perahu     string
	keterangan string
	muatan     int
}

func (m muatanPerahu) hitungMuatan() int {
	if m.perahu == "motor" {
		return m.muatan * 1
	} else if m.perahu == "pesiar" {
		return m.muatan * 2
	} else if m.perahu == "layar" {
		return m.muatan * 3
	} else {
		return 0
	}
}

func (i inputKategory) perahuMotor() muatanPerahu {
	var data muatanPerahu
	data.perahu = "motor"
	data.keterangan = "Perahu motor"
	data.muatan = i.jmlahMuatan

	data.muatan = data.hitungMuatan()

	return data
}

func (i inputKategory) perahuPesiar() muatanPerahu {
	var data muatanPerahu
	data.perahu = "pesiar"
	data.keterangan = "Perahu motor"
	data.muatan = i.jmlahMuatan
	data.muatan = data.hitungMuatan()
	return data
}

func (i inputKategory) perahuLayar() muatanPerahu {
	var data muatanPerahu
	data.perahu = "layar"
	data.keterangan = "Perahu layar"
	data.muatan = i.jmlahMuatan
	data.muatan = data.hitungMuatan()
	return data
}

func main() {
	var muatan hitungMuatan
	muatan = inputKategory{2}
	fmt.Println("total muatan motor ", muatan.perahuMotor())
	fmt.Println("total muatan pesiar ", muatan.perahuPesiar())
	fmt.Println("total muatan layar ", muatan.perahuLayar())
}
