package multiple

import (
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/anguloc/zet/internal/pkg/console"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	var err error

	num, err := cmd.Flags().GetInt("num")
	if num <= 0 || err != nil {
		console.Errorf("num参数错误,num:%d\n", num)
		return
	}
	count, err := cmd.Flags().GetInt("count")
	if count <= 0 || err != nil {
		console.Errorf("count参数错误,count:%d\n", count)
		return
	}
	yes, _ := cmd.Flags().GetBool("yes")
	commandType, _ := cmd.Flags().GetString("shell")
	if commandType != "bash" && commandType != "sh" {
		commandType = "bash"
	}
	if len(args) == 0 {
		console.Error("缺失执行命令")
		return
	}
	command := args[0]
	if command == "" {
		console.Error("缺失command参数")
		return
	}

	console.Infof("开始执行命令:[%s],shell类型为:[%s],总执行次数:[%d],协程数:[%d]\n", command, commandType, num, count)

	if !yes {
		var scanRsp string
		console.Infof("输出[y]确认执行")
		_, err = fmt.Scanln(&scanRsp)
		if err != nil {
			console.Info("中断执行")
			return
		}
		if _, ok := map[string]struct{}{"y": {}, "Y": {}, "yes": {}, "YES": {}}[scanRsp]; !ok {
			console.Info("中断执行")
			return
		}
	}

	_ = runTask(command, commandType, num, count)
	return
}

type threadItem struct {
	num      int
	taskNum  int
	costTime time.Duration
}

func runTask(command, commandType string, num, count int) error {
	wg := &sync.WaitGroup{}

	startChan := make(chan struct{})
	taskChan := make(chan struct{}, num)
	resultChan := make(chan threadItem, count)

	console.Info("开始加载任务")
	startTime := time.Now()
	for i := 0; i < num; i++ {
		taskChan <- struct{}{}
	}
	close(taskChan)
	for i := 0; i < count; i++ {
		k := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-startChan
			goStartTime := time.Now()
			taskNum := 0
			for {
				if _, ok := <-taskChan; ok {
					taskNum++
					console.Debugf("存量任务数:[%d]\n", len(taskChan))
					cmd := exec.Command(commandType, "-c", command)
					err := cmd.Run()
					if err != nil {
						console.Infof("执行命令异常,err:%s\n", err)
					}
				} else {
					break
				}
			}
			resultChan <- threadItem{
				num:      k,
				taskNum:  taskNum,
				costTime: time.Since(goStartTime),
			}
		}()
	}
	costTime := time.Since(startTime)
	console.Infof("任务加载完毕,耗时:%s\n", costTime.String())

	console.Info("开始执行任务")
	startTime = time.Now()
	close(startChan)
	wg.Wait()
	close(resultChan)
	costTime = time.Since(startTime)

	printResult(resultChan, count)
	console.Infof("任务执行完毕,耗时:[%s],平均:[%s]\n", costTime, costTime/time.Duration(num))

	return nil
}

func printResult(c <-chan threadItem, count int) {
	threadMap := make(map[int]threadItem, count)
	for i := 0; i < count; i++ {
		item := <-c
		threadMap[item.num] = item
	}
	type item struct {
		taskNum  int
		costTime time.Duration
		avaTime  time.Duration
	}
	s := make([]item, 0, count)
	for i := 0; i < count; i++ {
		v := threadMap[i]
		avaTime := time.Duration(0)
		if v.taskNum > 0 {
			avaTime = v.costTime / time.Duration(v.taskNum)
		}
		s = append(s, item{
			taskNum:  v.taskNum,
			costTime: v.costTime,
			avaTime:  avaTime,
		})
	}
	for i, v := range s {
		console.Infof("协程序号:[%d],耗时:[%s],执行任务数:[%d],平均耗时:[%s]\n", i, v.costTime, v.taskNum, v.avaTime)
	}
}

