from django.shortcuts import render
from .models import Note
# Create your views here.
from django.http import HttpResponseRedirect, HttpResponse

# A decorator in Python is a function that takes another function as its argument, and returns yet another function.
'''def add_num(fnc):
    def inner(tuple_list):
            return [fnc(tuple[0],tuple[1]) for tuple in tuple_list]
    return inner'''


def session_check(fnc):
    def inner(request):
        if request.COOKIES.get('username'):
            request.session['username'] = request.COOKIES.get('username')
            request.session['userid'] = request.COOKIES.get('userid')
            return fnc(request)
        else:
            return HttpResponseRedirect('/index')

    return inner


@session_check
def add_note(request):
    if request.method == 'GET':
        return render(request, 'note/add_note.html')
    elif request.method == 'POST':
        user_id = request.session.get('userid')
        title = request.POST.get('title')
        content = request.POST.get('content')
        Note.objects.create(user_id=user_id,title=title,content=content)
        return HttpResponseRedirect('/note/add_note')
