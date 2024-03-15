package main

import "fmt"

func main() {
    const magicNumber = 42
    var number int

    // Contoh penggunaan magic number
    number = magicNumber + 5
    fmt.Println("Hasil penambahan magic number dan 5:", number)

    // Penggunaan magic number tanpa penjelasan
    number = 1000
    fmt.Println("Magic number tanpa penjelasan:", number)

    // Penggunaan magic number di dalam loop
    for i := 0; i < 5; i++ {
        number = 42 * i
        fmt.Println("Hasil perkalian magic number dan", i, ":", number)
    }
}
