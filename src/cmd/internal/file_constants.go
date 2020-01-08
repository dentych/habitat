package internal

const TmuxConfFileName = ".tmux.conf"
const TmuxConf = `
#############################################
# Nice tmux conf by Dentych.                #
# Everyone may use it.                      #
# It's created for my own pleasure          #
#############################################

# Rebind prefix
#unbind C-b
#set -g prefix C-a
#bind C-a send-prefix

# Start index from 1 instead of 0
set -g base-index 1

# Enable mouse control (clickable windows, panes, resizable panes)
#set -g mouse-select-window on
#set -g mouse-select-pane on
#set -g mouse-resize-pane on
# Enable mouse mode (tmux 2.1 and above)
set -g mouse on

# Set terminal to 256 colors, which is awesome
set -g default-terminal "screen-256color"

# Toggle mouse functions on/off
bind m set -g mouse-select-window on \; set -g mouse-select-pane \; set -g mouse-resize-pane

# Kill screen style
bind k confirm-before "kill-window"

# Split window, keep current directory
bind c new-window -c '#{pane_current_path}'
bind % split-window -h -c '#{pane_current_path}'
bind '"' split-window -c '#{pane_current_path}'

# switch panes using Alt-arrow without prefix
bind -n M-Left select-pane -L
bind -n M-Right select-pane -R
bind -n M-Up select-pane -U
bind -n M-Down select-pane -D

# don't rename windows automatically
set-option -g allow-rename off
`

const VimConfFileName = ".vimrc"
const VimConf = `
set tabstop=4
set shiftwidth=4
set nu
set mouse=n
set bg=dark
set expandtab
set nobackup
set noundofile
set backspace=indent,eol,start

syntax on

filetype plugin indent on
autocmd BufRead,BufNewFile /etc/nginx/sites-*/* setfiletype conf

hi DiffAdd ctermbg=darkblue
hi DiffChange ctermbg=darkblue
hi DiffDelete ctermbg=darkred
hi difftext cterm=none ctermfg=white ctermbg=22
`
