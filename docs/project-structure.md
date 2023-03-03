# 🗄️ Project Structure

Most of the code lives in the `internal` folder and looks like this:

```sh
api # Documentation of the API.
|
cmd # Main entry points for the application.
|
configs # Configuration files for the application.
|
deploy # Scripts and configuration files for deploying the application.
|
docs # Additional documentation for the application.
|
internal # Internal implementation of the application.
|
+--+-- client # implementation of used client API.
|  |-- configs # Configuration for the API.
|  |-- ent # Entity definitions and database schema.
|  |-- handler # Implementation of HTTP handlers.
|  |-- primitive # Implementation of domain models.
|  |-- service # Implementation of business logic.
|  |-- storage # Implementation of data access layer.
|  |-- tracing # Implementation of distributed tracing.
|  |-- utils # Utility functions used by other components.
|
pkg # Packages that can be used by other applications.
```
