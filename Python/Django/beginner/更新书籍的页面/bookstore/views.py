from django.shortcuts import render
from django.http import HttpResponseRedirect
from .models import Books


def all_book(request):
    all_book = Books.objects.all()
    return render(request, 'bookstotre/show_book.html', locals())


def update(request, page):
    obj = Books.objects.get(id=page);
    return render(request, 'bookstotre/update_info.html', locals());


def submit(request):
    obj = Books.objects.get(id=request.GET['id']);
    obj.price = request.GET['price'];
    obj.market_price = request.GET['market_price'];
    obj.save();
    return HttpResponseRedirect('/bookstore/all_book/');


def delete(request,page):
    Books.objects.get(id=page).delete();
    return HttpResponseRedirect('/bookstore/all_book/');

