# go-container

A container runtime implementation in Go that demonstrates core containerization concepts including process isolation, namespace management, and filesystem virtualization.

## Project Overview

Core Features:
- Process isolation using Linux namespaces (UTS, PID, MNT)
- Root filesystem management
- Container lifecycle management
- Basic process execution

## Development Setup

### Prerequisites

#### Apple Silicon (M1/M2/M3) Mac:
```bash
brew install --cask vmware-fusion
brew install vagrant
vagrant plugin install vagrant-vmware-desktop
```

#### Intel Mac/Linux:
```bash
# Mac
brew install virtualbox
brew install vagrant

# Linux (Ubuntu/Debian)
sudo apt-get install virtualbox vagrant
```

#### Windows:
1. Install [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
2. Install [Vagrant](https://www.vagrantup.com/downloads)

### Setup Notes

#### Apple Silicon Mac:
Initial `vagrant up` process:
1. Downloads Ubuntu ARM64 box (~1GB)
2. System package installation (~70MB)
3. VM provisioning takes 5-10 minutes
4. "Waiting for VM to receive an address" message appears during normal boot sequence

#### Intel Mac:
1. Requires Oracle/VirtualBox approval in System Settings → Security & Privacy
2. Downloads Ubuntu x86_64 box (~1GB)

#### Windows/Linux:
1. Requires virtualization enabled in BIOS
2. Downloads Ubuntu x86_64 box (~1GB)

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
├── Makefile            # Build commands
├── README.md
└── rootfs/            # Container root filesystem
```

### Development Workflow

1. Clone repository:
```bash
git clone https://github.com/yourusername/go-container.git
cd go-container
```

2. Start VM:
```bash
vagrant up
```

3. Connect to VM:
```bash
vagrant ssh
cd ~/go-container
```

4. Build and run:
```bash
make build
sudo ./bin/container run /bin/sh
```

### VM Management Commands
```bash
vagrant halt     # Stop VM
vagrant destroy  # Remove VM
vagrant up       # Start VM
vagrant ssh      # Connect to VM
vagrant reload   # Restart VM with new configuration
```

### Troubleshooting

1. VM Start Failures:
- Verify virtualization enabled in BIOS
- Check security settings for VirtualBox (Intel Mac)
- Confirm >2GB free RAM, >10GB disk space

2. Shared Folder Issues:
- Check file permissions
- Verify VMware Tools/VirtualBox Guest Additions installation

## Implementation Details

1. Process Isolation:
- UTS Namespace: Hostname isolation
- PID Namespace: Process ID isolation
- Mount Namespace: Filesystem isolation

2. Root Filesystem Management:
- Root filesystem setup
- Mount point handling
- Filesystem isolation

3. Process Execution:
- Isolated command execution
- Resource management
- Security boundaries

## Contributing

1. Fork repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Create Pull Request

## License

MIT License - See LICENSE file

## Acknowledgments

Based on container runtime fundamentals from Docker and Linux container implementations.