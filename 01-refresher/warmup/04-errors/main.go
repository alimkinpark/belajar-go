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
	// TODO 1: loop id 1-4, panggil cari(id), print error-nya

	// TODO 2: pakai errors.Is untuk cek apakah error sama dengan ErrNotFound
	//  (perhatikan: id 3 tetap cocok walau dibungkus %w)

	// TODO 3: pakai errors.As untuk ekstrak *ValidationError dari error id 2,
	//  print e.Field
}