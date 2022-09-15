module example.com/hello

go 1.19

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	example.com/module1 v0.0.0-00010101000000-000000000000
	example.com/module2 v0.0.0-00010101000000-000000000000
)

replace example.com/module1 => ../module1

replace example.com/module2 => ../module2
