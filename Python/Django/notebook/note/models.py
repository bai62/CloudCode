from django.db import models
from user.models import User
# Create your models here.

class Note(models.Model):
    title=models.CharField('文章标题',max_length=100,default='')
    content=models.TextField('内容',default='')
    created=models.DateTimeField('创建时间',auto_now_add=True)
    modified=models.DateTimeField('修改时间',auto_now=True)
    user=models.ForeignKey(User,on_delete=models.CASCADE)#外键字段名：user_id