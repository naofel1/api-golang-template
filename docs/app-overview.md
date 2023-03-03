# ✨ Application Overview

The application is a School Manager. Its primary goal is to manage student. It achieves this goal by providing a set of features and functionalities that enable the users to get students, create new one, and give an authentication layer.

## 🌸 Features

The following are the key features of the application:

- Authentication endpoint

Feel free to add new features 🦊

**Technical features:**

- Authentication Middleware
- Token management with JWT as provider
- Pagination system
- Usefull client like Mailgun client (mailer), Discord (webhook), S3 Bucket (File management)

## ⛰️ Get started

To install the application, you need to follow the following steps:

**Prerequisites:**

1. Golang 1.18+
2. Docker

To set up the app execute the following commands.

```sh
git clone https://github.com/naofel1/api-golang-template.git
cd api-golang-template
## Remplace the .env value with your credentials
cp .env.example .env
## To generate new certificate for token creation and validation
make certs
make local.docker
```
