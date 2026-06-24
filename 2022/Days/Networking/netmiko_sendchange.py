from netmiko import ConnectHandler

SW2 = {
    "device_type": "cisco_ios",
    "host": "192.168.169.193",
    "username": "admin",
    "password": "access123",
    "secret": "access123",
}

core_sw_config = ["int range gig0/1 - 2", "switchport trunk encapsulation dot1q",
                  "switchport mode trunk", "switchport trunk allowed vlan 1,2"]

print("########## Connecting to Device {0} ############".format(SW2))
net_connect = ConnectHandler(**SW2)
net_connect.enable()

print("***** Sending Configuration to Device *****")
net_connect.send_config_set(core_sw_config)
