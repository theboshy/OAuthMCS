package midleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"strings"
	"github.com/robbert229/jwt"
)
const SECRECT_KEY_CU  = "xnyAqBz3ERocwP1lxi12KUyx0zJ76zwxTicap3QWd59Hwyd19Tq3LIAMInCHkLq.775806DFA54B06C277CCF4D355F42F831D7D9226BB247B6A5163EE4A505AD0E1"

func AuthMidleware() gin.HandlerFunc {
	//-

	//--
	return func(c *gin.Context) {
		as :=c.Request.Header["Authorization"]
		fmt.Println(as)
		fmt.Println(strings.Replace(string(as[0]), "Bearer", "", -1))
		c.Next()
	}
}

func JwtGin()  {
	secret := SECRECT_KEY_CU
	algorithm := jwt.HmacSha256(secret)

	claims := jwt.NewClaim()
	claims.Set("Role", "Admin")
	claims.Set("ID", "alv")
	claims.SetTime("exp", time.Now().Add(time.Minute))

	token, err := algorithm.Encode(claims)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Token: %s\n", token)

	if algorithm.Validate(token) != nil {
		panic(err)
	}

	loadedClaims, err := algorithm.Decode(token)
	if err != nil {
		panic(err)
	}

	role, err := loadedClaims.Get("Role")
	if err != nil {
		panic(err)
	}

	roleString, ok := role.(string)
	if !ok {
		panic(err)
	}

	if strings.Compare(roleString, "Admin") == 0 {
		//user is an admin
		fmt.Println("User is an admin")
	}
}