ARG GO_VERSION=1.22.3

# Usa la imagen de Go basada en Alpine
FROM golang:${GO_VERSION}-alpine as builder

RUN go env -w GOPROXY=direct

# Instala git y certificados
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src

# Copia los archivos de go.mod y go.sum
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copia el resto de los archivos
COPY ./ ./

# Compila el proyecto
RUN CGO_ENABLED=0 go build \
     -installsuffix 'static' \
     -o /rest-ws 

# Usa una imagen vacía como runner
FROM scratch AS runner 

# Copia los certificados y el binario
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY .env ./
COPY --from=builder /rest-ws /rest-ws

EXPOSE 5050

# Ejecuta la aplicación
ENTRYPOINT [ "/rest-ws" ]
