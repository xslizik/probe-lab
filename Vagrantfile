# -*- mode: ruby -*-
# vi: set ft=ruby :

ansible_groups = {
  "servers" => ["server0", "server1" "server2"]
}

server_ip_list = ["10.10.3.19", "10.10.46.58", "10.10.170.185"]
kali_ip = "10.10.10.66"
cidr_netmask = "255.255.0.0"  

timestamp = Time.now.strftime("%Y%m%d%H%M%S")

Vagrant.configure("2") do |config|
  
  (0..2).each do |i|
    config.vm.define "server#{i}" do |device|
      device.vm.hostname = "server#{i}"
      device.vm.box = "generic/debian12"
      device.vm.provider "virtualbox" do |vb|
        vb.name = "probe-lab-server#{i}-#{timestamp}"
        vb.memory = 2048
        vb.cpus = 1           
      end
      device.vm.synced_folder ".",
        "/vagrant",
        type: "rsync",
        rsync__exclude: ".git/"
      device.vm.network "private_network",
        virtualbox__intnet: "internal-network",
        ip: "#{server_ip_list[i]}",
        netmask: "#{cidr_netmask}"
      device.vm.provision "ansible_local" do |ansible|
        ansible.playbook = "provisioning/playbook.yml"
        ansible.config_file = "provisioning/ansible.cfg"
        ansible.galaxy_role_file = "provisioning/requirements.yml"
        ansible.groups = ansible_groups
        ansible.limit = "server#{i}"
      end
    end
  end

  config.vm.define "kali" do |device|
    device.vm.hostname = "kali"
    device.vm.box = "kalilinux/rolling"
    device.vm.provider "virtualbox" do |vb|
      vb.name = "probe-lab-kali-#{timestamp}"
      vb.gui = true
      vb.memory = 4096
      vb.cpus = 2
      vb.customize ["modifyvm", :id, "--clipboard", "bidirectional"]
      vb.customize ["modifyvm", :id, "--draganddrop", "bidirectional"]
      vb.customize ["modifyvm", :id, "--vram", "128"]
    end
    device.vm.synced_folder ".",
        "/vagrant",
        type: "rsync",
        rsync__exclude: ".git/"
    device.vm.network "private_network",
      virtualbox__intnet: "internal-network",
      ip: "#{kali_ip}",
      netmask: "#{cidr_netmask}"
    device.vm.provision "ansible_local" do |ansible|
      ansible.playbook = "provisioning/playbook.yml"
      ansible.config_file = "provisioning/ansible.cfg"
      ansible.galaxy_role_file = "provisioning/requirements.yml"
      ansible.groups = ansible_groups
      ansible.limit = "kali"
    end
  end

end