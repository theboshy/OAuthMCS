package main

import (
	/*"net/http"
	"os"*/
	"time"
	"github.com/gin-gonic/gin"
	"../utilities"
	"../dao/factory"
	"../models"
	"../models/custom/structfaces"
	"../models/custom"
	"database/sql"
	"os"
	"net/http"
	"io/ioutil"
	"bytes"
	"github.com/appleboy/gin-jwt"
)

const SECRECT_KEY  = "xnyAqBz3ERocwP1lxi12KUyx0zJ76zwxTicap3QWd59Hwyd19Tq3LIAMInCHkLq.775806DFA54B06C277CCF4D355F42F831D7D9226BB247B6A5163EE4A505AD0E1"


func main()  {



	port := os.Getenv("PORT")
	r := gin.Default()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "strangebug.com",
		Key:        []byte(SECRECT_KEY),
		Timeout:    time.Hour*24,
		MaxRefresh: time.Hour,


		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			result,_ := FindUsuario(userId,password)
			if result.Valid == true && result.Int64 >0{
				return string(result.Int64), true
			}
			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			claims := jwt.ExtractClaims(c)
			if claims["role"] == "admin" {
				return true
			}
			return false
		},
		PayloadFunc: func(userID string) map[string]interface{} {
			//result,_ := FindUsuario(userID,"")

			return map[string]interface{}{
				"dependencia" : "un cafe",
				"role": "admin",
				//--reque / def
				"iat" : utilities.UnixNow(),
				"jti" : utilities.GetUUID(),
			}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:Authorization",
		TokenHeadName: "Bearer",//Bearer

		TimeFunc: time.Now,
	}



	/*r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}*/


	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/signup", SingUpUser)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/test", HapinnesHandlerRouter)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	http.ListenAndServe(":"+port, r)

}


func Authentication(con *sql.DB) gin.HandlerFunc {
	query := "SELECT user_id FROM users WHERE token = ?"

	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Token")
		var user_id int
		err := con.QueryRow(query, tokenHeader ).Scan(&user_id)
		if err != nil {
			// UNAUTHORISED ERROR ERROR ERROR
		}
		c.Set("UserID", user_id)
		c.Next()

	}
}

func FindUsuario(id string,pass string) (sql.NullInt64,error){
	config, err := utilities.GetConfiguration()
	if err != nil {
		panic(err)
	}

	type rque struct {
		idUsuario string
		password  string
	}

	/*var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	ss := new(rque)
	error := c.Bind(&ss)
fmt.Println("error : ",error)*/

	usuarioDao := factory.FactoryDaoUsuario(config.Engine)
	affect, errror := usuarioDao.FindUsuarioLogin(id, pass)

return affect,errror
	//c.JSON(200, radicaResult)
}

func SingUpUser(c *gin.Context)  {
	response :=_struct.Response{}
	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	usuario := new(models.Usuarios)
	error := c.Bind(usuario)
	if error !=nil{
		panic(error)
	}
	//fmt.Println(usuario)

	if usuario !=nil {
		usuario.Contrasena,_ = utilities.EncrypthUni(usuario.Contrasena)
		usuario.FechaCreacionUsuario = time.Now()
		usuario.EstadoUsuario = custom.DefineSystemValue(1)
		usuario.Conexion = custom.DefineSystemValue(0)
		//--

		config, err := utilities.GetConfiguration()
		if err != nil {
			panic(err)
		}


		usuarioDao := factory.FactoryDaoUsuario(config.Engine)
		usarioresult, errror := usuarioDao.CreateUsuario(*usuario)

		if errror != nil || usarioresult ==0{
			response.Status = 500
			response.Message ="insertion error in usuarios #"+usuario.NombreIdUsuario+"]"
			response.Error = errror
			//fmt.Println(error)
			//c.JSON(200, gin.H{"creationError": error})
		}else{
			response.Status = 200
			response.Message = "usuarios #"+usuario.NombreIdUsuario+" created"
			response.Error = nil
		}


	}else{
		response.Status = 500
		response.Message ="insertion error in usuarios #"+usuario.NombreIdUsuario+"]"
		response.Error = error

	}
	c.JSON(response.Status, response)
}

func HapinnesHandlerRouter(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "i just wannaa say hello",
	})
}


