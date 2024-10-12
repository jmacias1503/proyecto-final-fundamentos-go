# Proyecto Final Fundamentos de Go

<p align="center">
<img src="https://github.com/egonelbre/gophers/blob/master/.thumb/animation/gopher-dance-long-3x.gif">
</p>

Este proyecto tiene como finalidad simular un sistema de control escolar

## Contenidos

- [Requerimientos](#requerimientos)
    - [Estudiantes](#estudiantes)
    - [Materias](#materias)
    - [Calificaciones](#calificaciones)
    - [Diagrama de entidad relación](#diagrama-de-entidad-relación)
- [Stack](#stack)
- [Despliegue](#despliegue)
    - [Contenedor de Docker](#contenedor-de-docker)
    - [Proyecto](#proyecto)
- [Workflow](#workflow)
    - [Estructura de archivos](#estructura-de-archivos)
    - [Estructura de commits](#estructura-de-commits)

## Requerimientos

El proyecto tiene que consistir de un CRUD que guarde los datos en una base de datos, y que estos puedan ser accedidos por las siguientes rutas

### Estudiantes

- `GET` a `/api/students`: Obtiene todos los estudiantes
- `POST` a `/api/students`: Crea un nuevo estudiante
- `DELETE` a `/api/students/:student_id`: Elimina a un estudiante por medio del ID.
- `PUT` a `/api/students/:student_id`: Actualiza todos los datos de un estudiante.
- `GET` a `/api/students/:student_id`: Obtienes los datos de un estudiante por medio de su ID

### Materias

- `POST` a `/api/subjects`: Crear una nueva materia
- `PUT` a `/api/subjects/:subject_id`: Actualizar completamente una materia
- `GET` a `/api/subjects/:subject_id`: Obtener los datos de una materia por medio de su ID
- `DELETE` a `/api/subjects/:subject_id`: Eliminar una materia pasando su ID

### Calificaciones

- `POST` a `/api/grades`: Crear una nueva calificacion
- `PUT` a `/api/grades/:grade_id`: Actualizar una nueva calificacion pasando su ID
- `DELETE` a `/api/grades/:grade_id`: Eliminar una calificacion pasando su ID
- `GET` a `/api/grades/:grade_id/student/:student_id`: Obtener la calificacion de una materia de un estudiante
- `GET` a `/api/grades/student/:student_id`: Obetener todas las calificaciones de un sestudiante

### Diagrama de entidad relación

Para poder tener una mejor visión sobre las entidades que se van a necesitar en el CRUD, se diseñó un [diagrama de entidad relación](https://dbdiagram.io/d/erd-proyecto-final-fundamentos-go-670aadf997a66db9a3c21c56)

## Stack

- GoLang (lenguaje principal back)
- Gorm (ORM)
- Gin (Manejo de rutas)
- *Por definir* (Base de datos)
- Docker (Contenerización de la base de datos)

## Despliegue

### Contenedor de Docker

En caso de no contar con un contenedor de la base de datos, correr un contenedor de docker con la imagen de *base de datos a definir*

```bash
$ docker run --name db-proyecto-final-fundamentos-go -d dbAdefinir
```

### Proyecto

Entrando a la carpeta `cmd/`, ejecutar

```bash
$ go run main.go
```

## Workflow

Se trabajará con forks, cada quien trabajará a su respectiva comodidad, y al momento de hacer Pull Request, se enviarán a la rama `dev` hasta ser ese mismo código testeado

### Estructura de archivos

Se seguirá la siguiente estructura de archivos

```
- go.mod
- go.sum
- Dockerfile
cmd/
    - main.go
    controllers/
        - controllerName.go
    templates/
        - index.html
    config/
        - env.go
```

### Estructura de commits

Cada commit será lo más atómico posible, para poder tener mayor trazabilidad y rastrea de bugs. Al momento de hacer una Pull Request, esta se deberá de hacer respecto a un solo concepto, agrupando todos estos commits que interactuan entre estos. Tomar como referencia [este artículo sobre cómo escribir commits](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13)
