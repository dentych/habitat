#!/bin/bash

source settings.sh

# Copy files to homedir
for file in $FILESTOCOPY; do
    if [ ! -f ~/$file ]; then
        echo "Copying $file to homedir..."
        cp $file ~/$file
    fi
done

# Create setup bash script
echo "Generating bash setup script..."
cat << EOF > $OUTFILE
#!/bin/bash

# Setup aliases
echo "Setting up aliases..."
alias GIT="cd $HOMEDIR/$GITDIR"
alias gs="git status"
alias ls="ls --color"
alias vi="vim"
alias gc="git clean -f && git clean -f -d"

# Git initial setup
if [ ! -f $HOMEDIR/.gitconfig ]; then
    echo "Setting up git username and email..."
    git config --global user.name "$GITUSERNAME"
    git config --global user.email "$GITEMAIL"
fi

# Git stuff
echo "Setting up git aliases..."
git config --global alias.cp "cherry-pick"
git config --global alias.co "checkout"
git config --global alias.cl "clone"
git config --global alias.ci "commit"
git config --global alias.st "status -sb"
git config --global alias.br "branch"
git config --global alias.dc "diff --cached"
git config --global alias.p "pull -p"
git config --global alias.f "fetch -p"
git config --global alias.b "branch"

# Download git-prompt if not exists
if [ ! -f ~/git-prompt.sh ]; then
    echo "Downloading git-prompt.sh..."
    wget https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh > /dev/null 2>&1
fi

# Set PS1
# Download link: https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh
echo "Setting up PS1..."
source git-prompt.sh
PS1='\[\e[32m\]\u@\h \[\e[33m\]\w\[\e[92m\]\$(__git_ps1 " (%s)")\[\e[00m\] $ '
EOF

echo "Environment installation complete!"
