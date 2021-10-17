module backFunction

go 1.14

require (
	connectDB v0.0.0
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/gorilla/sessions v1.2.1
	newFunction v0.0.0
)

replace (
	connectDB => ../../../golang/connectDB
	newFunction => ../../../golang/newFunction
)
