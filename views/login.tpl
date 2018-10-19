<!DOCTYPE html>

<html>
<head>
  <title>ipp 智能巡检平台</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="shortcut icon" href="static/img/1234.ico" type="image/x-icon" />
  <link rel="stylesheet" href="static/css/base.css">
</head>


<body>
<div class="background">
<div class="loginBox">
            <h1><a href="javascript:;" class="logo"></a><span>| 智能巡检平台领导者</span></h1>
            <div class="lgWrap" id="Login">
                <h5 style="font-weight: 400;font-size: 1.4em;">智能巡检系统平台用户登录</h5>
                <div class="lgBox">
                    <span id="warn_message"></span>
                    <ul style="margin-top:5%;margin-bottom:5%;">
					 <li class="userName">
					<lable style="display:inline-block;width:85px;">数据库地址：</lable>
                            <input type="text" class="text" id="hostip" name="hostip">
                        </li>
							 <li class="userName">
							<lable style="display:inline-block;width:85px;">数据库端口：</lable>
                            <input type="text" class="text" id="hostport" name="hostport">
                        </li>
						 <li class="userName"><lable style="display:inline-block;width:85px;">服务名称：</lable>
                            <input type="text" class="text" id="servicename" name="servicename">
                        </li>
                        <li class="userName"><lable style="display:inline-block;width:85px;">账号：</lable>
                            <input type="text" class="text" id="username" name="username">
                        </li>
                        <li class="passWord"><lable style="display:inline-block;width:85px;">密码：</lable>
                            <input class="text" type="password" id="password" name="password">
                        </li>
                        <li style="margin-left: 170px;font-size: 10px">
                            <input id="remember" class="sm" type="checkbox">
                            记住密码
                        </li>
                        <li class="submit">
                            <input id="loginbtn" class="denglu" type="button" value="登录">
                        </li>
                    </ul>
                    <span class="tishi">提示：为了获得更好的体验建议使用火狐、谷歌或者IE浏览器进行管理</span>
                </div>
            </div>
            <div class="lgWrap" id="active" style="display: none">
                <h5 style="font-weight: 400;font-size: 1.4em;margin-left: 28px">激活客户体验管理系统</h5>
                <div class="lgBox">
                    <span></span>
                    <ul style="margin-top:5%;margin-bottom:7%;">
                        <li class="registeredCompany" style="margin-left: 104PX;">请求码：
                            <input type="text" class="text" readonly="true" id="textAreas">
                        </li>
                        <li class="activationCode" style="margin-left: 104PX;">激活码：
                            <input class="text" type="text" id="sign">
                        </li>
                        <span id="alertCode" style="display: none" class="activeTishi">激活码错误请重新输入!</span>
                        <li class="submit">
                            <input class="denglu" type="button" id="activeButton" value="激活" onclick="active()">
                        </li>
                    </ul>
                    <span class="tishi">提示：为了获得更好的体验建议使用火狐、谷歌或者IE浏览器进行管理</span>
                </div>
            </div>
        </div>
</div>
 
  
  

  <script src="/static/js/jquery-3.2.1.min.js"></script>
<script src="/static/js/login.js"></script>
</body>
</html>
