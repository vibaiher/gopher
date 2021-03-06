# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure(2) do |config|
  # Use Ubuntu 14.04
  config.vm.box = "ubuntu/trusty64"

  # SSH
  config.ssh.forward_agent = true

  # Private network for NFS
  config.vm.network :private_network, ip: "192.168.33.14"

  # NFS
  config.vm.synced_folder ".", "/vagrant", type: "nfs"

  # VirtualBox
  config.vm.provider "virtualbox" do |vb|
    vb.customize ["modifyvm", :id, "--memory", "2048"]
    vb.customize ["modifyvm", :id, "--cpus", 2]
  end

  # Provision with ansible
  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "provisioning/playbook.yml"
    ansible.inventory_path = "provisioning/default"
    ansible.limit = "functional_programming"
  end
end
