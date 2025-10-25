# CodigoFacilito-FantasyFUTLeague
### FantasyFUTLeague es una aplicación web desarrollada en Go que permite a los usuarios crear equipos de fútbol y participar en torneos, poniendo a prueba sus conocimientos de fútbol
#### Tecnologías Utilizadas:
- MySQL
- Go v1.24.5
- Gin
- GORM
- Docker
- HTML
- CSS
- JS
- Bootstrap v5.3.8


#### Recursos:

- [Documento Descriptivo del Proyecto](https://drive.google.com/file/d/1ttLDvZNhPOjZ3KeqOgL5v0HRUttHYi9B/view?usp=sharing)

- [Diagrama Entidad Relación](https://lucid.app/lucidchart/64f63ea2-13c1-47ac-a77c-148002265621/edit?viewport_loc=327%2C-3%2C1707%2C838%2C0_0&invitationId=inv_bbea07bb-8f97-4eee-b473-82346bf3e7c9)

---

#### ¿Cómo ejecutar el proyecto?

Asegurate de tener instalada la versión de Go **1.24.5 o superior** y MySQL en tu sistema

Luego corroborar tener creada una base de datos mySQL, y guardar su usuario, contraseña, host, puerto y nombre para incluirlos en el archivo .env
En el directorio base se encuentra un .env.example con los campos que deberías completar, entre ellos estos mencionados de la base de datos

Además, en el .env hay que completar la variable JWT_SECRET, que se usa para firmar los tokens de autenticación y manejar los inicios de sesión en el sitio

JWT_SECRET puede ser cualquier cadena aleatoria, por ejemplo:
JWT_SECRET=mysecretkey123  
Una buena práctica puede ser generarla online con algún generador

Clonar el repositorio, y por consola, desde el directorio base del proyecto, ejecutar los comandos:
- `go mod tidy` Para descargar dependencias
- `go build .\cmd\main.go` Para compilar el proyecto
- `go run .\cmd\main.go` Para ejecutar el proyecto

La primera ejecución generará las migraciones y creará las tablas en la base de datos

Posteriormente, se recomienda cargar la base con datos iniciales, para esto se provee un script .sql en el directorio base del proyecto, en el archivo `demo_data.sql`

Ejecutar ese script y se cargará la base de datos con:
- Usuarios (nombre de usuario, correo, contraseña, rol):
  - (user1, correo1@gmail.com, 1234, ADMIN)
  - (user2, correo2@gmail.com, 12345, USER)
  - (user3, correo3@gmail.com, 123456, USER)
  - (user4, correo4@gmail.com, 1234567, USER)
  - (user5, correo5@gmail.com, 1234678, USER)
- Cuatro equipos, pertenecientes a los 4 usuarios, todos completos con 11 jugadores cada uno
- Jugadores, algunos ya pertenecientes a los equipos de los usuarios, otros libres esperando equipo
- 3 Torneos, 2 finalizados, otro esperando equipos para comenzar
- Varios partidos, los pertenecientes a los 2 torneos finalizados cargados

Ejecución con Docker:

Para ejecutar el proyecto localmente en Docker hay que seguir los siguientes pasos:
- Asegurarse de tener cargadas las variables de entorno, para que el `docker-compose.yml` las reconozca correctamente
- Ejecutar el comando:
  - `docker compose up --build`: Para construir la imagen de la aplicación Go desde el Dockerfile, crear los contenedores de la app y de la DB, y copiar los datos iniciales de la DB que están en el script `demo_data.sql`
- Luego estos comandos serán útiles también:
  - `docker compose down`: Para detener los contenedores, manteniendo la DB guardada y los volúmenes
  - `docker compose down -v`: Para detener los contenedores, borrando además los datos de la DB y los volúmenes
  - `docker compose up`: Para volver a levantar los contenedores y poder seguir corriendo el proyecto, con los datos actualizados del último uso


**Asegurarse de utilizar las funcionalidades del proyecto de manera acorde, para evitar crear inconsistencias en el mismo**

### Para probar de forma correcta el proyecto, se recomienda leer las instrucciones de uso en el documento PDF del proyecto
---

#### Endpoints Principales:

**Rutas accesibles para todos los usuarios**

- `GET` `/`: Ruta base, vista Index  
- `GET` `POST` `/register`: Registro de nuevos usuarios  
- `GET` `POST` `/login`: Login de usuarios  
- `GET` `/logout`: Logout de usuarios  

**Rutas accesibles para usuarios autenticados en el sitio:**

- `GET` `/profile`: Perfil del usuario  
- `GET` `/teams`: Visualización de los equipos de la plataforma  
- `GET` `POST` `/teams/create`: Crear un equipo  
- `GET` `/tournaments`: Visualización de los torneos de la plataforma  
- `GET` `/tournaments/:id`: Detalle de un torneo  
- `POST` `/tournaments/:id/join`: Unirse a un torneo  
- `GET` `/players`: Visualización de los jugadores de la plataforma  
- `POST` `/players/:id/buy`: Comprar un jugador  
- `GET` `/matches`: Visualización de los partidos de la plataforma  
- `GET` `/matches/:id`: Detalle de un partido

**Rutas accesibles para usuarios administradores en el sitio:**

- `GET` `POST` `/tournaments/create`: Crear torneos  
- `POST` `/tournaments/:id/finish`: Finalizar un torneo  
- `GET` `POST` `/players/create`: Crear jugador  
- `GET` `POST` `/matches/create`: Crear un partido  
- `POST` `/matches/:id/simulate`: Simular un partido

#### Acerca De
Proyecto desarrollado por **Franco Smuraglia** como proyecto final del Bootcamp de Go Profesional de CodigoFacilito.
