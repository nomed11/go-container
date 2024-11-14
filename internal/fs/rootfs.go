package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// RootFS handles container root filesystem operations
type RootFS struct {
	Path string
}

// NewRootFS creates a new root filesystem handler
func NewRootFS(path string) *RootFS {
	return &RootFS{
		Path: path,
	}
}

// Setup prepares the root filesystem
func (r *RootFS) Setup() error {
	// Ensure root filesystem directory exists
	if err := os.MkdirAll(r.Path, 0755); err != nil {
		return fmt.Errorf("failed to create rootfs directory: %v", err)
	}

	// Setup basic mount points
	mounts := []struct {
		source string
		target string
		fstype string
		flags  uintptr
		data   string
	}{
		{"proc", "/proc", "proc", 0, ""},
		{"sysfs", "/sys", "sysfs", 0, ""},
		{"tmpfs", "/dev", "tmpfs", 0, ""},
	}

	for _, m := range mounts {
		target := filepath.Join(r.Path, m.target)
		
		// Create mount point
		if err := os.MkdirAll(target, 0755); err != nil {
			return fmt.Errorf("failed to create mount point %s: %v", target, err)
		}

		// Perform mount
		if err := syscall.Mount(m.source, target, m.fstype, m.flags, m.data); err != nil {
			return fmt.Errorf("failed to mount %s: %v", target, err)
		}
	}

	return nil
}

// Cleanup handles filesystem cleanup
func (r *RootFS) Cleanup() error {
	// Unmount in reverse order
	mounts := []string{"/dev", "/sys", "/proc"}
	for _, m := range mounts {
		target := filepath.Join(r.Path, m)
		if err := syscall.Unmount(target, 0); err != nil {
			return fmt.Errorf("failed to unmount %s: %v", target, err)
		}
	}

	return nil
}

// Exists checks if the root filesystem exists
func (r *RootFS) Exists() bool {
	_, err := os.Stat(r.Path)
	return err == nil
}