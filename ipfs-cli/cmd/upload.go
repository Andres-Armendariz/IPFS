package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

// UploadFile uploads a file to IPFS and adds it to MFS
func UploadFile(filePath string, mfsPath string) error {

	// Convert file paths to a format compatible with the OS
	filePath = filepath.ToSlash(filePath)

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
		destPath = filepath.ToSlash(filepath.Join(mfsPath, filepath.Base(filePath)))
	}

	// Check if file already exists and generate a new name if necessary
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

// getUniqueMFSPath checks if the file exists in the MFS path and generates a unique path
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
