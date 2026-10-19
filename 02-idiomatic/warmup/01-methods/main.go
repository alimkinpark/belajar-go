package main

import "fmt"

type Counter struct {
	n int
}

//	 1: method dengan VALUE receiver — n kalikan 2
//
//		(perhatikan: value receiver tidak mengubah Counter asli)
func (c Counter) Double() Counter {
	c.n *= 2
	return c
}

//	 2: method dengan POINTER receiver — tambah 1 ke n
//
//		(pointer receiver mengubah Counter asli)
func (c *Counter) Add() {
	c.n++
}

// 3: method String() string — biarkan Counter bisa dipakai fmt.Println langsung
func (c Counter) String() string {
	return fmt.Sprintf("Counter: %d", c.n)
}

// 4: method dengan receiver pointer — reset n ke 0
func (c *Counter) Reset() {
	c.n = 0
}

func main() {
	c := Counter{}
	//  5: panggil Add 3x, print c — lihat hasilnya
	c.Add()
	c.Add()
	c.Add()
	fmt.Println(c)
	//  6: panggil Double, print c lagi — cek apakah n berubah (kenapa?)
	c.Double()
	fmt.Println(c)
	// untuk benar-benar mengubah c, kita harus assign hasil Double() ke c
	// c = c.Double()
	// fmt.Println(c)
	//  7: print c via fmt.Println — String() kepake otomatis
	fmt.Println(c)
	//  8: panggil Reset, print c
	c.Reset()
	fmt.Println(c)
}
