# Setup aliases
if [[ $SETUPECHO = true ]]; then
    echo "Setting up aliases..."
fi
alias cdgit="cd $GITDIR"
alias cdhome="cd $HOMEDIR"
alias gs="git status"
alias ls="ls --color"
alias vi="vim"
alias gc="git clean -f && git clean -f -d"
alias gca="git clean -f && git clean -f -d && git checkout -f"
alias gdb="git branch | grep -v \"*\" | xargs git branch -d"
alias dc="docker-compose"
alias drma="docker rm -f \\\$(docker ps -aq)"
