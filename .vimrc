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

