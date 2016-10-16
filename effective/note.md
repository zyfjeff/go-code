* gofmt 格式化代码
* godoc 生成文档，需要自己编排好格式 -goroot 指定搜索的package的路径
* 声明错误，类似于C++的枚举错误类型。

```
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ....
    )
```

* 包名要小写，单个单词,不需要下划线，也不需要驼峰标识。

* go不支持getter和setter，通常的做法是有个小写字母开头的成员变量，然后设置一个大小字母开头与其对应的方法
来作为getter，然后再设置一个名为"Set`成员变量名`" 作为setter。

```
owner := obj.Owner()
if owner != user {
        obj.SetOwner(user)
}
```

* 方法的接口是方法名称加er后缀

* 驼峰标识别 MixedCaps或mixedCaps

* 避免变量重声明

```
f,err := os.Open(name)
d,err := f.Stat()
```
* 
