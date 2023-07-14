package handle

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"unicode"

	"github.com/anguloc/zet/pkg/console"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hd",
		Short: "handle",
		Long:  "handlinginteractions",
		Run:   Run,
	}
	return cmd
}

var isRun = false

var cancel context.CancelFunc

func Run(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()
	ctx, cancel = context.WithCancel(ctx)
	isRun = true

	console.Debug("start")
	_, err := it.open()
	if err != nil {
		console.Error("tty打开错误")
		return
	}
	defer it.close()
	closeWg := &sync.WaitGroup{}
	sigFn(closeWg)

	closeWg.Add(1)
	go func() {
		defer closeWg.Done()
		<-ctx.Done()
		it.close()
	}()

	readCh := make(chan rune, 1024)

	readConsole(readCh)

	handleConsole(ctx, closeWg, readCh)

	<-ctx.Done()
	closeWg.Wait()
	console.Debug("end")
	return
	c, err := newCompleter()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Please select table.")
	p := prompt.New(
		executor,
		c.Func,
		prompt.OptionTitle("zet-prompt: interactive handler"),
		prompt.OptionPrefix(">>> "),
		//prompt.OptionInputTextColor(prompt.Yellow),
		//prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		//prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		//prompt.OptionSuggestionBGColor(prompt.DarkGray),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}

func sigFn(closeWg *sync.WaitGroup) {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, []os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM}...)

	go func() {
		s := <-sig
		console.Debug("recv quit sig")
		closeWg.Add(1)
		go func() {
			defer closeWg.Done()
			cancel()
			it.close()
			console.Debug("close end")
		}()
		<-sig
		os.Exit(128 + int(s.(syscall.Signal)))
	}()
}

func readConsole(ch chan<- rune) {
	go func() {
		for {
			if it.tty == nil {
				break
			}
			r, rerr := it.tty.ReadRune()
			if it.tty == nil {
				break
			}
			if rerr != nil {
				console.Errorf("tty异常")
				cancel()
				break
			}
			ch <- r
		}
	}()
}

func handleConsole(ctx context.Context, closeWg *sync.WaitGroup, ch <-chan rune) {
	var (
		r  rune
		rs []rune
	)
	closeWg.Add(1)
	go func() {
		defer closeWg.Done()
		for {
			select {
			case r = <-ch:
				console.Debug(r, ":", string(r))
				switch r {
				case 13:
					_, _ = it.tty.Output().WriteString("\n")
					fmt.Println(rs)
					rs = nil
				case 8, 127:
					if len(rs) > 0 {
						rs = rs[:len(rs)-1]
						_, _ = it.tty.Output().WriteString("\b \b")
					}
				case 4:
					cancel()
					return
				default:
					if unicode.IsPrint(r) {
						rs = append(rs, r)
						_, _ = it.tty.Output().WriteString(string(r))
					}
				}
			case <-ctx.Done():
				console.Debug("handleConsole end")
				return
			}
		}
	}()
}

func completer1(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func executor(in string) {
	strings.TrimSpace(in)
	if in == "" {
		return
	} else if in == "quit" || in == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}
}

type c struct {
}

func newCompleter() (*c, error) {
	return &c{}, nil
}

func (c *c) Func(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	// args := strings.Split(d.TextBeforeCursor(), " ")
	// w := d.GetWordBeforeCursor()
	//    _ = w

	// for i := range args {
	// 	if args[i] == "|" {
	// 		return []prompt.Suggest{}
	// 	}
	// }

	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	return nil
}
