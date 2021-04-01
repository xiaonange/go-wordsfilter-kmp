# go-wordsfilter-kmp
A high performance text filter for KMP
敏感词/关键字过滤之KMP算法实现
注意：大多数情况下比BM算法要慢
## Download & Install
```shell
go get github.com/xiaonange/go-wordsfilter-kmp
```

## Quick Start
```go
import (
    "github.com/xiaonange/go-wordsfilter-kmp"
)

func main() {
	searchWord := `作用域`
	kmp, err := new(go_wordsfilter_kmp.WordsFilter).Create().
		SetStripSpace(true).
		SetPlaceholder("*******").
		SetText([]string{"作用域", "禁止"}).
		Add("敏感词").
		ReadWithFile(`D:\Gopath\src\kmp\go-wordsfilter-kmp\word.txt`)
	if err != nil {
		fmt.Println(err)
	}
	kmp.Contains(searchWord)
	kmp.Replace(searchWord)
	kmp.Remove(searchWord)

    // 支持三种设置敏感词模式
		kmp.
		SetText([]string{"作用域", "禁止"}).
		Add("敏感词").
		ReadWithFile(`D:\Gopath\src\kmp\go-wordsfilter-kmp\word.txt`)

    // Contains
    c1 := kmp.Contains("情不自禁")
    // c1: false
    c2 := kmp.Contains("作用")
    // c2: true

    // Remove
    c3 := kmp.Replace("什么作用域")
    // c3: 什么"*******"

    // Replace
    r1 := kmp.Remove("什么作用域")
    // r1: "什么"
}
```

## Apis
```go
待补充
