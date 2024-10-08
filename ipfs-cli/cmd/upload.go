package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

// UploadFile uploads a file to IPFS and adds it to MFS using a derived MFS path
func UploadFile(filePath string) error {
	
	// Convert file paths to a format compatible with the OS
	filePath = filepath.ToSlash(filePath)

	// Extract the base file name without extension
	fileName := filepath.Base(filePath)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

	// Generate the MFS path based on the file name
	mfsPath := generateMFSDirPath(fileName)

	// Connect to the local IPFS node
	sh := shell.NewShell("localhost:5001")

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Use IPFS command line to add the file with CIDv1
	cmd := exec.Command("ipfs", "add", "--cid-version=1",  "--hash=sha2-256", "--chunker=size-262144","-Q", filePath)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("could not add file to IPFS with CIDv1: %v", err)
	}

	cid := strings.TrimSpace(string(output))

	// Create a context
	ctx := context.Background()

	// Check and create the MFS path if it doesn't exist
	err = createMFSPath(ctx, sh, mfsPath)
	if err != nil {
		return fmt.Errorf("could not create MFS path: %v", err)
	}

	// Construct the destination path in MFS
	destPath := "/" + filepath.ToSlash(filepath.Join(mfsPath, filepath.Base(filePath)))

	// Check if the file already exists in MFS and generate a unique path if necessary
	destPath, err = getUniqueMFSPath(ctx, sh, destPath)
	if err != nil {
		return fmt.Errorf("could not get unique MFS path: %v", err)
	}

	// Add the file to MFS
	err = sh.FilesCp(ctx, "/ipfs/"+cid, destPath)
	if err != nil {
		return fmt.Errorf("could not copy file to MFS: %v", err)
	}

	fmt.Printf("File added to MFS at %s with CID %s\n", destPath, cid)
	return nil
}

// generateMFSDirPath generates an MFS directory path based on the file name
func generateMFSDirPath(fileName string) string {
	// Split the file name into parts
	parts := strings.Split(fileName, ".")

	// Join parts with slashes to create MFS directory path
	return strings.Join(parts, "/")
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

// getUniqueMFSPath checks if the file exists in the MFS path and generates a unique path if necessary
func getUniqueMFSPath(ctx context.Context, sh *shell.Shell, destPath string) (string, error) {
	dir := filepath.Dir(destPath)
	base := filepath.Base(destPath)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	for i := 1; ; i++ {
		newName := fmt.Sprintf("%s_version%d%s", name, i, ext)
		newPath := filepath.ToSlash(filepath.Join(dir, newName))
		if _, err := sh.FilesStat(ctx, newPath); err != nil {
			if strings.Contains(err.Error(), "file does not exist") {
				return newPath, nil
			}
			return "", err
		}
	}
}
