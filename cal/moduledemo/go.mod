module moduledemo

go 1.17

require "src" v0.0.0  //调用 src 下的包, 配置
replace "src" => "../src" // replace指令

