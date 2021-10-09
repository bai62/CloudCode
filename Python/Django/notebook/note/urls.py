from django.urls import path
from . import views

urlpatterns=[
    path('add_note',views.add_note),
]