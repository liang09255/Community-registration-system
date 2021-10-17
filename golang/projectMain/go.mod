module projectMain

go 1.14

require (
	connectDB v0.0.0
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/thedevsaddam/gojsonq v2.3.0+incompatible // indirect
	newFunction v0.0.0
	backFunction v0.0.0
)

replace (
	connectDB => ../connectDB
	newFunction => ../newFunction
	backFunction => ../../backstage/golang/backFunction
)
