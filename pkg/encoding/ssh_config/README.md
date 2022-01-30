# ssh_config

## Behaviour to look out for
### Backslash transformations
upstream: ErrWarnSingleBackslashTransformed:

any backslashes, where the next char is not `\`, `'`, `"`; an extra backslash is added/seen/read.

2 backslashes are parsed as 2 backslashes: `Hostname foo\\bar` == `Hostname foo\bar` in a config literally means SSH will connect to `foo\\bar`, not `foo\bar`, or `foober`.

There are exactly 0 ways to express a single backslash as a value.

### Automatic quoting
On encoding, if a value begins or ends with a space, and is not specified to be quoted, it is quoted with single quotes.

### Comment levels
To aid custom configuration (simliar to `X-` headers), expressed as comments, comments may be either in a root (TLD) or a sub (keyword / children) section.

On reading/writing, comments may either be root object, or a child of an another TLD (`Host` or `Match`).

Comments and newlines are child objects until the last non-comment key on the (root) tree, where they change to root objects on the first newline.  
See [layer1b_test.go](layer1b_test.go) for an example. There basic schema is here as well:
 ```
 # C1 (root)

 # C2 (root)
 Host foo (root)
   KVx (sub)
   # C3 (sub)
        (TBD, result: sub)
   # C4 (sub)
   KVx  (sub)
   # C5 (sub)
        (TBD, result: root)
 # C6 !deep
      !deep
 # C7 !deep
 Hostx deep
```

### Comments and newlines
Comments in ssh_config are defined as empty lines or space after an unquoted `#` character.
 - `# hello` in ssh_config is under `*.Comment` as ` hello` (mind the space)
 - `Hostname foo #`, ` #` will be trimmed, while `Hostname foo # ` Comment = ` `
   - `#\n` will be trimmed to `\n`.

## Status
tl;dr: pivoting, no typing support planned.

***

Frankly, it is a headache to do ssh_config with typing in go. I see multiple paths, yet none can keep metadata gracefully.

I want to do:
```go
Keywords.Hostname = "hello"
Keywords.Hostname.Comment = " world"
```

I can't. There is no root value.

I'd be fine with `Keywords.Hostname.Value`, but then everything in `types_keywords.go` would have to have a custom type: instead of:
```go
type Keyword struct {
	HashKnownHosts                   *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	HostbasedAuthentication          *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	IdentitiesOnly                   *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
}
```

It'd be:
```go
type Keyword struct {
	HashKnownHosts                   *HasKnownHosts `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	HostbasedAuthentication          *HostbasedAuthentication `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	IdentitiesOnly                   *IdentitiesOnly `minArgs:"1" maxArgs:"-2" definition:"Flag"`
}

type HashKnownHosts struct {
    Value bool
    Comment *string
    EncodingKVSeperatorIsEquals bool
}
type HostbasedAuthentication struct {
    Value bool
    Comment *string
    EncodingKVSeperatorIsEquals bool
}
type IdentitiesOnly struct {
    Value bool
    Comment *string
    EncodingKVSeperatorIsEquals bool
}
```

I just want to say:
```go
type Keyword struct {
	HashKnownHosts                   *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	HostbasedAuthentication          *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	IdentitiesOnly                   *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
}
type Keyword.* struct {
    Comment *string
    EncodingKVSeperatorIsEquals *bool
}
```

Using a map with metadata, and adding methods to each and every one would also be an option, yet it has the same drawbacks as the first option, and if not using path, it mixes up metadata of same-named keys.

### I give up
I don't need it, I've sunken enough time. The library will provide just raw objects. No typing.

Thus, `layer1b.go`. I'm not planning on actively adding anything. PRs welcome, upstream changes for both ssh_config and golang are encouraged.
