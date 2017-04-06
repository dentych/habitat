#!/bin/bash

source settings.sh

if [ -f ~/git-prompt.sh ]; then
    echo "Removing git-prompt.sh from homedir..."
    rm ~/git-prompt.sh
fi

for file in $FILESTOCOPY; do
    if [ -f ~/$file ]; then
        echo "Removing $file from homedir..."
        rm ~/$file
    fi
done

echo "Clean complete!"
