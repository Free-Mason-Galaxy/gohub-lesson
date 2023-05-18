// Package chatgpt
// ChatGPT 示例代码
// author fm
// date 2023/3/1 13:47
package chatgpt

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

var httpClientPool sync.Pool

func init() {
	httpClientPool = sync.Pool{
		New: func() interface{} {
			return &http.Client{Timeout: 5 * time.Second}
		},
	}
}

func DoGet(url string) (string, error) {
	client := httpClientPool.Get().(*http.Client)
	defer httpClientPool.Put(client)

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// httpRequestWithTimeout
// 在使用 os.CreateTemp 函数时，有可能会出现并发性能问题。因为该函数在创建文
// 件时，使用了随机名称和当前时间作为前缀和后缀，这可能会导致大量的文件名冲突。
// 此外，该函数还会调用 os.MkdirAll 函数创建目录，这可能会导致大量的磁盘 I/O
// 操作，从而影响性能。
//
// 为了避免这种情况，我们可以使用 ioutil.TempFile 函数来创建临时文件。与
// os.CreateTemp 不同，ioutil.TempFile 函数不需要指定前缀和后缀，并且
// 它使用操作系统的默认临时目录
//
// 在处理 HTTP 响应时，有两种常见的方式：将响应内容直接保存到内存中，
// 或将响应内容保存到文件中。在一些场景下，直接将响应内容保存到内存中是可行的，
// 比如响应内容很小，或者内存资源较为充足。但是在一些情况下，响应内容可能非常大，
// 这时将响应内容保存到内存中可能会导致程序占用过多的内存资源，甚至可能导致程序崩溃。
//
// 为了避免这种情况，我们可以将响应内容保存到文件中。这样可以避免程序占用过多的
// 内存资源，同时还可以减小程序的崩溃风险。此外，将响应内容保存到文件中，还可以
// 避免一些潜在的安全问题，比如防止攻击者通过构造恶意请求来占用程序的内存资源，
// 从而导致程序崩溃或变得不稳定。
//
// 在示例中，将响应内容保存到了一个临时文件中，因为它是一种比较常见的做法。由于我
// 们不确定响应内容的大小，所以使用一个文件来保存响应内容是比较保险的。此外，在使
// 用完响应内容后，我们还可以使用 os.Remove 函数删除文件，从而释放磁盘空间。
func httpRequestWithTimeout(url string, timeout time.Duration) (string, error) {
	// 创建一个 http 客户端
	// client := &http.Client{}
	// // 创建一个 context，设置超时时间
	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()
	//
	// // 创建一个 GET 请求，并设置 context
	// req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	//
	// if err != nil {
	// 	return "", err
	// }
	//
	// // 发送请求并获取响应
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return "", err
	// }

	// 使用 context.WithTimeout 函数可以为 HTTP 请求设置超时时间，
	// 这是一种很好的做法。但是在这个例子中，我使用了 &http.Client{Timeout: timeout}
	// 来创建 http.Client，而没有使用 context.WithTimeout 函数的主要原因是简化代码。
	// 使用 &http.Client{Timeout: timeout} 可以直接设置 http.Client 的超时时间，
	// 从而避免了创建 context.Context 和取消请求的步骤。这样可以减少代码的复杂度，使代
	// 码更加简洁。同时，由于这个函数只是下载文件，没有其他需要使用 context.Context 的
	// 场景，因此我认为使用 &http.Client{Timeout: timeout} 更加合适。但是，在一些需
	// 要使用 context.Context 的场景下，使用 context.WithTimeout 是一种很好的做法
	client := &http.Client{Timeout: timeout}

	// 发起 GET 请求
	resp, err := client.Get(url)

	// 关闭响应体
	defer resp.Body.Close()

	// 将响应体中的内容复制到一个文件中
	file, err := os.CreateTemp("", "http-response-")
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Name())

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	// 返回文件路径
	return file.Name(), nil
}

// httpRequestWithTimeoutExample
// 调用示例
func httpRequestWithTimeoutExample() {
	resp, err := httpRequestWithTimeout("http://example.com", time.Second*5)
	if err != nil {
		// 处理错误
	}
	data, err := os.ReadFile(resp)
	if err != nil {
		// 处理错误
	}
	fmt.Println(string(data))
}

// GetServerIP 获取服务器 IP
// 注意，这里的实现只能获取公网 IP 地址，如果需要获取内网 IP 地址，
// 可以将 isPublicIP 函数中的判断逻辑去掉即可。
func GetServerIP() (ip string) {
	interfaces, err := net.Interfaces()

	if err != nil {
		return
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()

		if err != nil {
			return
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)

			// 公网 IP
			// if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil && isPublicIP(ipnet.IP) {
			//				return ipnet.IP.String(), nil
			//			}

			// 内网 IP
			if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return
}

// isPublicIP 是否公网 IP
func isPublicIP(ip net.IP) bool {
	if ip.IsPrivate() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}

	if ip4 := ip.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}

	return false
}
