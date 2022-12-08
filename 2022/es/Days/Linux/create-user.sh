#! /usr/bin/bash
if [ -z "${1}" ]
then 
   echo "¿Qué nombre de usuario quieres?\n"
   read  username
   echo "¿Cuál es el password?"
   echo "(Cuando teclees no se verá nada por aquí)\n"
   read -s password

   # Un usuario puede ser pasado como un argumento en la línea de comandos
   echo "\nSe está creando la cuenta de usuario $username\n"

   # Se crea un usuario con el nombre del argumento.
   sudo useradd -m $username

   # Se puede pasar una contraseña como argumento en la línea de comandos.
   # Se establece la contraseña para el usuario.
   sudo chpasswd <<< $username:$password

   sleep 2 
   echo "Si quieres borrar el usuario entonces pasa 'del' y nombre de usuario en el argumento de la línea de comandos. e.g: ./create-user.sh del username"

else
   sudo userdel -rf "${2}"
   sleep 2	
   echo "La cuenta de usuario ${2} se ha eliminado con éxito."
   exit 0
fi

