package file_svr

import (
	"net/http"
)

type index struct {
	dir string
}

func (i *index) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	i.dir = "C:\\Users\\anguloc\\Desktop\\t1"
	fs := http.FileServer(http.Dir(i.dir))
	fs.ServeHTTP(writer, req)
	// doc, err := i.main(i.dir)
	// if err != nil {
	// 	doc = []byte(fmt.Sprintf("服务异常,%v", err))
	// }
	// _, _ = writer.Write(doc)
}

// func (i *index) main(dir string) ([]byte, error) {
// 	buf := bytes.Buffer{}
//
// 	buf.Write("")
//
// 	return buf.Bytes(), nil
// }
