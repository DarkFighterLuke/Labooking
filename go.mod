module Labooking

go 1.17

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/smartystreets/goconvey v1.6.4
)

replace github.com/beego/beego/v2 => github.com/DarkFighterLuke/beego/v2 v2.0.2-0.20211001174519-05b8436ed5cd

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/wagslane/go-password-validator v0.3.0
)
