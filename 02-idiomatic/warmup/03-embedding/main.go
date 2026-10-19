package main

type Base struct {
	ID int
}

// TODO 1: method Describe() string untuk Base — "ID: <id>"

// Logger embedded: method-nya "dipromosikan" ke struct luar
type Service struct {
	Base // embedding: tanpa nama field, langsung tipe
	Name string
}

func main() {
	s := Service{Base{1}, "svc"}

	// TODO 2: panggil s.Describe() langsung — method dari Base kepake
	//  tanpa s.Base.Describe() (method promotion)

	// TODO 3: print s.ID — field dari Base juga dipromosikan

	// Komposisi: Service "punya-a" Base, bukan "adalah-a" Base.
	// Tidak ada inheritance seperti Java/PHP.

	// TODO 4: buat Service kedua, override perilaku:
	//  buat method Describe di Service juga (shadow method Base),
	//  panggil dari s2 — cek mana yang kepake (shadowing)
	_ = s
}