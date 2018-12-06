  <script src="/static/js/jquery-3.2.1.min.js"></script>
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
