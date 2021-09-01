package magnetinfo

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Contains all magnet link parameters
type MagnetInfo struct {
	DisplayNames      []string // DN
	ExactLength       uint64   // XL
	ExactTopics       []string // XT
	WebSeeds          []string // WS
	AcceptableSources []string // AS
	ExactSources      []string // XS
	KeywordTopics     []string // KT
	ManifestTopics    []string // MT
	AddressTrackers   []string // TR
}

// Parses a magnet link and returns its info
func Parse(magnetLink string) (*MagnetInfo, error) {
	// split a magnet link identifier and key-value pairs
	parts := strings.Split(magnetLink, ":?")

	if len(parts) == 1 || parts[0] != "magnet" {
		return nil, fmt.Errorf("given string does not contain a \"magnet\" identifier")
	}

	// get all key-value pairs
	pairs := strings.Split(parts[1], "&")

	var info MagnetInfo
	for _, pair := range pairs {
		// get key and its value
		keyval := strings.Split(pair, "=")

		key := keyval[0]
		value := keyval[1]

		switch strings.ToLower(key) {
		case "dn":
			value, err := url.QueryUnescape(value)
			if err != nil {
				return nil, err
			}
			info.DisplayNames = append(info.DisplayNames, value)

		case "xl":
			length, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			info.ExactLength = length

		case "xt":
			info.ExactTopics = append(info.ExactTopics, value)

		case "ws":
			info.WebSeeds = append(info.WebSeeds, value)

		case "as":
			value, err := url.QueryUnescape(value)
			if err != nil {
				return nil, err
			}
			info.AcceptableSources = append(info.AcceptableSources, value)

		case "xs":
			info.ExactSources = append(info.ExactSources, value)

		case "kt":
			value, err := url.QueryUnescape(value)
			if err != nil {
				return nil, err
			}
			info.KeywordTopics = append(info.KeywordTopics, value)

		case "mt":
			value, err := url.QueryUnescape(value)
			if err != nil {
				return nil, err
			}
			info.ManifestTopics = append(info.ManifestTopics, value)

		case "tr":
			value, err := url.QueryUnescape(value)
			if err != nil {
				return nil, err
			}
			info.AddressTrackers = append(info.AddressTrackers, value)
		}
	}

	return &info, nil
}
