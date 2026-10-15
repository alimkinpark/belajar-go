# Catatan Format Verb `fmt`

Verb yang umum dipakai dalam pengembangan Go. Semua contoh runnable:
simpan ke file, tambah `import "fmt"`, jalankan.

## Tabel verb

| Verb   | Untuk                                | Contoh                                    | Hasil                        |
|--------|--------------------------------------|-------------------------------------------|------------------------------|
| `%d`   | integer (desimal)                    | `fmt.Sprintf("id %d", 3)`                  | `id 3`                       |
| `%s`   | string                               | `fmt.Sprintf("halo %s", "liam")`           | `halo liam`                  |
| `%v`   | value apa pun, format default        | `fmt.Sprintf("%v", User{"liam", 25})`      | `{liam 25}`                  |
| `%+v`  | struct + nama field                  | `fmt.Sprintf("%+v", User{"liam", 25})`     | `{Nama:liam Umur:25}`        |
| `%#v`  | Go syntax literal (debug)            | `fmt.Sprintf("%#v", User{"liam", 25})`     | `main.User{Nama:"liam", Umur:25}` |
| `%q`   | string ter-quoted + escape           | `fmt.Sprintf("%q", "baris\nbaru")`         | `"baris\nbaru"`              |
| `%t`   | boolean                              | `fmt.Sprintf("ok %t", true)`               | `ok true`                    |
| `%f`   | float                                | `fmt.Sprintf("%f", 3.14159)`               | `3.141590` (6 digit default) |
| `%.2f` | float + presisi 2 digit              | `fmt.Sprintf("%.2f", 3.14159)`             | `3.14`                       |
| `%e`   | float notasi ilmiah                  | `fmt.Sprintf("%e", 12345.6789)`            | `1.234568e+04`               |
| `%x`   | hex (bytes, hash)                    | `fmt.Sprintf("%x", []byte("go"))`          | `676f`                       |
| `%p`   | pointer address (debug)              | `fmt.Sprintf("%p", &x)`                    | `0xc000012345`               |
| `%%`   | literal persen                       | `fmt.Sprintf("diskon %d%%", 20)`           | `diskon 20%`                 |
| `%w`   | **wrap error** (HANYA di Errorf)     | `fmt.Errorf("cari: %w", ErrNotFound)`      | chain utuh utk `errors.Is/As` |

## Detail per kasus pemakaian umum

### `%v` vs `%+v` vs `%#v` — print struct

```go
type User struct {
    Nama string
    Umur int
}

u := User{"liam", 25}
fmt.Println(u)   // {liam 25}          <- %v implisit
fmt.Printf("%+v\n", u)  // {Nama:liam Umur:25}  <- HANYA untuk debug
fmt.Printf("%#v\n", u)  // main.User{Nama:"liam", Umur:25}
```

Logs/HTTP response? Jangan `%#v` — tidak stabil dan bocor nama package/field
internal. Response API eksplisit marshal ke JSON.

### `%w` — error chain (paling penting untuk backend)

```go
var ErrNotFound = errors.New("data tidak ditemukan")

func cari(id int) error {
    dbErr := fmt.Errorf("query id %d gagal: %w", id, ErrNotFound)
    return dbErr
}

err := cari(3)
errors.Is(err, ErrNotFound)  // true — chain utuh walau dibungkus 2x
```

- `%w` → chain utuh, bisa `errors.Is`/`errors.As`
- `%v` di Errorf → chain PUTUS, `errors.Is` = false
- Aturan: error harus bisa dicek pemanggil → `%w`; detail disembunyikan → `%v`

### `%.2f` — angka rupiah/desimal

```go
harga := 15000.555
fmt.Printf("Rp %.2f\n", harga)  // Rp 15000.56
```

### `%q` — debugging string terlihat karakter tersembunyi

```go
input := "liam\t	gojek\n"
fmt.Printf("input=%q\n", input)
// input="liam\t	goje\n"  <- whitespace terlihat eksplisit
```

Berguna saat validasi input user — spasi/tab tersembunyi langsung kelihatan.

### Width & presisi

```go
fmt.Printf("[%8.2f]\n", 3.14)    // [    3.14]   <- padding spasi kiri, total 8 char
fmt.Printf("[%08.2f]\n", 3.14)   // [00003.14]   <- padding nol
fmt.Printf("[%-8d]\n", 42)       // [42      ]   <- rata kiri
fmt.Printf("[%8d]\n", 42)        // [      42]   <- rata kanan (default)
```

Pakai saat output tabel kolom di CLI.

## Yang jarang kepakai (skip, lihat `go doc fmt` kalau perlu)

`%b` (binary), `%o` (octal), `%c` (rune char), `%U` (unicode code point),
`%g` (compact float), `%T` (tipe variabel — debug tipe).

## Gotcha

1. `Println` vs `Printf`: `fmt.Println("x=", x)` — spasi otomatis antar arg,
   `fmt.Printf("x=%d\n", x)` — format eksplisit, **tambah `\n` sendiri**.
2. Mismatch verb/tipe tidak panic, cuma output aneh: `fmt.Printf("%d", "abc")`
   → `%!d(string=abc)`. Cek output `%!` = tanda verb salah.
3. `Errorf` + `%w` cuma berlaku satu `%w` per error (Go <1.20). Go 1.20+ boleh
   beberapa `%w`.

## Referensi

- `go doc fmt`
- https://pkg.go.dev/fmt