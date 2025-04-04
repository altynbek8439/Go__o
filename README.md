Букмекерский сайт

Описание проекта

"Betting-site" — это веб-приложение, которое позволяет пользователям просматривать спортивные события, делать ставки и отслеживать свои ставки. Проект состоит из backend, написанного на Go с использованием фреймворка **Gin** и ORM **GORM**, и frontend, реализованного на React. Основной акцент в разработке был сделан на эффективное использование Gin для создания API и GORM для работы с базой данных PostgreSQL, что позволило создать надёжный и масштабируемый сервис.

### Основные функции:

- **Просмотр событий:** Пользователи могут видеть список доступных спортивных событий с датой и коэффициентами.
- **Добавление событий:** Администратор может добавлять новые события через форму.
- **Создание ставок:** Пользователи могут делать ставки на события, указывая сумму и предполагаемый исход.
- **Просмотр ставок:** Пользователи могут посмотреть свои ставки по ID.

## Использование GORM и Gin

### GORM: Работа с базой данных

В проекте я использовал **GORM** — мощную ORM-библиотеку для Go, которая упростила взаимодействие с базой данных PostgreSQL. GORM был применён для следующих задач:

1. **Определение моделей данных:**

   - В файле `internal/models/models.go` я определил структуры `Event`, `Bet` и `User`, которые представляют таблицы в базе данных. Каждая структура включает `gorm.Model` для автоматического добавления полей `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`.
     ```go
     type Event struct {
         gorm.Model
         ID       uint    `json:"id"`
         Name     string  `json:"name"`
         Date     string  `json:"date"`
         OddsWin1 float32 `json:"odds_win1"`
         OddsDraw float32 `json:"odds_draw"`
         OddsWin2 float32 `json:"odds_win2"`
     }
     ```
   - Благодаря GORM, мне не пришлось вручную писать SQL-запросы для создания таблиц — GORM автоматически генерирует их на основе структур.

2. **Миграция базы данных:**

   - В `main.go` я использовал метод `db.AutoMigrate` для автоматического создания и обновления таблиц в базе данных:
     ```go
     err = db.AutoMigrate(&models.Event{}, &models.Bet{}, &models.User{})
     if err != nil {
         log.Fatal("Error on migrating to the DB", err)
     }
     ```
   - Это позволило мне быстро настроить схему базы данных и сосредоточиться на бизнес-логике.

3. **Работа с данными:**

   - В репозиториях (`internal/repository/bet_repo.go` и `event_repo.go`) я использовал GORM для выполнения операций с базой данных:
     - Создание записей: `r.db.Create(bet)`.
     - Получение записей: `r.db.Where("user_id = ?", userID).Find(&bets)`.
     - GORM автоматически преобразует результаты запросов в структуры Go, что упростило работу с данными.
   - Например, в `bet_repo.go`:

     ```go
     func (r *BetRepository) Create(bet *models.Bet) error {
         return r.db.Create(bet).Error
     }

     func (r *BetRepository) GetBetsByUserID(userID int) ([]models.Bet, error) {
         var bets []models.Bet
         err := r.db.Where("user_id = ?", userID).Find(&bets).Error
         return bets, err
     }
     ```

4. **Результаты использования GORM:**
   - GORM позволил мне быстро настроить взаимодействие с PostgreSQL без написания сложных SQL-запросов.
   - Автоматическая миграция сэкономила время на этапе разработки.
   - Удобная работа с данными через структуры Go сделала код более читаемым и поддерживаемым.

### Gin: Создание API

Для создания REST API я использовал **Gin** — высокопроизводительный веб-фреймворк для Go. Gin был применён для следующих задач:

1. **Настройка роутера:**

   - В `main.go` я создал экземпляр роутера Gin с помощью `gin.Default()`:
     ```go
     r := gin.Default()
     ```
   - Gin автоматически добавляет middleware для логирования и обработки ошибок, что упростило отладку.

2. **Настройка CORS:**

   - Я использовал `github.com/gin-contrib/cors` для настройки CORS, чтобы фронтенд на `http://localhost:3000` мог отправлять запросы к backend на `http://localhost:8080`:
     ```go
     r.Use(cors.New(cors.Config{
         AllowOrigins:     []string{"http://localhost:3000"},
         AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
         AllowHeaders:     []string{"Content-Type"},
         AllowCredentials: true,
     }))
     ```
   - Это решило проблему с CORS и позволило фронтенду взаимодействовать с API.

3. **Определение маршрутов:**

   - В `routes/routes.go` я настроил маршруты API с помощью Gin:
     ```go
     api := r.Group("/api/v1")
     {
         api.GET("/events", eventHandler.GetEvents)
         api.POST("/events", eventHandler.CreateEvent)
         api.POST("/bets", betHandler.CreateBet)
         api.GET("/bets/user/:user_id", betHandler.GetBetsByUser)
     }
     ```
   - Gin позволил мне легко организовать маршруты в группу `/api/v1`, что сделало API более структурированным.

4. **Обработка запросов:**

   - В хендлерах (`internal/delivery/bet_handler.go` и `event_handler.go`) я использовал Gin для обработки входящих запросов и отправки ответов:
     - Получение данных из JSON: `c.ShouldBindJSON(&req)`.
     - Отправка ответов: `c.JSON(http.StatusCreated, newBet)`.
   - Например, в `bet_handler.go`:

     ```go
     func (h *BetHandler) CreateBet(c *gin.Context) {
         type BetRequest struct {
             UserID  int     `json:"user_id"`
             EventID int     `json:"event_id"`
             Amount  float32 `json:"amount"`
             Outcome string  `json:"outcome"`
         }

         var req BetRequest
         if err := c.ShouldBindJSON(&req); err != nil {
             c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
             return
         }

         newBet, err := h.service.Create(req.UserID, req.EventID, req.Amount, req.Outcome)
         if err != nil {
             c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bet"})
             return
         }

         c.JSON(http.StatusCreated, newBet)
     }

     ```

5. **Логирование запросов:**

   - Я добавил middleware для логирования всех входящих запросов:
     ```go
     r.Use(func(c *gin.Context) {
         log.Printf("Входящий запрос: %s %s", c.Request.Method, c.Request.URL.Path)
         log.Printf("Заголовки: %v", c.Request.Header)
         c.Next()
     })
     ```
   - Это помогло мне отладить проблемы, такие как дублирование ставок, и понять, какие данные приходят от фронтенда.

6. **Результаты использования Gin:**
   - Gin позволил мне быстро создать REST API с минимальным количеством кода.
   - Встроенные middleware (логирование, CORS) упростили настройку и отладку.
   - Высокая производительность Gin обеспечила быструю обработку запросов, что важно для букмекерского сайта.

### Итог

Благодаря GORM и Gin я смог:

- Быстро настроить взаимодействие с базой данных и API.
- Создать надёжный backend, который обрабатывает запросы от фронтенда и хранит данные в PostgreSQL.
- Сфокусироваться на бизнес-логике, а не на низкоуровневых деталях работы с базой данных и HTTP.

## Технологии

### Backend:

- **Go** — язык программирования для серверной части.
- **Gin** — веб-фреймворк для создания API.
- **GORM** — ORM для работы с базой данных.
- **PostgreSQL** — реляционная база данных.

### Frontend:

- **React** — библиотека для создания интерфейса.
- **Axios** — библиотека для HTTP-запросов.
- **CSS** — стили для оформления.

## Установка

### Требования:

- Go (версия 1.16 или выше).
- Node.js и npm.
- PostgreSQL (версия 13 или выше).
- pgAdmin.

### Шаги установки:

1. **Клонируйте репозиторий:**

   ```bash
   git clone <URL_репозитория>
   cd betting-site
   ```

2. **Настройте базу данных:**

   - В pgAdmin создайте базу данных `betting_site`.
   - Убедитесь, что пользователь `postgres` имеет пароль `9801042` (или измените пароль в `main.go`).

3. **Установите зависимости для backend:**

   ```bash
   cd betting-site
   go mod tidy
   ```

4. **Установите зависимости для frontend:**

   ```bash
   cd frontend
   npm install
   npm install axios
   ```

5. **Запустите backend:**

   ```bash
   cd ..
   go run main.go
   ```

6. **Запустите frontend:**
   ```bash
   cd frontend
   npm start
   ```

## Использование

1. **Просмотр событий:**

   - В разделе "События" отображаются все доступные события.

2. **Добавление события:**

   - В разделе "Добавить событие" заполните форму и нажмите "Добавить событие".

3. **Создание ставки:**

   - В разделе "Сделать ставку" укажите ID пользователя, ID события, сумму и исход, затем нажмите "Сделать ставку".

4. **Просмотр ставок:**
   - В разделе "Ставки пользователя" введите ID пользователя и нажмите "Показать ставки".

## Структура проекта

```
betting-site/
├── frontend/                 # Frontend (React)
│   ├── src/
│   │   ├── components/
│   │   │   ├── EventList.js
│   │   │   ├── CreateEvent.js
│   │   │   ├── CreateBet.js
│   │   │   ├── BetList.js
│   │   ├── App.js
│   │   ├── App.css
│   ├── package.json
├── internal/                 # Backend (Go)
│   ├── delivery/
│   │   ├── bet_handler.go
│   │   ├── event_handler.go
│   ├── models/
│   │   ├── models.go
│   ├── repository/
│   │   ├── bet_repo.go
│   │   ├── event_repo.go
│   ├── services/
│   │   ├── bet_service.go
│   │   ├── event_service.go
├── routes/
│   ├── routes.go
├── main.go
├── go.mod
```
