package balok

type Balok struct {
	Panjang int
	Lebar   int
	Tinggi  int
}

func (b *Balok) LuasPermukaan() int {
	result := (2 * b.Panjang * b.Lebar) + (2 * b.Lebar * b.Tinggi) + (2 * b.Panjang * b.Tinggi)
	return result
}

func (b *Balok) Volume() int {
	result := b.Panjang * b.Lebar * b.Tinggi
	return result
}

func (b *Balok) Upsize(s int) {
	b.Panjang = s * b.Panjang
	b.Lebar = s * b.Lebar
	b.Tinggi = s * b.Tinggi
}
