from django.db import models

class Books(models.Model):
    title=models.CharField('书名',max_length=100,default='')
    pub=models.CharField('出版社',max_length=100,default='')
    price=models.DecimalField('价格',max_digits=7,decimal_places=2,default=0)
    market_price=models.DecimalField('市场价',max_digits=7,decimal_places=2,default=0)

    class Meta:
        db_table='books'

    def __str__(self):
        return "%s_%s_%s"%(self.title,self.pub,self.price)

class book(models.Model):
    title=models.CharField('书名',max_length=50,default='')
    pub=models.CharField('出版社',max_length=100,null=False,default='')
    price=models.DecimalField('价格',max_digits=7,decimal_places=2,default=0.0)
    market_price=models.DecimalField('图书零售价',max_digits=7,decimal_places=2,default=0.0)
    info=models.CharField(max_length=100,default='')
    class Meta:
        db_table='book'

class author(models.Model):
    name=models.CharField(max_length=10,default='ma')
    age=models.IntegerField(default=1)
    email=models.EmailField(null=True)
    class Meta:
        db_table='author' # 可改变当前模型类对应的表

