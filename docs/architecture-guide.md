# 🎃 Architecture guide

## What is Hexagonal Architecture?

Hexagonal architecture in Golang is a software development method that wraps business components in an abstraction layer. It allows developers to build applications that can be easily maintained and evolved.

### Handlers

**Handlers** are functions that interpret HTTP requests and manage request parameters, data validation, and data transformation into output.

### Services

**Services** are classes that contain business logic and interact with **repositories** to retrieve and store data.

### Repositories

**Repositories** are classes that interact with data storage and can be used to retrieve, modify, add, and delete data. These can be implemented with relational databases, NoSQL databases, or JSON files.
