package main

import (
	"bufio"
	"flag"
	"os"
)

type stats struct {
	words int
	lines int
}

func main() {
	w := flag.Bool("w", false, "hitung kata saja")
	l := flag.Bool("l", false, "hitung baris saja")
	flag.Parse()

	var s stats
	sc := bufio.NewScanner(os.Stdin)
	// ponytail: blank assign biar stub compile, hapus kalau TODO 1 sudah pakai sc
	_ = sc

	// TODO 1: loop sc.Scan(), tiap baris tambah s.lines
	//  hitung kata: strings.Fields (ponytail: paling sederhana, upgrade ke
	//  unicode spatial kalau nanti butuh—cukup strings.Fields untuk sekarang)
	// TODO 2: cek sc.Err() setelah loop, print error ke stderr kalau ada

	// TODO 3: print hasil:
	//  flag -w -> hanya kata, -l -> hanya baris, tanpa flag -> keduanya
	//  pakai *w dan *l (flag.Bool balikin pointer)
	_ = s
	_ = w
	_ = l
}