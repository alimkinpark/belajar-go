package main

import "fmt"

func main() {
	// bentuk 1: for klasik
	for i := 0; i < 3; i++ {
		fmt.Println("klasik", i)
	}

	// bentuk 2: for seperti while
	n := 3
	for n > 0 {
		fmt.Println("while", n)
		// TODO 1: turunkan n, kalau tidak loop-nya infinite (Ctrl+C untuk stop)
	}

	// bentuk 3: for range
	buahly := []string{"go", "lagi", "asik"}
	// ponytail: blank assign biar stub compile, hapus kalau TODO 2 sudah pakai var-nya
	_ = buahly

	// TODO 2: range buahly, print index dan nilainya

	// if dengan init statement
	// TODO 3: tulis if yang cek len(buahly), deklarasi variabel di init statement-nya:
	//  if v := len(buahly); ...

	// switch
	// TODO 4: switch tanpa kondisi (switch { case ... }) — cek angka dari variabel,
	//  case pertama, kedua, default
}