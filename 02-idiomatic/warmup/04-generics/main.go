package main

import "fmt"

// Fungsi generik: type parameter [T], constraint comparable
func Pertama[T any](s []T) (T, bool) {
	// TODO 1: isi: kalau kosong balik zero value + false, kalau tidak elemen pertama + true
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[0], true
}

// TODO 2: fungsi generik Ada[T comparable](s []T, v T) bool
//
//	— balikin true kalau v ada di s (pakai ==, makanya comparable)
func Ada[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

// Constraint custom: gabung beberapa tipe
type Angka interface {
	~int | ~float64
}

// TODO 3: fungsi generik Jumlah[T Angka](s []T) T
//
//	— jumlah semua elemen (kenapa ~int, bukan int? ~ = termasuk tipe turunan)
func Jumlah[T Angka](s []T) T {
	var total T
	for _, e := range s {
		total += e
	}
	return total
}

func main() {
	// TODO 4: panggil Pertama dengan []int dan []string, print hasil + bool-nya
	a, ok := Pertama([]int{1, 2, 3})
	fmt.Println(a, ok)
	b, ok := Pertama([]string{"a", "b", "c"})
	fmt.Println(b, ok)
	// TODO 5: panggil Ada dengan []int cari 3 (true) dan 99 (false)
	fmt.Println(Ada([]int{1, 2, 3}, 3))  // true
	fmt.Println(Ada([]int{1, 2, 3}, 99)) // false
	// TODO 6: panggil Jumlah dengan []int dan []float64
	fmt.Println(Jumlah([]int{1, 2, 3}))
	fmt.Println(Jumlah([]float64{1.1, 2.2, 3.3}))
}
