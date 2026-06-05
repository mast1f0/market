# Market

Монорепозиторий интернет-магазина: REST API, микросервисы аутентификации и работы с изображениями, веб-интерфейс на React.

## Состав проекта

| Сервис | Порт | Описание |
|--------|------|----------|
| **market** | 8080 | Каталог, корзина, заказы |
| **auth-service** | 8081 | Регистрация, вход, профиль |
| **images-service** | 8082 | Загрузка и выдача файлов через MinIO |
| **nginx** | 80 | Единая точка входа: фронтенд + проксирование API |
| **postgres** | 5433 | БД магазина |
| **auth-postgres** | — | БД аутентификации |
| **minio** | 9000 / 9001 | S3-хранилище и консоль |

## Возможности

- Каталог товаров и категорий
- Корзина и оформление заказов
- JWT-аутентификация с ролями: `buyer`, `seller`, `admin`
- Загрузка изображений товаров в MinIO
- Веб-интерфейс (React + Vite + Tailwind CSS)

## Стек

| Слой | Технологии |
|------|------------|
| Backend | Go 1.25, Chi, GORM, slog |
| Auth | Go, Chi, JWT |
| Images | Go, MinIO SDK |
| Frontend | React 18, TypeScript, Vite, Tailwind CSS |
| Инфраструктура | PostgreSQL 16, MinIO, Docker, Nginx |

---

## Быстрый старт

### 1. Клонирование

Репозиторий использует git submodules — клонируйте с флагом `--recurse-submodules`:

```bash
git clone --recurse-submodules https://github.com/mast1f0/market.git
cd market
```

Если репозиторий уже склонирован без submodules:

```bash
git submodule update --init --recursive
```

### 2. Запуск через Docker

```bash
docker compose up -d --build
```

После старта доступны:

| Адрес | Назначение |
|-------|------------|
| http://localhost | Веб-интерфейс |
| http://localhost/api/ | API магазина |
| http://localhost/auth/ | API аутентификации |
| http://localhost/minio/ | API изображений |
| http://localhost:9001 | Консоль MinIO |

Тестовые учётные записи (создаются при старте auth-service):

| Логин | Пароль | Роль |
|-------|--------|------|
| admin | admin123 | admin |
| seller | seller123 | seller |
| buyer | buyer123 | buyer |

### 3. Локальный запуск API

Для разработки без Docker поднимите только PostgreSQL:

```bash
docker compose up -d postgres
```

Создайте `.env` в корне проекта:

```env
DB_HOST=localhost
DB_PORT=5433
DB_USER=store_usr
DB_PASSWORD=password132
DB_NAME=store_db
JWT_SECRET=superSecret
```

Примените миграции и запустите сервис:

```bash
# миграции (нужен migrate CLI)
migrate -path migrations -database "postgres://store_usr:password132@localhost:5433/store_db?sslmode=disable" up

# запуск с начальным наполнением БД
go run ./cmd/srv -seed
```

---

## API

Все эндпоинты ниже указаны с префиксом nginx. Для защищённых маршрутов передавайте заголовок `Authorization: Bearer <token>`.

### Аутентификация — `/auth/`

| Метод | URL | Доступ |
|-------|-----|--------|
| POST | `/auth/register` | Публичный |
| POST | `/auth/login` | Публичный |
| GET | `/auth/profile` | JWT |
| GET | `/auth/users` | JWT |
| PUT | `/auth/users/{id}` | JWT |
| DELETE | `/auth/users/{id}` | JWT |

### Каталог — `/api/`

| Метод | URL | Доступ |
|-------|-----|--------|
| GET | `/api/products` | Публичный |
| GET | `/api/products/{id}` | Публичный |
| GET | `/api/categories` | Публичный |
| GET | `/api/categories/{id}` | Публичный |
| POST | `/api/products` | JWT · seller, admin |
| PUT | `/api/products/{id}` | JWT · seller, admin |
| DELETE | `/api/products/{id}` | JWT · seller, admin |
| POST | `/api/categories` | JWT · seller, admin |
| PUT | `/api/categories/{id}` | JWT · seller, admin |
| DELETE | `/api/categories/{id}` | JWT · seller, admin |

### Корзина и заказы — `/api/`

| Метод | URL | Доступ |
|-------|-----|--------|
| GET | `/api/cart` | JWT · buyer, seller, admin |
| POST | `/api/cart/items` | JWT · buyer, seller, admin |
| PUT | `/api/cart/items` | JWT · buyer, seller, admin |
| DELETE | `/api/cart/items` | JWT · buyer, seller, admin |
| GET | `/api/orders` | JWT · buyer, seller, admin |
| GET | `/api/orders/{id}` | JWT · buyer, seller, admin |
| POST | `/api/orders` | JWT · buyer, seller, admin |
| PUT | `/api/orders/{id}` | JWT · seller, admin |

### Изображения — `/minio/`

| Метод | URL | Доступ |
|-------|-----|--------|
| POST | `/minio/upload` | Публичный |
| GET | `/minio/view/{id}` | Публичный |
| GET | `/minio/file/{id}` | Публичный |
| DELETE | `/minio/file/{id}` | Публичный |
| GET | `/minio/files` | Публичный |
| POST | `/minio/files` | Публичный |
| DELETE | `/minio/files` | Публичный |

---

## Структура проекта

```text
market/
├── cmd/srv/                  # Точка входа API магазина
├── internal/
│   ├── adapters/
│   │   ├── http/             # HTTP-обработчики, роутер, middleware
│   │   ├── jwt/              # JWT-утилиты
│   │   └── storage/postgres/ # Репозитории
│   ├── core/
│   │   ├── domain/           # Доменные модели
│   │   ├── ports/            # Интерфейсы репозиториев
│   │   └── service/          # Бизнес-логика
│   └── engine/
│       ├── config/           # Загрузка конфигурации
│       ├── logger/           # Логирование (slog)
│       └── seed/             # Начальные данные
├── migrations/               # SQL-миграции магазина
├── nginx/                    # Конфигурация reverse proxy
├── services/
│   ├── auth/                 # Микросервис аутентификации (submodule)
│   └── images/               # Микросервис изображений (submodule)
├── web/market-frontend/      # React SPA
└── docker-compose.yml
```

---

## Тестирование

```bash
go test ./...
```

## Лицензия

Проект в разработке.
