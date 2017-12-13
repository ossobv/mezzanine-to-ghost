package main

import (
	"errors"
	"net/url"
)

func IsURL(val interface{}) error {
	// Should be a string
	if _, ok := val.(string); !ok {
		return errors.New("Not a string")
	}

	link := val.(string)
	u, err := url.Parse(link)

	// link should conform RFC 3986 and have a scheme
	if err != nil {
		//return errors.New("Not a valid url")
		return err
	}

	// scheme should be "http" or "https"
	if len(u.Scheme) > 0 {
		if u.Scheme != "http" && u.Scheme != "https" {
			return errors.New("Not a valid scheme")
		}
	}

	// TODO: check gTLD?

	return nil
}
