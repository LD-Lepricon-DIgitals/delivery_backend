# Handlers

## Опис модуля

Модуль `handlers` відповідає за обробку HTTP-запитів. Він реалізує основну бізнес-логіку для роботи з користувачами, авторизації, замовленнями та іншими функціями застосунку. У цьому модулі використовуються сервіси для взаємодії з базою даних і виконання операцій.

---

## Основні компоненти

### Структура `Handlers`

```go
type Handlers struct {
services *service.Services // Сервіси для виконання основних операцій
}
```

- **`services`**: Інтерфейс для взаємодії з логікою сервісного рівня (робота з БД, авторизація тощо).

### Конструктор `NewHandlers`

Ініціалізує `Handlers` з наданими сервісами:

```go
func NewHandlers(services *service.Services) *Handlers
```

---

## Основні функції

## User handlers
### `RegisterUser`

Реєструє нового користувача, створює токен і зберігає його в cookie.

```go
func (h *Handlers) RegisterUser(c fiber.Ctx) error
```

- **Маршрут**: `POST /auth/register`
- **Опис**: Реєстрація нового користувача з переданими обліковими даними.
- **Параметри**:
    - JSON-запит типу `models.UserReg`.
- **Відповіді**:
    - `200 OK`: Успішна реєстрація.
    - `400 Bad Request`: Некоректні дані.
    - `409 Conflict`: Користувач вже існує.

---

### `LoginUser`

Авторизує існуючого користувача, створює токен і зберігає його в cookie.

```go
func (h *Handlers) LoginUser(c fiber.Ctx) error
```

- **Маршрут**: `POST /auth/login`
- **Опис**: Авторизація користувача за обліковими даними.
- **Параметри**:
    - JSON-запит типу `LoginPayload`.
- **Відповіді**:
    - `200 OK`: Успішний вхід.
    - `400 Bad Request`: Некоректні дані.
    - `401 Unauthorized`: Неправильні облікові дані.

---

### `ChangeUserCredentials`

Оновлює особисті дані користувача.

```go
func (h *Handlers) ChangeUserCredentials(c fiber.Ctx) error
```

- **Маршрут**: `PATCH /api/user/change`
- **Опис**: Оновлення деталей облікового запису користувача.
- **Параметри**:
    - JSON-запит типу `models.ChangeUserCredsPayload`.
- **Відповіді**:
    - `200 OK`: Успішне оновлення.
    - `400 Bad Request`: Некоректні дані.
    - `401 Unauthorized`: Користувач не авторизований.

---

### `ChangeUserPassword`

Змінює пароль користувача.

```go
func (h *Handlers) ChangeUserPassword(c fiber.Ctx) error
```

- **Маршрут**: `PATCH /api/user/change_password`
- **Опис**: Зміна пароля для облікового запису користувача.
- **Параметри**:
    - JSON-запит типу `ChangePasswordPayload`.
- **Відповіді**:
    - `200 OK`: Успішна зміна пароля.
    - `400 Bad Request`: Некоректні дані.
    - `401 Unauthorized`: Неправильний старий пароль.

---

### `LogoutUser`

Вихід користувача із системи.

```go
func (h *Handlers) LogoutUser(c fiber.Ctx) error
```

- **Маршрут**: `POST /api/user/logout`
- **Опис**: Видаляє токен авторизації, очищуючи cookie.
- **Відповіді**:
    - `200 OK`: Успішний вихід.

---

### `DeleteUser`

Видаляє обліковий запис користувача.

```go
func (h *Handlers) DeleteUser(c fiber.Ctx) error
```

- **Маршрут**: `DELETE /api/user/delete`
- **Опис**: Видаляє обліковий запис користувача за його ID.
- **Відповіді**:
    - `200 OK`: Успішне видалення.
    - `401 Unauthorized`: Користувач не авторизований.

---

### `GetUserInfo`

Отримує інформацію про користувача.

```go
func (h *Handlers) GetUserInfo(c fiber.Ctx) error
```

- **Маршрут**: `GET /api/user/info`
- **Опис**: Повертає деталі облікового запису користувача.
- **Відповіді**:
    - `200 OK`: Успішно отримано інформацію.
    - `401 Unauthorized`: Користувач не авторизований.

---

### `UpdatePhoto`

Оновлює фото профілю користувача.

```go
func (h *Handlers) UpdatePhoto(c fiber.Ctx) error
```

- **Маршрут**: `PATCH /api/user/photo`
- **Опис**: Оновлює фото профілю користувача.
- **Параметри**:
    - JSON-запит типу `ChangePhotoPayload`.
- **Відповіді**:
    - `200 OK`: Фото успішно оновлено.
    - `400 Bad Request`: Некоректні дані.
    - `401 Unauthorized`: Користувач не авторизований.

---

## Dish handlers
### `GetDishes`

Отримує всі доступні страви.

```go
func (h *Handlers) GetDishes(ctx fiber.Ctx) error
```

- **Маршрут**: `GET /api/dishes`
- **Опис**: Повертає список усіх доступних страв.
- **Відповіді**:
    - `200 OK`: Список страв.
    - `404 Not Found`: Страви не знайдено.
    - `500 Internal Server Error`: Помилка сервера.

---

### `AddDish`

Додає нову страву.

```go
func (h *Handlers) AddDish(ctx fiber.Ctx) error
```

- **Маршрут**: `POST /api/dishes/admin/add`
- **Опис**: Додає нову страву з переданими даними.
- **Параметри**:
    - JSON-запит типу `AddDishPayload`.
- **Відповіді**:
    - `201 Created`: Страву успішно додано.
    - `400 Bad Request`: Некоректні дані.
    - `403 Forbidden`: Немає доступу.
    - `500 Internal Server Error`: Помилка сервера.

---

### `DeleteDish`

Видаляє страву за ID.

```go
func (h *Handlers) DeleteDish(ctx fiber.Ctx) error
```

- **Маршрут**: `DELETE /api/dishes/admin/delete/{id}`
- **Опис**: Видаляє страву за її ID.
- **Відповіді**:
    - `200 OK`: Страву успішно видалено.
    - `400 Bad Request`: Некоректний ID страви.
    - `403 Forbidden`: Немає доступу.
    - `500 Internal Server Error`: Помилка сервера.

---

### `ChangeDish`

Оновлює дані страви.

```go
func (h *Handlers) ChangeDish(ctx fiber.Ctx) error
```

- **Маршрут**: `PUT /api/dishes/admin/update`
- **Опис**: Оновлює деталі страви.
- **Параметри**:
    - JSON-запит типу `ChangeDishPayload`.
- **Відповіді**:
    - `200 OK`: Страву успішно оновлено.
    - `400 Bad Request`: Некоректні дані.
    - `403 Forbidden`: Немає доступу.
    - `500 Internal Server Error`: Помилка сервера.

---

### `GetDishesByCategory`

Отримує страви за категорією.

```go
func (h *Handlers) GetDishesByCategory(ctx fiber.Ctx) error
```

- **Маршрут**: `POST /api/dishes/by_category`
- **Опис**: Повертає список страв за вказаною категорією.
- **Параметри**:
    - JSON-запит типу `GetDishesByCategoryPayload`.
- **Відповіді**:
    - `200 OK`: Список страв.
    - `400 Bad Request`: Некоректні дані.
    - `404 Not Found`: Страви не знайдено.
    - `500 Internal Server Error`: Помилка сервера.

---

### `GetDishById`

Отримує страву за її ID.

```go
func (h *Handlers) GetDishById(ctx fiber.Ctx) error
```

- **Маршрут**: `GET /api/dishes/by_id/{dish_id}`
- **Опис**: Повертає деталі страви за її ID.
- **Відповіді**:
    - `200 OK`: Деталі страви.
    - `400 Bad Request`: Некоректний ID страви.
    - `404 Not Found`: Страву не знайдено.
    - `500 Internal Server Error`: Помилка сервера.

---

### `SearchByName`

Шукає страви за назвою.

```go
func (h *Handlers) SearchByName(ctx fiber.Ctx) error
```

- **Маршрут**: `GET /api/dishes/search/{name}`
- **Опис**: Шукає страви за їхньою назвою.
- **Відповіді**:
    - `200 OK`: Список знайдених страв.
    - `400 Bad Request`: Некоректні дані.
    - `404 Not Found`: Страви не знайдено.
    - `500 Internal Server Error`: Помилка сервера.

---

### `AddCategory`

Додає нову категорію страв.

```go
func (h *Handlers) AddCategory(ctx fiber.Ctx) error
```

- **Маршрут**: `POST /api/categories`
- **Опис**: Додає нову категорію страв за переданими даними.
- **Параметри**:
    - JSON-запит типу `AddDishCategoryPayload`.
- **Відповіді**:
    - `201 Created`: Категорію успішно додано.
    - `400 Bad Request`: Некоректні дані.
    - `403 Forbidden`: Немає доступу.
    - `500 Internal Server Error`: Помилка сервера.

---

### `GetCategories`

Отримує всі доступні категорії страв.

```go
func (h *Handlers) GetCategories(ctx fiber.Ctx) error
```

- **Маршрут**: `GET /api/categories`
- **Опис**: Повертає список усіх доступних категорій страв.
- **Відповіді**:
    - `200 OK`: Список категорій.
    - `500 Internal Server Error`: Помилка сервера.

## Order handlers
### `CreateOrder`

Створює нове замовлення.

```go
func (h *Handlers) CreateOrder(ctx fiber.Ctx) error
```

- **Маршрут**: `POST /api/orders`
- **Опис**: Створює замовлення для авторизованого користувача або вказаного клієнта.
- **Параметри**:
    - JSON-запит типу `models.CreateOrder`.
- **Відповіді**:
    - `200 OK`: Замовлення успішно створено.
    - `400 Bad Request`: Некоректний запит.
    - `401 Unauthorized`: Недійсний або відсутній токен.
    - `403 Forbidden`: ID користувача не відповідає ID клієнта у замовленні.
    - `500 Internal Server Error`: Помилка сервера.

---

### `GetOrders`

Отримує всі замовлення для авторизованого користувача з роллю "worker".

```go
func (h *Handlers) GetOrders(ctx fiber.Ctx) error
```

- **Маршрут**: `GET /api/orders`
- **Опис**: Повертає список усіх замовлень, пов'язаних із користувачем.
- **Відповіді**:
    - `200 OK`: Список замовлень успішно отримано.
    - `401 Unauthorized`: Недійсний або відсутній токен.
    - `403 Forbidden`: Користувач не має необхідної ролі.
    - `500 Internal Server Error`: Помилка сервера.

---

### `GetOrderDetails`

Отримує деталі замовлення за його ID.

```go
func (h *Handlers) GetOrderDetails(ctx fiber.Ctx) error
```

- **Маршрут**: `GET /api/orders/{order_id}`
- **Опис**: Повертає детальну інформацію про замовлення. Доступно лише для користувачів з роллю "worker".
- **Відповіді**:
    - `200 OK`: Деталі замовлення успішно отримано.
    - `400 Bad Request`: Некоректний ID замовлення.
    - `401 Unauthorized`: Недійсний або відсутній токен.
    - `403 Forbidden`: Користувач не має необхідної ролі.
    - `500 Internal Server Error`: Помилка сервера.

---

### `ConfirmOrder`

Підтверджує замовлення за його ID.

```go
func (h *Handlers) ConfirmOrder(ctx fiber.Ctx) error
```

- **Маршрут**: `POST /api/orders/confirm/{order_id}`
- **Опис**: Підтверджує замовлення, позначаючи його як прийняте користувачем з роллю "worker".
- **Відповіді**:
    - `200 OK`: Замовлення успішно підтверджено.
    - `400 Bad Request`: Некоректний ID замовлення.
    - `401 Unauthorized`: Недійсний або відсутній токен.
    - `403 Forbidden`: Користувач не має необхідної ролі.
    - `500 Internal Server Error`: Помилка сервера.

---

### `FinishOrder`

Завершує замовлення за його ID.

```go
func (h *Handlers) FinishOrder(ctx fiber.Ctx) error
```

- **Маршрут**: `POST /api/orders/finish/{order_id}`
- **Опис**: Позначає замовлення як завершене користувачем з роллю "worker".
- **Відповіді**:
    - `200 OK`: Замовлення успішно завершено.
    - `400 Bad Request`: Некоректний ID замовлення.
    - `401 Unauthorized`: Недійсний або відсутній токен.
    - `403 Forbidden`: Користувач не має необхідної ролі.
    - `500 Internal Server Error`: Помилка сервера.

---

## Валідація

### `verifyUserToken`

Функція валідації токена авторизованого користувача. Використовується у всіх обробниках для перевірки автентичності запитувача.

```go
func verifyUserToken(ctx fiber.Ctx) (int, string, error)
```

- **Повертає**:
    - ID користувача.
    - Роль користувача.
    - Помилку, якщо токен недійсний або користувач не авторизований.

