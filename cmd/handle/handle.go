package handle

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/anguloc/zet/internal/pkg/console"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/mattn/go-tty"
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

func Run(cmd *cobra.Command, args []string) {
	isRun = true
	sigFn()

	fmt.Println("start")

	t, err := tty.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = t.Close()
		fmt.Println("close end")
	}()
	for {
		//if !isRun {
		//	fmt.Println("rec")
		//	break
		//}
		fmt.Println("r")
		//r, rerr := t.ReadRune()
		r, rerr := t.ReadString()
		fmt.Println("ro")
		if rerr != nil {
			fmt.Println(rerr)
			return
		}
		fmt.Println(r)
	}

	fmt.Println("end")

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

func sigFn() {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, []os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM}...)

	go func() {
		s := <-sig
		console.Info("recv quit sig")
		go func() {
			isRun = false
		}()
		<-sig
		os.Exit(128 + int(s.(syscall.Signal)))
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
