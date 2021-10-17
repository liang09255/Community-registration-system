module newFunction

go 1.14

require (
	connectDB v0.0.0
	github.com/alibabacloud-go/darabonba-openapi v0.1.5
	github.com/alibabacloud-go/dysmsapi-20170525/v2 v2.0.1
	github.com/alibabacloud-go/tea v1.1.16
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/robfig/cron v1.2.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

replace connectDB => ../connectDB
