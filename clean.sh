#!/bin/bash

if [ ! -f ./settings.sh ]; then
    echo "!! ERROR WHILE RUNNING CLEAN.SH !!"
    echo "No settings.sh found!"
    exit 1
fi

. ./settings.sh

echo "Starting: cleanup!"

if [ -f ~/git-prompt.sh ]; then
    echo " - Removing git-prompt.sh from homedir..."
    rm ~/git-prompt.sh
fi

for file in $FILESTOCOPY; do
    if [ -f ~/$file ]; then
        echo " - Removing $file from homedir..."
        rm -f ~/$file
    fi
done

if [ "$CYGWINPROMPT" = "true" ]; then
    . ./cygwin/run.sh clean
fi

echo "Cleanup done!"
