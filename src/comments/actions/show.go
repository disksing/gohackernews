package commentactions

import (
	"github.com/disksing/gohackernews/src/comments"
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"
)

// HandleShow displays a single comment
func HandleShow(context router.Context) error {

	// Find the comment
	comment, err := comments.Find(context.ParamInt("id"))
	if err != nil {
		return router.InternalError(err)
	}

	// No auth as all are public - if we restricted by status we might need to authorise here

	// Render the template
	view := view.New(context)
	view.AddKey("comment", comment)

	return view.Render()
}
