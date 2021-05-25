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

## other api
```go
method GET
https://translate.googleapis.com/translate_a/single?client=gtx&sl=zh&tl=en&dt=t&q=xxxx
```
注： 此方式最大文本长度1814(中文，其他未测)
在python中测试

```python
from urllib.parse import quote
import requests
s = "国" * 1814  # the largest length
res = requests.get(
    "https://translate.googleapis.com/translate_a/single?client=gtx&sl=zh&tl=en&dt=t&q={}".format(quote(s)),
    # verify=False
    )
print(res.status_code)  # 200
```