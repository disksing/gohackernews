package useractions

import (
	"strings"

	"github.com/disksing/gohackernews/src/lib/authorise"
	"github.com/disksing/gohackernews/src/lib/status"
	"github.com/disksing/gohackernews/src/users"
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"
)

// HandleCreateShow handles GET /users/create
func HandleCreateShow(context router.Context) error {

	// No auth as anyone can create users in this app

	// Setup
	view := view.New(context)
	user := users.New()
	view.AddKey("user", user)

	// Serve
	return view.Render()
}

// HandleCreate handles POST /users/create from the register page
func HandleCreate(context router.Context) error {

	// Check csrf token
	err := authorise.AuthenticityToken(context)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Setup context
	params, err := context.Params()
	if err != nil {
		return router.InternalError(err)
	}

	// Check for email duplicates
	email := params.Get("email")
	if len(email) > 0 {

		if len(email) < 3 || !strings.Contains(email, "@") {
			return router.InternalError(err, "邮箱格式不正确", "请填入正确的邮箱，或者留空。")
		}

		count, err := users.Query().Where("email=?", email).Count()
		if err != nil {
			return router.InternalError(err)
		}
		if count > 0 {
			return router.NotAuthorizedError(err, "已注册过的邮箱", "抱歉，该邮箱已经被使用。")
		}
	}

	// Check for invalid or duplicate names
	name := params.Get("name")
	if len(name) < 2 {
		return router.InternalError(err, "用户名过短", "请输入至少2字符。")
	}

	count, err := users.Query().Where("name=?", name).Count()
	if err != nil {
		return router.InternalError(err)
	}
	if count > 0 {
		return router.NotAuthorizedError(err, "用户名被占用", "抱歉，该用户名已经被占用，请选择其他用户名。")
	}

	// Set some defaults for the new user
	params.SetInt("status", status.Published)
	params.SetInt("role", users.RoleReader)
	params.SetInt("points", 1)

	allows := append(users.AllowedParams(), "status", "role", "points")

	// Now try to create the user
	id, err := users.Create(params.Clean(allows))
	if err != nil {
		return router.InternalError(err, "错误", "创建用户记录的过程中发生错误。")
	}

	context.Logf("#info Created user id,%d", id)

	// Find the user again so we can save login
	user, err := users.Find(id)
	if err != nil {
		context.Logf("#error parsing user id: %s", err)
		return router.NotFoundError(err)
	}

	// Save the fact user is logged in to session cookie
	err = loginUser(context, user)
	if err != nil {
		return router.InternalError(err)
	}

	// Redirect to root
	return router.Redirect(context, "/?message=welcome")
}
