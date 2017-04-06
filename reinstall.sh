#!/bin/bash

if [[ ! -f settings.sh ]]; then
    echo "!! ERROR WHILE RUNNING REINSTALL.SH !!"
    echo "No settings.sh found!"
    exit 1
fi

echo "Starting: reinstall!"

source clean.sh
source install.sh

echo "Reinstall done!"
