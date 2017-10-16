# Setup aliases
if [[ $SETUPECHO = true ]]; then
    echo "Setting up aliases..."
fi
alias GIT="cd $GITDIR"
alias HOME="cd $HOMEDIR"
alias gs="git status"
alias ls="ls --color"
alias vi="vim"
alias gc="git clean -f && git clean -f -d"
alias gca="git clean -f && git clean -f -d && git checkout -f"
alias gdb="git branch | grep -v \"*\" | xargs git branch -d"
