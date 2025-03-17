package main

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "net/http"
    "time"
)

// Clave secreta para firmar los tokens (¡mantenla segura y privada!)
var jwtSecret = []byte("mi-super-clave-secreta")

// Estructura de los claims del token
type Claims struct {
    UserID string `json:"user_id"` // Puedes añadir más campos según tu necesidad
    jwt.RegisteredClaims
}

// Función para generar un JWT
func generateJWT(userID string) (string, error) {
    // Configurar los claims del token
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expira en 24 horas
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    // Crear el token firmado
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// Middleware para validar el JWT en las solicitudes protegidas
func jwtMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Obtener el token del encabezado Authorization
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
            c.Abort()
            return
        }

        // Parsear y validar el token
        token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })

        // Comprobar si el token es válido
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
            c.Abort()
            return
        }

        // Extraer los claims si el token es válido
        claims, ok := token.Claims.(*Claims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudieron obtener los claims"})
            c.Abort()
            return
        }

        // Guardar el UserID en el contexto para su uso posterior
        c.Set("user_id", claims.UserID)
        c.Next()
    }
}

func main() {
    r := gin.Default()

    // Ruta para iniciar sesión y obtener un token
    r.POST("/login", func(c *gin.Context) {
        // Supongamos que el usuario se autentica correctamente
        userID := "12345" // Esto normalmente vendría de una base de datos

        // Generar un JWT
        token, err := generateJWT(userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
            return
        }

        // Responder con el token
        c.JSON(http.StatusOK, gin.H{
            "token": token,
        })
    })

    // Ruta protegida que requiere un token válido
    r.GET("/protegido", jwtMiddleware(), func(c *gin.Context) {
        // Obtener el UserID desde el contexto
        userID := c.GetString("user_id")
        c.JSON(http.StatusOK, gin.H{
            "mensaje": "Acceso permitido",
            "user_id": userID,
        })
    })

    r.Run(":8080") // Ejecutar el servidor en el puerto 8080
}
