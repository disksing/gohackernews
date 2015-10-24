package commentactions

import (
	"github.com/disksing/gohackernews/src/comments"
	"github.com/disksing/gohackernews/src/lib/authorise"
	"github.com/fragmenta/router"
)

// HandleDestroy handles a DESTROY request for comments
func HandleDestroy(context router.Context) error {

	// Find the comment
	comment, err := comments.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise destroy comment
	err = authorise.Resource(context, comment)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Destroy the comment
	comment.Destroy()

	// Redirect to comments root
	return router.Redirect(context, comment.URLIndex())
}
