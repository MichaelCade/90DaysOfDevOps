#!/usr/bin/python

import paramiko
import time

Channel = paramiko.SSHClient()

Channel.set_missing_host_key_policy(paramiko.AutoAddPolicy())
Channel.connect(hostname="192.168.169.115", username='admin', password='access123', look_for_keys=False, allow_agent=False)

shell = Channel.invoke_shell()  # This will set interactive shell
shell.send("enable\n")
shell.send("access123\n")
shell.send("terminal length 0\n")
shell.send("show ip int b\n")
shell.send("show arp \n")
time.sleep(2)

print(shell.recv(5000))
# This will receive everything from the buffer, if you need to receive specifc output, the you should execute the command
# and immediately receive the output before executing the 2nd command, Also you should sleep a little
Channel.close()