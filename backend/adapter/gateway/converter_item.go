package gateway

import (
	"encoding/base64"
	"strings"
)

func ToGatewayItemHolderSortKey(domainSortKey, defaultKey string) string {
	sortKey := defaultKey
	// TODO: camel -> snake 関数を探す！
	switch domainSortKey {
	case "firstName":
		sortKey = "first_name"
	case "lastName":
		sortKey = "last_name"
	case "nickame":
		sortKey = "nickname"
	}
	return sortKey
}

func DecodeCursor(cursor *string) (string, string, error) {
	if cursor == nil {
		return "", "", nil
	}

	b, err := base64.StdEncoding.DecodeString(*cursor)
	if err != nil {
		return "", "", err
	}
	splits := strings.Split(string(b), "=")
	return splits[0], splits[1], nil
}
