package svcinfra

func RequireLogin(c *Context) {
	if c.User == nil {
		c.RequireLogin("")
		return
	}
	c.Next()
}
