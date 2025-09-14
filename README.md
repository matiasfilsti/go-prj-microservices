# Microservicios en Go

Este proyecto es una arquitectura de microservicios desarrollada en Go que implementa un sistema completo con autenticación, backend, frontend y capacidades de IA.

## Tecnologías Principales

- **Lenguaje**: Go (Golang)
- **Bases de Datos**:
  - PostgreSQL para datos de autenticación
  - Redis para caché
- **Frontend**: Aplicación web interactiva
- **IA/ML**: Integración con Ollama para modelos de lenguaje
- **Mensajería**: RabbitMQ (actualmente comentado en la configuración)
- **Proxy Inverso**: Nginx
- **Contenedorización**: Docker y Docker Compose

## Arquitectura

El sistema está compuesto por los siguientes servicios:

1. **Authentication Service**: Manejo de autenticación y autorización
2. **Backend Service**: Lógica principal de negocio
3. **Frontend Service**: Interfaz de usuario web
4. **AI Operator**: Servicio de inteligencia artificial
5. **Redis Cache**: Almacenamiento en caché
6. **PostgreSQL**: Base de datos relacional
7. **Nginx**: Proxy inverso y balanceador de carga

## Requisitos Previos

- Docker y Docker Compose instalados
- Git para clonar el repositorio
- Al menos 4GB de RAM disponibles para los contenedores

## Cómo Levantar el Proyecto Localmente

1. **Clonar el repositorio**:
   ```bash
   git clone <url-del-repositorio>
   cd go-prj-microservices
   ```

2. **Iniciar los servicios con Docker Compose**:
   ```bash
   docker-compose up -d
   ```

3. **Verificar que los contenedores estén en ejecución**:
   ```bash
   docker-compose ps
   ```

4. **Acceder a los servicios**:
   - nginx: http://localhost:8080  ( Acceso Principal )
   - Frontend: http://localhost:8085
   - Backend API: http://localhost:8081
   - Autenticación: http://localhost:8082
   - AI Operator: http://localhost:8086
   - Redis Admin: http://localhost:8080

## Variables de Entorno

Los servicios utilizan las siguientes variables de entorno (configuradas en el archivo `docker-compose.yml`):

- `DB_HOSTNAME`: Host de la base de datos
- `MQ_HOSTNAME`: Host de RabbitMQ
- `AUTH_HOSTNAME`: Host del servicio de autenticación
- `REDIS_HOSTNAME`: Host del servicio Redis
- `LLM_URL`: URL del servicio de modelos de lenguaje

## Estructura del Proyecto

```
go-prj-microservices/
├── AIOperator/          # Servicio de operaciones de IA
├── authentication/      # Servicio de autenticación
├── backend/             # Servicio principal del backend
├── consumer/            # Consumidor de mensajes
├── frontend/            # Interfaz de usuario web
├── nginx/               # Configuración de Nginx
├── redisai-data/        # Datos de Redis para IA
└── redisAuth-data/      # Datos de autenticación en Redis
```

## Despliegue

El proyecto está configurado para ser desplegado directamente con Docker Compose. Para entornos de producción, se recomienda:

1. Configurar volúmenes persistentes para las bases de datos
2. Implementar HTTPS con certificados SSL
3. Configurar monitoreo y logs centralizados
4. Establecer políticas de seguridad adecuadas

## Notas Adicionales

- El servicio de RabbitMQ está actualmente comentado en el `docker-compose.yml`
- Se recomienda configurar las credenciales de acceso en producción
- Verificar los puertos expuestos para evitar conflictos

## Contribución

1. Haz un fork del proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Haz commit de tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Haz push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.