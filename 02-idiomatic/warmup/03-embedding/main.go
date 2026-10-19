package main

import "fmt"

type Base struct {
	ID int
}

// TODO 1: method Describe() string untuk Base — "ID: <id>"
func (b Base) Describe() string {
	return fmt.Sprintf("ID : %d", b.ID)
}

// Logger embedded: method-nya "dipromosikan" ke struct luar
type Service struct {
	Base // embedding: tanpa nama field, langsung tipe
	Name string
}

func (s Service) Describe() string {
	return fmt.Sprintf("ID: %d, %s", s.ID, s.Name)
}

func main() {
	s := Service{Base{1}, "svc"}

	// TODO 2: panggil s.Describe() langsung — method dari Base kepake
	//  tanpa s.Base.Describe() (method promotion)
	fmt.Println(s.Describe()) // ID: 1, svc (shadow Service menang)

	// TODO 3: print s.ID — field dari Base juga dipromosikan
	fmt.Println(s.ID) // 1

	// Komposisi: Service "punya-a" Base, bukan "adalah-a" Base.
	// Tidak ada inheritance seperti Java/PHP.

	// TODO 4: buat Service kedua, override perilaku:
	//  buat method Describe di Service juga (shadow method Base),
	//  panggil dari s2 — cek mana yang kepake (shadowing)
	s2 := Service{Base{2}, "svc2"}
	fmt.Println(s2.Describe()) // ID: 2, svc2
}
