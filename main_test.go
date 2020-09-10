package main

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestParsePortalForm(t *testing.T) {
	data, err := ParsePortalForm(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title>上网认证</title>
<link rel="stylesheet" href="/portal/images/css.css" type="text/css" media="screen" />
<!-- <script src="/portal/js/hmac-sha1.js"></script>
<script src="/master/include/base64.js"></script> -->
<link rel="stylesheet" type="text/css" href="/portal/css/pop.css">
<link rel="stylesheet" type="text/css" href="/portal/css/button.css" />
<!--jquery-->
<script type="text/javascript" src="/master/include/jquery.js?rand=1599535977759"></script>
<script type="text/javascript">
		var noteicejson =[{"noticeForm1":"","noticeForm2":"","noticeForm3":"","noticeForm4":"","noticeForm5":"","authNoticeForm":"","platNoticeForm":"","authNoticeFormList":[],"commonNoticeList":[],"muchNoticeList":[],"muchNoticeList1":[],"muchNoticeList2":[]}];
		var userid="";
</script>
<script type="text/javascript" src="/portal/notice.js?rand=1599535977759"></script>
<!--通用的js-->
<script type="text/javascript" src="/portal/ch.js"></script>
<script type="text/javascript" src="/portal/portal_new.js?rand=1599535977759"></script>
<script type="text/javascript" src="/portal/js/banner2.js?rand=1599535977759"></script>
<script type="text/javascript" src="/portal/js/base.js"></script>
<script type="text/javascript" src="/master/include/bootstrap.min.js"></script>
<div class="modal hide" id="questionnaire">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h4 class="modal-title">调查问卷</h4>
      </div>
      <div class="modal-body" >

        <form id="form_questionnaire" action="/questionnaireAction.do" method="post">
          <input type="hidden" name="act" value="edit_action">
          <input type="hidden" name="answeruserid" id="answeruserid" value="">
          <input type="hidden" name="class_id"  value="0">
          <input type="hidden" name="group_id" value="0">
          
        </form>

      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-primary" id="quest_submit">问题还没有回答完</button>
      </div>
    </div>
  </div>
</div><script type="text/javascript" src="/portal/js/bootstrap.min.js"></script>
<link rel="stylesheet" href="/portal/css/bootstrap.min.css">
<script type="text/javascript">
	$(function(){		

	  

	})
</script><SCRIPT LANGUAGE="JavaScript">
<!--
	
	 $(function(){
	$("#testbutton").click(function() {
    $.ajax({
        type: "POST",
        url: "/master/TestJinJianAction.do",
        contentType: "application/json; charset=utf-8",
        
        dataType: "json",
        success: function (message) {
            
                alert(message);
            
        },
        error: function (message) {
        	alert("提交数据失败！");
            
        }
    });
});
});
//-->
</SCRIPT>
<!-- <script type="text/javascript">
            function genkey() {
                var userName="GET&http%3A%2F%crmtest.chinacloudapp.cn%2Fsrv%2Fget%2FGetMember&MemberID%3D31%26Mobile%3D%26oauth_consumer_key%3D48668584%26oauth_nonce%3D04fff526-feba-4159-a50a-151021cb72a4%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1475133078%26oauth_version%3D1.0"
                var password="5600ecdfde5b692d100a5de2";
                var hash = CryptoJS.HmacSHA1(userName, password);
                alert(hash);
                 var b = new Base64();
                var base64 = b.encode(hash);
                alert(base64);
            };
        </script> -->
</head>
<body  onload="onbodyload()">
<br />
<div class="container2">
<!-- head -->
<div class="main">
<div class="head"><div style='float:left;display:block;width:935px; height:63px;margin-top:5px;overflow:hidden;'><img  src="/portal/userimages/20190823/20190823144705_15.jpg"/></div>
<div class="nav_top"><!--baner---></div>
</div>
</div>
<!-- head -->
<div class="main">
<div class="banner">
<div class="l" id="tname_1_picpath2"><img src="/portal/userimages/20191021/20191021162222_160.jpg" width="935" height="241" border="0">

</div>
</div>
</div>
<div class="height20"></div>
<div class="main">
<div class="side_l">
<div class="box">
<div class="title title_bg1">登录区</div>
<div class="content">
<div class="login">
	<form action="/portalAuthAction.do" method="POST" onsubmit="return setMD5Passwd();"  name="portalForm">
<input type="hidden" name="wlanuserip" id="wlanuserip" value='10.10.10.10' >
  <input type="hidden" name="wlanacname" id="wlanacname" value='NFV-BASE' >
  <input type="hidden" name="chal_id" id="chal_id" value='' >
  <input type="hidden" name="chal_vector" id="chal_vector" value='' >
  <input type="hidden" name="auth_type" id="auth_type" value='PAP' >
  <input type="hidden" name="seq_id" id="seq_id" value='' >
  <input type="hidden" name="req_id" id="req_id" value='' >
  <input type="hidden" name="wlanacIp" id="wlanacIp" value='10.10.10.1' >
  <input type="hidden" name="ssid" id="ssid" value='' >
  <input type="hidden" name="vlan" id="vlan" value='0' >
  <input type="hidden" name="mac" id="mac" value='00:00:00:00:00:00' >
  <input type="hidden" name="message" id="message" value='' >
  <input type="hidden" name="bank_acct" id="bank_acct">
  <INPUT TYPE="hidden" name="isCookies">
  <input type="hidden" name="version" id="version" value='0'>
  <input type="hidden" name="authkey" id="authkey" value='itellin'>
  <input type="hidden" name="url" id="url" value='http://baidu.com/' >
  <input type="hidden" name="usertime" id="usertime" value='0' >
  <input type="hidden" name="listpasscode" id="listpasscode" value='0' ><!--是否获取验证码-->
  <input type="hidden" name="listgetpass" id="listgetpass" value='0' ><!--获取密码方式：1短信，2邮件-->
  <input type="hidden" name="getpasstype" id="getpasstype" value='0' ><!--是否开户：0开户，1仅查询-->
  <input type="hidden" name="randstr" id="randstr" value='8800' >
  <input type="hidden" name="domain" id="domain" value=''>
  <input type="hidden" name="isRadiusProxy" id="isRadiusProxy" value='false' >
  <input type="hidden" name="usertype" id="usertype" value="0">
  <input type="hidden" name="isHaveNotice" id="isHaveNotice" value='0' >
  <input type="hidden" name="times" id="times" value='12'>
  <input type="hidden" name="weizhi" id="weizhi" value='0' >
  <input type="hidden" name="smsid" id="smsid" value=''>
  <input type="hidden" name="freeuser" id="freeuser" value=''>
  <input type="hidden" name="freepasswd" id="freepasswd" value=''>
  <input type="hidden" name="listwxauth" id="listwxauth" value='0'>
  <input type="hidden" name="templatetype" id="templatetype" value='1'><!--终端类型-->
  <input type="hidden" name="tname" id="tname" value='1'><!--模板-->
  <input type="hidden" name="logintype" id="logintype" value="0"><!--临时存储的字段，用来判断当前是哪个登陆，0是账号，1是微信，默认是0-->
  <input type="hidden" name="act" id="act" value=''>
  <input type="hidden" name="is189" id="is189" value='false'>
  <input type="hidden" name="terminalType" id="terminalType" value='' >
  <input type="hidden" name="checkterminal" id="checkterminal" value="true"/>
  <input type="hidden" name="portalpageid" id="portalpageid" value="1">
  <input type="hidden" name="listfreeauth" id="listfreeauth" value="0">
  <input type="hidden" name="viewlogin" id="viewlogin" value="1">
  <input type="hidden" name="userid" id="userid" value=''>
  <input type="hidden" name="wisprpasswd" id="wisprpasswd" value=''>
  <input type="hidden" name="twocode" id="twocode" value="">
  <input type="hidden" name="authGroupId" id="authGroupId" value="">
  <INPUT type="hidden" name="alipayappid" id="alipayshopid" value="">
  <INPUT type="hidden" name="wlanstalocation" id="wlanstalocation" value="">
  <INPUT type="hidden" name="wlanstamac" id="wlanstamac" value="">
  <INPUT type="hidden" name="wlanstaos" id="wlanstaos" value="">
  <INPUT type="hidden" name="wlanstahardtype" id="wlanstahardtype" value="">
  <INPUT type="hidden" name="smsoperatorsflat" id="smsoperatorsflat" value="">
  <INPUT type="hidden" name="reason" id="reason" value="">
  <INPUT type="hidden" name="res" id="res" value="">
  <INPUT type="hidden" name="userurl" id="userurl" value="">
  <INPUT type="hidden" name="challenge" id="challenge" value="">
  <INPUT type="hidden" name="uamip" id="uamip" value="">
  <INPUT type="hidden" name="uamport" id="uamport" value="">
  <INPUT type="hidden" name="toqrcode" id="toqrcode" value="">
  <INPUT type="hidden" name="isIOSPortal" id="isIOSPortal" value="false"><table width="370" border="0" align="center" cellpadding="0" cellspacing="0" id="ids0" style="margin-top:10px;">
    <tr id="userLoginForm">
		

      <td valign="top" ><table width="90%" border="0" align="center" cellpadding="0" cellspacing="0"   style="display:">
		<tr>
           <td align="right" width="30%">登录名&nbsp;&nbsp;</td ><td align="left" width="70%"><input type="text" name="useridtemp" id="useridtemp" class="suru"  value=''  onkeyup="javascript:if(event.keyCode!=37 && event.keyCode!=39){ value=value.replace(/[^0-9a-zA-Z@_:.,-]/g,'');}" maxlength="32"/><!--<input type="hidden" id="userid" name="userid" value=''>--></td>
        </tr>
        <tr>
           <td align="right">密码&nbsp;&nbsp;</td><td align="left"><input type="password" name="passwd" id="passwd" class="suru"  value=''  maxlength="32"/><INPUT TYPE="checkbox" NAME="isRemind" id="isRemind"  align="absmiddle"    ><label for="isRemind">记密码</label></td>
        </tr>
		</table></td></tr>
    

	  <tr>
		<td width="100%" valign="middle" align="center">
          	
					<input  type="submit" value='登录' class="st_submit"/>
				<!-- <input  type="button" value='生成key' class="st_submit" onclick="genkey();"/> --></td>
				<!-- <td   align="center"><button id="testbutton" type="button" class="st_submit">调用</button></td> -->
    	  </tr>
		

	  <tr id="testuseForm" style="display:none;"><td width="100%" valign="middle" align="center" style="padding-top:10px;"></td></tr>
	  </table></td>
    </tr>
	</table>

  <table width="370" border="0" align="center" cellpadding="0" cellspacing="0" style="display:none" id="ids1" style="margin-top:10px;">
    <tr>
      <td valign="top" ><table width="90%" border="0" align="center" cellpadding="0" cellspacing="0">
		<Tr id="wx1">
			<td   align="left" height="40px" valign="middle">&nbsp;&nbsp;&nbsp;&nbsp;请先关注此公众号:<span class="fontdl"></span>,并且发送"wifi"信息,即可获取上网密码</td>
		</Tr>
        <Tr id="wx3">
			<td   align="center"><input type="button"  tabindex="3" class="btnlogin"  style="cursor: pointer;" title=''  value='已获得上网密码' onclick="gotonext();"/></td>
		</Tr>

		<Tr id="wx2">
			<td   align="center" height="40px" valign="middle"></td>
		</Tr>
			
		<Tr id="wx4"  style="display:none">
			<td   align="center" height="40px" valign="middle">输入上网密码&nbsp;<INPUT TYPE="text" NAME="wxuser" id="wxuser"  class="suru">&nbsp;<font color="red">*</font></td>
		</Tr>

		<Tr id="wx5" style="display:none">
			<td   align="center"><input type="submit"  tabindex="3" class="st_submit"  style="cursor: pointer;"  value='登录' /></td>
			
		</Tr>

		

      </table></td>
     
    </tr>
	</table>
</form>
</div>

<!--<div class="sq_side">
<table width="280" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td>&nbsp;</td>
  </tr>
</table></div>-->
<p id="logintext">从即日起，如需要小蝴蝶客户端请点击安装右侧免安装认证客户端（下载后解压便可使用），如不安装客户端也可使用浏览器输入任意网址会自动转到portal登录页面，输入原来的账号密码即可完成登录，请大家相互转告</p>

</div>
</div>
</div>
<div class="side_r">
<div class="box">
<div class="title title_bg2"><table width="99%" border="0" cellspacing="0" cellpadding="0">
  <tr>
    <td>最新资讯</td>
    <td align="right"> <!--<a href="/portal/noticeAll.jsp?wlanacname=NFV-BASE" target="_blank">更多>></a>--> </td>
  </tr>
</table></div>
<div class="content">
<ul class="new_list">
</ul>
</div>
</div>
<div class="height25" align="center"><table border="0" cellspacing="0" cellpadding="0"><tr><td width="100%" valign="middle" height="25">&nbsp;</td></tr></table></div>
<div><a href="https://net.lnut.edu.cn/ljy/LNUTWEB.zip" id="tname_1_picpathurl3" target="_blank">

	<img src="/portal/userimages/20191204/20191204111528_128.jpg" width="391" height="59" id="tname_1_picpath3"></a></div>
</div>
</div>
</div>
<div style="height:30px;line-height:30px;margin-right: auto;margin-left: auto;text-align:center;"></div>

<div id="bg_dialog"></div>

<div id="zadbox">
<div id="time_tip_dialog"></div>
  <div id="adconten">
	  <div class="box2">
	  	<div id="container2">
	  		<ul class="box_ul2">
	  			<!--<li>
	  				<img name="cal2" id="img2_1" src="/portal/images/bg_middle.png" width="320px" height="240px"/>  
	  			</li>
	  			<li>
	  				<img name="cal2" id="img2_2" src="/portal/images/bg_middle.png" width="320px" height="240px"/>  
	  			</li>
	  			<li>
	  				<img name="cal2" id="img2_2" src="/portal/images/bg_middle.png" width="320px" height="240px"/>  
	  			</li>
	  			<li>
	  				<img name="cal2" id="img2_3" src="/portal/images/bg_middle.png" width="320px" height="240px"/>  
	  			</li>-->
				

	  		</ul>
	  	</div>
	  	<div id="imgIndex2"></div>
	  </div>
  </div>
</div>
</body>
</html>
`)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	expected := url.Values{}
	expected.Set("wlanuserip", "10.10.10.10")
	expected.Set("wlanacname", "NFV-BASE")
	expected.Set("chal_id", "")
	expected.Set("chal_vector", "")
	expected.Set("auth_type", "PAP")
	expected.Set("seq_id", "")
	expected.Set("req_id", "")
	expected.Set("wlanacIp", "10.10.10.1")
	expected.Set("ssid", "")
	expected.Set("vlan", "0")
	expected.Set("mac", "00:00:00:00:00:00")
	expected.Set("message", "")
	expected.Set("bank_acct", "")
	expected.Set("isCookies", "")
	expected.Set("version", "0")
	expected.Set("authkey", "itellin")
	expected.Set("url", "http://baidu.com/")
	expected.Set("usertime", "0")
	expected.Set("listpasscode", "0")
	expected.Set("listgetpass", "0")
	expected.Set("getpasstype", "0")
	expected.Set("randstr", "8800")
	expected.Set("domain", "")
	expected.Set("isRadiusProxy", "false")
	expected.Set("usertype", "0")
	expected.Set("isHaveNotice", "0")
	expected.Set("times", "12")
	expected.Set("weizhi", "0")
	expected.Set("smsid", "")
	expected.Set("freeuser", "")
	expected.Set("freepasswd", "")
	expected.Set("listwxauth", "0")
	expected.Set("templatetype", "1")
	expected.Set("tname", "1")
	expected.Set("logintype", "0")
	expected.Set("act", "")
	expected.Set("is189", "false")
	expected.Set("terminalType", "")
	expected.Set("checkterminal", "true")
	expected.Set("portalpageid", "1")
	expected.Set("listfreeauth", "0")
	expected.Set("viewlogin", "1")
	expected.Set("userid", "")
	expected.Set("wisprpasswd", "")
	expected.Set("twocode", "")
	expected.Set("authGroupId", "")
	expected.Set("alipayappid", "")
	expected.Set("wlanstalocation", "")
	expected.Set("wlanstamac", "")
	expected.Set("wlanstaos", "")
	expected.Set("wlanstahardtype", "")
	expected.Set("smsoperatorsflat", "")
	expected.Set("reason", "")
	expected.Set("res", "")
	expected.Set("userurl", "")
	expected.Set("challenge", "")
	expected.Set("uamip", "")
	expected.Set("uamport", "")
	expected.Set("toqrcode", "")
	expected.Set("isIOSPortal", "false")
	expected.Set("useridtemp", "")
	expected.Set("passwd", "")
	expected.Set("isRemind", "")
	expected.Set("wxuser", "")
	assert.Equal(t, expected, data)
}
