#!/bin/bash

# Download air
curl -fLo air https://raw.githubusercontent.com/cosmtrek/air/master/install.sh
chmod +x air

# Move air to /usr/local/bin with sudo
sudo mv air /usr/local/bin/
#sudo ./install-air.sh
