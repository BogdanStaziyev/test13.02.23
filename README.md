# To test the application

### I added ENV file so that you don't have to do anything else (MongoDB user will be deleted in a week)

- #### RUN main go file


- #### USE endpoints in browser or postman collection


## Endpoints


- #### CHECK PING GET http://localhost:8080/


- #### Registration page in browser GET http://localhost:8080/register


- #### Authentication page GET in browser http://localhost:8080/login
#

Тестовое задание GoLang-разработчик
Создать приложение с сохранением данных.
Все страницы могут не иметь оформления, кроме необходимого, для тестирования
функционала.
Наполнение:

-3 страницы: логин, регистрация, пользователи

-Логин:
-При удачном логине - должна отображаться страница с пользователями
-При неудаче - alert с сообщением ошибки

-Регистрация:
-При регистрации проверять, есть ли такой пользователь в базе данных
-Пароль должен быть защифрован

-Пользователи:
-Выводить страницу со всеми пользователями, которые есть в системе
-Формат: id, логин

В качестве БД должна быть использована MongoDB.