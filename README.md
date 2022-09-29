<div align="center">
   <a href="https://go.dev/"><img src = https://img.shields.io/badge/GO-v1.17.6-blue></a>
   <a href="https://echo.labstack.com/"><img src = https://img.shields.io/badge/Echo-v4.7.0-blue></a>
   <a href="https://hub.docker.com/"><img src = https://img.shields.io/badge/Deploy-%20Docker-blue></a>
</div>

<div align="center">

  <h3 align="center">Rogerdev Titanic Test Backend</h3>
  <p align="center">
   Aplikasi Backend ini dibuat untuk memenuhi golang test dari CV Kabayan Insan Informatika
  </p>
</div>
<br />

## Endpoint
<!-- ENDPOINT -->
Person
   - POST http://localhost:8081/person/post            untuk filter data person format json (backend-titanic-test.json)
   - GET http://localhost:8081/person/post/histories      untuk melihat history filter data person
   
User
   - POST http://localhost:8081/users/register            untuk daftar account data yang dikirimkan format json {name,email,password,address,gender}
   - GET http://localhost:8081/users/me                   untuk melihat detail profil (perlu untuk memasukan token yang di dapat dari login pada Header Authorization Barer Token)
   - PUT http://localhost:8081/users/me            untuk update account data yang dikirimkan format json {name,email,password,address,gender} perlu token


Auth
   - POST http://localhost:8081/users/login            untuk login ke aplikasi dan mendapatkan token access data format json {email,password}
   - POST http://localhost:8081/google/login            untuk login ke aplikasi melalui acount google dan mendapatkan token access
   - POST http://localhost:8081/users/logout            untuk keluar dari aplikasi

<!-- TABLE OF CONTENTS -->
## Table of Contents
1. [About the Project](#about-the-project)
2. [Feature](#feture)
3. [Tech Stack](#tech-stack)
7. [Authors](#authors)

<!-- ABOUT THE PROJECT -->
## About The Project
   Aplikasi Backend ini dibuat untuk memenuhi golang test dari CV Kabayan Insan Informatika
   Aplikasi ini dapat di clone langsung dari github atau dapat juga di pull imagenya dari dockerhub


<p align="right">(<a href="#top">back to top</a>)</p>

## Feature
-  Login and Logout
-  Create account registration

As visitor
-  filter person data
-  View history filter data

<p align="right">(<a href="#top">back to top</a>)</p>

## Tech Stack
### Framework
- [Echo (Go Web Framework)](https://echo.labstack.com/)

### Build With
- [Golang (Language)](https://go.dev/) 
- [Testify (Unit Test)](https://github.com/stretchr/testify)

### Deployment
- [Docker (Container - image)](https://hub.docker.com/)

<p align="right">(<a href="#top">back to top</a>)</p>

## Structure
``` bash
Rogerdev Titanic Test Backend
  ├── configs                
  │     └──config.go                  # Configs files
  ├── delivery                        # Endpoints handlers or controllers
  │     ├──controllers
  │     │   └── user
  │     │     ├── formatter.go        # Default response format for spesific controllers
  │     │     └── user.go            # Spesific controller
  │     │   ├── auth
  │     │     ├── formatter.go        # Default response format for spesific controllers
  │     │     └── auth.go            # Spesific controller 
  │     │   └── person
  │     │     ├── formatter.go        # Default response format for spesific controllers
  │     │     └── person.go            # Spesific controller
  │     ├──middlewares
  │     │   └── jwtMiddleware.go      # Middlewares Function
  │     └──routes  
  │         └── routes.go             # Endpoints list
  ├── deployment               
  │     └── app-deployment.yaml       # Deployment kubernetes
  │     └── db_secret.yaml            # secret for kubernetes implementation
  ├── entities                
  │     └── user.go                   # Database model user
  │     └── person.go                 # Database model person
  ├── repository
  │     │   └── user
  │     │   │     ├── interface.go              # Repository Interface for controllers
  │     │   │     └── user.go                  # Spesific Repository
  │     │   └── auth
  │     │   │     ├── interface.go              # Repository Interface for controllers
  │     │   │     └── auth.go                  # Spesific Repository
  │     │   └── person
  │     │   │     ├── interface.go              # Repository Interface for controllers
  │     │   │     └── person.go                  # Spesific Repository
  ├── utils                 
  │     └── mysqldriver.go            # Database driver 
  │     └── logingoogle.go            # function for handle login with google account
  ├── Dockerfile                      # command to build app
  ├── go.mod                  
  ├── go.sum                  
  ├── main.go                         # Main Program
  └── README.md    
```
<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
## Authors
* Ade Mawan - [Github](https://github.com/ademawan) · [LinkedIn](https://www.linkedin.com/in/ade-mawan-527657177/)

<p align="right">(<a href="#top">back to top</a>)</p>
