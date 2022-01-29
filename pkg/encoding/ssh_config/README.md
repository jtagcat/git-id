# ssh_config

## Status

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
