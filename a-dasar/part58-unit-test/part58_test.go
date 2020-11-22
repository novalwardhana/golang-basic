package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var kubus Kubus = Kubus{5}
var volumeRight float64 = 125
var luasPermukaanRight float64 = 150
var kelilingright float64 = 60

//go test part58.go part58_test.go -v
func TestHitungLuasPermukaan(t *testing.T) {
	if kubus.LuasPermukaan() != luasPermukaanRight {
		t.Errorf("Luas permukaan seharusnya: %.2f", luasPermukaanRight)
	}
}

func TestHitungLuasPermukaanAssert(t *testing.T) {
	assert.Equal(t, kubus.LuasPermukaan(), luasPermukaanRight, "Hasil belum sama")
}

func TestHitungVolume(t *testing.T) {
	if kubus.Volume() != volumeRight {
		t.Errorf("Volume seharusnya: %2.f", volumeRight)
	}
}

func TestHitungVolumeAssert(t *testing.T) {
	assert.Equal(t, kubus.Volume(), volumeRight, "Hasil belum sama")
}

func TestKeliling(t *testing.T) {
	if kubus.Keliling() != kelilingright {
		t.Errorf("Keliling seharusnya: %2.f", kelilingright)
	}
}

func TestKelilingAssert(t *testing.T) {
	assert.Equal(t, kubus.Keliling(), kelilingright, "Hasil belum sama")
}

//go test part58.go part58_test.go -bench=.
func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.LuasPermukaan()
	}
}

func BenchmarkHitungVolume(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Volume()
	}
}

func BenchmarkHitungKeliling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Keliling()
	}
}
