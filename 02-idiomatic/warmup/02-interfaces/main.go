package main

import (
	"fmt"
	"strings"
)

// Interface kecil = idiom Go. Tanpa kata "implements" — implisit.
type Speaker interface {
	Speak() string
}

type Dog struct{}
type Robot struct {
	Name string
}

// TODO 1: method Speak() string untuk Dog — balikin "guk"
// TODO 2: method Speak() string untuk Robot — balikin nama + ": bip"

// TODO 3: fungsi yang terima []Speaker, loop & print tiap Speak()
//  (perhatikan: tahu Dog maupun Robot tanpa import masing-masing)

func main() {
	zoo := []Speaker{}
	// TODO 4: isi zoo dengan Dog{} dan Robot{Name: "R2"}, panggil fungsi TODO 3

	// Pattern io.Reader — interface paling umum di Go stdlib
	// TODO 5: buat io.Reader dari string dengan strings.NewReader
	//  lalu io.ReadAll, print hasilnya (ini cara baca apapun di Go:
	//  file, network, HTTP body — semua lewat io.Reader)
	r := strings.NewReader("")
	_ = r
	_ = zoo
	_ = fmt.Sprint
	
	// TODO 5 petunjuk: fmt.Sprint di atas hanya placeholder biar import fmt tidak error.
	//  Hapus baris `_ = fmt.Sprint` setelah TODO 1/2 pakai fmt.Println.
}