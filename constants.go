package proxylist

type Key string

const (
	KeyExceptionsList            Key = "ExceptionsList"
	KeyExcludeSimpleHostnames        = "ExcludeSimpleHostnames"
	KeyFTPEnable                     = "FTPEnable"
	KeyFTPPassive                    = "FTPPassive"
	KeyFTPPort                       = "FTPPort"
	KeyFTPProxy                      = "FTPProxy"
	KeyGopherEnable                  = "GopherEnable"
	KeyGopherPort                    = "GopherPort"
	KeyGopherProxy                   = "GopherProxy"
	KeyHTTPEnable                    = "HTTPEnable"
	KeyHTTPPort                      = "HTTPPort"
	KeyHTTPProxy                     = "HTTPProxy"
	KeyHTTPSEnable                   = "HTTPSEnable"
	KeyHTTPSPort                     = "HTTPSPort"
	KeyHTTPSProxy                    = "HTTPSProxy"
	KeyRTSPEnable                    = "RTSPEnable"
	KeyRTSPPort                      = "RTSPPort"
	KeyRTSPProxy                     = "RTSPProxy"
	KeySOCKSEnable                   = "SOCKSEnable"
	KeySOCKSPort                     = "SOCKSPort"
	KeySOCKSProxy                    = "SOCKSProxy"
	KeyProxyAutoConfigEnable         = "ProxyAutoConfigEnable"
	KeyProxyAutoConfigJavaScript     = "ProxyAutoConfigJavaScript"
	KeyProxyAutoConfigURLString      = "ProxyAutoConfigURLString"
	KeyProxyAutoDiscoveryEnable      = "ProxyAutoDiscoveryEnable"
)
