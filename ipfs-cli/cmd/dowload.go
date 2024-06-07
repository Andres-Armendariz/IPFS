package cmd

import (
	"fmt"
	"io"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var outputPath = "/home/aaioet/" //cambiar con la ruta donde se desea descargar

// DownloadFileFromCID downloads a file from IPFS using its CID and saves it locally
func DownloadFile(cid string) error {
	// Connect to the local IPFS node
	sh := shell.NewShell("localhost:5001")

	// Get the file from IPFS
	reader, err := sh.Cat(cid)
	if err != nil {
		return fmt.Errorf("could not get file from IPFS: %v", err)
	}
	defer reader.Close()

	// Create the output file
	outFile, err := os.Create(outputPath+cid+".txt")
	if err != nil {
		return fmt.Errorf("could not create output file: %v", err)
	}
	defer outFile.Close()

	// Copy the content from IPFS to the local file
	_, err = io.Copy(outFile, reader)
	if err != nil {
		return fmt.Errorf("could not copy content to output file: %v", err)
	}

	fmt.Printf("File saved locally at %s\n", outputPath)
	return nil
}
