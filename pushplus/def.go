package pushplus

// 发送渠道
const (
	Channel_Wechat  = "wechat"  // 微信公众号
	Channel_Webhook = "webhook" // 第三方webhook；HiFlow连接器、企业微信、钉钉、飞书、server酱、IFTTT
	Channel_Cp      = "cp"      // 企业微信应用
	Channel_Mail    = "mail"    // 邮箱
	Channel_Sms     = "sms"     // 短信；成功发送1条短信需要10积分（0.1元）
)

// 模板
const (
	Template_Html         = "html"         // 默认模板，支持html文本
	Template_Txt          = "txt"          // 纯文本展示，不转义html
	Template_JSON         = "json"         // 内容基于json格式展示
	Template_Markdown     = "markdown"     // 内容基于markdown格式展示
	Template_CloudMonitor = "cloudMonitor" // 阿里云监控报警定制模板
	Template_Jenkins      = "jenkins"      // jenkins插件定制模板
	Template_Route        = "route"        // 路由器插件定制模板
	Template_Pay          = "pay"          // 支付成功通知模板
)

type Message struct {
	// Token     string `json:"token"`     // (必填) 用户令牌，可直接加到请求地址后
	Title     string `json:"title"`     // 消息标题
	Content   string `json:"content"`   // (必填) 具体消息内容，根据不同template支持不同格式
	Topic     string `json:"topic"`     // 群组编码，不填仅发送给自己；channel为webhook时无效
	Template  string `json:"template"`  // 发送模板。默认值: html
	Channel   string `json:"channel"`   // 发送渠道。默认值: wechat
	Webhook   string `json:"webhook"`   // webhook编码
	Callback  string `json:"callback"`  // 发送结果回调地址
	Timestamp string `json:"timestamp"` // 毫秒时间戳。格式如：1632993318000。服务器时间戳大于此时间戳，则消息不会发送
	To        string `json:"to"`        // 好友令牌，微信公众号渠道填写好友令牌，企业微信渠道填写企业微信用户id
}

// 响应内容
type Response struct {
	Code int    `json:"code"` // 响应码
	Msg  string `json:"msg"`  // 响应说明
	Data string `json:"data"` // 响应数据
}

// TODO: 回调内容
type Callback struct {
}
