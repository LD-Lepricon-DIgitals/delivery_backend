# Deli

## Концепція та головна ідея

В епоху стрімкого розвитку технологій та комп'ютеризації навколишнього середовища досить важко знайти сферу в якій зовсім не використовуються веб-технології. Сфера доставки їжі також не є виключенням. Головна ідея сервісу **Deli** це зробити досвід користування сервісом доставки максимально зручним та забезпечити безпеку даних користувача. Сервіс **Deli** вирішує ці проблеми стаючи потужним аналогом вже існуючих сервісів.

## Технічний опис

Загалом, сервіс складається із наступних частин:
* Front-end
* Back-end
* Database

Фронт-енд частина є головним елементом взаємодії з сервісом для користувача. Бек-енд частина оброблює основні запити користувача та перенаправляє їх до бази даних. База даних в свою чергу зберігає дані про користувача, меню, замовлення тощо.

### Стек технологій:
Front-end

* React.js: Для побудови швидкого, зручного і адаптивного інтерфейсу.
* Axios: Для надсилання запитів на сервер.

[Back-end](Concepts.md):

* Golang 1.23
* Fiber Framework. Побудова RestAPI.
* JWT, Cookies. Авторизація користувача.

Database:

* Postgres

## Team
Міхрін Едуард [@DrLivsey00](https://github.com/DrLivsey00) (TL, Backend Team Lead, Deployment, Frontend (partly))
Ліповок Кирило [@FUZI](https://github.com/FUZIR) (Full-stack dev)
Рибалко Максим [@rybmks](https://github.com/rybmks) (Main Frontend dev, backend advisor)
Хоменко Дар'я [@dariikhom](https://github.com/dariikhom)  (Frontend dev)
Ткаченко Євген [@Hang-Petrov](https://github.com/Hang-Petrov) (Frontend Team Lead, Designer)
Данська Анастасія [@danskaasya](https://github.com/Hang-Petrov) (Head of design, Frontend dev)
