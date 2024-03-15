package main

import (
    "fmt"
    "os"
)

func DetectMagicNumber(fileName string) (bool, error) {
    file, err := os.Open(fileName)
    if err != nil {
        return false, err
    }
    defer file.Close()

    var buffer [4]byte
    _, err = file.Read(buffer[:])
    if err != nil {
        return false, err
    }

    // Cek apakah buffer sesuai dengan magic number yang diinginkan
    magicNumber := [4]byte{0x2A, 0x00, 0x00, 0x00} // Representasi hexa dari 42 (0x2A)
    if buffer == magicNumber {
        return true, nil
    }

    return false, nil
}

func main() {
    fileName := "file.txt"
    isMagic, err := DetectMagicNumber(fileName)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if isMagic {
        fmt.Println("Magic number 42 found in", fileName)
    } else {
        fmt.Println("Magic number not found in", fileName)
    }
}
