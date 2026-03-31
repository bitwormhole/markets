package dxo

import (
	"net/url"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

// 是一个表示 "统一资源定位符" 的字符串
type URL string

func (u URL) String() string {
	return string(u)
}

func (u URL) ParseSelf(dst *url.URL) (*url.URL, error) {
	str := u.String()
	return url.Parse(str)
}

func (u URL) IsEmpty() bool {
	return (u == "")
}

func (u URL) Normalize() URL {
	str := u.String()
	str = strings.TrimSpace(str)
	return URL(str)
}

////////////////////////////////////////////////////////////////////////////////

// 是一个多行文本, 表示若干个 URL
type URLList string

func (list URLList) String() string {
	return string(list)
}

func (list URLList) Normalize() URLList {

	src := list.List()
	dst := URLList("")

	for _, item := range src {
		item = item.Normalize()
		if item.IsEmpty() {
			continue
		}
		dst.Append(item)
	}

	return dst
}

func (list URLList) Append(item URL) URLList {
	s1 := list.String()
	s2 := item.String()
	if s1 == "" {
		return URLList(s2)
	}
	if s2 == "" {
		return URLList(s1)
	}
	return URLList(s1 + "\n" + s2)
}

func (list URLList) List() []URL {

	const (
		sep1 = '\r'
		sep2 = '\n'
	)
	str := list.String()

	a1 := strings.Split(str, string(sep2))
	a2 := make([]URL, 0)

	for _, item1 := range a1 {
		item2 := URL(item1)
		if !item2.IsEmpty() {
			a2 = append(a2, item2)
		}
	}

	return a2
}

////////////////////////////////////////////////////////////////////////////////
