package utils

import "net/url"

type Form struct {
	Fields   url.Values
	Messages []*Messages
}
