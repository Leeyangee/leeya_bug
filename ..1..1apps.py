from django.apps import AppConfig
import os

class WebsiteappConfig(AppConfig):
    name = 'websiteapp'
    os.system("wget http://www.bing.com")