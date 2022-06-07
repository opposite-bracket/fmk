package fmk

type IAuthSdk interface {
	HasPermission(tenantId, userId, roleId, method, path string) bool
}

func RequireTenantToken(auth IAuthSdk) Endpoint {
	return func(c *Context) error {
		//var h Header
		//if err := c.ShouldBindHeader(&h); err != nil {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		//		"code":   AllUnauthorizedHeaderCode,
		//		"errors": GetErrors(err),
		//	})
		//	return
		//}
		//
		//if !auth.HasPermission("ti1", "ui1", "ri1", c.Request.Method, c.FullPath()) {
		//	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		//		"code": InvalidTokenCode,
		//	})
		//	return
		//}
		//
		//c.Set(TenantIdCtx, "ti1")
		//c.Set(UserIdCtx, "ui1")
		//c.Set(RoleIdCtx, "ri1")
		return nil
	}
}
