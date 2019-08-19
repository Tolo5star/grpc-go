#!/bin/bash
#demo
gnome-terminal -e "sudo docker run -it -p 8090:8090 --rm server-scratch"
read -p "Opening the server in a new terminal.Press enter after entering password in the new terminal to continue"
sudo docker run -it --net=host --rm client-scratch