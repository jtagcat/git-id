# [git-id](https://github.com/jtagcat/git-id)
Portable dumb git identity management.

NON-FUNCTIONAL, under development.

```sh
$ git id init # ~/.ssh/config → ~/.ssh/base.conf
              # + touch ~/.ssh/git-id{,_default}.conf
              # Import git-id{,_default}.conf
$ git id origin add gh github.com
$ git id add gh default ~/.ssh/id_rsa --username jtagcat --email 'user@domain.tld'
$ git id add gh work ~/.ssh/workyubi_sk --username irlname --email 'irl@work.biz' --description 'Evilcorp'

$ git id clone work git@github.com:evilcorp/private_repo.git # clone using id:work
$ cd private_repo
$ git id default # switch from id:work to id:default
$ git push
ERROR: Permission to evilcorp/private_repo.git denied to jtagcat.
```

TODO: instructions for zsh/othershell automatic `alias "git clone"="git id clone <id>"` (can be done with func: case: "$@")
