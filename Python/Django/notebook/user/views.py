from django.shortcuts import render
from django.http import HttpResponse, HttpResponseRedirect
from .models import User
import hashlib


# Create your views here.

def session_update(request, username, userid):
    request.session['username'] = username
    request.session['userid'] = userid


def signup(request):
    if request.method == 'GET':
        return render(request, 'user/signup.html')
    elif request.method == 'POST':
        username = request.POST['username']
        password = request.POST['password']
        retype_password = request.POST['retype_password']

        if password != retype_password:
            return HttpResponse('<h1>两次密码不一致</h1>')
        # A efficient approach to check if a particular object exists or not
        elif User.objects.filter(username=username).exists():
            return HttpResponse('<h1>用户名已存在</h1>')

        # md5 use a cryptographic hash function that expresses as a 32 digit hexadecimal number
        md5 = hashlib.md5()
        md5.update(password.encode())
        password_m = md5.hexdigest()
        try:
            User.objects.create(username=username, password=password_m)
        except Exception as e:
            print('error %s'%(e))
            return HttpResponse('<h1>用户名已存<p>%s</p></h1>'%(e))

        # password-free login for one day
        session_update(request, username, User.objects.get(username=username).id)

        return HttpResponseRedirect('/index')


def signin(request):
    if request.method == 'GET':
        if request.session.get('username') and request.session.get('userid'):
            username=request.session['username']
            session_update(request, username, User.objects.get(username=username).id)
            return HttpResponseRedirect('/index')
        elif request.COOKIES.get('username') and request.session.get('userid'):
            username = request.COOKIES['username']
            session_update(request, username, User.objects.get(username=username).id)
            return HttpResponseRedirect('/index')
        return render(request, 'user/signin.html')
    elif request.method == 'POST':
        username = request.POST['username']
        password = request.POST['password']

        if not User.objects.filter(username=username).exists():
            return HttpResponse('<h1>用户名或密码错误</h1>')

        md5 = hashlib.md5()
        md5.update(password.encode())
        password_m = md5.hexdigest()
        if password_m != User.objects.get(username=username).password:
            return HttpResponse('<h1>用户名或密码错误</h1>')

        session_update(request, username, User.objects.get(username=username).id)

        resp = HttpResponseRedirect('/index')

        # if checkbox is checked, there is a remember_me field in POST dictionary
        if 'remember_me' in request.POST:
            resp.set_cookie('username', username, 60 * 60 * 24 * 3)
            resp.set_cookie('userid', User.objects.get(username=username).id, 60 * 60 * 24 * 3)

        return resp

def logout(request):
    if request.session.get('username'):
        del request.session['username']
    if request.session.get('userid'):
        del request.session['userid']

    resp=HttpResponseRedirect('/index')
    if request.COOKIES.get('username'):
        resp.delete_cookie('username')
    if request.COOKIES.get('userid'):
        resp.delete_cookie('userid')
    return resp