# go-container

A lightweight container runtime implementation in Go, demonstrating the fundamentals of containerization by building a simplified Docker-like system from scratch.

## Development Setup

### Prerequisites
- [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
- [Vagrant](https://www.vagrantup.com/downloads)

### Setting Up Development Environment

1. Clone the repository
```bash
git clone https://github.com/yourusername/go-container.git
cd go-container
```

2. Start the Vagrant VM
```bash
vagrant up
```

3. SSH into the VM
```bash
vagrant ssh
```

4. Navigate to project directory
```bash
cd ~/go-container
```

### Project Structure
```
go-container/
├── cmd/
│   └── container/
│       └── main.go      # Entry point
├── internal/
│   ├── container/
│   │   ├── container.go # Core container logic
│   │   └── namespace.go # Namespace operations
│   └── fs/
│       └── rootfs.go    # Filesystem operations
├── Vagrantfile          # Vagrant configuration
├── go.mod
├── Makefile            # Build and run commands
├── README.md
└── rootfs/            # Directory for container root filesystem
```

### Building and Running

Inside the Vagrant VM:

```bash
# Build the project
make build

# Run a command in container
sudo ./bin/container run /bin/bash
```

### Development Workflow

1. Edit code on your local machine using your preferred editor
2. Code will automatically sync to the VM through the shared folder
3. Build and test inside the VM
4. Use `vagrant halt` to stop the VM when done
5. Use `vagrant destroy` to remove the VM completely

## Implementation Details

The container runtime demonstrates:
- Process isolation using Linux namespaces (UTS, PID, MNT)
- Root filesystem management
- Basic process execution