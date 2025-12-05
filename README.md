****Users API****

Простой REST API на Go с использованием Gin и PostgreSQL.
Проект демонстрирует базовые CRUD-операции и работу со слоями handler/service/storage.

****Технологии****

Go, Gin, PostgreSQL, sqlx

****Запуск****

Поднять PostgreSQL (локально или через Docker):

docker run --name users-db -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres


Создать базу:

CREATE DATABASE usersdb;

Создать таблицу:

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT,
    age INT
);


Запустить приложение:

go run ./cmd

****Эндпоинты****

POST   /users

GET    /users

GET    /users/:id

PATCH  /users/:id

DELETE /users/:id
