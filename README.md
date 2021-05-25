# google-translate


### go get
```shell
go get -u github.com/chnls/google-translate
```

# example

#### 1. pronounce is false
```go
s, p, err := google_translate.Translate("zh", "en", "hello, world. this is google translate", false)
if err != nil {
    fmt.Println(err)
}
fmt.Println("s: ", s, " p: ", p)
```
```go
s:  你好，世界。这是谷歌翻译  p:  
```

#### 2. pronounce is true
```go
s, p, err := google_translate.Translate("zh", "en", "hello, world. this is google translate", true)
if err != nil {
    fmt.Println(err)
}
fmt.Println("s: ", s, " p: ", p)
```
```go
s:  你好，世界。这是谷歌翻译  p:  Nǐ hǎo, shìjiè. Zhè shì gǔgē fānyì
```
