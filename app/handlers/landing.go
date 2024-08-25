package handlers

import (
	"lazymeal/app/views/home"

	"github.com/anthdm/superkit/kit"
)

func HandleLandingIndex(kit *kit.Kit) error {
	return kit.Render(home.Index())
}
