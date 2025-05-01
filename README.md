# go-prj-microservices
microservices project

# Log into DB
psql -d auth-api -U authapi



TTL <nombre_de_la_clave>
HGETALL <nombre_del_hash>
GET <nombre_de_la_clave>
KEYS *

ver usuarios 
acl list
ACL GETUSER <nombre_del_usuario>





Rabbit url
http://localhost:8082/#/queues

package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "log"
    "io/ioutil"
)

func main() {
    r := gin.Default()

    r.Use(AuthMiddleware())

    r.GET("/someEndpoint", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Authorized!"})
    })

    r.Run(":8080")
}



func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Crear cliente HTTP
        client := &http.Client{}
        
        // Crear petición HTTP
        req, err := http.NewRequest("GET", "http://localhost:8081/checkAuth", nil)
        if err != nil {
            log.Fatal(err)
        }

        // Establecer encabezado de autenticación básica
        req.SetBasicAuth("user", "password")

        // Enviar petición
        resp, err := client.Do(req)
        if err != nil {
            log.Fatal(err)
        }
        defer resp.Body.Close()

        // Leer respuesta
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }

        if resp.StatusCode != http.StatusOK {
            c.JSON(http.StatusUnauthorized, gin.H{"error": string(body)})
            c.Abort()
            return
        }

        c.Next()
    }
}



-----
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    
    r.GET("/checkAuth", func(c *gin.Context) {
        // Obtener credenciales del encabezado de autenticación
        user, password, hasAuth := c.Request.BasicAuth()
        if !hasAuth || user != "user" || password != "password" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "User authorized"})
    })

    r.Run(":8081")
}

----------
package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "io/ioutil"
)

func main() {
    r := gin.Default()

    r.Use(AuthMiddleware())

    r.GET("/someEndpoint", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Authorized!"})
    })

    r.Run(":8080")
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, password, hasAuth := c.Request.BasicAuth()
        if !hasAuth {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        
        log.Printf("Received credentials - User: %s, Password: %s", user, password)

        // Crear cliente HTTP
        client := &http.Client{}
        
        // Crear petición HTTP
        req, err := http.NewRequest("GET", "http://localhost:8081/checkAuth", nil)
        if err != nil {
            log.Fatal(err)
        }

        // Establecer encabezado de autenticación básica
        req.SetBasicAuth(user, password)

        // Enviar petición
        resp, err := client.Do(req)
        if err != nil {
            log.Fatal(err)
        }
        defer resp.Body.Close()

        // Leer respuesta
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }

        if resp.StatusCode != http.StatusOK {
            c.JSON(http.StatusUnauthorized, gin.H{"error": string(body)})
            c.Abort()
            return
        }

        c.Next()
    }
}



--------------------

user, password, hasAuth := c.Request.BasicAuth()
if !hasAuth {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
    c.Abort()
    return
}
log.Printf("Received credentials - User: %s, Password: %s", user, password)
