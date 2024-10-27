# Usando a imagem oficial do Go como base
FROM golang:1.22 AS builder

# Definindo o diretório de trabalho
WORKDIR /app

# Copiando os arquivos do código-fonte para o contêiner
COPY . .

# Instalando dependências
RUN go mod tidy

# Construindo a aplicação
RUN CGO_ENABLED=0 go build -o server cmd/main.go

# Criando uma nova imagem para o runtime
FROM alpine:latest

# Instalando glibc (se necessário)
RUN apk add --no-cache libc6-compat

# Copiando o executável da imagem de builder
COPY --from=builder /app/server /server

# Expondo a porta que a aplicação irá rodar
EXPOSE 8001

# Comando para executar a aplicação
CMD ["/server"]