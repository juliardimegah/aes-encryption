package main

import (
    "crypto/aes"
    "crypto/cipher"
    "database/sql"
    "fmt"
)

func main() {
    // Buka koneksi ke database
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/mydatabase")
    if err != nil {
        panic(err)
    }

    // Buat cipher AEP dengan mode GCM
    key := []byte("rahasia")
    iv := []byte("iniiv")
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err)
    }

    // Enkripsi data
    plaintext := []byte("data yang akan dienkripsi")
    ciphertext := gcm.Seal(nil, iv, plaintext, nil)

    // Simpan data terenkripsi ke database
    stmt, err := db.Prepare("UPDATE users SET data = ? WHERE id = 1")
    if err != nil {
        panic(err)
    }
    _, err = stmt.Exec(ciphertext)
    if err != nil {
        panic(err)
    }

    // Dekripsi data
    plaintext, err = gcm.Open(nil, iv, ciphertext, nil)
    if err != nil {
        panic(err)
    }

    // Tampilkan data didekripsi
    fmt.Println(plaintext)

    // Tutup koneksi ke database
    db.Close()
}
