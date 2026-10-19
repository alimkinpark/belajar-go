package main

import (
	"fmt"
	"io"
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

// 1: method Speak() string untuk Dog — balikin "guk"
func (d Dog) Speak() string {
	return "guk"
}

// 2: method Speak() string untuk Robot — balikin nama + ": bip"
func (r Robot) Speak() string {
	return fmt.Sprintf("%s: bip", r.Name)
}

//	 3: fungsi yang terima []Speaker, loop & print tiap Speak()
//
//		(perhatikan: tahu Dog maupun Robot tanpa import masing-masing)
func SpeakAll(speakers []Speaker) {
	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
	}
}

func main() {
	zoo := []Speaker{}
	// TODO 4: isi zoo dengan Dog{} dan Robot{Name: "R2"}, panggil fungsi TODO 3
	zoo = append(zoo, Dog{}, Robot{Name: "R2"})
	SpeakAll(zoo)

	// Pattern io.Reader — interface paling umum di Go stdlib
	//  5: buat io.Reader dari string dengan strings.NewReader
	//  lalu io.ReadAll, print hasilnya (ini cara baca apapun di Go:
	//  file, network, HTTP body — semua lewat io.Reader)
	r := strings.NewReader("apa yang kamu dengar?")
	data, err := io.ReadAll(r)

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("data: %s\n", data)
	}

	// belajar tambahan assertion Speaker ke Dog, Robot, dan type lain
	var s Speaker = Dog{}
	// s.Speak() // ini aman, karena Dog punya Speak()
	v, ok := s.(Robot) // assertion ke Robot, tapi s adalah Dog
	if ok {
		fmt.Println("s adalah Robot, Speak():", v.Speak())
	} else {
		fmt.Println("s bukan Robot, tapi:", s.Speak())
	}
}
