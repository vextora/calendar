package securityheader

type Options struct {
	XFrameOptions           string
	XContentTypeOptions     string
	XXSSProtection          string
	ContentSecurityPolicy   string
	ReferrerPolicy          string
	StrictTransportSecurity string
}

func ApplySecurityHeaders(headers map[string]string, opts Options) {
	if opts.XFrameOptions != "" {
		headers["X-Frame-Options"] = opts.XFrameOptions
	}
	if opts.XContentTypeOptions != "" {
		headers["X-Content-Type-Options"] = opts.XContentTypeOptions
	}
	if opts.XXSSProtection != "" {
		headers["X-XSS-Protection"] = opts.XXSSProtection
	}
	if opts.ContentSecurityPolicy != "" {
		headers["Content-Security-Policy"] = opts.ContentSecurityPolicy
	}
	if opts.ReferrerPolicy != "" {
		headers["Referrer-Policy"] = opts.ReferrerPolicy
	}
	if opts.StrictTransportSecurity != "" {
		headers["Strict-Transport-Security"] = opts.StrictTransportSecurity
	}
}
