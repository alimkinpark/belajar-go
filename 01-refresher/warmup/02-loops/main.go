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
		// 1: turunkan n, kalau tidak loop-nya infinite (Ctrl+C untuk stop)
		n--
	}

	// bentuk 3: for range
	buahly := []string{"go", "lagi", "asik"}

	// 2: range buahly, print index dan nilainya
	for i, v := range buahly {
		fmt.Println(i, v)
	}

	// if dengan init statement
	// 3: tulis if yang cek len(buahly), deklarasi variabel di init statement-nya:
	//  if v := len(buahly); ...
	if v := len(buahly); v > 0 {
		fmt.Println("len(buahly) > 0", v)
	}

	// switch
	// 4: switch tanpa kondisi (switch { case ... }) — cek angka dari variabel,
	//  case pertama, kedua, default
	switch {
	case n == 0:
		fmt.Println("n == 0")
	case n > 0:
		fmt.Println("n > 0")
	default:
		fmt.Println("n < 0")
	}
}
