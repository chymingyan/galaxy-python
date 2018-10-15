$('#password').focus(function () {
        var user = $('#username').val();
        if (getCookie('user') == user){
            var pwd = getCookie('pswd');
            $('#password').val(pwd);
            $('input[type="checkbox"]').attr('checked','checked')
        }
    }
);

//cookie
$('#remember').onchange = function(){
    if(!this.checked){
        delCookie('user');
        delCookie('pswd')
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
    if (userName === ''||userName === undefined||password === ''||password === undefined) {
        alert("");
        return false
    }
    $.ajax({
        url:url.login,
        data:{userName:userName,password:password},
        type:"post",
        dataType:"json",
        success:function (root) {
            if(root === 'success'){
                window.location.href='index.html';
                if($('input[type="checkbox"]:checked')){
                    setCookie('user',userName,7); //
                    setCookie('pswd',password,7) //
                }
            }else if(root === ""){
                alert(root);
                $('#Login').attr('style','display:none');
                $('#active').attr('style','display:block');
                getCode()
            }else if(root === ""){
                window.location.href='index.ptl';
                if($('input[type="checkbox"]:checked')){
                    setCookie('user',userName,7); //
                    setCookie('pswd',password,7) //
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