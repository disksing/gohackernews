package useractions

import (
	"github.com/disksing/gohackernews/src/lib/authorise"
	"github.com/disksing/gohackernews/src/users"
	"github.com/fragmenta/router"
)

// HandleDestroy responds to POST /users/1/destroy
func HandleDestroy(context router.Context) error {

	// Set the user on the context for checking
	user, err := users.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise
	err = authorise.ResourceAndAuthenticity(context, user)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Destroy the user
	user.Destroy()

	// Redirect to users root
	return router.Redirect(context, user.URLIndex())
}
