#!/bin/bash

if [ ! -f ./settings.sh ]; then
    echo "!! ERROR WHILE RUNNING REINSTALL.SH !!"
    echo "No settings.sh found!"
    exit 1
fi

echo "Starting: reinstall!"

. ./clean.sh
. ./install.sh

echo "Reinstall done!"
