package file_svr

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/console"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	// 默认端口41230
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		console.Warn("输入的端口号异常，使用默认的41230")
		port = 41230
	}

	// 获取执行时，用户所在目录，读取下面的文件
	dir, err := os.Getwd()
	if err != nil {
		console.Errorf("获取执行目录文件失败:[%s]\n", err)
		return
	}
	console.Info("准备开启http文件服务")

	hosts, err := getLocalIPs()
	if err != nil {
		console.Warnf("获取内网ip失败:[%s]\n", err)
	}

	for _, host := range hosts {
		console.Infof("服务地址:[%s:%d]\n", host, port)
	}

	mux := http.NewServeMux()
	mux.Handle("/", &index{dir: dir})
	mux.Handle("/file", &file{dir: dir})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	zs := application.NewZSig()
	zs.Wait(ctx, func(ctx context.Context) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 启动服务
			if err = server.ListenAndServe(); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					console.Info("服务正常关闭")
				} else {
					console.Warnf("服务异常关闭, %v", err)
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				// 退出信号
				if sErr := server.Shutdown(context.Background()); sErr != nil {
					fmt.Printf("http服务退出异常:%s\n", sErr)
				}

			}
		}()

		wg.Wait()
	})

	console.Info("执行完毕")
}

// 获取所有内网 IP 地址
func getLocalIPs() ([]string, error) {
	// 用于存储内网 IP 地址
	var localIPs []string

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// 遍历每个接口
	for _, iface := range interfaces {
		// 排除回环接口和非启用接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口的地址列表
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		// 遍历每个地址
		for _, addr := range addrs {
			// 检查地址类型是否为 IPNet
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			// 检查是否为 IPv4 地址
			ip := ipNet.IP.To4()
			if ip == nil {
				continue
			}

			// 检查是否为私有 IP 地址
			if isPrivateIP(ip) {
				localIPs = append(localIPs, ip.String())
			}
		}
	}

	if len(localIPs) == 0 {
		return nil, fmt.Errorf("no private IP addresses found")
	}

	return localIPs, nil
}

// 检查是否为私有 IP 地址
func isPrivateIP(ip net.IP) bool {
	// 私有 IP 地址范围：
	// 10.0.0.0 - 10.255.255.255
	// 172.16.0.0 - 172.31.255.255
	// 192.168.0.0 - 192.168.255.255
	if ip[0] == 10 {
		return true
	}
	if ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31 {
		return true
	}
	if ip[0] == 192 && ip[1] == 168 {
		return true
	}
	return false
}
