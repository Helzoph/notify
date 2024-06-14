# Notify  

目前仅支持以下通知方式:  
- [PushPlus](https://www.pushplus.plus)  

## 使用方法  

```
go get github.com/helzoph/notify
```

```go
config, err := LoadConfig("config.yaml")
if err != nil {
    panic(err)
}

bot := NewPushPlusBot(config)
if bot.Token == "" {
    panic(err)
}

_, err = bot.Send(&pushplus.Message{
    Content: "Hello, World!",
})
if err != nil {
    panic(err)
}
```

`config.yaml` 配置文件格式如下：  

```yaml
pushplus:
    token:
```