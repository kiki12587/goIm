require(['tim','cosjs','generateTestUserSig'],function (TIM,COS) {

// 创建 SDK 实例，TIM.create() 方法对于同一个 SDKAppID 只会返回同一份实例
    let options = {
        SDKAppID: 1400362640 // 接入时需要将0替换为您的即时通信应用的 SDKAppID
    }

    let tim = TIM.create(options) // SDK 实例通常用 tim 表示
// 设置 SDK 日志输出级别，详细分级请参考 setLogLevel 接口的说明
    tim.setLogLevel(0) // 普通级别，日志量较多，接入时建议使用

// 注册 COS SDK 插件
    tim.registerPlugin({ 'cos-js-sdk': COS });

// 监听事件，如：
    tim.on(TIM.EVENT.SDK_READY, function (event) {

        console.log(event, 2222222)
        // 发送文本消息，Web 端与小程序端相同
        // 1. 创建消息实例，接口返回的实例可以上屏
        // let message = tim.createTextMessage({
        //     to: '1',
        //     conversationType: TIM.TYPES.CONV_C2C,
        //     // 消息优先级，用于群聊（v2.4.2起支持）。如果某个群的消息超过了频率限制，后台会优先下发高优先级的消息，详细请参考：https://cloud.tencent.com/document/product/269/3663#.E6.B6.88.E6.81.AF.E4.BC.98.E5.85.88.E7.BA.A7.E4.B8.8E.E9.A2.91.E7.8E.87.E6.8E.A7.E5.88.B6)
        //     // 支持的枚举值：TIM.TYPES.MSG_PRIORITY_HIGH, TIM.TYPES.MSG_PRIORITY_NORMAL（默认）, TIM.TYPES.MSG_PRIORITY_LOW, TIM.TYPES.MSG_PRIORITY_LOWEST
        //     // priority: TIM.TYPES.MSG_PRIORITY_NORMAL,
        //     payload: {
        //         text: 'Hello world!'
        //     }
        // });
        // // 2. 发送消息
        // let promise = tim.sendMessage(message);
        // promise.then(function (imResponse) {
        //     // 发送成功
        //     console.log(imResponse);
        // }).catch(function (imError) {
        //     // 发送失败
        //     console.warn('sendMessage error:', imError);
        // });



        // 收到离线消息和会话列表同步完毕通知，接入侧可以调用 sendMessage 等需要鉴权的接口
        // event.name - TIM.EVENT.SDK_READY
    });


    tim.on(TIM.EVENT.MESSAGE_RECEIVED, function (event) {

        console.log(event)
        event.data.forEach((item, index, arr) => {

            //收到图片消息
            if (item.type == "TIMImageElem") {
                let html = `<img src=${item.payload.imageInfoArray[index].imageUrl} alt=""/>`
                $("#ul").append(html);
            }

            //收到文本消息
            if (item.type == "TIMTextElem") {
                let html = `<li>文本消息:${item.payload.text}</li>`
                $("#ul").append(html);
            }

            //收到自定义消息
            if (item.type == "TIMCustomElem") {
                let html = `<li>自定义消息:${item.payload.description}</li>`
                $("#ul").append(html);
            }

            //收到文件消息
            if (item.type == "TIMFileElem") {

            }



            console.log(item)

        });

        // for(var item of event.data) {
        //     let html = `<li>${item.}</li>`
        //     $("#ul").append(html);
        // }


        // 收到推送的单聊、群聊、群提示、群系统通知的新消息，可通过遍历 event.data 获取消息列表数据并渲染到页面
        // event.name - TIM.EVENT.MESSAGE_RECEIVED
        // event.data - 存储 Message 对象的数组 - [Message]
    });

    tim.on(TIM.EVENT.MESSAGE_REVOKED, function (event) {
        // 收到消息被撤回的通知。使用前需要将SDK版本升级至v2.4.0或以上。
        // event.name - TIM.EVENT.MESSAGE_REVOKED
        // event.data - 存储 Message 对象的数组 - [Message] - 每个 Message 对象的 isRevoked 属性值为 true
    });

    tim.on(TIM.EVENT.CONVERSATION_LIST_UPDATED, function (event) {
        // 收到会话列表更新通知，可通过遍历 event.data 获取会话列表数据并渲染到页面
        // event.name - TIM.EVENT.CONVERSATION_LIST_UPDATED
        // event.data - 存储 Conversation 对象的数组 - [Conversation]
    });

    tim.on(TIM.EVENT.GROUP_LIST_UPDATED, function (event) {
        // 收到群组列表更新通知，可通过遍历 event.data 获取群组列表数据并渲染到页面
        // event.name - TIM.EVENT.GROUP_LIST_UPDATED
        // event.data - 存储 Group 对象的数组 - [Group]
    });

    tim.on(TIM.EVENT.PROFILE_UPDATED, function (event) {
        // 收到自己或好友的资料变更通知
        // event.name - TIM.EVENT.PROFILE_UPDATED
        // event.data - 存储 Profile 对象的数组 - [Profile]
    });

    tim.on(TIM.EVENT.BLACKLIST_UPDATED, function (event) {
        // 收到黑名单列表更新通知
        // event.name - TIM.EVENT.BLACKLIST_UPDATED
        // event.data - 存储 userID 的数组 - [userID]
    });

    tim.on(TIM.EVENT.ERROR, function (event) {
        // 收到 SDK 发生错误通知，可以获取错误码和错误信息
        // event.name - TIM.EVENT.ERROR
        // event.data.code - 错误码
        // event.data.message - 错误信息
    });

    tim.on(TIM.EVENT.SDK_NOT_READY, function (event) {
        // 收到 SDK 进入 not ready 状态通知，此时 SDK 无法正常工作
        // event.name - TIM.EVENT.SDK_NOT_READY
    });

    tim.on(TIM.EVENT.KICKED_OUT, function (event) {
        // 收到被踢下线通知
        // event.name - TIM.EVENT.KICKED_OUT
        // event.data.type - 被踢下线的原因，例如 :
        //   - TIM.TYPES.KICKED_OUT_MULT_ACCOUNT 多实例登录被踢
        //   - TIM.TYPES.KICKED_OUT_MULT_DEVICE 多终端登录被踢
        //   - TIM.TYPES.KICKED_OUT_USERSIG_EXPIRED 签名过期被踢（v2.4.0起支持）。
    });

    tim.on(TIM.EVENT.NET_STATE_CHANGE, function (event) {
        // 网络状态发生改变（v2.5.0 起支持）。
        // event.name - TIM.EVENT.NET_STATE_CHANGE
        // event.data.state 当前网络状态，枚举值及说明如下：
        //   - TIM.TYPES.NET_STATE_CONNECTED - 已接入网络
        //   - TIM.TYPES.NET_STATE_CONNECTING - 连接中。很可能遇到网络抖动，SDK 在重试。接入侧可根据此状态提示“当前网络不稳定”或“连接中”
        //   - TIM.TYPES.NET_STATE_DISCONNECTED - 未接入网络。接入侧可根据此状态提示“当前网络不可用”。SDK 仍会继续重试，若用户网络恢复，SDK 会自动同步消息
    });


    let login = tim.login({ userID: '1', userSig: 'eJyrVgrxCdYrSy1SslIy0jNQ0gHzM1NS80oy0zLBwoZQweKU7MSCgswUJStDEwMDYzMjMxMDiExqRUFmUSpQ3NTU1MjAACpakpkLFrOwMDczNjIxhpqSmQ40M83AJMok29MxqjLDryhX2zciOCxZ27A02csgJK-cJDi0vNwk1Sjcotip0tFWqRYA-gwvPg__' });
    login.then(function (imResponse) {

        console.log(imResponse.data,555555555); // 登录成功
    }).catch(function (imError) {
        console.warn('login error:', imError); // 登录失败的相关信息
    });
})