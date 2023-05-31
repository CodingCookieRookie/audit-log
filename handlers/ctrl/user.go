package ctrl

import (
	"github.com/CodingCookieRookie/audit-log/constants"
	"github.com/CodingCookieRookie/audit-log/log"
)

const ()

func PostUser(email, apiKey string) error {

	if apiKey != constants.GetAPIKey() {
		log.Infof("api key does not match, api key provided: %v", apiKey)
		return nil
	}

	return nil
}
