from django.urls import path, re_path
from inspection import views

urlpatterns = [
    # Matches any html file - to be used for gentella
    # Avoid using your .html in your resources.
    # Or create a separate django app.
    re_path(r'^.*\.html', views.galaxy_html, name='inspection'),

    # The home page

    path('', views.login, name='login'),
    path('index', views.index, name='index'),
]