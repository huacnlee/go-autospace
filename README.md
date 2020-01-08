# autospace

Automatically add spaces between Chinese and English words.

This is a go version of [auto-correct](https://github.com/huacnlee/auto-correct).

## Usage

```go
package main

import "gthub.com/huacnlee/go-autospace"

func main() {
  autospace.Format("长桥LongBridge App下载")
  // => "长桥 LongBridge App 下载"

  autospace.Format("Ruby 2.7版本第1次发布")
  // => "Ruby 2.7 版本第 1 次发布"

  autospace.Format("于3月10日开始")
  // => "于 3 月 10 日开始"

    autospace.Format("包装日期为2013年3月10日")
  // => "包装日期为2013年3月10日"
}
```

## Benchmark

| Total chars | Duration |
| ----- | ------- |
| 50  | 0.09 ms |
| 100  | 0.14 ms |
| 400  | 0.39 ms |

## License

This project under MIT license.
