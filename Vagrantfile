# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  # use Ubuntu 22.04 LTS
  config.vm.box = "ubuntu/jammy64"

  # configure VM resources
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "2048"
    vb.cpus = 2
    vb.name = "container-dev"
  end

  # sync project directory
  config.vm.synced_folder ".", "/home/vagrant/go-container"

  # install necessary packages and Go
  config.vm.provision "shell", inline: <<-SHELL
    apt-get update
    apt-get install -y build-essential git

    # install Go 1.21
    wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
    rm go1.21.6.linux-amd64.tar.gz

    # add Go to PATH for all users
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile

    # add Go to PATH for vagrant user
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/vagrant/.bashrc
    echo 'export GOPATH=/home/vagrant/go' >> /home/vagrant/.bashrc

    # create directory for root filesystem
    mkdir -p /home/vagrant/go-container/rootfs
  SHELL
end
