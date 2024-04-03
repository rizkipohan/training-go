package main

import "fmt"

type Ban int

const (
	BanKaret Ban = iota
	BanKayu
	BanBesi
)

type Roda struct {
	Ban Ban
}

type Pintu struct {
	bunyiKetuk string
	bunyiBuka  string
}

type Mobil struct {
	Roda       [4]Roda
	PintuKanan Pintu
	PintuKiri  Pintu
}

func (m *Mobil) GantiBan(indeksRoda int, jenisBan Ban) {
	if indeksRoda >= 0 && indeksRoda < 4 {
		m.Roda[indeksRoda].Ban = jenisBan
	} else {
		fmt.Println("Indeks roda tidak valid")
	}
}

func (p *Pintu) Bunyi(aksi string) {
	if aksi == "ketuk" {
		fmt.Println(p.bunyiKetuk)
	} else if aksi == "buka" {
		fmt.Println(p.bunyiBuka)
	} else {
		fmt.Println("Aksi tidak valid")
	}
}

func main() {

	var PintuKanan Pintu
	PintuKanan.bunyiKetuk = "Knock"
	PintuKanan.bunyiBuka = "Beep"

	var PintuKiri Pintu
	PintuKiri.bunyiKetuk = "Beep"
	PintuKiri.bunyiBuka = "Knock"

	var mobil Mobil
	mobil.PintuKanan = PintuKanan
	mobil.PintuKiri = PintuKiri

	mobil.GantiBan(0, BanKaret)

	mobil.GantiBan(1, BanKayu)

	mobil.GantiBan(2, BanBesi)

	mobil.GantiBan(3, BanKaret)

	mobil.PintuKanan.Bunyi("ketuk")
	mobil.PintuKanan.Bunyi("buka")

	mobil.PintuKiri.Bunyi("ketuk")
	mobil.PintuKiri.Bunyi("buka")
}
