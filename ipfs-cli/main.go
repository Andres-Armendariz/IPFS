package main

import (
    "os"
    "github.com/spf13/cobra"
    "fmt"
    "github.com/ipfs/go-ipfs-api"
)


var rootCmd = &cobra.Command{
    Use:   "ipfs-cli",
    Short: "CLI para IPFS",
}

var path = "/home/edvard/Pictures/1330715.png"

func Upload(filename string) {
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
}

var uploadCmd = &cobra.Command{
    Use:   "upload [file]",
    Short: "Sube un archivo a IPFS",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        filename := args[0]
        Upload(filename)
    },
}

func init() {
    rootCmd.AddCommand(uploadCmd)
    // Agrega más comandos
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

// pasar la direccion de donde esta el archivo como variable
// llamar al metodo sin argumento 
// crear metodo para descargar archivo
// crear metodo para listar 
// al momennto de crear el archivo, debe recibir un parametro (que puede ser otra funcion) para crear un directorio de como debe irse guardando
// crear metodo para instanciar el daemon