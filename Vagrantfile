# -*- mode: ruby -*-
# vi: set ft=ruby :

#server_ip_list = ["10.8.3.19", "10.86.46.58", "10.244.170.185"]
#kali_ip = "192.168.1.77" 

ansible_groups = {
  "hosts" => ["server0", "server1", "server2", "kali", "router"],
  "servers" => ["server0", "server1", "server2"]
}

server_mac_list = ["08002709E0F0", "080027AB8FDD", "08002764FD0B"]
kali_mac = "08002709E0FC"

server_cidr_netmask = "255.0.0.0"  
kali_cidr_netmask = "255.255.255.0"     

timestamp = Time.now.strftime("%Y%m%d%H%M%S")

def provision_ansible(device, ansible_groups, host)
  device.vm.provision "ansible_local" do |ansible|
    ansible.playbook = "provisioning/playbook.yml"
    ansible.config_file = "provisioning/ansible.cfg"
    ansible.galaxy_role_file = "provisioning/requirements.yml"
    ansible.groups = ansible_groups
    ansible.limit = host
  end
end

Vagrant.configure("2") do |config|
  
    # Define Kali VM
  config.vm.define "kali" do |device|
    device.vm.hostname = "kali"
    device.vm.box = "munikypo/kali"
    device.vm.provider "virtualbox" do |vb|
      vb.name = "probe-lab-kali-#{timestamp}"
      vb.gui = false
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
    device.vm.network "public_network", 
      type: "dhcp", 
      bridge: "enp5s0", 
      mac: "#{kali_mac}" 

    provision_ansible(device, ansible_groups, "kali")
  end
  
  # Define server VMs
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
      device.vm.network "public_network", 
        type: "dhcp", 
        bridge: "eno1", 
        mac: "#{server_mac_list[i]}" 
      
      provision_ansible(device, ansible_groups, "server#{i}")
    end
  end

end
