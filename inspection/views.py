import base64

from django.template import loader
from django.shortcuts import  HttpResponse

# Create your views here.
from inspection import models


# 首页
def index(request):
    context = {}
    if request.method=="POST":
        username= request.POST.get("username")
        password=request.POST.get("password")
        user_list = models.Users.objects.all()
        for user in user_list:
            pwd = base64.b64decode(user.password).decode("utf-8")
            if user.username==username and pwd==password:
                context={"username":username}
                template = loader.get_template('inspection/index.html')
                return HttpResponse(template.render(context, request))
            else:
                template = loader.get_template('inspection/page_500.html')
                return HttpResponse(template.render(context, request))
    else:
        template = loader.get_template('inspection/index.html')
        return HttpResponse(template.render(context, request))

# 登陆
def login(request):
    context = {}
    template = loader.get_template('inspection/login.html')
    return HttpResponse(template.render(context, request))

# 注册
def register(request):
    context = {}
    if request.method=="POST":
        username= request.POST.get("username")
        password=request.POST.get("password")
        repassword=request.POST.get("repassword")
        email=request.POST.get("email")
        if password!=repassword:
            print("两次输入密码不一致。")
            template = loader.get_template('/#signup')
            return HttpResponse(template.render(context, request))
        # 判断数据库中是否存在重复名称和邮箱
        user_list = models.Users.objects.all()

        for user in user_list:
            if user.username==username or user.email==email:
                print("用户名或邮箱重复")
                template = loader.get_template('/#signup')
                return HttpResponse(template.render(context, request))
        pwd = password.encode("utf-8")
        ciphertext = base64.b64encode(pwd).decode()  # 被编码的参数必须是二进制数据

        models.Users.objects.create(username=username, password=ciphertext, repassword=ciphertext,
                                    email=email)
        template = loader.get_template('inspection/login.html')
        return HttpResponse(template.render(context, request))

    else:
        context = {}
        template = loader.get_template('/#signup')
        return HttpResponse(template.render(context, request))

# html
def galaxy_html(request):
    context = {}
    # The template to be loaded as per gentelella.
    # All resource paths for gentelella end in .html.

    # Pick out the html file name from the url. And load that template.
    load_template = request.path.split('/')[-1]
    template = loader.get_template('inspection/' + load_template)
    return HttpResponse(template.render(context, request))