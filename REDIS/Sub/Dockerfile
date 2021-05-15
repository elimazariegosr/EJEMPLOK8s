# Obtener la imagen que necesitamos
FROM golang

# Asignar el workdir de la api
WORKDIR /

# Copiar todos los archivos al workdir
COPY . .

#ENV HOST=0.0.0.0:50051

# Correr 'npm install' para instalar todas las dependencias que necesitemos
RUN go mod download
# Aca pueden existir muchos mas comandos "run"

# Correr NodeJS
CMD ["go", "run", "sub.go"]