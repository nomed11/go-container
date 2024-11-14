package container

import (
	"syscall"
)

// NamespaceOp handles container namespace operations
type NamespaceOp struct {
	// Namespace flags for the container
	flags uintptr
}

// NewNamespaceOp creates a new namespace operator with default flags
func NewNamespaceOp() *NamespaceOp {
	return &NamespaceOp{
		flags: syscall.CLONE_NEWUTS | // New UTS namespace (hostname)
			syscall.CLONE_NEWPID | // New PID namespace
			syscall.CLONE_NEWNS | // New mount namespace
			syscall.CLONE_NEWNET, // New network namespace
	}
}

// Flags returns the namespace flags
func (n *NamespaceOp) Flags() uintptr {
	return n.flags
}

// AddFlag adds a namespace flag
func (n *NamespaceOp) AddFlag(flag uintptr) {
	n.flags |= flag
}

// RemoveFlag removes a namespace flag
func (n *NamespaceOp) RemoveFlag(flag uintptr) {
	n.flags &^= flag
}

// HasFlag checks if a specific flag is set
func (n *NamespaceOp) HasFlag(flag uintptr) bool {
	return n.flags&flag == flag
}