package ratelimit

import (
	logs "calendarapi/pkg/logutil"
	"calendarapi/pkg/ratelimiter"
	"calendarapi/pkg/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiterHandler struct {
	Limiter *ratelimiter.Limiter
}

func NewRateLimiterHandler(limiter *ratelimiter.Limiter) *RateLimiterHandler {
	return &RateLimiterHandler{Limiter: limiter}
}

func (h *RateLimiterHandler) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ratelimiter.IsExcludePath(c.Request.URL.Path) {
			c.Next()
			return
		}

		key := c.ClientIP() + ":" + c.FullPath()
		limiterInstance := h.Limiter.GetLimiter(key)

		logs.Debug("client ip : %v", key)
		//logs.Debug("Limiter config => Limit: %d", limiterInstance.Limit(), " Burst: %v", limiterInstance.Burst())
		//logs.Debug("Limiter info for %v", key, " | Tokens available: %v", limiterInstance.Tokens())

		reservation := limiterInstance.Reserve()
		if !reservation.OK() {
			logs.Debug("Rate limit not allowed for %v", key)
			response.SendError(c, http.StatusTooManyRequests, "Too many requests. Please try again later.")
			c.Abort()
			return
		}

		delay := reservation.Delay()
		logs.Debug("Rate limit delay : %v", delay)

		if delay > time.Second {
			reservation.Cancel()
			retryAfter := int(delay.Seconds())
			//logs.Debug("Rate limit exceeded for %v", key, " | Retry after: %v", retryAfter, "s")
			response.SendError(c, http.StatusTooManyRequests, "Too many requests. Please try again in "+strconv.Itoa(retryAfter)+" seconds.")
			c.Abort()
			return
		}

		c.Next()
	}
}
