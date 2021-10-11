from django.http import HttpResponseRedirect, HttpResponse
from django.shortcuts import render

from django.contrib.auth.models import User
from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required

def register_view(request):
    #注册
    if request.method == 'GET':
        return render(request, 'register.html')
    elif request.method == 'POST':
        username = request.POST['username']
        password_1 = request.POST['password_1']
        password_2 = request.POST['password_2']
        if password_1 != password_2:
            return HttpResponse('---两次密码输入不一致---')
        #TODO 查询用户名是否已注册

        #务必使用create_user创建用户
        user = User.objects.create_user(username=username,password=password_1)
        #如果需要注册后 免登录
        #login(request,user)
        #return HttpResponseRedirect('/index')
        return HttpResponseRedirect('/login')


def login_view(request):
    #登录
    if request.method == 'GET':
        return render(request, 'login.html')
    elif request.method == 'POST':
        username = request.POST['username']
        password = request.POST['password']

        user = authenticate(username=username, password=password)
        if not user:
            #用户名 或 密码错误
            return HttpResponse('--用户名或密码错误--')
        else:
            #校验成功
            #记录会话状态
            login(request, user)
            return HttpResponseRedirect('/index')


def logout_view(request):
    #退出
    logout(request)
    return HttpResponse('---已退出')


@login_required
def index_view(request):
    #首页， 必须登录才能访问，未登录跳转至settings.LOGIN_URL
    user = request.user
    return HttpResponse('欢迎 %s 来到 测试内部验证的首页'%(user.username))


