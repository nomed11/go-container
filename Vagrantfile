# -*- mode: ruby -*-
# vi: set ft=ruby :

# Determine host architecture
host_arch = `uname -m`.strip

Vagrant.configure("2") do |config|
  # Detect architecture and set appropriate box
  if host_arch == "arm64"
    # M1/M2/M3 Macs
    config.vm.box = "bento/ubuntu-22.04-arm64"

    config.vm.provider "vmware_desktop" do |vmware|
      vmware.memory = "2048"
      vmware.cpus = 2
      vmware.gui = false
    end
  else
    # Intel machines
    config.vm.box = "ubuntu/jammy64"

    config.vm.provider "virtualbox" do |vb|
      vb.memory = "2048"
      vb.cpus = 2
      vb.name = "container-dev"
    end
  end

  # Sync the project directory
  config.vm.synced_folder ".", "/home/vagrant/go-container"

  # Install necessary packages and Go
  config.vm.provision "shell", inline: <<-SHELL
    apt-get update
    apt-get install -y build-essential git wget curl

    # Detect architecture for Go installation
    ARCH=$(uname -m)
    GO_ARCH="amd64"
    if [ "$ARCH" = "aarch64" ]; then
      GO_ARCH="arm64"
    fi

    # Install Go 1.21
    wget "https://go.dev/dl/go1.21.6.linux-$GO_ARCH.tar.gz"
    rm -rf /usr/local/go && tar -C /usr/local -xzf "go1.21.6.linux-$GO_ARCH.tar.gz"
    rm "go1.21.6.linux-$GO_ARCH.tar.gz"

    # Add Go to PATH for all users
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile

    # Add Go to PATH for vagrant user
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/vagrant/.bashrc
    echo 'export GOPATH=/home/vagrant/go' >> /home/vagrant/.bashrc

    # Create directory for root filesystem
    mkdir -p /home/vagrant/go-container/rootfs
  SHELL
end
