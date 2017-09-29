# Set PS1
# Download link: https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh
if [[ $SETUPECHO = true ]]; then
    echo "Setting up PS1..."
fi
source ~/git-prompt.sh
PS1='\[\e[32m\]\u@\h \[\e[33m\]\w\[\e[92m\]\$(__git_ps1 " (%s)")\[\e[00m\] $ '
if [[ $SETUPECHO = true ]]; then
    echo "Bash setup complete!"
fi
