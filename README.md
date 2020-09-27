# GoFiber ToDo

[GoFiber/v2](https://github.com/gofiber/fiber) based backend for a ToDo Application (No Frontend)

## Contents

- [Installation](#installation)
- [Requirements](#requirements)
- [Usage](#usage)
- [Rules](#rules)
- [Dependencies](#dependencies)

## Installation

```sh
$ git clone https://github.com/adnanbrq/fiber-todo.git
```

## Requirements
- Postgres Database (or change the Code to use a dif. db)

## Usage

```sh
cd {path} && go run .
```

## Testing
- You can test the API with Paw (similar to Postman) by using fiber-todo.paw

## Dependencies

- [github.com/gofiber/fiber/v2 - v2.0.2](github.com/gofiber/fiber/v2)
Server Engine
- [gorm.io/gorm v1.20.1](gorm.io/gorm)
DB ORM
- [gorm.io/driver/postgres v1.0.0](gorm.io/driver/postgres)
Postgres Driver
- [github.com/adnanbrq/validation - v1.1.1](https://github.com/adnanbrq/validation)
Validation
- [github.com/stretchr/testify - v1.6.1](https://github.com/stretchr/testify)
Assertions