package ssh_config

//TODO: transform case, strip 2+ spaces
// Compatible with OpenSSH 8.8
type TopLevel struct {
	Key string // enum(4): "" (comment / empty line), Host, Match,
	// Import: not recursed, nothing is done (no Children)
	Values []Value
	// "# foobar" → " foobar", note the leading space
	Comment string

	Children Keywords // named []Keyword
}

type Keyword struct {
	// enum(TODO): https://man.openbsd.org/ssh_config
	Key    string
	Values []Value // when key set, len(Values) >= 1
	// "# foobar" → " foobar", note the leading space
	Comment string // at the end of same line as Key

	EncodingKVSeperatorIsEquals bool // "Key=Value" instead of "Key Value"
}
type Value struct {
	Value     string
	Quoted    int    // enum: 0: not, 1: single, 2: double
	ValueType string // enum TBD TODO:

	// UNDOCUMENTED UPSTREAM: escapes
	// https://ftp.openbsd.org/pub/OpenBSD/OpenSSH/openssh-8.8.tar.gz misc.c#1889: Copy the token in, removing escapes
	// both single and double quotes are allowed
	// \ → \\; \\ → \\; \\\ → \\\\ (2 sets)
	// \" → "; \' → ' (after escapechars are processed)
	// \\" → invalid (\\ + non-escaped ")
	// va"lue → invalid
	// "value" → value
}

// https://ftp.openbsd.org/pub/OpenBSD/OpenSSH/openssh-8.8.tar.gz readconf.c#984: switch (opcode) {

type Keywords struct {
}

//var keywordIndex = map[string]string{}

// https://ftp.openbsd.org/pub/OpenBSD/OpenSSH/openssh-8.8.tar.gz readconf.c#792: Multistate option parsing
// "string", "int",
var enumIndex = map[string][]enumValues{
	"Flag": {{"", nil}, {"true", true}, {"false", false}, {"yes", true}, {"no", false}},
	"YesNoAsk": {{"", nil}, {"true", true}, {"false", false}, {"yes", true}, {"no", false},
		{"ask", "ask"}},
	"YesNoAskConfirm": {{"", nil}, {"true", true}, {"false", false}, {"yes", true}, {"no", false},
		{"ask", "ask"}, {"confirm", "confirm"}},
	"ControlMaster": {{"", nil}, {"true", true}, {"false", false}, {"yes", true}, {"no", false},
		{"auto", "auto"}, {"ask", "ask"}, {"autoask", "autoask"}},
	"StrictHostkey": {{"", nil}, {"true", true}, {"false", false}, {"yes", true}, {"no", false}, {"off", false},
		{"ask", "ask"}, {"accept-new", "accept-new"}},
	"CanonicalizeHostname": {{"", nil}, {"true", true}, {"false", false}, {"yes", true}, {"no", false},
		{"always", "always"}},
	"RequestTTY": {{"", nil}, {"true", true}, {"false", false}, {"yes", true}, {"no", false},
		{"force", "force"}, {"auto", "auto"}},
	"Tunnel": {{"", nil}, {"true", true}, {"false", false}, {"yes", true}, {"no", false},
		{"ethernet", "ethernet"}, {"point-to-point", "point-to-point"}},
	"Compression":   {{"", nil}, {"yes", true}, {"no", false}},
	"AddressFamily": {{"", nil}, {"inet", "inet"}, {"inet6", "inet6"}, {"any", "any"}},
	"SessionType":   {{"", nil}, {"none", "none"}, {"subsystem", "subsystem"}, {"default", "default"}},
	"LogLevel":      {{"", nil}, {"quiet", "quiet"}, {"fatal", "fatal"}, {"error", "error"}, {"info", "info"}, {"verbose", "verbose"}, {"debug", "debug"}, {"debug1", "debug1"}, {"debug2", "debug2"}, {"debug3", "debug3"}},
	"LogFacility":   {{"", nil}, {"daemon", "daemon"}, {"user", "user"}, {"auth", "auth"}, {"local0", "local0"}, {"local1", "local1"}, {"local2", "local2"}, {"local3", "local3"}, {"local4", "local4"}, {"local5", "local5"}, {"local6", "local6"}, {"local7", "local7"}},
	"Hash":          {{"", nil}, {"md5", "md5"}, {"sha256", "sha256"}},
}

type keywordIndexType struct {
	minArgs int // non-negative int
	maxArg  int // non-negative int or negative for unlimited
	//	repeatable bool   // Whether slices in ssh_config is achieved with multiple declarations
	definition string // type to follow
}

type enumValues struct {
	stringName string
	value      interface{}
}

// TODO: have an overwrite detector for duplicate(ish) key-values
