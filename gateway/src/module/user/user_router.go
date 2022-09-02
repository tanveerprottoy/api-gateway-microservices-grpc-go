package user

/* import (
	"txp/web-service-gin/src"
	"txp/web-service-gin/src/util"
)

func RegisterUserRoutes(
	router *src.Router,
	version string,
) {
	handler := &UserHandler{}
	handler.InitDependencies()
	group := router.Engine.Group(
		util.ApiPattern + version + util.UsersPattern,
	)
	{
		group.GET(
			util.RootPattern,
			handler.FindUsers,
		)
		group.GET(
			util.RootPattern + ":id",
			handler.FindUsers,
		)
	}
} */
