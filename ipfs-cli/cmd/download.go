package cmd

import (
	"fmt"
	"io"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var outputPath = "/home/aaioet/ioet/EPN/"
var gatewayURL = "http://127.0.0.1:8081/ipfs/" 

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
	outFile, err := os.Create(outputPath+cid+".tif")
	if err != nil {
		return fmt.Errorf("could not create output file: %v", err)
	}
	defer outFile.Close()

	// Copy the content from IPFS to the local file
	_, err = io.Copy(outFile, reader)
	if err != nil {
		return fmt.Errorf("could not copy content to output file: %v", err)
	}

	// Print the output file path and the URL to view the file on the local gateway
	docURL := gatewayURL + cid
	fmt.Printf("File saved locally at %s\n", outputPath)
	fmt.Printf("View the file on the local gateway:  %s\n", docURL)
	return  nil
}
