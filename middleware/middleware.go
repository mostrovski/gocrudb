package middleware

import (
	"gocrudb/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	rateLimiter := createRateLimiter()

	return func(c *gin.Context) {
		if !rateLimiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests: try again later"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func createRateLimiter() *rate.Limiter {
	rps, _ := strconv.Atoi(config.Get("rate_limiter_requests_per_second"))
	rbs, _ := strconv.Atoi(config.Get("rate_limiter_requests_burst_size"))
	return rate.NewLimiter(rate.Limit(rps), rbs)
}
