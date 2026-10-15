package main

import "fmt"

type User struct {
	Nama string
	Umur int
}

func main() {
	arr := [3]int{1, 2, 3}
	slc := []int{1, 2, 3}

	//  1: print len(arr) dan len(slc), lalu append angka 4 ke slc,
	//  print lagi — cek perbedaan array (fixed) vs slice (dinamis)
	fmt.Println(len(arr), len(slc))
	slc = append(slc, 4)
	fmt.Println(len(arr), len(slc))

	m := map[string]int{"a": 1}

	//  2: tambah key "b" dengan nilai 2, lalu baca key "c" yang tidak ada,
	//  print hasilnya — lihat zero value map
	m["b"] = 2
	fmt.Println(m["c"]) // zero value int = 0
	fmt.Println(m)

	//  3: cek keberadaan key dengan bentuk dua nilai:
	//  v, ok := m["c"]
	v, ok := m["c"]
	fmt.Println(v, ok)

	u := User{"liam", 25}

	//  4: print struct u, lalu ubah Umur-nya jadi 26 dan print lagi
	fmt.Println(u)
	u.Umur = 26
	fmt.Println(u)
}
