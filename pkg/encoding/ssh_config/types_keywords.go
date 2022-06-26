package ssh_config

import (
	"time"
)

// for sake of documentation
type RootWords struct {
	Host    *[]string `minArgs:"1" maxArgs:"-2" definition:"stringSlice"`
	Match   *[]string `minArgs:"1" maxArgs:"-2" definition:"match"` // [] of string=csvStringSlice, or string (key only)
	Include *[]string `minArgs:"1" maxArgs:"-2" definition:"stringSlice"`
}

// https://ftp.openbsd.org/pub/OpenBSD/OpenSSH/openssh-8.8.tar.gz readconf.c#984: switch (opcode) {
type Keywords struct {
	// minArgs, maxArgs: nzint (disable check: negative) of arguments demanded by OpenSSH
	// definition: what standard to follow for reading-mapping-writing and deeper type checking

	// ##### native types #####
	CanonicalizeHostname             *bool `minArgs:"1" maxArgs:"1" definition:"CanonicalizeHostname"`
	Compression                      *bool `minArgs:"1" maxArgs:"1" definition:"Compression"`
	ControlMaster                    *bool `minArgs:"1" maxArgs:"1" definition:"ControlMaster"`
	KbdInteractiveAuthentication     *bool `minArgs:"1" maxArgs:"1" definition:"Flag" aliases:"ChallengeResponseAuthentication,TISAuthentication"`
	PubkeyAuthentication             *bool `minArgs:"1" maxArgs:"1" definition:"Flag" aliases:"DSAAuthentication"`
	BatchMode                        *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	CanonicalizeFallbackLocal        *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	CheckHostIP                      *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	ClearAllForwardings              *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	EnableSSHKeysign                 *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	ExitOnForwardFailure             *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	ForkAfterAuthentication          *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	ForwardX11                       *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	ForwardX11Trusted                *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	GatewayPorts                     *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	GssAuthentication                *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	GssDelegateCreds                 *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	HashKnownHosts                   *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	HostbasedAuthentication          *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	IdentitiesOnly                   *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	NoHostAuthenticationForLocalhost *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	PasswordAuthentication           *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	PermitLocalCommand               *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	ProxyUseFdpass                   *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	StdinNull                        *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	StreamLocalBindUnlink            *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	TCPKeepAlive                     *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	VisualHostKey                    *bool `minArgs:"1" maxArgs:"-2" definition:"Flag"`
	RequestTTY                       *bool `minArgs:"1" maxArgs:"-2" definition:"RequestTTY"`
	StrictHostKeyChecking            *bool `minArgs:"1" maxArgs:"-2" definition:"StrictHostkey"`
	Tunnel                           *bool `minArgs:"1" maxArgs:"-2" definition:"Tunnel"`
	// definition csvStringSlice: all slice items are in ssh_config[0] (len == 1), seperated by comma
	GlobalKnownHostsFile     *[]string `minArgs:"1" maxArgs:"-2" definition:"csvStringSlice"`
	IgnoreUnknown            *[]string `minArgs:"1" maxArgs:"-2" definition:"csvStringSlice"`
	KbdInteractiveDevices    *[]string `minArgs:"1" maxArgs:"-2" definition:"csvStringSlice"`
	PreferredAuthentications *[]string `minArgs:"1" maxArgs:"-2" definition:"csvStringSlice"`
	UserKnownHostsFile       *[]string `minArgs:"1" maxArgs:"-2" definition:"csvStringSlice"`
	// definition canonicalizeCNAMEs: csvStringSlice where each set is 2tuplet x:y
	CanonicalizePermittedCNAMEs *CanonicalizeCNAMEs `minArgs:"1" maxArgs:"-2" definition:"canonicalizeCNAMEs"`
	/* definition duration Format is a sequence of:
	*      time[qualifier]...
	*
	* Valid time qualifiers are:
	*      <none>  seconds
	*      s|S     seconds
	*      m|M     minutes
	*      h|H     hours
	*      d|D     days
	*      w|W     weeks
	* Examples:
	*      90m     90 minutes
	*      1h30m   90 minutes
	 */
	ConnectTimeout    *time.Duration `minArgs:"1" maxArgs:"-2" definition:"duration"`
	ForwardX11Timeout *time.Duration `minArgs:"1" maxArgs:"-2" definition:"duration"`
	// definition dynamicForward: [host:]port, may have [ipv6]
	DynamicForward *DynamicForward `minArgs:"1" maxArgs:"-2" definition:"dynamicForward"`
	// definition indifferentString Key="echo" "hi" are joined to "echo hi" on parse
	KnownHostsCommand *string `minArgs:"1" maxArgs:"-2" definition:"indifferentString"`
	LocalCommand      *string `minArgs:"1" maxArgs:"-2" definition:"indifferentString"`
	ProxyCommand      *string `minArgs:"1" maxArgs:"-2" definition:"indifferentString"`
	RemoteCommand     *string `minArgs:"1" maxArgs:"-2" definition:"indifferentString"`
	// definition permoctal: 0777 (4 digits in base8)
	StreamLocalBindMask *int `minArgs:"1" maxArgs:"-2" definition:"permoctal"` // 0777
	/* definition rekeyLimit: 1st: "default" or int[qualifier], qualifier:K/M/G; 2nd (optional): duration (or 'none')
	* 1st: 1..15: RekeyLimit too small */
	RekeyLimit *RekeyLimit `minArgs:"1" maxArgs:"-2" definition:"rekeyLimit"`
	// stringSlice: slice in go, slice in conf
	CanonicalDomains *[]string `minArgs:"1" maxArgs:"-2" definition:"stringSlice"`
	Include          *[]string `minArgs:"1" maxArgs:"-1" definition:"stringSlice"`
	LogVerbose       *[]string `minArgs:"1" maxArgs:"-2" definition:"stringSlice"`
	SendEnv          *[]string `minArgs:"1" maxArgs:"-2" definition:"stringSlice"`
	SetEnv           *[]string `minArgs:"1" maxArgs:"-2" definition:"stringSlice"`
	// definition multiDefineStringSlice: specifying key-value multiple times as a way of expressing slice
	CertificateFile *[]string `minArgs:"1" maxArgs:"-2" definition:"multiDefineStringSlice"`
	IdentityFile    *[]string `minArgs:"1" maxArgs:"-2" definition:"multiDefineStringSlice"`
	// definition tunnelDevice: local_tun[:remote_tun]
	TunnelDevice TunnelDevice `minArgs:"1" maxArgs:"-2" definition:"tunnelDevice"`
	// definition string: also includes keywords where type checking isn't possible outside of config time of use
	BindAddress         *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	BindInterface       *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	ControlPath         *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	EscapeChar          *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	HostKeyAlias        *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	Hostname            *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	IdentityAgent       *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	PKCS11Provider      *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	RevokedHostKeys     *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	SecurityKeyProvider *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	User                *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	XAuthLocation       *string `minArgs:"1" maxArgs:"-2" definition:"string"`
	Port                *uint16 `minArgs:"1" maxArgs:"-2" definition:"uint16"`

	// ##### enums #####
	AddressFamily    *AddressFamily   `minArgs:"1" maxArgs:"-2" definition:"AddressFamily"`
	FingerprintHash  *Hash            `minArgs:"1" maxArgs:"-2" definition:"Hash"`
	LogFacility      *LogFacility     `minArgs:"1" maxArgs:"-2" definition:"LogFacility"`
	LogLevel         *LogLevel        `minArgs:"1" maxArgs:"-2" definition:"LogLevel"`
	SessionType      *SessionType     `minArgs:"1" maxArgs:"-2" definition:"SessionType"`
	UpdateHostkeys   *YesNoAsk        `minArgs:"1" maxArgs:"-2" definition:"YesNoAsk"`
	VerifyHostKeyDNS *YesNoAsk        `minArgs:"1" maxArgs:"-2" definition:"YesNoAsk"`
	AddKeysToAgent   *YesNoAskConfirm `minArgs:"1" maxArgs:"-2" definition:"YesNoAskConfirm"`
	/* definition csvStringSlice, but begins with "" (set/override), + (append), ^ (preappend), - (subtract regex)
	*  - supports wildtags (*), NEEDINFO:{-2, -2, are wildcards supported on set/additions?
	* ciphers shall not be enummed, as they are independent of the config, and change often */
	CASignatureAlgorithms       *Cipher `minArgs:"1" maxArgs:"-2" definition:"cipher"`
	Ciphers                     *Cipher `minArgs:"1" maxArgs:"-2" definition:"cipher"`
	HostbasedAcceptedAlgorithms *Cipher `minArgs:"1" maxArgs:"-2" definition:"cipher"`
	HostKeyAlgorithms           *Cipher `minArgs:"1" maxArgs:"-2" definition:"cipher"`
	KexAlgorithms               *Cipher `minArgs:"1" maxArgs:"-2" definition:"cipher"`
	PubkeyAcceptedAlgorithms    *Cipher `minArgs:"1" maxArgs:"-2" definition:"cipher"`
	// definition permitRemoteOpen: none / any / host:port / :port
	PermitRemoteOpen *PermitRemoteOpen `minArgs:"1" maxArgs:"-2" definition:"permitRemoteOpen"`

	// ##### custom types #####
	// definition boolString: bool or string
	ForwardAgent *ForwardAgent `minArgs:"1" maxArgs:"-2" definition:"boolString"`
	// definition controlPersist: Flag or duration
	ControlPersist *ControlPersist `minArgs:"1" maxArgs:"-2" definition:"controlPersist"`
	/* definition: Accepted values are af11, af12, af13, af21, af22, af23, af31, af32, af33, af41, af42, af43,
	* cs0, cs1, cs2, cs3, cs4, cs5, cs6, cs7, ef, le, lowdelay, throughput, reliability,
	* a numeric value, or none to use the operating system default. */
	IPQoS *Ipqos `minArgs:"1" maxArgs:"2" definition:"ipqos"`
	// definition nzint32: non-zero int32:{-2, -2, 0..0x7fffffff (2147483647)
	CanonicalizeMaxDots     *Nzint32 `minArgs:"1" maxArgs:"1" definition:"nzint32"`
	ConnectionAttempts      *Nzint32 `minArgs:"1" maxArgs:"1" definition:"nzint32"`
	NumberOfPasswordPrompts *Nzint32 `minArgs:"1" maxArgs:"1" definition:"nzint32"`
	ServerAliveCountMax     *Nzint32 `minArgs:"1" maxArgs:"1" definition:"nzint32"`
	ServerAliveInterval     *Nzint32 `minArgs:"1" maxArgs:"1" definition:"nzint32"`
	/* definition forward:
	* dynamicfwd == 0
	 *   [listenhost:]listenport|listenpath:connecthost:connectport|connectpath
	 *   listenpath:connectpath
	 * dynamicfwd == 1
	 *	[listenhost:]listenport */
	LocalForward  *Forward `minArgs:"1" maxArgs:"-2" definition:"forward"` // type checking very hard, falling back to string for now
	RemoteForward *Forward `minArgs:"1" maxArgs:"-2" definition:"forward"`

	// ##### not availiable ##### (treated as string)
	// definition deprecated: deprecated by OpenSSH, maybe return friendly error?
	Cipher                *string `minArgs:"1" maxArgs:"-2" definition:"deprecated"`
	FallbackToRsh         *string `minArgs:"1" maxArgs:"-2" definition:"deprecated"`
	GlobalKnownHostsFile2 *string `minArgs:"1" maxArgs:"-2" definition:"deprecated"`
	RhostsAuthentication  *string `minArgs:"1" maxArgs:"-2" definition:"deprecated"`
	UsePrivilegedPort     *string `minArgs:"1" maxArgs:"-2" definition:"deprecated"`
	UserKnownHostsFile2   *string `minArgs:"1" maxArgs:"-2" definition:"deprecated"`
	UseRoaming            *string `minArgs:"1" maxArgs:"-2" definition:"deprecated"`
	UseRsh                *string `minArgs:"1" maxArgs:"-2" definition:"deprecated"`
	// definition deprecatedHidden: no friendly error :(
	Protocol *string `minArgs:"1" maxArgs:"-2" definition:"deprecatedHidden"`
	// unsupported by OpenSSH, maybe return friendly error
	AFSTokenPassing         *string `minArgs:"1" maxArgs:"-2" definition:"unsupported"`
	CompressionLevel        *string `minArgs:"1" maxArgs:"-2" definition:"unsupported"`
	KerberosAuthentication  *string `minArgs:"1" maxArgs:"-2" definition:"unsupported"`
	KerberosTGTPassing      *string `minArgs:"1" maxArgs:"-2" definition:"unsupported"`
	RhostsRSAAuthentication *string `minArgs:"1" maxArgs:"-2" definition:"unsupported"`
	RSAAuthentication       *string `minArgs:"1" maxArgs:"-2" definition:"unsupported"`
	// ??????
	// TODO: "Macs undocumented  // NEEDINFO?, cipher?
}

// native structs
type CanonicalizeCNAMEs struct {
	sourceDomains []string
	targetDomains []string
}

type DynamicForward struct {
	bindAddress string // "" or "*" for all addresses
	port        uint16
}

type TunnelDevice struct {
	local  string
	remote *string
}

// #### enums ####
type (
	AddressFamily   string // enum: "inet", "inet6", "any"
	Hash            string // enum: "md5", "sha256"
	LogFacility     string // enum: "daemon", "user", "auth", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7"
	LogLevel        string // enum: "quiet", "fatal", "error", "info", "verbose", "debug", "debug1", "debug2", "debug3"
	SessionType     string // enum: "none", "subsystem", "default"
	YesNoAsk        string // enum: "false", "true", "ask"
	YesNoAskConfirm string // enum: "false", "true", "ask", "confirm"
)

type Cipher struct {
	mode    string   // enum: "" (set/override), "+" (append), "^" (preappend), "-" (subtract regex)
	ciphers []string // supports wildcards (*)
}

type PermitRemoteOpen struct {
	permit string  // enum: "none", "any", "defined"
	host   *string // optional for "defined"
	port   *uint16 // required for "defined"
}

// https://ftp.openbsd.org/pub/OpenBSD/OpenSSH/openssh-8.8.tar.gz readconf.c#792: Multistate option parsing

var enumIndex = map[string][]enumValues{
	"Flag": {{"true", true}, {"false", false}, {"yes", true}, {"no", false}},
	"YesNoAsk": {
		{"true", true},
		{"false", false},
		{"yes", true},
		{"no", false},
		{"ask", "ask"},
	},
	"YesNoAskConfirm": {
		{"true", true},
		{"false", false},
		{"yes", true},
		{"no", false},
		{"ask", "ask"},
		{"confirm", "confirm"},
	},
	"ControlMaster": {
		{"true", true},
		{"false", false},
		{"yes", true},
		{"no", false},
		{"auto", "auto"},
		{"ask", "ask"},
		{"autoask", "autoask"},
	},
	"StrictHostkey": {
		{"true", true},
		{"false", false},
		{"yes", true},
		{"no", false},
		{"off", false},
		{"ask", "ask"},
		{"accept-new", "accept-new"},
	},
	"CanonicalizeHostname": {
		{"true", true},
		{"false", false},
		{"yes", true},
		{"no", false},
		{"always", "always"},
	},
	"RequestTTY": {
		{"true", true},
		{"false", false},
		{"yes", true},
		{"no", false},
		{"force", "force"},
		{"auto", "auto"},
	},
	"Tunnel": {
		{"true", true},
		{"false", false},
		{"yes", true},
		{"no", false},
		{"ethernet", "ethernet"},
		{"point-to-point", "point-to-point"},
	},
	"Compression":   {{"yes", true}, {"no", false}},
	"AddressFamily": {{"inet", "inet"}, {"inet6", "inet6"}, {"any", "any"}},
	"SessionType":   {{"none", "none"}, {"subsystem", "subsystem"}, {"default", "default"}},
	"LogLevel":      {{"quiet", "quiet"}, {"fatal", "fatal"}, {"error", "error"}, {"info", "info"}, {"verbose", "verbose"}, {"debug", "debug"}, {"debug1", "debug1"}, {"debug2", "debug2"}, {"debug3", "debug3"}},
	"LogFacility":   {{"daemon", "daemon"}, {"user", "user"}, {"auth", "auth"}, {"local0", "local0"}, {"local1", "local1"}, {"local2", "local2"}, {"local3", "local3"}, {"local4", "local4"}, {"local5", "local5"}, {"local6", "local6"}, {"local7", "local7"}},
	"Hash":          {{"md5", "md5"}, {"sha256", "sha256"}},
}

type enumValues struct {
	stringName string
	value      interface{}
}

// #### custom types
type (
	ForwardAgent   string // "false", "true" is treated as bools
	ControlPersist string // bool or time.Duration
	Ipqos          struct {
		// enum: TODO: see definition
		interactive    string
		nonInteractive *string
	}
)

type Nzint32 int32 // must not be negative

type RekeyLimit struct {
	bytes     int     // 0 for "default", int >= 16
	bytesUnit *string // for non-zero enum: "K", "M", "G"
	time      *time.Duration
}

type Forward string
