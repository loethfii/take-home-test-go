# Take Home TEST

---

## Cara penggunaan

- Download repository ini atau clone
- Buka CMD / terminal pastikan path berada di `.../take-home-test-go-main`
- Ketikan perintah `docker-compose up --build`

\*jika terdapat error pada build bisa ketikan `docker-compose down` dan ulangi lagi perintah di atas

---

## Testing API

Setelah program berjalan pengecekan api bisa dilakukan dengan postman atau menggunakan swagger

### Postman

-- Import file `api_collection.json` pada postman lalu

### Swagger

-- Buka browser lalu masuk ke host `http://localhost:3000/swagger`

---

## Teknologi

- Go
- Fiber
- Docker
- Swagger
- Postgresql
- Validation GO
- Jwt Authentication

---

## Fitur dan Service

### Employee

- Fetch Employee [GET]
- Detail Employee [GET]
- Add Employee [POST]

### User

- Check Authentication [GET]
- Login [POST]
- Register [POST]

### Api Gateway

- Fetch Employee [GET]
- Add Employee [POST]
- Detail Employee[GET]
- Register [POST]
- Login [POST]

---

## API Spek

### Service User

- Register User
- Login User
- Check Authentication

#### Register User 👤

- URL : `http://localhost:3001/api/v1/user/register`
- Method : POST
- Required Token : No
- Request Body :

```sh
{
    "email" : "john@mail.com",
    "password" : "john333"
}
```

- Response Body :

```sh
{
    "code": 201,
    "status": "created,
    "data": "john@mail.com"
}
```

#### Login User 👤

- URL : `http://localhost:3001/api/v1/user/login`
- Method : POST
- Required Token : No
- Request Body :

```sh
{
    "email" : "john@mail.com",
    "password" : "john333"
}
```

- Response Body :

```sh
{
   "code": 200,
    "status": "ok",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5AbWFpbC5jb20ifQ._o4lsYNgHiICJJfQvKfZ6ScUD6dIPidrk-oFq6hMYwc",
    "data": {
        "id": 25,
        "email": "john@mail.com",
        "password": "$2a$10$pCNzaNRa5g.j5Zdi5x2freufWwT.VgnKmGAkPKBxNJz1L33WdEcWq"
    }
}
```

#### Check Authentication 👤

- URL : `http://localhost:3001`
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "auth status": "Authentication OK",
    "message": "Hello from user service",
    "who_login": "john@mail.com"
}
```

### Service Employee

- Api list
- Fetch employee
- Detail employee
- Add employee

#### Api List 👷

- URL : `http://localhost:3002`
- Method : GET
- Required Token : No
- Request Body : No
- Response Body :

```sh
{
   "List API": {
        "GET http://localhost:3002/api/v1/employees": "Fetch Employee",
        "GET http://localhost:3002/api/v1/employees/:id": "Detail Employee",
        "POST http://localhost:3002/api/v1/employees": "Add Employee"
    }
}
```

#### Api List 👷

- URL : `http://localhost:3002`
- Method : GET
- Required Token : No
- Request Body : No
- Response Body :

```sh
{
   "List API": {
        "GET http://localhost:3002/api/v1/employees": "Fetch Employee",
        "GET http://localhost:3002/api/v1/employees/:id": "Detail Employee",
        "POST http://localhost:3002/api/v1/employees": "Add Employee"
    }
}
```

#### Fetch Employee 👷

- URL : `http://localhost:3002/api/v1/employees`
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "code": 200,
    "status": "OK",
    "data": [
        {
            "id": 11,
            "name": "ajam"
        },
        {
            "id": 10,
            "name": "sifa"
        }
    ]
}
```

#### Detail Employee 👷

- URL : `http://localhost:3002/api/v1/employees/:id`
- Param : id
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
  "code": 200,
        "status": "OK",
        "data": {
            "id": 10,
            "name": "sifa"
        }
}
```

#### Add Employee 👷

- URL : `http://localhost:3002/api/v1/employees`
- Method : POST
- Required Token : Yes
- Request Body :

```sh
    {
        "name" : "sifa"
    }
```

- Response Body :

```sh
{
  "code": 200,
    "status": "OK",
    "data": {
        "id": 10,
        "name": "sifa"
    }
}
```

### Api Gateway

- Fetch Employee
- Add Employee
- Detail Employee
- Register User
- Login User

#### Fetch Employee ⚡

- URL : `http://localhost:3000/api/v1/employees`
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
   "code": 200,
    "status": "OK",
    "data": [
        {
            "id": 11,
            "name": "ajam"
        },
        {
            "id": 10,
            "name": "sifa"
        }
    ]
}
```

#### Add Employee ⚡

- URL : `http://localhost:3000/api/v1/employees`
- Method : POST
- Required Token : Yes
- Request Body :

```sh
    {
        "name" : "sifa"
    }
```

- Response Body :

```sh
{
  "code": 200,
    "status": "OK",
    "data": {
        "id": 10,
        "name": "sifa"
    }
}
```

#### Detail Employee ⚡

- URL : `http://localhost:3000/api/v1/employees/:id`
- Param : id
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
  "code": 200,
    "status": "OK",
    "data": {
        "id": 10,
        "name": "sifa"
    }
}
```

#### Register User ⚡

- URL : `http://localhost:3000/api/v1/register`
- Method : POST
- Required Token : No
- Request Body :

```sh
    {
        "email": "john@mail.com",
        "password": "john333"
    }

```

- Response Body :

```sh
{
    "code": 201,
    "data": "john@gmail.com",
    "status": "created"
}
```

#### Login User ⚡

- URL : `http://localhost:3000/api/v1/login`
- Method : POST
- Required Token : No
- Request Body :

```sh
    {
        "email": "john@mail.com",
        "password": "john333"
    }

```

- Response Body :

```sh
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5AbWFpbC5jb20ifQ._o4lsYNgHiICJJfQvKfZ6ScUD6dIPidrk-oFq6hMYwc",
    "code": 200,
    "data": {
        "email": "john@mail.com",
        "id": 25,
        "password": "$2a$10$pCNzaNRa5g.j5Zdi5x2freufWwT.VgnKmGAkPKBxNJz1L33WdEcWq"
    },
    "status": "ok"
}
```
