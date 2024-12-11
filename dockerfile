# Utiliza una imagen base oficial de Go
FROM golang:1.20-alpine as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum para instalar las dependencias
COPY go.mod go.sum ./

# Instala las dependencias
RUN go mod download

# Copia el código fuente del proyecto al contenedor
COPY . .

# Compila la aplicación Go
RUN go build -o server .

# Etapa final
FROM alpine:latest

# Instala las dependencias necesarias para ejecutar la aplicación Go (en este caso solo ca-certificates)
RUN apk --no-cache add ca-certificates

# Copia el binario compilado desde la etapa anterior
COPY --from=builder /app/server /server

# Expón el puerto que utilizará la aplicación (8080 en este caso)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/server"]
