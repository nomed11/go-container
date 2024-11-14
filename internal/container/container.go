package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Container represents a Linux container instance
type Container struct {
	ID          string
	Command     []string
	RootFS      string
	NamespaceOp *NamespaceOp
}

// NewContainer creates a new container instance
func NewContainer(command []string) (*Container, error) {
	return &Container{
		ID:          generateID(),
		Command:     command,
		RootFS:      "rootfs",
		NamespaceOp: NewNamespaceOp(),
	}, nil
}

// Run executes the container process
func (c *Container) Run() error {
	fmt.Printf("Creating container [%s]\n", c.ID)

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, c.Command...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: c.NamespaceOp.Flags(),
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running container: %v", err)
	}

	return nil
}

// Child executes the actual container process
func (c *Container) Child() error {
	fmt.Printf("Setting up container environment [%s]\n", c.ID)

	if err := c.setupRootFS(); err != nil {
		return fmt.Errorf("error setting up rootfs: %v", err)
	}

	// Set hostname to container ID
	if err := syscall.Sethostname([]byte(c.ID)); err != nil {
		return fmt.Errorf("error setting hostname: %v", err)
	}

	cmd := exec.Command(c.Command[0], c.Command[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing command in container: %v", err)
	}

	return nil
}

// setupRootFS configures the container's root filesystem
func (c *Container) setupRootFS() error {
	// First, ensure we're in a new mount namespace
	must(syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""))

	// Mount our rootfs
	must(syscall.Mount(c.RootFS, c.RootFS, "", syscall.MS_BIND|syscall.MS_REC, ""))

	// Create temporary directory for old root
	must(os.MkdirAll(c.RootFS+"/oldrootfs", 0700))

	// Do pivot_root
	must(syscall.PivotRoot(c.RootFS, c.RootFS+"/oldrootfs"))

	// Change to new root
	must(os.Chdir("/"))

	// Mount proc filesystem
	must(os.MkdirAll("/proc", 0755))
	must(syscall.Mount("proc", "/proc", "proc", 0, ""))

	// Mount sys filesystem
	must(os.MkdirAll("/sys", 0755))
	must(syscall.Mount("sysfs", "/sys", "sysfs", 0, ""))

	// Mount dev filesystem
	must(os.MkdirAll("/dev", 0755))
	must(syscall.Mount("tmpfs", "/dev", "tmpfs", 0, ""))

	// Unmount old root
	must(syscall.Unmount("/oldrootfs", syscall.MNT_DETACH))
	must(os.RemoveAll("/oldrootfs"))

	return nil
}

// generateID creates a unique container ID
func generateID() string {
	return fmt.Sprintf("container-%d", os.Getpid())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}