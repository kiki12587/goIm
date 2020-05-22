require.config({
baseUrl:'../static/',
paths: {
        jquery: "https://libs.baidu.com/jquery/2.0.0/jquery.min",
        layui: "layui/layui",
        bootstrap: "bootstrap/bootstrap.min" ,
        gVerify :"js/gVerify",
        amazeui :"js/amazeui.min",
        wechat:"js/wechat",
         zUI :"js/zUI",
        jquerymin:"js/jquery.min",
         cosjs:"js/cos-js-sdk-v5.min",
         generateTestUserSig:"js/GenerateTestUserSig",
         tim:"js/tim-js",
         index:"js/index",
    },
    shim:{
        wechat:{
            deps:['jquerymin']
        },
        zUI:{
            deps:['jquerymin']
        }

    }

});