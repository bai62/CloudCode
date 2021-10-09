from django.db import models

# Create your models here.

class User(models.Model):
    username=models.CharField('用户名',max_length=15,unique=True,default='')
    # md5 generates a string that longer than password can accommodate.
    password=models.CharField('密码',max_length=32,default='')
    created=models.DateTimeField('创建时间',auto_now_add=True)
    modified=models.DateTimeField('修改时间',auto_now=True)

    def __str__(self):
        return '用户名 %s'%(self.username)