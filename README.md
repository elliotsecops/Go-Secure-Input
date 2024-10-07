**Go-Secure-Input**

Un servidor web en Go que valida y sanea las entradas de usuario para prevenir ataques de inyección de código y Cross-Site Scripting (XSS).

**Descripción**

Este proyecto es un servidor web en Go que se encarga de validar y sanear las entradas de usuario para prevenir ataques de inyección de código y Cross-Site Scripting (XSS). El servidor utiliza una expresión regular para validar la entrada del usuario y elimina cualquier espacio en blanco al principio y al final de la cadena.

**Características**

* Validación de entradas de usuario utilizando una expresión regular
* Saneamiento de entradas de usuario para eliminar espacios en blanco
* Protección contra ataques de inyección de código y Cross-Site Scripting (XSS)
* Servidor web en Go para recibir y procesar solicitudes HTTP

**Uso**

Para utilizar este proyecto, simplemente clona el repositorio y ejecuta el comando `go run main.go` para iniciar el servidor. Luego, puedes utilizar el comando `curl` para enviar solicitudes HTTP al servidor y probar su funcionalidad.

**Ejemplo de uso**

```bash
curl -X POST -d "user_input=HolaMundo" http://localhost:8080/input
```

**Licencia**

Este proyecto se distribuye bajo la licencia MIT.

**Autor**

* Elliot SecOps (elliotsecops)

---

**Go-Secure-Input**

A Go web server that validates and sanitizes user input to prevent code injection and Cross-Site Scripting (XSS) attacks.

**Description**

This project is a Go web server that validates and sanitizes user input to prevent code injection and Cross-Site Scripting (XSS) attacks. The server uses a regular expression to validate user input and removes any whitespace from the beginning and end of the string.

**Features**

* User input validation using a regular expression
* Input sanitization to remove whitespace
* Protection against code injection and Cross-Site Scripting (XSS) attacks
* Go web server to receive and process HTTP requests

**Usage**

To use this project, simply clone the repository and run the command `go run main.go` to start the server. Then, you can use the `curl` command to send HTTP requests to the server and test its functionality.

**Example Usage**

```bash
curl -X POST -d "user_input=HelloWorld" http://localhost:8080/input
```

**License**

This project is distributed under the MIT License.

**Author**

* Elliot SecOps (elliotsecops)
