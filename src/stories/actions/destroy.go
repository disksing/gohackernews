package storyactions

import (
	"github.com/disksing/gohackernews/src/lib/authorise"
	"github.com/disksing/gohackernews/src/stories"
	"github.com/fragmenta/router"
)

// HandleDestroy handles a DESTROY request for stories
func HandleDestroy(context router.Context) error {

	// Find the story
	story, err := stories.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise destroy story
	err = authorise.Resource(context, story)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Destroy the story
	story.Destroy()

	// Redirect to stories root
	return router.Redirect(context, story.URLIndex())
}
