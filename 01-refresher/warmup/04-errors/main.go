package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("data tidak ditemukan")

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validasi gagal di %s", e.Field)
}

func cari(id int) error {
	if id == 1 {
		return ErrNotFound
	}
	if id == 2 {
		// ponytail: pointer receiver karena errors.As butuh *ValidationError
		return &ValidationError{Field: "id"}
	}
	if id == 3 {
		return fmt.Errorf("gagal cari id %d: %w", id, ErrNotFound)
	}
	return nil
}

func main() {
	//  1: loop id 1-4, panggil cari(id), print error-nya
	for id := 1; id <= 4; id++ {
		err := cari(id)
		if err != nil {
			fmt.Println("id", id, "error:", err)
		} else {
			fmt.Println("id", id, "tidak ada error")
		}
	}

	//  2: pakai errors.Is untuk cek apakah error sama dengan ErrNotFound
	//  (perhatikan: id 3 tetap cocok walau dibungkus %w)
	for id := 1; id <= 4; id++ {
		err := cari(id)
		if errors.Is(err, ErrNotFound) {
			fmt.Println("id", id, "error is ErrNotFound")
		} else {
			fmt.Println("id", id, "error is NOT ErrNotFound")
		}
	}

	//  3: pakai errors.As untuk ekstrak *ValidationError dari error id 2,
	//  print e.Field
	for id := 1; id <= 4; id++ {
		err := cari(id)
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println("id", id, "error is ValidationError, field:", ve.Field)
		} else {
			fmt.Println("id", id, "error is NOT ValidationError")
		}
	}
}
