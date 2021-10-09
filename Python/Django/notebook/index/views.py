from django.shortcuts import render

# Create your views here.

def index_view(request):
    # request is passed to render as a parameter
    return render(request,'index/index.html')