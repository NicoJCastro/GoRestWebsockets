
# Usames la imagen de PostgresSQL versión 10.3

FROM postgres:10.3

# Copiamos el archivo up.sql a la carpeta /docker-entrypoint-initdb.d/1.sql dentro del contenedor.

COPY up.sql /docker-entrypoint-initdb.d/1.sql

# Ejecutamos el comando postgres para iniciar el contenedor.

CMD ["postgres"]