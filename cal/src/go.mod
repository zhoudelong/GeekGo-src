module src // 在统一目录下的go.mod 配置

go 1.17  // 版本

require  "moduledemo" v0.0.0
replace "moduledemo" => "../moduledemo" // replace指令
