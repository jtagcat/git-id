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
	Value    string
	isQuoted bool // UNDOCUMENTED UPSTREAM: escapes
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

var keywordIndex = map[string]string{
	"ForwardX11Trusted":                "Flag",
	"ForwardX11Timeout":                "duration",
	"GatewayPorts":                     "Flag",
	"ExitOnForwardFailure":             "Flag",
	"PasswordAuthentication":           "Flag",
	"KbdInteractiveAuthentication":     "Flag", // has aliases
	"challengeresponseauthentication":  "Flag", // alias KbdInteractiveAuthentication
	"skeyauthentication":               "Flag", // alias KbdInteractiveAuthentication
	"tisauthentication":                "Flag", // alias KbdInteractiveAuthentication
	"KbdInteractiveDevices":            "csvstring",
	"PubkeyAuthentication":             "Flag", // has aliases
	"DSAAuthentication":                "Flag", // alias PubkeyAuthentication
	"HostbasedAuthentication":          "Flag",
	"GssAuthentication":                "Flag",
	"GssDelegateCreds":                 "Flag",
	"BatchMode":                        "Flag",
	"CheckHostIP":                      "Flag",
	"TCPKeepAlive":                     "Flag",
	"NoHostAuthenticationForLocalhost": "Flag",
	"NumberOfPasswordPrompts":          "nint",
	"XAuthLocation":                    "string",
	"Hostname":                         "string",
	"HostKeyAlias":                     "string",
	"PreferredAuthentications":         "csvstring",
	"BindAddress":                      "string",
	"BindInterface":                    "string",
	"PKCS11Provider":                   "string",
	"SecurityKeyProvider":              "string",
	"ClearAllForwardings":              "Flag",
	"EnableSSHKeysign":                 "Flag",
	"IdentitiesOnly":                   "Flag",
	"ServerAliveInterval":              "time",
	"ServerAliveCountMax":              "nint",
	"ControlPath":                      "string",
	"HashKnownHosts":                   "Flag",
	"PermitLocalCommand":               "Flag",
	"VisualHostKey":                    "Flag",
	"StdinNull":                        "Flag",
	"ForkAfterAuthentication":          "Flag",
	"IgnoreUnknown":                    "csvstring",
	"ProxyUseFdpass":                   "Flag",
	"CanonicalizeMaxDots":              "nint",
	"CanonicalizeFallbackLocal":        "Flag",
	"StreamLocalBindUnlink":            "Flag",
	"RevokedHostKeys":                  "string",
	"UserKnownHostsFile":               "csvstring",
	"GlobalKnownHostsFile":             "csvstring",
	"ConnectTimeout":                   "duration",
	"ForwardX11":                       "Flag",
	"ForwardAgent":                     "string", // bool or string
	"IdentityAgent":                    "string",
	"VerifyHostKeyDNS":                 "YesNoAsk",
	"StrictHostKeyChecking":            "StrictHostkey",
	"Compression":                      "Compression",
	"RekeyLimit":                       "rekeyLimit",
	"IdentityFile":                     "multiDefineStringSlice",
	"CertificateFile":                  "multiDefineStringSlice",
	"User":                             "string",
	"ProxyCommand":                     "indifferentString",
	"KnownHostsCommand":                "indifferentString",
	"LocalCommand":                     "indifferentString",
	"RemoteCommand":                    "indifferentString",
	"Port":                             "port",
	"ConnectionAttempts":               "nint",
	"LocalForward":                     "free",
	"RemoteForward":                    "free",
	"AFSTokenPassing":                  "unsupported",
	"KerberosAuthentication":           "unsupported",
	"KerberosTGTPassing":               "unsupported",
	"RSAAuthentication":                "unsupported",
	"RhostsRSAAuthentication":          "unsupported",
	"CompressionLevel":                 "unsupported",
	"Protocol":                         "deprecatedHidden",
	"Cipher":                           "deprecated",
	"FallbackToRsh":                    "deprecated",
	"GlobalKnownHostsFile2":            "deprecated",
	"RhostsAuthentication":             "deprecated",
	"UserKnownHostsFile2":              "deprecated",
	"UseRoaming":                       "deprecated",
	"UseRsh":                           "deprecated",
	"UsePrivilegedPort":                "deprecated",
	"AddressFamily":                    "AddressFamily",
	"ControlMaster":                    "ControlMaster",
	"Tunnel":                           "Tunnel",
	"RequestTTY":                       "RequestTTY",
	"SessionType":                      "SessionType",
	"CanonicalizeHostname":             "CanonicalizeHostname",
	"UpdateHostkeys":                   "YesNoAsk",
	"AddKeysToAgent":                   "YesNoAskConfirm",
	"LogLevel":                         "LogLevel",
	"LogFacility":                      "LogFacility",
	"LogVerbose":                       "stringSlice",
	"Ciphers":                          "cipher",
	"KexAlgorithms":                    "cipher",
	"Macs":                             "undocumented", // NEEDINFO?, cipher?
	"HostKeyAlgorithms":                "cipher",
	"CASignatureAlgorithms":            "cipher",
	"HostbasedAcceptedAlgorithms":      "cipher",
	"PubkeyAcceptedAlgorithms":         "cipher",
	"FingerprintHash":                  "Hash",
	"StreamLocalBindMask":              "permoctal",
	"CanonicalizePermittedCNAMEs":      "canonicalizeCNAMEs",
	"CanonicalDomains":                 "stringSlice",
	"IPQoS":                            "ipqos",
	"TunnelDevice":                     "tunnelDevice",
	"ControlPersist":                   "controlPersist",
	"SetEnv":                           "stringSlice",
	"SendEnv":                          "stringSlice",
	"EscapeChar":                       "string",
	"PermitRemoteOpen":                 "permitRemoteOpen",
	"DynamicForward":                   "dynamicForward",
}

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

type enumValues struct {
	stringName string
	value      interface{}
}

/* "duration":
* Format is a sequence of:
*      time[qualifier]
*
* Valid time qualifiers are:
*      <none>  seconds
*      s|S     seconds
*      m|M     minutes
*      h|H     hours
*      d|D     days
*      w|W     weeks
*
* Examples:
*      90m     90 minutes
*      1h30m   90 minutes
*      2d      2 days
*      1w      1 week */

// rekeyLimit: 1st: default or int[qualifier], format=K,M,G; 2nd (optional): duration (or 'none')

// controlPersist: Flag, or duration

// indifferentString: Key="echo" "hi" are joined to "echo hi"

// multiDefineStringSlice: specifying same key-value multiple times expresses a slice

// stringSlice: slice in go, slice in conf
// csvstring: comma seperated string in the same thing, slice in go
// cipher: csvstring, but begins with "" (set), + (append), - (subtract), or ^ (preappend)
//   wildcards (*) are supported in subtraction
//   Ciphers will not be enummed nor runtime-checked (list available ciphers) as they change often

// permoctal: 0777 (4 digits in base8)

// canonicalizeCNAMEs: csvstring, but each set is 2tuplet x:y (colon-seperated)

// permitRemoteOpen: none / any / host:port (/ *:port) / :port

// dynamicForward: [host:]port

/* ipqos: Specifies the IPv4 type-of-service or DSCP class for connections.  Accepted values are af11,
af12, af13, af21, af22, af23, af31, af32, af33, af41, af42, af43, cs0, cs1, cs2, cs3, cs4, cs5,
cs6, cs7, ef, le, lowdelay, throughput, reliability, a numeric value, or none to use the operat‐
ing system default.  This option may take one or two arguments, separated by whitespace.  If one
argument is specified, it is used as the packet class unconditionally.  If two values are speci‐
fied, the first is automatically selected for interactive sessions and the second for non-inter‐
active sessions.  The default is af21 (Low-Latency Data) for interactive sessions and cs1 (Lower
Effort) for non-interactive sessions. */

// port: int 0..65535

// nint: int 0..0x7fffffff

// tunnelDevice: local_tun[:remote_tun], probably single arg/string?

// free: no type, string just forwarded
// unsupported: unsupported by OpenSSH, mby return friendly error; use free
// deprecated: deprecated, mby return friendly error; use free
// deprecatedHidden: deprecated, but don't warn

// TODO: have an overwrite detector for duplicate(ish) key-values

//TODO: test: set everything and see what fails; humans make mistakes