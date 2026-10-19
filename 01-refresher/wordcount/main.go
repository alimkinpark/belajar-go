package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

	// TODO 1: loop sc.Scan(), tiap baris tambah s.lines
	//  hitung kata: strings.Fields (ponytail: paling sederhana, upgrade ke
	//  unicode spatial kalau nanti butuh—cukup strings.Fields untuk sekarang)

	for sc.Scan() {
		s.lines++
		s.words += len(strings.Fields(sc.Text()))
	}
	// TODO 2: cek sc.Err() setelah loop, print error ke stderr kalau ada
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// TODO 3: print hasil:
	//  flag -w -> hanya kata, -l -> hanya baris, tanpa flag -> keduanya
	//  pakai *w dan *l (flag.Bool balikin pointer)
	if *w {
		fmt.Println(s.words)
	} else if *l {
		fmt.Println(s.lines)
	} else {
		fmt.Printf("%d %d\n", s.lines, s.words)
	}
}
