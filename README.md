# ByteBridge

ByteBridge es un prototipo en Go para transferir archivos por TCP dentro de una red local.

El proyecto contiene un servidor que escucha conexiones entrantes y un cliente que envia un archivo al servidor usando un encabezado JSON seguido del contenido binario del archivo.

## Estado del proyecto

Este repositorio esta en desarrollo activo. Actualmente soporta el flujo basico de envio de un archivo desde el cliente al servidor.

Pendiente o en progreso:

- Envio de multiples archivos.
- Seleccion de usuarios/dispositivos destino.
- Deteccion automatica de otros clientes ByteBridge en la red.
- Uso completo de `config.yml`.
- Verificacion de integridad con SHA.
- Barra de progreso.

## Requisitos

- Go `1.26.2` o compatible con la version indicada en `go.mod`.
- `make` para usar los comandos incluidos en el `Makefile`.

## Estructura

```text
.
|-- cmd/
|   |-- client/          # Entrada del cliente
|   `-- server/          # Entrada del servidor
|-- internals/
|   |-- client/          # Conexion TCP y envio de archivos
|   |-- dispatcher/      # Despacho de conexiones entrantes
|   |-- jobs/            # Jobs de envio/recepcion
|   `-- serverconn/      # Pool de conexiones del servidor
|-- config.yml           # Configuracion planeada del proyecto
|-- Makefile             # Comandos de desarrollo
|-- PLAN.md              # Ideas y objetivos del proyecto
`-- go.mod
```

## Uso

Primero crea el directorio donde el servidor guarda los archivos recibidos:

```bash
mkdir -p files/server
```

Levanta el servidor:

```bash
make server
```

Por defecto el servidor escucha en el puerto `4000`.

En otra terminal, envia un archivo con el cliente:

```bash
make client file=/ruta/al/archivo
```

El archivo recibido se guarda en:

```text
files/server/received_<nombre-del-archivo>
```

## Comandos disponibles

```bash
make fmt
```

Formatea el codigo con `go fmt ./...`.

```bash
make vet
```

Ejecuta `go vet ./...` despues de formatear.

```bash
make server
```

Ejecuta el servidor con:

```bash
go run ./cmd/server -port=4000 -env="development"
```

```bash
make client file=/ruta/al/archivo
```

Ejecuta el cliente y envia el archivo indicado.

```bash
make buildserver
make buildclient
```

Compilan los binarios en `bin/server` y `bin/client`.

## Configuracion

Existe un archivo `config.yml` con valores iniciales:

```yaml
network:
  serverport: 4000
  clientport: 4001

storage:
  upload_path: "./files"
```

Nota: el uso completo de esta configuracion todavia esta pendiente. En el estado actual, el servidor recibe el puerto por flag y el cliente conecta a `:4000`.

## Protocolo actual

El cliente envia:

1. Un encabezado JSON terminado en salto de linea.
2. El contenido binario del archivo.

Ejemplo de encabezado:

```json
{
  "type": "ReceiveFileJob",
  "size": 12345,
  "filename": "example.txt",
  "userdata": "testJob"
}
```

El servidor interpreta el campo `type` para construir el job correspondiente y guardar el archivo recibido.

## Desarrollo

Antes de enviar cambios, corre:

```bash
make vet
```

Esto aplica formato y valida el codigo con las herramientas estandar de Go.
