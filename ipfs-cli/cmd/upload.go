package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

// UploadFileToMFS uploads a file to IPFS and adds it to MFS
func UploadFile(filePath string, mfsPath string) error {
	// Connect to the local IPFS node
	sh := shell.NewShell("localhost:5001")

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Add the file to IPFS
	cid, err := sh.Add(file)
	if err != nil {
		return fmt.Errorf("could not add file to IPFS: %v", err)
	}

	// Create a context
	ctx := context.Background()

	// Check and create the MFS path if it doesn't exist
	err = createMFSPath(ctx, sh, mfsPath)
	if err != nil {
		return fmt.Errorf("could not create MFS path: %v", err)
	}

	// Append the file name to the MFS path if it is a directory
	destPath := mfsPath
	if isDir(ctx, sh, mfsPath) {
		destPath = filepath.Join(mfsPath, filepath.Base(filePath))
	}

	// Add the file to MFS
	err = sh.FilesCp(ctx, "/ipfs/"+cid, destPath)
	if err != nil {
		return fmt.Errorf("could not copy file to MFS: %v", err)
	}

	fmt.Printf("File added to MFS at %s with CID %s\n", destPath, cid)
	return nil
}

// createMFSPath creates the specified MFS path if it doesn't exist
func createMFSPath(ctx context.Context, sh *shell.Shell, mfsPath string) error {
	parts := strings.Split(mfsPath, "/")
	currentPath := ""

	for _, part := range parts {
		if part == "" {
			continue
		}
		currentPath += "/" + part
		if _, err := sh.FilesStat(ctx, currentPath); err != nil {
			if err := sh.FilesMkdir(ctx, currentPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// isDir checks if the given MFS path is a directory
func isDir(ctx context.Context, sh *shell.Shell, mfsPath string) bool {
	stat, err := sh.FilesStat(ctx, mfsPath)
	if err != nil {
		return false
	}
	return stat.Type == "directory"
}
