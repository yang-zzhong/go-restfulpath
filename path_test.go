package restfulpath

import (
	"testing"
)

func TestSimple(t *testing.T) {

	p := &Path{"/users"}
	matched, _ := p.Match("/users")
	if !matched {
		t.Error("/users Error")
	}
}

func TestParams(t *testing.T) {
	p := &Path{"/users/:name"}
	matched, params := p.Match("/users/yang-zzhong")
	if !matched {
		t.Error("/users/:name Not Matched")
	}
	if params["name"] != "yang-zzhong" {
		t.Error("/users/:name Params Not Got")
	}
}

func TestComplex(t *testing.T) {
	p := &Path{"/users/:name/articles"}
	matched, params := p.Match("/users/yang-zzhong/articles")
	if !matched {
		t.Error("/users/:name/articles Not Matched")
	}
	if params["name"] != "yang-zzhong" {
		t.Error("/users/:name/articles Params Not Got")
	}
}

func TestAllParams(t *testing.T) {
	p := &Path{"/:name/:article-name"}
	matched, params := p.Match("/yang-zzhong/Me+And+My+Broken+Heart")
	if !matched {
		t.Error("/:name/:article-name Not Matched")
	}
	para := make(map[string]string)
	para["name"] = "yang-zzhong"
	para["article-name"] = "Me+And+My+Broken+Heart"
	for k, v := range params {
		if para[k] != v {
			t.Error("/:name/:article-name Params Not Got")
			return
		}
	}
}
