module example.com/hello

go 1.19

replace example.com/greetings => ../greetings

replace example.com/module1 => ../module1

replace example.com/module2 => ../module2

replace example.com/module3 => ../module3

require example.com/module4 v0.0.0-00010101000000-000000000000

replace example.com/module4 => ../module4
