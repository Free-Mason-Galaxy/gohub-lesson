// Package test
// descr
// author fm
// date 2022/11/16 10:45
package test

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"gohub-lesson/app/models/test"
	"gohub-lesson/app/models/user"
	"gohub-lesson/pkg/captcha"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/database"
	"gohub-lesson/pkg/logger"
	"gohub-lesson/pkg/response"
	"gohub-lesson/pkg/sms"
	"gohub-lesson/pkg/verifycode"
)

type TestController struct {
}

func Test() {
	panic("这是panic测试代码")
}

type TestInterface interface {
	Testing() string
}

type T struct {
}

func (class T) Testing() string {
	return "testing"
}

type T2 struct {
	T
}

func fn(t TestInterface) {
	t.Testing()
}

// userIdAndUuid null +6
// 频率限制 +2

func worker2(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}

func (class *TestController) Any(ctx *gin.Context) {
	{
	}
	{
		return
	}
	{
		ch := make(chan struct{})
		fmt.Println("1")
		close(ch)

		response.Data(ctx, gin.H{"len": "l"})
		return
	}
	{
		var (
			a []string
		)

		a = append(a, "a")

		b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

		b2 := b[1:4]
		b3 := b[0]
		fmt.Println(string(b2))
		fmt.Println(b)
		fmt.Println(b2)
		fmt.Println(string(111))

		i := []int{1, 2, 3, 4, 5}

		s := make([]byte, 5)

		fmt.Println("len:", len(s))
		fmt.Println("cap:", cap(s))

		s = s[2:4]

		fmt.Println("len:", len(s))
		fmt.Println("cap:", cap(s))

		response.Data(ctx, gin.H{"a": a, "b": string(b2), "b3": b3, "i": i[1:4], "i2": i[1:5], "b4": string(b[:2]), "b5": string(b[2:6])})

		return
	}
	{

		err := &net.OpError{Err: &os.SyscallError{Err: errors.New("broken pipe")}}

		var a any
		a = err
		fmt.Println("a:", a)
		var s *os.SyscallError
		if errors.As(err, &s) {
			// yes
			fmt.Println("SyscallError")
		}

		return
	}
	{
		var t test.Test
		// database.DB.Transaction(func(tx *gorm.DB) (err error) {
		// 	ctx.Set("transaction", tx)
		// 	t.Transaction = tx
		// 	t.Title = "t1"
		// 	t.TxCreate()
		//
		// 	t.Title = "t2"
		// 	t.TxCreate()
		//
		// 	return nil
		// })
		t.Transaction = database.DB.Begin()
		ctx.Set("transaction", t.Transaction)
		t.Title = "t1"
		t.TxCreate()

		t.Title = "t2"
		if t.TxCreate().Error != nil {
			t.Transaction.Rollback()
			return
		}

		t.Transaction.Commit()

		// t.Transaction.Commit()

		return
	}
	{
		var a = map[string]string{
			"k1": "v1",
			"k2": "v2",
		}
		for k, v := range a {
			fmt.Println("k_addr:", &k, k)
			fmt.Println("v_addr:", &v, v)
		}

		return

	}
	{
		var u user.User
		u.Name = "name1"
		u.Email = "name@qq.com"
		u.Password = "admin"
		u.Create()

		response.JSON(ctx, gin.H{
			"data": u,
		})
		return
	}
	{
		var t T2
		fn(t)
	}
	{
		isSuccess := verifycode.NewVerifyCode().SendSMS(ctx.Query("key"))
		response.JSON(ctx, gin.H{"isSuccess": isSuccess})
		return
	}
	{
		sms.NewSMS().Send("17602118840", sms.Message{
			Template: config.GetString("sms.aliyun.template_code"),
			Data:     map[string]string{"code": "23456"},
		})
		return
	}
	{
		response.CJSON(ctx, 200, nil)
		return
	}
	{
		logger.Dump(captcha.NewCaptcha().VerifyCaptcha(ctx.Query("key"), ctx.Query("value")))
		return
	}
	Test()
}
