package main

type User struct {
	Nama string
	Umur int
}

func main() {
	arr := [3]int{1, 2, 3}
	slc := []int{1, 2, 3}
	// ponytail: blank assign biar stub compile, hapus kalau TODO 1 sudah pakai var-nya
	_, _ = arr, slc

	// TODO 1: print len(arr) dan len(slc), lalu append angka 4 ke slc,
	//  print lagi — cek perbedaan array (fixed) vs slice (dinamis)

	m := map[string]int{"a": 1}
	// ponytail: blank assign biar stub compile, hapus kalau TODO 2 sudah pakai m
	_ = m

	// TODO 2: tambah key "b" dengan nilai 2, lalu baca key "c" yang tidak ada,
	//  print hasilnya — lihat zero value map

	// TODO 3: cek keberadaan key dengan bentuk dua nilai:
	//  v, ok := m["c"]

	u := User{"liam", 25}
	// ponytail: blank assign biar stub compile, hapus kalau TODO 4 sudah pakai u
	_ = u

	// TODO 4: print struct u, lalu ubah Umur-nya jadi 26 dan print lagi
}