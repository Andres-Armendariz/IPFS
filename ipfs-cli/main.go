package main

import (
	"fmt"
	"os"
	"ipfs-cli/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ipfs-cli",
	Short: "CLI para IPFS",
}

var uploadCmd = &cobra.Command{
	Use:   "upload [file] [mfs-path]",
	Short: "Sube un archivo a IPFS y lo agrega a MFS",
	Args:  cobra.ExactArgs(2),
	Run: func(command *cobra.Command, args []string) {
		filePath := args[0]
		mfsPath := args[1]
		err := cmd.UploadFile(filePath, mfsPath)
		if err != nil {
			fmt.Printf("Error al subir el archivo: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	// Agrega más comandos aquí si es necesario
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}



// pasar la direccion de donde esta el archivo como variable
// llamar al metodo sin argumento 
// crear metodo para descargar archivo
// crear metodo para listar 
// al momennto de crear el archivo, debe recibir un parametro (que puede ser otra funcion) para crear un directorio de como debe irse guardando
// crear metodo para instanciar el daemon