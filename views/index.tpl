
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <title>智能巡检平台</title>
    <link rel="shortcut icon" href="static/img/1234.ico">
    <link rel="stylesheet" href="static/css/base.css">
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
            <li class="nav-item nav-munu" id="menu" title="最新概览">
                <a class="nav-link" id="sidenavToggler" href="#" >
                    <i class="fa-fw fa-mun"></i>
            <!--        <span class="nav-link-text" >最新概览</span>-->
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right" id="overview" title="常规巡检" style="position:relative;">
                <a class="nav-link" href="overview.html" target="mainCon" >
                    <i class="fa-fw fa-dashboards"></i>
                    <!--<span class="nav-link-text nav_text select">常规巡检</span>-->
                    <span class="nav-red"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right" id="performance" title="虚拟化巡检" style="position:relative;" >
                <a class="nav-link" href="performance.html" target="mainCon" >
                    <i class="fa-fw fa-area-chart"></i>
              <!--      <span class="nav-link-text nav_text">虚拟化巡检</span>-->
                    <span class="nav-red" style="display: none;"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right" id="operate" title="导出报告" style="position:relative;">
                <a class="nav-link" href="operate.html" style="position:relative;" target="mainCon"  >
                    <i class="fa-fw fa-table"></i>
                    <!--<span class="nav-link-text nav_text">导出报告</span>-->
                    <span class="nav-red" style="display: none;"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right"  id= "dataBase" title="虚拟化报告" style="position:relative;">
                <a class="nav-link" href="database.html" style="position:relative;" target="mainCon"  >
                    <i class="fa-fw fa-data"></i>
           <!--         <span class="nav-link-text nav_text">虚拟化报告</span>-->
                    <span class="nav-red" style="display: none;"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right"  id= "report" title="巡检配置" style="position:relative;">
                <a class="nav-link" href="report.html" style="position:relative;" target="mainCon"  >
                    <i class="fa-fw fa-report"></i>
                 <!--   <span class="nav-link-text nav_text">巡检配置</span>-->
                    <span class="nav-red" style="display: none;"></span>
                </a>
            </li>
            <li class="nav-item" data-toggle="tooltip" data-placement="right" id="config" title="规则配置" style="position:relative;">
                <a class="nav-link" href="set.html" style="position:relative;" target="mainCon"  >
                    <i class="fa-fw fa-wrench"></i>
                 <!--   <span class="nav-link-text nav_text">规则配置</span>-->
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

<div class="maincon">
sdfjsdkljfsdkfklsdlfk
  
</div>

<script src="Default/public/js/jquery.min.js"></script>
<!--<script src="Default/public/js/bootstarp/bootstrap.bundle.js"></script>
<script src="Default/public/js/sb-admin.js"></script>
<script src="Default/public/js/multiple-select.js"></script>
<script src="Default/public/calendar/moment.js"></script>
<script src="Default/public/calendar/daterangepicker.js"></script>
<script src="Default/public/calendar/calendar.js"></script>
<script src="Default/public/js/ZSelect.js"></script>
<script src="Default/public/js/nav.js"></script>-->

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