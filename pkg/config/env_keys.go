package config

const (
	Environment = "ENV"
	AppName     = "APP_NAME"
	AppPort     = "APP_PORT"

	PostgreDbHost     = "POSTGRE_DB_HOST"
	PostgreDbPort     = "POSTGRE_DB_PORT"
	PostgreDbUser     = "POSTGRE_DB_USER"
	PostgreDbPassword = "POSTGRE_DB_PASSWORD"
	PostgreDbSsl      = "POSTGRE_DB_SSL"
	PostgreDbName     = "POSTGRE_DB_NAME"

	SentryDsn = "SENTRY_DSN"

	JwtSecret       = "JWT_SECRET"
	JwtTokenExpired = "JWT_TOKEN_EXPIRED"

	RateLimitPerSecond  = "RATE_LIMIT_PER_SECOND"
	RateLimitBurst      = "RATE_LIMIT_BURST"
	RateLimitTtlSeconds = "RATE_LIMIT_TTL_SECONDS"

	TracerServiceName = "TRACER_SERVICE_NAME"
	TracerEnv         = "TRACER_ENV"
	TracerProtocol    = "TRACER_PROTOCOL"
	TracerEndpoint    = "TRACER_ENDPOINT"

	ZapOutputPath      = "ZAP_OUTPUT_PATH"
	ZapErrorOutputPath = "ZAP_ERROR_OUTPUT_PATH"

	ShXFrameOptions           = "SH_X_FRAME_OPTIONS"
	ShXContentTypeOptions     = "SH_X_CONTENT_TYPE_OPTIONS"
	ShXxssProtection          = "SH_XXSS_PROTECTION"
	ShContentSecurityPolicy   = "SH_CONTENT_SECURITY_POLICY"
	ShReferrerPolicy          = "SH_REFERRER_POLICY"
	ShStrictTransportSecurity = "SH_STRICT_TRANSPORT_SECURITY"
)
