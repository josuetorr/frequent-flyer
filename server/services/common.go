package services

import (
	"github.com/josuetorr/frequent-flyer/server/data"
	"github.com/josuetorr/frequent-flyer/server/models"
)

type (
	ID             = models.ID
	User           = models.User
	UserRepository = data.Repository[User]
)
