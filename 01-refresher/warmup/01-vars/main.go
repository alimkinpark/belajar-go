package main

import "fmt"

func main() {
	var i int
	var s string
	var b bool
	var f float64

	// TODO 1: print semua var di atas, lihat zero value tiap tipe
	fmt.Println(i, s, b, f)

	x := 10
	const y = 20
	fmt.Println(x, y)

	// TODO 2: tambah x dengan y, simpan ke var baru, print hasilnya
	sum := x + y
	fmt.Println(sum)

	// TODO 3: buat var string dengan := , lalu deklarasi ulang satu var pakai (
	//  var name string di dalam blok, buktikan compile error kalau redeclare dengan :=
	//  (hapus baris yang bikin error setelah lihat pesannya)
	name := "Liam"
	// name := "x"
	// {
	// 	var name string
	// 	fmt.Println(name)
	// }
	fmt.Println(name)
}
