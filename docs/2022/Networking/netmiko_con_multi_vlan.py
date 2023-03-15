#!/usr/bin/env python
from netmiko import ConnectHandler
from getpass import getpass

#password = getpass()

SW1 = {
    "device_type": "cisco_ios",
    "host": "192.168.169.178",
    "username": "admin",
    "password": "access123",
}

SW2 = {
    "device_type": "cisco_ios",
    "host": "192.168.169.193",
    "username": "admin",
    "password": "access123",
}

SW3 = {
    "device_type": "cisco_ios",
    "host": "192.168.169.125",
    "username": "admin",
    "password": "access123",
}

SW4 = {
    "device_type": "cisco_ios",
    "host": "192.168.169.197",
    "username": "admin",
    "password": "access123",
}
command = "show int trunk"
for device in (SW1, SW2, SW3, SW4):
    net_connect = ConnectHandler(**device)
    print(net_connect.find_prompt())
    print(net_connect.send_command(command))
    net_connect.disconnect()