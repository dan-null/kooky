package kooky

import (
	"testing"
	"time"
)

// d18f6247db68045dfbab126d814baf2cf1512141391
func TestReadSafariCookies(t *testing.T) {
	cookies, err := ReadSafariCookies("testdata/small-safari-cookie-db.binarycookies", "", "", time.Time{})
	if err != nil {
		t.Fatal(err)
	}

	domain := "news.ycombinator.com"
	name := "user"
	cookie := findCookie(domain, name, cookies)
	if cookie == nil {
		t.Fatalf("Found no cookie with domain=%q, name=%q", domain, name)
	}
	wantValue := "zellyn&EdK9mzRM38fGtIZQTiqCyAeWg93RDjdo"
	if cookie.Value != wantValue {
		t.Errorf("Want cookie value %q; got %q", wantValue, cookie.Value)
	}

	wantExpires := time.Date(2038, 01, 17, 19, 14, 07, 0, time.UTC)
	if !cookie.Expires.Equal(wantExpires) {
		t.Errorf("Want cookie.Expires=%v; got %v", wantExpires, cookie.Expires)
	}

	wantCreation := time.Date(2017, 12, 16, 23, 23, 19, 0, time.UTC)
	if !cookie.Creation.Equal(wantCreation) {
		t.Errorf("Want cookie.Creation=%v; got %v", wantCreation, cookie.Creation)
	}
}
