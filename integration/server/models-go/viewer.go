package models

import "github.com/geneva/gqlgen/integration/server/remote_api"

type Viewer struct {
	User *remote_api.User
}
