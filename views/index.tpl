
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <title>智能巡检平台</title>
    <link rel="shortcut icon" href="static/img/1234.ico">
    <link rel="stylesheet" href="static/css/base.css">
	<!--<link rel="stylesheet" href="static/css/iconfont.css">-->
    <link rel="stylesheet" href="static/css/bootstrap.css">
    <link rel="stylesheet" href="static/css/sb-admin.css">
    <link rel="stylesheet" href="static/css/style.css">
</head>
<body class="page-top" id="page-top">
<!-- Navigation-->
<nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top" id="mainNav">
    <a href="/"><P class="m-navimg"></P></a>
    <a class="navbar-brand" href="#"></a>
    <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarResponsive">
        <ul class="navbar-nav navbar-sidenav one" id="exampleAccordion">
            <li class="nav-item nav-munu" id="menu">
                <a class="nav-link" id="sidenavToggler" >
                    <i class="fa-fw fa-mun"></i>
                    <span class="nav-link-text" >最新概览</span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right" id="overview" title="概览" style="position:relative;">
                <a class="nav-link" href="overview.html" target="mainCon" >
                    <i class="fa-fw fa-dashboards"></i>
                    <span class="nav-link-text nav_text select">常规巡检</span>
                    <span class="nav-red"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right" id="performance" title="性能&故障" style="position:relative;" >
                <a class="nav-link" href="performance.html" target="mainCon" >
                    <i class="fa-fw fa-area-chart"></i>
                    <span class="nav-link-text nav_text">虚拟化巡检</span>
                    <span class="nav-red" style="display: none;"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right" id="operate" title="运营决策" style="position:relative;">
                <a class="nav-link" href="operate.html" style="position:relative;" target="mainCon"  >
                    <i class="fa-fw fa-table"></i>
                    <span class="nav-link-text nav_text">导出报告</span>
                    <span class="nav-red" style="display: none;"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right"  id= "dataBase" title="数据库" style="position:relative;">
                <a class="nav-link" href="database.html" style="position:relative;" target="mainCon"  >
                    <i class="fa-fw fa-data"></i>
                    <span class="nav-link-text nav_text">虚拟化报告</span>
                    <span class="nav-red" style="display: none;"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right"  id= "report" title="报表" style="position:relative;">
                <a class="nav-link" href="report.html" style="position:relative;" target="mainCon"  >
                    <i class="fa-fw fa-report"></i>
                    <span class="nav-link-text nav_text">巡检配置</span>
                    <span class="nav-red" style="display: none;"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right" id="config" title="设置" style="position:relative;">
                <a class="nav-link" href="set.html" style="position:relative;" target="mainCon"  >
                    <i class="fa-fw fa-wrench"></i>
                    <span class="nav-link-text nav_text">规则配置</span>
                    <span class="nav-red" style="display: none;"></span>
                </a>
            </li>
        </ul>
      
        <ul class="navbar-nav ml-auto">
            <li class="nav-item returns" style="position: relative;left: 16px">
                <a href="/" class="nav-link" style="color: #fff;font-size: 14px;">
                    <div class="dropdown return " style="display: inline-block;">
                    <i class="fa fa-fw fa-sign-index" style="margin:2px 0px 0px 4px;"></i>
                    <div class="dropdown" style="display: inline-block;">
                <ul class="dropdown-menu dropdown-return" aria-labelledby="dLabel">
                    <i class="adminicno" style="left: 10px;top: -7px;"></i>
                    <li data-toggle="modal" data-target="#myModal" class="navadmin">
                        首页
                    </li>
                </ul>
                    </div>
                    </div>
                </a>
            </li>
            <!--<li class="nav-item nav-li"></li>-->
            <li class="nav-item abouts" style="position: relative;left: 10px">
                <a class="nav-link" style="color: #fff;font-size: 14px;" id="pertaining">
                    <div class="dropdown about" style="display: inline-block;">
                    <i class="fa fa-fw fa-sign-out" style="margin:5px 0px 2px 6px"></i>
                    <div class="dropdown" style="display: inline-block;">
                        <ul class="dropdown-menu dropdown-about" aria-labelledby="dLabel">
                            <i class="adminicno" style="left: 10px;top: -7px;"></i>
                            <li data-toggle="modal" data-target="#myModal11" class="navadmin">
                               关于
                            </li>
                        </ul>
                    </div>
                    </div>
                    </a>
            </li>
           <!-- <li class="nav-item nav-li"></li>-->
            <li class="nav-item itemli">
                <a class="nav-link" data-toggle="modal" data-target="#exampleModal" style="color: #fff;font-size: 14px;cursor: pointer;margin-right:20px;">
                    <div class="dropdown navAdminColor" style="display: inline-block;">
                    <i class="fa fa-fw fa-admin"></i>
                        <button id="dLabel" type="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" class="adminButton" style="cursor: pointer;margin-bottom: 5px">
                            <i id="username" style="font-style:normal;font-size: 14px;">admin</i>
                            <span class="caret"></span>
                            <i class="fa fa-fw fa-admin-down" style="cursor: pointer;margin-right: 0px;"></i>
                        </button>
                        <ul class="dropdown-menu dropdown-admin" aria-labelledby="dLabel">
                            <i class="adminicno"></i>
                            <li data-toggle="modal" data-target="#myModal" id="userManagement" class="navadmin">
                                <i class="adminName"></i>
                                用户管理
                            </li>
                            <li id="about"  class="navadmin">
                                <i class="adminAbort"></i>
                                退出账户
                            </li>
                        </ul>

                    </div>

                </a>
            </li>
        </ul>
    </div>
</nav>

<!--关于模态框-->
<div  class="modal" data-backdrop="static" id="my-modal-modify">
    <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content modal-guanyu">
            <div class="modal-header modal-nav-header" style="  padding: 17px;">

                <p class="modal-title modal-titles" id="gridSystemModalLabel" style="font-size: 14px;"> 关于海天CEM管理系统v5.0.1</p>

                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            </div>
            <div class="modal-body">
                <ul>
                    <li>
                        <span >咨询：</span>
                        010-58701010 13501365250
                    </li>
                    <li>
                        <span>邮箱：</span>
                        raobing@hthorizon.com
                    </li>
                    <li>
                        <span>网址：</span>
                        http://www.hthorizon.com
                    </li>
                    <li>
                        <span>地址：</span>
                        北京朝阳区东大桥路8号尚都国际中心2806室
                    </li>
                    <li>
                        <p class="erweima"> </p>
                        <span>扫码关注微信公众号</span>
                    </li>
                </ul>
            </div>
            <div class="modal-footer modify-footer modifyGuanyu">
                <button type="button" class="btn btn-default" data-dismiss="modal" style="position: absolute;bottom: 10px;">确定</button>
            </div>
        </div>
    </div>
</div>
<!--退出-->
<div  class="modal" data-backdrop="static" id="my-modal-about">
    <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content modal-about">
            <div class="modal-body">
                <ul>
                    <li>
                        <p class="dropOut"> </p>
                        <span class="dropspan">确定退出账户吗？</span>
                    </li>
                </ul>
            </div>
            <div class="modal-footer modify-footer">
                <button type="button" class="btn btn-primary"  data-dismiss="modal" style="position: relative;left: -28px;">取消</button>
                <button type="button" class="btn btn-default" style="position: absolute;left: 31px;" onclick="exitAccount()">确定</button>
            </div>
        </div>
    </div>
</div>
<!--用户管理-->
<div class="modal right fade" id="myModal" data-backdrop="static" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog ">
        <div class="modal-content modal-content-s">
            <div class="modal-header">
                <p class="modal-title" id="myModalLabel">用户管理</p>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true" style="line-height: 50px;padding: 8px;">&times;</span></button>
            </div>
            <div class="modal-body tableUser">
               <button class="addName" data-toggle="modal" id="addUser">添加用户</button>
                <table class="table  table-condensed tableUser">
                    <thead>
                    <tr style="background-color: rgba(0, 0, 0, 0.05);">
                        <th>用户</th>
                        <th>角色</th>
                        <th>邮箱</th>
                        <th>电话</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody id="userList"></tbody>
                </table>
            </div>
        </div>
    </div>
</div>
<!--添加用户-->
<div class="modal right fade addModal" id="addUser-s" data-backdrop="static" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog ">
        <div class="modal-content addmodalUser">
            <div class="modal-header">
                <p class="modal-title" style="font-size: 14px">添加用户</p>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true" style="line-height: 50px;padding: 8px;">&times;</span></button>
            </div>
            <div class="modal-body userForm">
                    <ul>
                        <li>
                            <span>用户名：</span>
                            <input type="text" style="margin-left: 13px" id="addUserName">
                        </li>
                        <li>
                            <span>电话：</span>
                            <input type="text" style="margin-left: 27px" id="addPhone">
                        </li>
                        <li class="positionR">
                            <span>邮箱：</span>
                            <input type="text" style="margin-left: 27px" id="addEmail">
                        </li>
                        <li class="positionR" >
                            <span>密码：</span>
                            <input type="password" style="margin-left: 27px" id="addPassWord">
                        </li>
                        <li class="positionR">
                            <span>确认密码：</span>
                            <input type="password" id="replayPassWord">
                        </li>
                        <li class="positionR">
                            <span>权限：</span>
                            <select name="" id="add-select" multiple="multiple"></select>
                        </li>
                    </ul>
            </div>
            <div class="modal-footer modify-footer">
              <p style="margin: auto">
                  <button type="button" class="btn btn-default" onclick="addUser()">确定</button>
                  <button type="button" class="btn btn-primary "  data-dismiss="modal" onclick="removeAdd()">取消</button>
              </p>
            </div>
        </div>
    </div>
</div>
<!--修改用户-->
<div class="modal right fade addModal" id="userMod" data-backdrop="static" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog ">
        <div class="modal-content addmodalUser">
            <div class="modal-header">
                <p class="modal-title" style="font-size: 14px">用户修改</p>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true" style="line-height: 50px;padding: 8px;">&times;</span></button>
            </div>
            <div class="modal-body userForm">
                <ul>
                    <li>
                        <span>用户名：</span>
                        <input type="hidden" id="changeId">
                        <input type="text" readonly="readonly" style="margin-left: 13px;background:#CCCCCC" id="changeUserName">
                    </li>
                    <li>
                        <span>电话：</span>
                        <input type="text" style="margin-left: 27px" id="changePhone">

                    </li>
                    <li class="positionR">
                        <span>邮箱：</span>
                        <input type="text" style="margin-left: 27px" id="changeEmail">
                    </li>
                    <li class="positionR" id="authority">
                        <span>权限：</span>
                        <select id="change-select" multiple="multiple"></select>
                    </li>
                </ul>
            </div>
            <div class="modal-footer modify-footer">
                <p style="margin: auto">
                    <button type="button" class="btn btn-default" onclick="changeUser()">修改</button>
                    <button type="button" class="btn btn-primary "  data-dismiss="modal" onclick="removeChange()">取消</button>
                </p>
            </div>
        </div>
    </div>
</div>
<input type="hidden" id="hiddenId">
<!--删除用户-->
<div class="modal right fade addModal" id="userDelect" data-backdrop="static" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog ">
        <div class="modal-content modal-shanchu">
            <div class="modal-body modal-bodys">
                <ul>
                    <li>
                        <p class="dropOut"> </p>
                        <span class="dropspan">确定删除用户吗？</span>
                    </li>
                </ul>
            </div>
            <div class="modal-footer modify-footer">
                <button type="button" class="btn btn-primary "  data-dismiss="modal" style="position: relative;left: -28px;">取消</button>
                <button type="button" class="btn btn-default" style="position: absolute;left: 31px;" data-dismiss="modal" onclick="deleteUser()" id="deleteUser">确定</button>
            </div>
        </div>
        </div>
    </div>
<!--修改密码-->
<div class="modal right fade addModal" id="userPassword" data-backdrop="static" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog ">
        <div class="modal-content passwordUser">
            <div class="modal-header">
                <p class="modal-title">修改密码</p>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true" style="line-height: 50px;padding: 8px;">&times;</span></button>
            </div>
            <div class="modal-body userForm">
                <ul>
                    <li class="positionR" >
                        <span>密码：</span>
                        <input type="password" style="margin-left: 27px" id="changePassWord">
                    </li>
                    <li class="positionR">
                        <span>确认密码：</span>
                        <input type="password" id="reChangePassWord">
                    </li>
                </ul>
            </div>
            <div class="modal-footer modify-footer">
                <p style="margin: auto">
                    <button type="button" class="btn btn-default" onclick="changePassWord()">确定</button>
                    <button type="button" class="btn btn-primary "  data-dismiss="modal" onclick="removeDelete()">取消</button>
                </p>
            </div>
        </div>
    </div>
</div>

<div class="maincon">
    <iframe id="mainCon" name="mainCon" frameborder="0" style="overflow-y: scroll;width:100%;height:100%;margin-left:0;"></iframe>
</div>

<script src="Default/public/js/jquery.min.js"></script>
<!--<script src="Default/public/js/bootstarp/bootstrap.bundle.js"></script>
<script src="Default/public/js/sb-admin.js"></script>
<script src="Default/public/js/multiple-select.js"></script>
<script src="Default/public/calendar/moment.js"></script>
<script src="Default/public/calendar/daterangepicker.js"></script>
<script src="Default/public/calendar/calendar.js"></script>
<script src="Default/public/js/publicURL.js"></script>
<script src="Default/public/js/ZSelect.js"></script>
<script src="Default/public/js/nav.js"></script>
<script src="Default/js/user.js"></script>-->

<script>
    $(document).ajaxComplete(function(event,xhr){
        if(xhr.getResponseHeader("sessionstatus") =="timeOut"){
            window.location.href = 'login.html'
        }
    });
    changeHeight("mainCon")
</script>
</body>
</html>