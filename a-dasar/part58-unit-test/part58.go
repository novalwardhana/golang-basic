package main

type Kubus struct {
	Sisi float64
}

func (k Kubus) LuasPermukaan() float64 {
	hasil := k.Sisi * k.Sisi * 6
	return hasil
}

func (k Kubus) Volume() float64 {
	//hasil := math.Pow(k.Sisi, 3)
	hasil := k.Sisi * k.Sisi * k.Sisi
	return hasil
}

func (k Kubus) Keliling() float64 {
	hasil := float64(12) * k.Sisi
	return hasil
}
