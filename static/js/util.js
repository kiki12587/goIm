define([],function () {
return {

        //设置cookie
        //这是有设定过期时间的使用示例：
        //s20是代表20秒
        //h是指小时，如12小时则是：h12
        //d是天数，30天则：d30
        // 样例:setCookie("name","hayden","s20");
        setCookie:function (name,value,time)
        {
            var strsec = this.getsec(time);
            var exp = new Date();
            exp.setTime(exp.getTime() + strsec*1);
            document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
        },

        getsec:function (str) {
            var str1 = str.substring(1, str.length) * 1;
            var str2 = str.substring(0, 1);
            if (str2 == "s") {
                return str1 * 1000;
            } else if (str2 == "h") {
                return str1 * 60 * 60 * 1000;
            } else if (str2 == "d") {
                return str1 * 24 * 60 * 60 * 1000;
            }

        },
        //读取cookie
        getCookie:function (name) {
            var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
            if (arr = document.cookie.match(reg))
                return unescape(arr[2]);
            else
                return null;
        },


        //删除cookie
        delCookie:function (name) {
            var exp = new Date();
            exp.setTime(exp.getTime() - 1);
            var cval = this.getCookie(name);
            if (cval != null)
                document.cookie = name + "=" + cval + ";expires=" + exp.toGMTString();
        },


       //校验是不是邮箱
       checkMail:function(email){
           let myreg = /^(\w-*\.*)+@(\w-?)+(\.\w{2,})+$/;

           if (!myreg.test(email)) {
               return false;
           } else {
               return true;
           }
       },
    // 判断字符串是否为数字和字母组合
    checkS:function (str){
        var reg=new RegExp(/[A-Za-z].*[0-9]|[0-9].*[A-Za-z]/);
        if(reg.test(str)){
            return true;
        }else{
            return false;
        }
    },


       }

});