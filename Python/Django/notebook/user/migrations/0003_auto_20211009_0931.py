# Generated by Django 3.2.8 on 2021-10-09 01:31

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('user', '0002_rename_uesr_user'),
    ]

    operations = [
        migrations.AlterField(
            model_name='user',
            name='password',
            field=models.CharField(default='', max_length=32, verbose_name='密码'),
        ),
        migrations.AlterField(
            model_name='user',
            name='username',
            field=models.CharField(default='', max_length=15, unique=True, verbose_name='用户名'),
        ),
    ]
