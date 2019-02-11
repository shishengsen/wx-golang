## Golang 实现的微信公众号开发工具

![License](https://img.shields.io/github/license/Weixin-Golang/wx-golang.svg)
![stars](https://img.shields.io/github/stars/Weixin-Golang/wx-golang.svg)
![forks](https://img.shields.io/github/forks/Weixin-Golang/wx-golang.svg)
![building](https://img.shields.io/badge/status-building-green.svg?longCache=true&style=plastic)

该开发包的实现，主要参考了目前流行的java版的微信开发包——[Wechat-Group/weixin-java-tools](https://github.com/Wechat-Group/weixin-java-tools)

#### DONE

- [x] 微信配置实现以及相关实体的构建
- [x] 微信accesstoken的获取、存储、刷新
- [x] 微信自定义菜单、个性化菜单相关操作的实现
- [x] 微信客服相关接口
- [x] 微信消息通知的路由分发操作实现
- [x] `http`接口实现
- [x] 微信素材管理
- [x] 用户标签管理
- [x] 带参数二维码
- [x] 长链接转换
- [x] 数据统计
- [x] 用户管理
- [x] 账户管理

#### TODO

- [ ] 图文消息留言管理
- [ ] 新版客服功能

#### example

1. 如何使用消息路由

    ```go
    package service
    
    import (
    	"wx-golang/weixin-mp/enpity"
    )
    
    // 继承该接口实现微信消息的处理
    type MsgHandler interface {
    	Handler(enpity.WxMessage)
    }    
    ```

    ```go
    package example
    
    import (
    	"fmt"
    	"testing"
    	"wx-golang/weixin-common/wxconsts"
    	"wx-golang/weixin-mp/enpity"
    	"wx-golang/weixin-mp/service"
    )

    type GuanZhu struct{}
    
    func (g *GuanZhu) Handler(enpity.WxMessage) {
    	fmt.Print("关注事件")
    }
    
    type SaoMa struct{}
    
    func (s *SaoMa) Handler(enpity.WxMessage) {
    	fmt.Print("扫码事件")
    }
    
    func TestRouter(test *testing.T) {
    	w := &service.WeChat{}
    	router := w.RouterInit()
    	g := GuanZhu{}
    	router.Start().
    		MsgType(wxconsts.MSG_TYPE_EVENT).
    		Event(wxconsts.EVENT_TYPE_SUBSCRIBE).
    		Handler(&g).
    		End().
    		Start().
    		MsgType(wxconsts.MSG_TYPE_EVENT).
    		Event(wxconsts.EVENT_TYPE_SCAN).
    		Handler(&SaoMa{}).End()
    	msg1 := enpity.WxMessage{MsgType:"event", Event:"subscribe"}
    	msg2 := enpity.WxMessage{MsgType:"event", Event:"SCAN"}
    	w.Route(msg1)
    	w.Route(msg2)
    }
    ```
    
    测试结果
    
    ```bash
    === RUN   TestRouter
    &service.MsgRouter{rules:[]*service.MsgRule{(*service.MsgRule)(0xc0001c7c00), (*service.MsgRule)(0xc0001c7c70)}}
    关注事件&service.MsgRouter{rules:[]*service.MsgRule{(*service.MsgRule)(0xc0001c7c00), (*service.MsgRule)(0xc0001c7c70)}}
    扫码事件--- PASS: TestRouter (0.00s)
    PASS
    
    Process finished with exit code 0
    ```