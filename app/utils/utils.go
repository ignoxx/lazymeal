package utils

import (
	"fmt"
	"strings"
)

func ImageURL(imageURL string) string {
	if strings.HasPrefix(imageURL, "https://") {
		return imageURL
	}

	return fmt.Sprintf("/public/%s", imageURL)
}
