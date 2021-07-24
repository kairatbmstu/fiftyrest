package main

type Config struct {
	// DEFAULT_CONNECTION_TIMEOUT int //  = 10000;
	// int DEFAULT_MAX_CONNECTIONS = 200;
	// int DEFAULT_MAX_PER_ROUTE = 20;
	// int DEFAULT_CONNECT_TIMEOUT = 10000;
	// int DEFAULT_SOCKET_TIMEOUT = 60000;

	// Client client;
	// private Optional<AsyncClient> asyncClient = Optional.empty();
	// private Optional<ObjectMapper> objectMapper = Optional.of(new JsonObjectMapper());

	// private List<HttpRequestInterceptor> apacheinterceptors = new ArrayList<>();
	// private Headers headers;
	// private Proxy proxy;
	ConnectionTimeout       int
	SocketTimeout           int
	MaxTotal                int
	MaxPerRoute             int
	FollowRedirects         bool
	CookieManagement        bool
	UseSystemProperties     bool   // default value is true
	defaultResponseEncoding string // = StandardCharsets.UTF_8.name();
	// private Function<Config, AsyncClient> asyncBuilder;
	// private Function<Config, Client> clientBuilder;
	RequestCompressionOn bool // default = true;
	AutomaticRetries     bool
	VerifySsl            bool // default = true;
	// private boolean addShutdownHook = false;
	// private KeyStore keystore;
	// private Supplier<String> keystorePassword = () -> null;
	// private String cookieSpec;
	// private UniMetric metrics = new NoopMetric();
	ttl int64 //default value = -1;
	// private SSLContext sslContext;
	// private String[] ciphers;
	// private String[] protocols;
	// private CompoundInterceptor interceptor = new CompoundInterceptor();
	// private HostnameVerifier hostnameVerifier;
	DefaultBaseUrl string
	// private CacheManager cache;

}

func NewDefaultConfig() {

}
