package handler

import "strings"

func keyGenerate(campaignID string, idType string, id ...string) string {
	return campaignID + ":" + idType + ":" + strings.Join(id, "/")
}
