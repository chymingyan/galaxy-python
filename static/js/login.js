$('#password').focus(function () {
        var user = $('#username').val();
        if (getCookie('user') == user){
            var pwd = getCookie('pswd');
	     var hostip=getCookie('hostip');
	     var hostport=getCookie('hostport');
	     var servicename=getCookie('servicename');
            $('#password').val(pwd);
	     $('#hostip').val(hostip);
	     $('#hostport').val(hostport);
	     $('#servicename').val(servicename);
            $('input[type="checkbox"]').attr('checked','checked')
        }
    }
);

//cookie
$('#remember').onchange = function(){
    if(!this.checked){
        delCookie('user');
        delCookie('pswd')
	 delCookie('hostip');
	delCookie('hostport');
	delCookie('servicename');
    }
};

$('#password').keydown(function (even) {
    if(even && even.keyCode==13){ // enter 
        login()
    }
});

$('#loginbtn').click(function () {
    login()
});

function login() {
    var userName = $('#username').val();
    var password = $('#password').val();
    var hostip=$('#hostip').val();
    var hostport=$('#hostport').val();
    var servicename=$('#servicename').val();
    if (userName === ''||userName === undefined||password === ''||password === undefined||hostip===''||hostip===undefined||hostport===''||hostport===undefined||servicename===''||servicename===undefined)
    {
        alert('userName:'+userName+'hostIp:'+hostip+'password:'+password+'hostPort:'+hostport+'serviceName:'+servicename);
        return false
    }
    $.ajax({
        url:"/login",
        data:{userName:userName,password:password,hostIp:hostip,hostPort:hostport,serviceName:servicename},
        type:"post",
        dataType:"json",
        success:function (root) {
            if(root === 'success'){
                window.location.href='index.tpl';
                if($('input[type="checkbox"]:checked')){
                    setCookie('user',userName,7); //
                    setCookie('pswd',password,7); //
		       setCookie('hostip',hostip,7);
			setCookie('hostport',hostport,7);
			setCookie('servicename',servicename,7)
                }
            }else if(root === ""){
                alert(root);
                $('#Login').attr('style','display:none');
                $('#active').attr('style','display:block');
                getCode()
            }else if(root === ""){
                window.location.href='index.tpl';
                if($('input[type="checkbox"]:checked')){
                    setCookie('user',userName,7); //
                    setCookie('pswd',password,7) //
			setCookie('hostip',hostip,7);
			setCookie('hostport',hostport,7);
			setCookie('servicename',servicename,7)
                }
            }else if(root === ""){
                alert(root);
                $('#Login').attr('style','display:none');
                $('#active').attr('style','display:block');
                getCode()
            }else {
                alert(root)
            }
        }
    })
}

//
function active() {
    var code = $('#textAreas').val();
    var sign = $('#activeCode').val();
    if (code === ''||code === undefined) {
        alert("");
        return false
    }
    $.ajax({
        url:url.active,
        data:{code:code,sign:sign},
        type:"post",
        dataType:"json",
        success:function (root) {
            if(root == 'success'){
                window.location.href='index.tpl'
            }else {
                alert(root)
            }
        }
    })
}

//
function getCode() {
    $.ajax({
        url:url.code,
        type:"post",
        dataType:"json",
        success:function (root) {
            $("#textAreas").val(root)
        }
    })
}

//cookie
function setCookie(name,value,day){
    var date = new Date();
    date.setDate(date.getDate() + day);
    document.cookie = name + '=' + value + ';expires='+ date
}

//cookie
function getCookie(name){
    var reg = RegExp(name+'=([^;]+)');
    var arr = document.cookie.match(reg);
    if(arr){
        return arr[1]
    }else{
        return ''
    }
}

//cookie
function delCookie(name){
    setCookie(name,null,-1)
}