package main

import (
	"fmt"
	"os"
	"ipfs-cli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: my-ipfs-app <command> [<args>]")
		fmt.Println("Commands:")
		fmt.Println("  upload <file-path>   Upload a file to IPFS and add it to MFS")
		fmt.Println("  download <cid> <output-path>   Download a file from IPFS using its CID")
        fmt.Println("  list <path> List CIDs from a path")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "upload":
		if len(os.Args) != 3 {
			fmt.Println("Usage: IPFS upload <file-path>")
			os.Exit(1)
		}
		filePath := os.Args[2]

		err := cmd.UploadFile(filePath)
		if err != nil {
			fmt.Printf("Error uploading file: %v\n", err)
			os.Exit(1)
		}

	case "download":
		if len(os.Args) != 3 {
			fmt.Println("Usage: IPFS download <cid>")
			os.Exit(1)
		}
		cid := os.Args[2]
		err := cmd.DownloadFile(cid)
		if err != nil {
			fmt.Printf("Error downloading file: %v\n", err)
			os.Exit(1)
		}

    case "list":
        if len(os.Args) < 2 {
            fmt.Println("Usage: IPFS list-CID <mfs-path>")
            os.Exit(1)
        }
        mfsPath := os.Args[2]

        cids, err := cmd.ListCID(mfsPath)
        if err != nil {
            fmt.Printf("Error listing path: %v\n", err)
            os.Exit(1)
        }
        if len(cids) == 0 {
            fmt.Println("No CIDs found in  ", mfsPath)
            return
        }

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Commands:")
		fmt.Println("  upload <file-path>  Upload a file to IPFS and add it to MFS")
		fmt.Println("  download <cid>  Download a file from IPFS using its CID")
        fmt.Println("  list <path> List CIDs from a path")
		os.Exit(1)
	}
}
