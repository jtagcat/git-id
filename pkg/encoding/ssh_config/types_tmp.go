package ssh_config

var keywordIndex = map[string]keywordIndexType{
	// -2: untested / unknown / TODO:
	// ##### enums #####
	"AddressFamily":                    {1, -2, "AddressFamily"},
	"CanonicalizeHostname":             {1, -2, "CanonicalizeHostname"},
	"Compression":                      {1, -2, "Compression"},
	"ControlMaster":                    {1, -2, "ControlMaster"},
	"KbdInteractiveAuthentication":     {1, -2, "Flag"}, // has aliases
	"challengeresponseauthentication":  {1, -2, "Flag"}, // alias KbdInteractiveAuthentication
	"skeyauthentication":               {1, -2, "Flag"}, // alias KbdInteractiveAuthentication
	"tisauthentication":                {1, -2, "Flag"}, // alias KbdInteractiveAuthentication
	"PubkeyAuthentication":             {1, -2, "Flag"}, // has aliases
	"DSAAuthentication":                {1, -2, "Flag"}, // alias PubkeyAuthentication
	"BatchMode":                        {1, -2, "Flag"},
	"CanonicalizeFallbackLocal":        {1, -2, "Flag"},
	"CheckHostIP":                      {1, -2, "Flag"},
	"ClearAllForwardings":              {1, -2, "Flag"},
	"EnableSSHKeysign":                 {1, -2, "Flag"},
	"ExitOnForwardFailure":             {1, -2, "Flag"},
	"ForkAfterAuthentication":          {1, -2, "Flag"},
	"ForwardX11":                       {1, -2, "Flag"},
	"ForwardX11Trusted":                {1, -2, "Flag"},
	"GatewayPorts":                     {1, -2, "Flag"},
	"GssAuthentication":                {1, -2, "Flag"},
	"GssDelegateCreds":                 {1, -2, "Flag"},
	"HashKnownHosts":                   {1, -2, "Flag"},
	"HostbasedAuthentication":          {1, -2, "Flag"},
	"IdentitiesOnly":                   {1, -2, "Flag"},
	"NoHostAuthenticationForLocalhost": {1, -2, "Flag"},
	"PasswordAuthentication":           {1, -2, "Flag"},
	"PermitLocalCommand":               {1, -2, "Flag"},
	"ProxyUseFdpass":                   {1, -2, "Flag"},
	"StdinNull":                        {1, -2, "Flag"},
	"StreamLocalBindUnlink":            {1, -2, "Flag"},
	"TCPKeepAlive":                     {1, -2, "Flag"},
	"VisualHostKey":                    {1, -2, "Flag"},
	"FingerprintHash":                  {1, -2, "Hash"},
	"LogFacility":                      {1, -2, "LogFacility"},
	"LogLevel":                         {1, -2, "LogLevel"},
	"RequestTTY":                       {1, -2, "RequestTTY"},
	"SessionType":                      {1, -2, "SessionType"},
	"StrictHostKeyChecking":            {1, -2, "StrictHostkey"},
	"Tunnel":                           {1, -2, "Tunnel"},
	"UpdateHostkeys":                   {1, -2, "YesNoAsk"},
	"VerifyHostKeyDNS":                 {1, -2, "YesNoAsk"},
	"AddKeysToAgent":                   {1, -2, "YesNoAskConfirm"},

	// #### custom types #####
	// Flag or string
	"ForwardAgent": {1, -2, "boolString"},
	// all slice items are in ssh_config[0] (len == 1), seperated by comma
	"GlobalKnownHostsFile":     {1, -2, "csvStringSlice"},
	"IgnoreUnknown":            {1, -2, "csvStringSlice"},
	"KbdInteractiveDevices":    {1, -2, "csvStringSlice"},
	"PreferredAuthentications": {1, -2, "csvStringSlice"},
	"UserKnownHostsFile":       {1, -2, "csvStringSlice"},
	// csvStringSlice, but each set is 2tuplet x:y (colon-seperated)
	"CanonicalizePermittedCNAMEs": {1, -2, "canonicalizeCNAMEs"},
	/* csvStringSlice, but begins with "" (set/override), + (append), ^ (preappend), - (subtract regex)
	*  - supports wildtags (*), NEEDINFO:{-2, -2, are wildcards supported on set/additions?
	* ciphers shall not be enummed, as they are independent of the config, and change often */
	"CASignatureAlgorithms":       {1, -2, "cipher"},
	"Ciphers":                     {1, -2, "cipher"},
	"HostbasedAcceptedAlgorithms": {1, -2, "cipher"},
	"HostKeyAlgorithms":           {1, -2, "cipher"},
	"KexAlgorithms":               {1, -2, "cipher"},
	"PubkeyAcceptedAlgorithms":    {1, -2, "cipher"},
	/* Format is a sequence of:
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
	*      1h30m   90 minutes */
	"ConnectTimeout":    {1, -2, "duration"},
	"ForwardX11Timeout": {1, -2, "duration"},
	// Flag or duration
	"ControlPersist": {1, -2, "controlPersist"},
	// [host:]port
	"DynamicForward": {1, -2, "dynamicForward"},
	// Key="echo" "hi" are joined to "echo hi"
	"KnownHostsCommand": {1, -2, "indifferentString"},
	"LocalCommand":      {1, -2, "indifferentString"},
	"ProxyCommand":      {1, -2, "indifferentString"},
	"RemoteCommand":     {1, -2, "indifferentString"},
	/* Accepted values are af11, af12, af13, af21, af22, af23, af31, af32, af33, af41, af42, af43,
	* cs0, cs1, cs2, cs3, cs4, cs5, cs6, cs7, ef, le, lowdelay, throughput, reliability,
	* a numeric value, or none to use the operating system default. */
	"IPQoS": {1, 2, "ipqos"}, // default []{"af21", "cs1"}
	// specifying key-value multiple times as a way of expressing slice
	"CertificateFile": {1, -2, "multiDefineStringSlice"},
	"IdentityFile":    {1, -2, "multiDefineStringSlice"},
	// non-zero int32:{-2, -2, 0..0x7fffffff (2147483647)
	"CanonicalizeMaxDots":     {1, -2, "nzint32"},
	"ConnectionAttempts":      {1, -2, "nzint32"},
	"NumberOfPasswordPrompts": {1, -2, "nzint32"},
	"ServerAliveCountMax":     {1, -2, "nzint32"},
	"ServerAliveInterval":     {1, -2, "nzint32"},
	// none / any / host:port / :port
	"PermitRemoteOpen": {1, -2, "permitRemoteOpen"},
	// permoctal: 0777 (4 digits in base8)
	"StreamLocalBindMask": {1, -2, "permoctal"},
	// 1st: "default" or int[qualifier], qualifier:K/M/G; 2nd (optional): duration (or 'none')
	// 1st: 1..15: RekeyLimit too small
	"RekeyLimit": {1, 2, "rekeyLimit"},
	// slice in go, slice in conf
	"CanonicalDomains": {1, -2, "stringSlice"},
	"LogVerbose":       {1, -2, "stringSlice"},
	"SendEnv":          {1, -2, "stringSlice"},
	"SetEnv":           {1, -2, "stringSlice"},
	// local_tun[:remote_tun]
	"TunnelDevice": {1, -2, "tunnelDevice"},

	// ##### standard types #####
	// string also includes keywords where type checking isn't possible outside of config time of use
	"BindAddress":         {1, -2, "string"},
	"BindInterface":       {1, -2, "string"},
	"ControlPath":         {1, -1, "string"},
	"EscapeChar":          {1, -2, "string"},
	"HostKeyAlias":        {1, -2, "string"},
	"Hostname":            {1, -2, "string"},
	"IdentityAgent":       {1, -2, "string"},
	"PKCS11Provider":      {1, -2, "string"},
	"RevokedHostKeys":     {1, -2, "string"},
	"SecurityKeyProvider": {1, -2, "string"},
	"User":                {1, -2, "string"},
	"XAuthLocation":       {1, -2, "string"},
	"Port":                {1, -2, "uint16"},
	/* dynamicfwd == 0
	 *   [listenhost:]listenport|listenpath:connecthost:connectport|connectpath
	 *   listenpath:connectpath
	 * dynamicfwd == 1
	 *	[listenhost:]listenport
	 */
	//
	"LocalForward":  {1, 2, "forward"}, // type checking very hard, almost alias to string
	"RemoteForward": {1, 2, "forward"}, // type checking very hard, almost alias to string

	// ##### not availiable ##### (treated as string)
	// deprecated by OpenSSH, maybe return friendly error?
	"Cipher":                {1, -2, "deprecated"},
	"FallbackToRsh":         {1, -2, "deprecated"},
	"GlobalKnownHostsFile2": {0, -2, "deprecated"},
	"RhostsAuthentication":  {1, -2, "deprecated"},
	"UsePrivilegedPort":     {1, -2, "deprecated"},
	"UserKnownHostsFile2":   {1, -2, "deprecated"},
	"UseRoaming":            {1, -2, "deprecated"},
	"UseRsh":                {1, -2, "deprecated"},
	// no friendly error :(
	"Protocol": {1, -2, "deprecatedHidden"},
	// unsupported by OpenSSH, maybe return friendly error
	"AFSTokenPassing":         {1, -2, "unsupported"},
	"CompressionLevel":        {1, -2, "unsupported"},
	"KerberosAuthentication":  {1, -2, "unsupported"},
	"KerberosTGTPassing":      {1, -2, "unsupported"},
	"RhostsRSAAuthentication": {1, -2, "unsupported"},
	"RSAAuthentication":       {1, -2, "unsupported"},
	// ??????
	//TODO: "Macs": {1, -2, "undocumented"}, // NEEDINFO?, cipher?
}
