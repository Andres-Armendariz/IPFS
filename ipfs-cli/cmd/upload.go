package cmd

import (
    "fmt"
    "os"
    "github.com/ipfs/go-ipfs-api"
)

sh := shell.NewShell("localhost:5001")
file, err := os.Open(filename)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error al abrir el archivo: %s\n", err)
    return
}
defer file.Close()

// Add the file to IPFS
hash, err := sh.Add(file)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error al subir el archivo: %s\n", err)
    return
}
fmt.Printf("Archivo subido con CID: %s\n", hash)
// }
