#!/bin/bash

if [[ ! -f settings.sh ]]; then
    echo "!! ERROR WHILE RUNNING INSTALL.SH !!"
    echo "No settings.sh file found!"
    exit 1
fi

source settings.sh

echo "Starting: installation!"

# Copy files to homedir
for file in $FILESTOCOPY; do
    if [[ ! -f ~/$file ]]; then
        echo " - Copying $file to homedir..."
        cp $file ~/$file
    fi
done

# Download git-prompt if not exists
if [[ ! -f ~/git-prompt.sh ]]; then
    echo " - Downloading git-prompt.sh..."
    wget https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh > /dev/null 2>&1
    echo " - Moving git-prompt.sh to homedir..."
    mv git-prompt.sh ~/git-prompt.sh
fi

# Create setup bash script
echo " - Generating bash setup script..."
echo "cat << EOF > $OUTFILE" > tmp.sh
echo "#!/bin/bash" >> tmp.sh
if [ $CYGWINPROMPT = "true" ]; then
    cat ./cygwin/setup.sh >> tmp.sh
fi

cat setup-files/alias.sh >> tmp.sh
cat setup-files/git.sh >> tmp.sh
cat setup-files/ps1.sh >> tmp.sh
echo "EOF" >> tmp.sh

source ./tmp.sh

rm ./tmp.sh

# Add source to bashrc, profile or whatever is setup
for file in $BASHPROFILEFILES; do
    if ! grep --quiet "source $OUTFILE" ~/$file; then
        echo " - Adding source to $file..."
        echo "source $OUTFILE" >> ~/$file
    fi
done

echo "Installation done!"
