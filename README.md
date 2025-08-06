# Go User API

Una API RESTful construida con Go, Gin Framework y GORM para gestión de usuarios y autenticación.

## 🚀 Tecnologías

- [Go](https://golang.org/) - Lenguaje de programación
- [Gin](https://gin-gonic.com/) - Framework web
- [GORM](https://gorm.io/) - ORM para Go
- [JWT](https://github.com/golang-jwt/jwt) - JSON Web Tokens para autenticación
- [SQLite](https://www.sqlite.org/) - Base de datos

## 📋 Requisitos Previos

- Go 1.21 o superior
- SQLite

## 🔧 Instalación

1. Clona el repositorio
```bash
git clone https://github.com/lumiere11/go-user-api.git
cd go-user-api
```

2. Instala las dependencias
```bash
go mod download
```

3. Configura las variables de entorno
```bash
export JWT_SECRET_KEY=tu_clave_secreta
```

4. Ejecuta el servidor
```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

## 🛣️ Endpoints

### Autenticación

#### POST /register
Registra un nuevo usuario
```json
{
    "name": "Usuario",
    "email": "usuario@ejemplo.com",
    "password": "contraseña123"
}
```

#### POST /login
Inicia sesión y obtiene un token JWT
```json
{
    "email": "usuario@ejemplo.com",
    "password": "contraseña123"
}
```

#### GET /profile
Obtiene el perfil del usuario autenticado
> Requiere token JWT en el header: `Authorization: Bearer <token>`

### Libros

#### GET /books
Obtiene la lista de libros
> Requiere autenticación

#### POST /books
Crea un nuevo libro
```json
{
    "title": "Título del Libro",
    "author": "Autor del Libro"
}
```
> Requiere autenticación

## 🔒 Autenticación

La API utiliza JWT (JSON Web Tokens) para la autenticación. Para acceder a rutas protegidas:

1. Obtén un token mediante el endpoint de login
2. Incluye el token en el header de tus peticiones:
```
Authorization: Bearer <tu_token>
```

## 🗄️ Estructura del Proyecto

```
go-user-api/
├── config/         # Configuración de la base de datos
├── controllers/    # Controladores de la API
├── middlewares/   # Middlewares (auth, etc.)
├── models/        # Modelos de datos
├── routes/        # Definición de rutas
└── main.go        # Punto de entrada
```

## 🚀 Despliegue

El proyecto está configurado para ser desplegado en Render.com usando GitHub Actions.

### Requisitos para el Despliegue
1. Cuenta en Render.com
2. Configurar los siguientes secretos en GitHub:
   - `RENDER_API_KEY`
   - `RENDER_SERVICE_ID`

## 📝 Variables de Entorno

- `JWT_SECRET_KEY`: Clave secreta para firmar los tokens JWT
- `PORT`: Puerto en el que se ejecutará el servidor (por defecto: 8080)

## 🛠️ Desarrollo

Para ejecutar en modo desarrollo:
```bash
go run main.go
```

Para ejecutar los tests:
```bash
go test ./...
```

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - mira el archivo [LICENSE](LICENSE) para detalles

## ✨ Contribuir

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request
