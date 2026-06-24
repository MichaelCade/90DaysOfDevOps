import sys
import time
import paramiko 
import os
import cmd
import datetime

now = datetime.datetime.now()
dt_string = now.strftime("%d/%m/%Y %H:%M:%S")
print("Your backup has started at", dt_string)	
tic = time.perf_counter()

#user = input("Enter username:")
#password = input("Enter Paswd:")
#enable_password = input("Enter enable pswd:")
user = "admin"
password = "access123"
enable_password = "access123"

port=22
f0 = open('backup.txt')
for ip in f0.readlines():
       ip = ip.strip()
       filename_prefix ='/Users/shambhu/Documents' + ip 
       ssh = paramiko.SSHClient()
       ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
       ssh.connect(ip,port, user, password, look_for_keys=False)
       chan = ssh.invoke_shell()
       time.sleep(2)
       chan.send('enable\n')
       chan.send(enable_password +'\n')
       time.sleep(1)
       chan.send('term len 0\n')
       time.sleep(1)
       chan.send('sh run\n')
       time.sleep(20)
       output = chan.recv(999999)
       filename = "%s_%.2i%.2i%i_%.2i%.2i%.2i" % (ip,now.year,now.month,now.day,now.hour,now.minute,now.second)
       f1 = open(filename, 'a')
       f1.write(output.decode("utf-8") )
       f1.close()
       ssh.close() 
       f0.close()
toc = time.perf_counter()
print("Congratulations You Have Backed Up Your 90DaysOfDevOps Lab")
print(f"Your backup duration was {toc - tic:0.4f} seconds")

dt_string = now.strftime("%d/%m/%Y %H:%M:%S")
print("Your backup completed at", dt_string)
