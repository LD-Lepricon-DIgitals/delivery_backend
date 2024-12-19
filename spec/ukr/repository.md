# Services

## Опис

Модуль `Services` визначає інтерфейси для різних функціональних частин системи, таких як користувачі, замовлення, страви, рецензії тощо. Це забезпечує чітку структуру та гнучкість у реалізації бізнес-логіки.

---

## UserServices

Інтерфейс для роботи з користувачами.

| Метод                           | Опис                                      |
|---------------------------------|-------------------------------------------|
| `CreateUser(user UserReg)`      | Створює нового користувача.              |
| `GetUserId(login string)`       | Отримує ID користувача за його логіном.  |
| `IsCorrectPassword(login, password string)` | Перевіряє правильність пароля.           |
| `IfUserExists(login string)`    | Перевіряє, чи існує користувач.          |
| `ChangeUserCredentials(id int, info ChangeUserCredsPayload)` | Оновлює дані користувача. |
| `ChangePassword(id int, password string)` | Змінює пароль користувача.              |
| `DeleteUser(id int)`            | Видаляє користувача.                     |
| `IsCorrectPasswordId(id int, password string)` | Перевіряє пароль за ID користувача.     |
| `GetUserInfo(id int)`           | Повертає інформацію про користувача.     |
| `UpdatePhoto(photo string, userId int)` | Оновлює фотографію користувача.        |
| `GetUserRole(userId int)`       | Отримує роль користувача.                |

---

## OrderServices

Інтерфейс для роботи із замовленнями.

| Метод                                    | Опис                                   |
|------------------------------------------|----------------------------------------|
| `CreateOrder(order CreateOrder)`         | Створює нове замовлення.              |
| `GetOrders(workerId int)`                | Отримує список замовлень.             |
| `FinishOrder(orderId, workerId int)`     | Позначає замовлення як завершене.     |
| `StartOrder(orderId, workerId int)`      | Починає виконання замовлення.         |
| `GetOrderDetails(orderId int)`           | Отримує деталі замовлення.            |

---

## DishServices

Інтерфейс для роботи зі стравами.

| Метод                                  | Опис                                   |
|----------------------------------------|----------------------------------------|
| `AddDish(name string, price, weight float64, description, photo string, category int)` | Додає нову страву. |
| `GetDishes()`                          | Отримує список усіх страв.            |
| `DeleteDish(id int)`                   | Видаляє страву за ID.                 |
| `ChangeDish(id int, name string, price, weight float64, description, photo string, category int)` | Оновлює дані страви. |
| `GetDishesByCategory(category string)` | Отримує страви за категорією.         |
| `GetDishById(id int)`                  | Отримує деталі страви за ID.          |
| `SearchByName(name string)`            | Шукає страви за назвою.               |
| `AddCategory(categoryName string)`     | Додає нову категорію страв.           |
| `GetCategories()`                      | Отримує список категорій.             |

---

## Repository

Об'єднує всі сервіси для використання в системі.

| Поле               | Тип            | Опис                                  |
|--------------------|----------------|---------------------------------------|
| `UserServices`     | `UserServices` | Сервіси для роботи з користувачами.  |
| `AdminServices`    | `AdminServices`| Сервіси для адміністративних функцій.|
| `WorkerServices`   | `WorkerServices` | Сервіси для роботи з працівниками.   |
| `OrderServices`    | `OrderServices` | Сервіси для роботи із замовленнями.  |
| `DishServices`     | `DishServices` | Сервіси для роботи зі стравами.      |
| `ReviewServices`   | `ReviewServices` | Сервіси для роботи з рецензіями.    |

---

Модуль `Repository` забезпечує чітке розмежування функціоналу та дозволяє реалізувати різні аспекти бізнес-логіки. Завдяки інтерфейсам, визначеним у цьому модулі, система залишається гнучкою, легко розширюваною та підтримуваною. Ця структура сприяє організації коду та зручності роботи з основними функціональними компонентами проєкту.

