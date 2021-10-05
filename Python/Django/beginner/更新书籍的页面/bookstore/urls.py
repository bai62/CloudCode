from django.urls import path
from . import views

urlpatterns=[
    path('all_book/',views.all_book),
    path('<int:page>/',views.update),
    path('submit/',views.submit),
    path('delete/<int:page>/',views.delete),
]