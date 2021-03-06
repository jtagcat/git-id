# [`git-id`](https://github.com/jtagcat/git-id)
Portable dumb identity management for Git, the [stupid content tracker](https://git-scm.com/docs/git#_name).


## Project status
Totally non-functional. I went on a tangent on implementing a library to interact with ssh_config, wanting to make it beautiful. ssh_config is really, quite ugly. 

***


```sh
$ git id init # Include git-id.conf in ~/.ssh/config
$ git id remote add gh github.com
$ git id add gh default ~/.ssh/id_rsa --username jtagcat --email 'user@domain.tld'
$ git id add gh work ~/.ssh/work_sk --username irlname --email 'irl@work.biz' --description 'Evilcorp'

$ git id clone work git@github.com:evilcorp/private_repo.git # clone using id:work
$ cd private_repo
$ git id default # switch from id:work to id:default
$ git push
ERROR: Permission to evilcorp/private_repo.git denied to jtagcat.
```

TODO: instructions for zsh/othershell automatic `alias "git clone"="git id clone <id>"` (can be done with func: case: "$@")

## How it works
```ascii
     PUSH github.com
Git ─────────────────► SSH             ~/.ssh/id_rsa SSH github.com
 jtagcat                │                ▲
 user@domain.tld        │                │
                        ▼                │
                ~/.ssh/git-id.conf       │
                  MATCH originalhost github.com
```
```ascii
    PUSH work.gh.git-id
Git ───────────────────► SSH  ┌──────► ~/.ssh/work_sk SSH github.com
 irlname                  │   │
 irl@work.biz             │   │
                          ▼
                  ~/.ssh/git-id.conf
                    HOST work.gh.git-id ─► work_sk                                    Created with
                    HOST    *.gh.git-id ─► github.com          https://github.com/lewish/asciiflow
```

So, `git-id` is here to implement basic git identity management without dependencies: something simple to use, something that just works when you need to push a commit from your work, other work, or h4ck3r account, to the same remote host.

The reliable™ hijack depends on `ssh_config`, and 2 git configurations:
 - The domain/host inside the git remote url is changed to a pseudohost (later picked up by `ssh_config`). This is essentially a coded message to SSH for 'please use this identity'
 - `core.sshCommand` - this is used to force git (repo) to use a specific `ssh_config` file(s) (removing the possibility that the default gets picked up)
 - `ssh_config` matches the pseudohost to a (requested) ssh key, and changes back to the real host.

Besides organizing the hijcak, `git-id` sets Git user.name and user.email, and keeps a list of identities, their purpose and their usage (for later mass-changes).
