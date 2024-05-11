# My pass

CLI (heh, why not a TUI at some point!) application to store secret.  
It __aim__ to be a [pass](https://www.passwordstore.org) clone, but written in Go as a personal project to learn this language.

# Pass TL;DR

Encrypt:
`gpg -r <key_id> -e path/to/file`

Dencrypt:
`gpg -e path/to/file`

Export public key:
`gpg --export (--armor) <key_id>`
> The above command output the exported public key to the standard output (STDOUT).</br>
> Redirect the output to a file by doing: `<previous command> > path/to/file`
