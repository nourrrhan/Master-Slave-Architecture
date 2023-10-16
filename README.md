# Master-Slave-Architecture

This project involves establishing communication among three or more devices, designated as master, slaves, and client.

Tools: Visual Studio Code.

Languages: GO.

The code assumes that the client initially connects with the master to receive the IP addresses of the slaves devices.
Subsequently, then the client connects with each slave device individually to receive a text file.

On the other hand, the master device connects with the client and sends the IP addresses and port numbers of each slave device.

Each slave device establishes a connection with the client using a specific port number and then sends a local text file to the client.

To successfully run this project, it is necessary to disable the Windows Firewall and configure the IP addresses of the devices
