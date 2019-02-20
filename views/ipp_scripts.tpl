 <!-- jQuery (Bootstrap 的所有 JavaScript 插件都依赖 jQuery，所以必须放在前边) -->
    <script src="static/js/jquery.min.js"></script>
    <!-- 加载 Bootstrap 的所有 JavaScript 插件。你也可以根据需要只加载单个插件。 -->
    <script src="static/js/bootstrap.min.js"></script>
	<script src="static/js/bootstrap-treeview.min.js"></script>
<script type="text/javascript">
    $(document).ready(function() {
       
    });
	function inspClick(){
		 $.ajax({
        url:"/insp",
        data:{template:"insp"},
        type:"post",
        dataType:"json",
		async: false,
        success:function (msg) {
			  if(msg.Val === 'success'){
           alert("ok");
			}
		}});
	}
	
		function overviewClick(){
		 $.ajax({
        url:"/chat",
        data:{template:"overview"},
        type:"post",
        dataType:"json",
		async: true,
        success:function (msg) {
			  if(msg.Val === 'success'){
            alert("ok");
			}
		}});
	}
	
</script>
