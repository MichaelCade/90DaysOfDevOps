#! /usr/bin/bash
if [ -z "${1}" ]
then 
   echo "What is your intended username?"
   read  username
   echo "What is your password"
   read  password

   #A user can be passed in as a command line argument
   echo "$username user account being created."

   #A user is created with the name of command line argument
   sudo useradd -m $username

   #A password can be parsed in as a command line argument.
   sudo chpasswd <<< $username:$password

   sleep 2 
   echo "If you want to delete the user then pass 'del' and username in command line argument. e.g: ./create-user.sh del username"

else
   sudo userdel -rf "${2}"
   sleep 2	
   echo "${2} user account successfully deleted."
   exit 0
fi

