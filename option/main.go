package main

import (
	"github.com/k0kubun/pp"
)

type Request struct {
	page    int
	perPage int
	sort    string
}

type Options func(request *Request)

func NewRequest(opts ...Options) *Request {
	r := &Request{page: 1, perPage: 30, sort: "desc"}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func Page(p int) Options {
	return func(r *Request) {
		if r != nil {
			r.page = p
		}
	}
}

func PerPage(pp int) Options {
	return func(r *Request) {
		if r != nil {
			r.perPage = pp
		}
	}
}

func Sort(s string) Options {
	return func(r *Request) {
		if r != nil {
			r.sort = s
		}
	}
}

func main() {
	r := NewRequest()
	pp.Println(r)
	r = NewRequest(Page(10))
	pp.Println(r)
	r = NewRequest(Page(10), PerPage(2), Sort("asc"))
	pp.Println(r)
}

