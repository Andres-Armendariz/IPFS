package cmd

import (
	"context"
	"fmt"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

// ListCID lists all CIDs within a given MFS path
func ListCID(mfsPath string) ([]string, error) {
	// Connect to the local IPFS node
	sh := shell.NewShell("localhost:5001")

	// Verify the connection
	if !sh.IsUp() {
		return nil, fmt.Errorf("IPFS node is not running")
	}

	// Create a context
	ctx := context.Background()

	// Ensure the path starts with a leading slash
	if !strings.HasPrefix(mfsPath, "/") {
		mfsPath = "/" + mfsPath
	}

	// List the directory in MFS
	entries, err := sh.FilesLs(ctx, mfsPath)
	if err != nil {
		return nil, fmt.Errorf("could not list directory in MFS: %v", err)
	}

	// Extract CIDs from the directory entries
	var cids []string
	for _, entry := range entries {
		stat, err := sh.FilesStat(ctx, mfsPath+"/"+entry.Name)
		if err != nil {
			return nil, fmt.Errorf("could not stat file %s: %v", entry.Name, err)
		}
		fmt.Printf("%s, CID: %s\n", entry.Name, stat.Hash)
		cids = append(cids, stat.Hash)
	}

	return cids, nil
}
