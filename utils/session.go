package utils

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	//gorm.Model
	Id        uint   // Standard field for the primary key
	UserName  string // A regular string field
	Password  string
	Salt      string
	Phone     string
	PhoneCode string
	Email     *string // A pointer to a string, allowing for null values
	//Age          uint8          // An unsigned 8-bit integer
	//Birthday     *time.Time     // A pointer to time.Time, can be null
	//MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	//ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	Time   time.Time // Automatically managed by GORM for creation time
	UpTime time.Time // Automatically managed by GORM for update time
}

var Session sessions.Session

func Init(c *gin.Context) {
	Session = sessions.Default(c)
}

func SessionUp(users []User, c *gin.Context) {
	gap := int64(60 * 60 * 24) // 单位 秒
	fmt.Println("gap: ", gap)
	fmt.Println("gap: ", gap)
	if Session == nil {
		Session = sessions.Default(c)
	}
	//userid := fmt.Sprintf("%v", users[0].Id)
	newTime := time.Now().Unix()
	if Session.Get("session_time") == nil || Session.Get("userId") == nil {
		if string(users[0].Id) != "" {
			Session.Set("session_time", newTime+gap)

			Session.Set("userId", users[0].Id)
		} else {
			fmt.Println("userID为空", users[0])
		}
	} else {
		sessionTime := Session.Get("session_time")
		fmt.Println("sessionTime: ", sessionTime)
		fmt.Println("sessionTime64: ", sessionTime.(int64))
		fmt.Println("过期差距时间", sessionTime.(int64)-newTime)
		if newTime > sessionTime.(int64) {
			fmt.Println("session过期")
			Session.Delete("userId")
			Session.Delete("session_time")
		} else {
			if sessionTime.(int64)-newTime < 10 {
				fmt.Println("session即将过期，自动刷新")
				Session.Set("session_time", newTime)
			}
		}
	}
	// 设置session过期时间为60s
	Session.Options(sessions.Options{MaxAge: int(gap)})
	Session.Save()
	return
}

func LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userId")
	session.Delete("session_time")
}
