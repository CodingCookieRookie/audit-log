package ctrl

import (
	"os"

	"github.com/CodingCookieRookie/audit-log/constants"
	"github.com/CodingCookieRookie/audit-log/log"
)

func PostUser(email, apiKey string) error {

	if apiKey == constants.GetAPIKey() {
		log.Infof("api: %v", os.Getenv("API_KEY"))
	}

	return nil
}
