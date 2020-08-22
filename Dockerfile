# Imagen base para trabajar
FROM golang:alpine

# Instalación de git para paquetes
RUN apk add git

# Creación de directorio en contenedor
ADD . /go/src/myapp

# Directorio principal
WORKDIR /go/src/myapp

# Obtenemos librerias
RUN go get myapp

# Instalacion de librerias
RUN go install

# Punto de entrada
ENTRYPOINT ["/go/bin/myapp"]
