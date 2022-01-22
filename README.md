# git-id
Portable and independant dumb git identity management.

## todo
1. how to seperate ssh configs
  ```sh
  mv ~/.ssh/config ~/.ssh/global_config
  echo 'Include global_base.conf' | tee ~/.ssh/config && chmod 600 ~/.ssh/config # can also be 644

  printf '\nHost github.com # default\n  IdentityFile ~/.ssh/gh_rsa' # keep in mind, that this may be uesd by random stuff on your system; system might behave weirdly if this can't be used noninteractively

  printf '\n\nHost *.gh\n  HostName github.com\n  IdentitiesOnly yes\n' | tee -a ~/.ssh/git_alts.conf

  chmod 600 ~/.ssh/git_alts.conf
  sed -i '1iInclude git_alts.conf' ~/.ssh/global_config
  ```
1. addhost / addorigin / addremote
1. addid cmd
  ```sh
  #TODO: a) use foo.gh.git-id   identifyiable, maybe more machine-manipulative
  #      b) use foo.github.com  maybe benefits? *.github.com? anything else?
  #      c) use foo.gh          shorter
  printf 'Host foo.gh.git-id # foobar\n  IdentityFile ~/.ssh/foobar_sk\n' | tee ~/.ssh/git_alts.conf
  ```
  - git username, email, core.sshCommand
1. removeid cmd
  - ssh config fallback? default to no fallback/alias
1. nonmvp: editid
1. changeid cmd
  - this is used in repos
  - can we use (default:no) fallback secondary ssh config (default)?
    - will it be used include-style or fallback-style
    - can we use ssh_config things to print something / execute git-id hidden command
  - keep track of where it is used? can we query this later (for editing identifier, username, email, bla, keeping things working)
1. clone with id cmd
1. anything else clone-like?
## Commands
### git id [profile] [directory]
 - `profile`: profile slug to switch to; defaults to no config (user's default without `git-id`)
 - `directory` - repo dir on what to set the identiy on; defaults to `./`

## Configuration store
| Item | Default | Example | Description |
| --- | --- | --- | --- |
| `$GIT_ID_CONFIG` | `~/.config/git-id` |     |     |
| `/id/` |     | `/id/foo/` | Profiles |
| `/id/$slug/ssh` |     | (`~/.ssh/config`) | sshconfig used with the identity |
| #TODO: |     |     | default values for core `git-id` operation |
| `/id/$slug/git` |     | (`~/.gitconfig`) | gitconfig used with the identity |
| `/config` |     | (yaml) | Configuration for `git-id` |
| `/config`:`sshConfig`\[^conf\_sshConfig\_global\] | Import from global git config, `ssh` | `/tmp/ssh -vvv` | `git-id` will append `-F ~/$GIT_ID_CONFIG/id/$slug/config` |
