package main

// Fungsi generik: type parameter [T], constraint comparable
func Pertama[T any](s []T) (T, bool) {
	// TODO 1: isi: kalau kosong balik zero value + false, kalau tidak elemen pertama + true
	var zero T
	return zero, false
}

// TODO 2: fungsi generik Ada[T comparable](s []T, v T) bool
//  — balikin true kalau v ada di s (pakai ==, makanya comparable)

// Constraint custom: gabung beberapa tipe
type Angka interface {
	~int | ~float64
}

// TODO 3: fungsi generik Jumlah[T Angka](s []T) T
//  — jumlah semua elemen (kenapa ~int, bukan int? ~ = termasuk tipe turunan)

func main() {
	// TODO 4: panggil Pertama dengan []int dan []string, print hasil + bool-nya
	// TODO 5: panggil Ada dengan []int cari 3 (true) dan 99 (false)
	// TODO 6: panggil Jumlah dengan []int dan []float64
}