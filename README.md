# API - Golang

API desarrollada en golang de un CRUD básico.

## Necesidad que cubre

El objetivo del CRUD es tener una agenda de contactos telefónicos, para lo cual todos tienen nombre y telefono.

## Consideraciones técnicas

* Colección adjuntada para importar en POSTMAN.
* API ejecutandose en puerto ```3000```
* Base de datos ejecutandose en memoria
* Aplicación Dockerizada con Alpine

## Prerequisitos

* Tener instalado docker

## Instalación

1. Importar colección en POSTMAN
2. Realizar clone de este proyecto
3. Ubicarse por consola en la carpeta clonada
4. Ejecutar: ```docker build . -t alias``` donde el alias puede ser cualquier nombre para la imagen
5. Ejecutar: ```docker run -p 3000:3000 alias``` donde el alias puede ser cualquier nombre para la imagen

Ya está todo listo para usar la colección importada.

## License
[MIT](https://choosealicense.com/licenses/mit/)