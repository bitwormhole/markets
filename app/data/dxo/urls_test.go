package dxo

import "testing"

func TestURL(t *testing.T) {

	u1 := URL("http://www.example.com/u1")
	u2 := URL("https://example.com/u2")
	u3 := URL("u3://abc/def")

	ul := URLList("")

	ul = ul.Append(u1)
	ul = ul.Append(u2)
	ul = ul.Append(u3)

	ul = ul.Normalize()

	all := ul.List()

	for i, u := range all {

		t.Logf("URL[%d] = %s", i, u)

		res, err := u.ParseSelf(nil)
		if err == nil {
			t.Logf(" url.parsed = %v", res)
		} else {
			t.Log(err.Error())
		}

	}

}
