# git-id
Portable and independant dumb git identity management.

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
