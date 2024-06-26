module github.com/Thwani47/go-roadmap/modules/hello

go 1.22.1

require (
	github.com/Thwani47/go-roadmap/modules/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote/v3 v3.1.0 // indirect
)

require (
	golang.org/x/text v0.16.0 // indirect
	rsc.io/sampler v1.3.1 // indirect
)

replace github.com/Thwani47/go-roadmap/modules/greetings => ../greetings
