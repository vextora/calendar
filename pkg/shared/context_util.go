package shared

import "github.com/gin-gonic/gin"

const TraceIDKey = "trace_id"
const ContextKeyUserID = "userID"

func GetTraceID(c *gin.Context) string {
	if traceID, exists := c.Get(TraceIDKey); exists {
		if idStr, ok := traceID.(string); ok {
			return idStr
		}
	}
	return ""
}
