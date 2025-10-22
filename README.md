# Microservices ⚙️

Go Microservices Platform — это распределённая система из нескольких микросервисов на **Golang (Gin)**.  
Центральным элементом выступает **Gateway**, через который пользователь взаимодействует со всеми сервисами — от авторизации до событий.

The Go Microservices Platform is a distributed system of several microservices built on Golang (Gin).
The central element is the Gateway, through which the user interacts with all services, from authorization to events.

---

## Features / Функционал

| Component / Компонент | Description (EN) | Описание (RU) |
| ---------------------- | ---------------- | -------------- |
| **Gateway** | Unified API entry point (routes all client requests) | Единая точка входа — маршрутизирует все запросы пользователя |
| **Auth Service** | Handles registration, login, and JWT | Отвечает за регистрацию, вход и проверку JWT |
| **Event Service** | Manages user events | Управляет событиями пользователей |
| **PostgreSQL** | Central database for all services | Центральная база данных PostgreSQL |
| **Docker Compose** | Multi-container orchestration | Оркестрация микросервисов через Docker Compose |

---

## Key Features / Основные особенности

* Unified **API Gateway** — single entry point for all clients  
  Единый **API Gateway** — центральная точка входа для всех пользователей  
* RESTful API via **Gin**  
  REST API реализован с помощью **Gin**  
* Secure **JWT authentication**  
  Безопасная авторизация через **JWT**  
* **PostgreSQL** as persistent storage  
  Хранилище данных — **PostgreSQL**  
* **Docker Compose** for local deployment  
  Простое развёртывание через **Docker Compose**

---

## Tech Stack / Технологии

* **Golang 1.25+** (основной язык)  
* **Gin** — HTTP framework  
* **PostgreSQL** — база данных  
* **AES256** — безопасность
* **Docker / Docker Compose** — деплой и окружение  
* **Makefile** *(опционально)* — автоматизация сборки  

---

## Project Structure / Структура проекта

```

services/
├── gateway/
│   ├── cmd/server.go
│   ├── internal/transport/handler.go  # Основная маршрутизация
│
├── auth-service/
│   ├── cmd/server.go
│   ├── internal/repository/           # Работа с PostgreSQL
│   ├── internal/transport/            # Обработчики
│
├── event-service/
│   ├── cmd/server.go                  
│   ├── internal/repository/           # Работа с PostgreSQL
│   ├── internal/transport/            # Обработчики
│
└── docker-compose.yaml

````

---

## Gateway Routes / Роуты Gateway (сердце проекта)

| Method | Path | Description (EN) | Описание (RU) |
|--------|------|------------------|----------------|
| `GET` | `/auth/ping` | Check auth service health | Проверка доступности сервиса авторизации |
| `GET` | `/event/ping` | Check event service health | Проверка доступности сервиса событий |
| `POST` | `/auth/register` | Register new user | Регистрация нового пользователя и получение token |
| `POST` | `/auth/login` | Login user and get JWT | Вход пользователя и получение token |
| `GET` | `/auth/me` | Get account info (JWT required) | Получить данные профиля по token |
| `POST` | `/events/register` | Register new event | Создание нового события |
| `DELETE` | `/events/:id` | Delete event by ID | Удаление события по ID |
| `GET` | `/events/:id` | Get event by ID | Получение события по ID |
| `GET` | `/events` | Get all events | Получить список всех событий |

---

## Quick Start / Быстрый старт

1. Clone the repository / Клонируем репозиторий:
```bash
git clone https://github.com/gox7/microservices.git
cd services
````

2. Run with Docker / Запуск через Docker:

```bash
docker-compose up --build
```

3. Services will be available at / Сервисы будут доступны по адресам:

```
Gateway:       http://localhost:8080
PostgreSQL:    localhost:5432
```

---

## License / Лицензия

MIT License — свободное использование и модификация проекта.

Хочешь, я сразу создам этот `README.md` как файл в `/services_extracted/services/README.md`, чтобы ты мог скачать его и положить в репозиторий?
```
