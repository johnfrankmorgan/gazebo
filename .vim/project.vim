augroup project
  autocmd!
  autocmd BufRead,BufNewFile *.hpp,*.cpp set filetype=cpp
  autocmd BufNewFile *.hpp HeaderguardAdd
augroup END

let &path.="include,"

nnoremap <C-M> :make!<CR>
nnoremap <Leader>h :CocCommand clangd.switchSourceHeader<CR>
