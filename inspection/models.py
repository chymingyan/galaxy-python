from django.db import models

# Create your models here.
class Users(models.Model):
    id=models.AutoField(primary_key=True)
    username=models.CharField(max_length=20)
    password=models.CharField(max_length=50)
    repassword=models.CharField(max_length=50)
    email=models.CharField(max_length=50)
    isenable=models.BooleanField(default=False)
    level=models.IntegerField(null=True)
    createdate=models.DateTimeField(null=True, auto_now_add=True)
    lastlogintime=models.DateTimeField(null=True,auto_now=True)

class UserLevel(models.Model):
    id=models.AutoField(primary_key=True)
    levelname=models.CharField(max_length=50)
