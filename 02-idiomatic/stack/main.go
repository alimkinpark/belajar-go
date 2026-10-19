package main

import "fmt"

// Latihan wajib Fase 2: Stack[T] generik
type Stack[T any] struct {
	items []T
}

// TODO 1: method Push(v T) — append ke items (pointer receiver)

// TODO 2: method Pop() (T, bool) — keluarkan item teratas.
//  Kalau kosong: zero value + false. Kalau tidak: item teratas + true.
//  Hati-hati: assign ulang slice, jangan cuma baca (s.items = items[:len-1])

// TODO 3: method Peek() (T, bool) — lihat teratas TANPA keluarkan

// TODO 4: method Len() int

func main() {
	s := &Stack[int]{}
	// TODO 5: Push 1, 2, 3 — print Len()
	// TODO 6: Pop 2x, print hasil + bool-nya (harusnya 3, true lalu 2, true)
	// TODO 7: Peek, print (harusnya 1, true)
	// TODO 8: Pop sampai kosong, Pop lagi — lihat zero value + false
	// TODO 9: buat Stack[string], Push "go", Pop — buktikan generik
	_ = s
	_ = fmt.Sprint
}
