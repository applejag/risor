module github.com/risor-io/risor/modules/cli

go 1.22

toolchain go1.22.0

replace github.com/risor-io/risor => ../..

require (
	github.com/risor-io/risor v1.3.2
	github.com/urfave/cli/v2 v2.27.1
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20231213231151-1d8dd44e695e // indirect
)
