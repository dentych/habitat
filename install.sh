#!/bin/bash

if [ ! -f settings.sh ]; then
    echo "!! ERROR WHILE RUNNING INSTALL.SH !!"
    echo "No settings.sh file found!"
    exit 1
fi

source settings.sh

echo "Starting: installation!"

# Copy files to homedir
for file in $FILESTOCOPY; do
    if [ ! -f ~/$file ]; then
        echo " - Copying $file to homedir..."
        cp $file ~/$file
    fi
done

# Download git-prompt if not exists
if [ ! -f ~/git-prompt.sh ]; then
    echo " - Downloading git-prompt.sh..."
    wget https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh > /dev/null 2>&1
    echo " - Moving git-prompt.sh to homedir..."
    mv git-prompt.sh ~/git-prompt.sh
fi

# Create setup bash script
echo " - Generating bash setup script..."
cat << EOF > $OUTFILE
#!/bin/bash

# Setup aliases
if [ $SETUPECHO = true ]; then
    echo "Setting up aliases..."
fi
alias GIT="cd $HOMEDIR/$GITDIR"
alias gs="git status"
alias ls="ls --color"
alias vi="vim"
alias gc="git clean -f && git clean -f -d"

# Git initial setup
if [ ! -f $HOMEDIR/.gitconfig ]; then
    if [ $SETUPECHO = true ]; then
        echo "Setting up git username and email..."
    fi
    git config --global user.name "$GITUSERNAME"
    git config --global user.email "$GITEMAIL"
fi

# Git stuff
if [ $SETUPECHO = true ]; then
    echo "Setting up git aliases..."
fi
git config --global alias.cp "cherry-pick"
git config --global alias.co "checkout"
git config --global alias.cl "clone"
git config --global alias.ci "commit"
git config --global alias.st "status -sb"
git config --global alias.br "branch"
git config --global alias.d "diff"
git config --global alias.dc "diff --cached"
git config --global alias.p "pull -p"
git config --global alias.f "fetch -p"
git config --global alias.b "branch"

# Set PS1
# Download link: https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh
if [ $SETUPECHO = true ]; then
    echo "Setting up PS1..."
fi
source git-prompt.sh
PS1='\[\e[32m\]\u@\h \[\e[33m\]\w\[\e[92m\]\$(__git_ps1 " (%s)")\[\e[00m\] $ '
if [ $SETUPECHO = true ]; then
    echo "Bash setup complete!"
fi
EOF

# Add source to bashrc, profile or whatever is setup
for file in $BASHPROFILEFILES; do
    if ! grep --quiet "source $OUTFILE" ~/$file; then
        echo " - Adding source to $file..."
        echo "source $OUTFILE" >> ~/$file
    fi
done

echo "Installation done!"
