<div class="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
 <h1 class="page-header">主机数据库配置</h1>
<ul class="nav nav-pills">
  <li role="presentation" class="active"><a href="/config">主机系统配置</a></li>
  <li role="presentation"><a href="/targetdbs">数据库配置</a></li>
</ul>

<!--表格-->
<div class="panel panel-default">
  <!-- Default panel head -->
  <div class="panel-heading">Panel heading</div>
<!-- Default panel head -->
<div class="panel-body">
   <form class="navbar-form navbar-right">
<!--主机编号-->
<div class="input-group">
  <!--<span class="input-group-addon" id="basic-addon1">主机编号:</span>-->
  <input type="text" class="form-control" placeholder="主机编号" aria-describedby="basic-addon1">
</div>
<!--主机物理IP-->
<div class="input-group">
  <!--<span class="input-group-addon" id="basic-addon1">主机物理IP:</span>-->
  <input type="text" class="form-control" placeholder="主机屋里IP" aria-describedby="basic-addon1">
</div>
<!--主机端口-->
<div class="input-group">
  <!--<span class="input-group-addon" id="basic-addon1">主机端口:</span>-->
  <input type="text" class="form-control" placeholder="主机端口" aria-describedby="basic-addon1">
</div>
<!--数据库软件用户名-->
<div class="input-group">
 <!-- <span class="input-group-addon" id="basic-addon1">数据库软件用户名:</span>-->
  <input type="text" class="form-control" placeholder="数据库软件用户名" aria-describedby="basic-addon1">
</div>
<!--数据库用户密码-->
<div class="input-group">
 <!-- <span class="input-group-addon" id="basic-addon1">数据库用户密码:</span>-->
  <input type="text" class="form-control" placeholder="数据库用户密码" aria-describedby="basic-addon1">
</div>
<!--集群软件用户名-->
<div class="input-group">
 <!-- <span class="input-group-addon" id="basic-addon1">集群软件用户名:</span>-->
  <input type="text" class="form-control" placeholder="集群软件用户名" aria-describedby="basic-addon1">
</div>
<!--集群软件用户密码-->
<div class="input-group">
 <!-- <span class="input-group-addon" id="basic-addon1">集群软件用户密码:</span>-->
  <input type="text" class="form-control" placeholder="集群软件用户密码" aria-describedby="basic-addon1">
</div>
<!--ROOT用户密码-->
<div class="input-group">
  <!--<span class="input-group-addon" id="basic-addon1">ROOT用户密码:</span>-->
  <input type="text" class="form-control" placeholder="ROOT用户密码" aria-describedby="basic-addon1">
</div>
<!--数据库软件目录-->
<div class="input-group">
 <!-- <span class="input-group-addon" id="basic-addon1">数据库软件目录:</span>-->
  <input type="text" class="form-control" placeholder="数据库软件目录" aria-describedby="basic-addon1">
</div>
<!--集群软件目录-->
<div class="input-group">
<!--  <span class="input-group-addon" id="basic-addon1">集群软件目录:</span>-->
  <input type="text" class="form-control" placeholder="集群软件目录" aria-describedby="basic-addon1">
</div>
<!--主机协议-->
主机协议:
<div class="btn-group" role="group" >
  <button type="button" class="btn btn-default">SSH</button>
</div>
<!--主机系统-->
主机系统:
<div class="btn-group" role="group" >
  <button type="button" class="btn btn-default">LINUX</button>
  <button type="button" class="btn btn-default">AIX</button>
</div>
<!--Oracle数据库版本-->
<div class="btn-group">
  <button class="btn btn-default dropdown-toggle" type="button" id="dropdownOracleVersion" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
    数据库版本
    <span class="caret"></span>
	</button>
  <ul class="dropdown-menu">
    <li><a href="#">ORACLE 9I</a></li>
    <li><a href="#">ORACLE 10G</a></li>
    <li><a href="#">ORACLE 11G</a></li>
    <li><a href="#">ORACLE 12C</a></li>
	<li><a href="#">ORACLE 18C</a></li>
  </ul>
</div>
<a class="btn btn-primary" href="#" role="button">测试链接</a>
<a class="btn btn-primary" href="#" role="button">保存主机</a>
<a class="btn btn-primary" href="#" role="button">清空</a>
</form>
  </div>
  <!-- Table -->
   <div class="table-responsive">
            <table class="table table-striped">
              <thead>
                <tr>
                  <th>编号</th>
                  <th>主机地址</th>
                  <th>SSH端口</th>
                  <th>Oracle用户</th>
				  <th>Grid用户</th>
                  <th>Oracle版本</th>
				<th>ORACLE_HOME</th>
				<th>GRID_HOME</th>
				<th>系统版本</th>
				<th>操作</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>1,001</td>
                  <td>Lorem</td>
                  <td>ipsum</td>
                  <td>dolor</td>
                  <td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td><button type="button" class="btn btn-default" aria-label="Left Align">
  <span class="glyphicon glyphicon-edit" aria-hidden="true"></span>
</button>
<button type="button" class="btn btn-default" aria-label="Left Align">
  <span class="glyphicon glyphicon-trash" aria-hidden="true"></span>
</button></td>
                </tr>
                <tr>
                  <td>1,002</td>
                  <td>amet</td>
                  <td>consectetur</td>
                  <td>adipiscing</td>
                  <td>elit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,003</td>
                  <td>Integer</td>
                  <td>nec</td>
                  <td>odio</td>
                  <td>Praesent</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,003</td>
                  <td>libero</td>
                  <td>Sed</td>
                  <td>cursus</td>
                  <td>ante</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,004</td>
                  <td>dapibus</td>
                  <td>diam</td>
                  <td>Sed</td>
                  <td>nisi</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,005</td>
                  <td>Nulla</td>
                  <td>quis</td>
                  <td>sem</td>
                  <td>at</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,006</td>
                  <td>nibh</td>
                  <td>elementum</td>
                  <td>imperdiet</td>
                  <td>Duis</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,007</td>
                  <td>sagittis</td>
                  <td>ipsum</td>
                  <td>Praesent</td>
                  <td>mauris</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,008</td>
                  <td>Fusce</td>
                  <td>nec</td>
                  <td>tellus</td>
                  <td>sed</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,009</td>
                  <td>augue</td>
                  <td>semper</td>
                  <td>porta</td>
                  <td>Mauris</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,010</td>
                  <td>massa</td>
                  <td>Vestibulum</td>
                  <td>lacinia</td>
                  <td>arcu</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,011</td>
                  <td>eget</td>
                  <td>nulla</td>
                  <td>Class</td>
                  <td>aptent</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,012</td>
                  <td>taciti</td>
                  <td>sociosqu</td>
                  <td>ad</td>
                  <td>litora</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,013</td>
                  <td>torquent</td>
                  <td>per</td>
                  <td>conubia</td>
                  <td>nostra</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,014</td>
                  <td>per</td>
                  <td>inceptos</td>
                  <td>himenaeos</td>
                  <td>Curabitur</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
                <tr>
                  <td>1,015</td>
                  <td>sodales</td>
                  <td>ligula</td>
                  <td>in</td>
                  <td>libero</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
				<td>sit</td>
                </tr>
              </tbody>
            </table>
          </div>

</div>    
</div>
