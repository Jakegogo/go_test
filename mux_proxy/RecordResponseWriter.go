package sexpr

import (
	"bytes"
	"net/http"
)

type RecordResponseWriter struct {
	writer     *http.ResponseWriter
	buffer     *bytes.Buffer
	statusCode int
}

func NewRecordResponseWriter(r *http.ResponseWriter, buffer *bytes.Buffer) (rw *RecordResponseWriter) {
	return &RecordResponseWriter{r, buffer, 200}
}

func (w *RecordResponseWriter) Header() http.Header {
	return (*(w.writer)).Header()
}

func (w RecordResponseWriter) Write(bytes []byte) (int, error) {
	w.buffer.Write(bytes)
	return (*(w.writer)).Write(bytes)
}

func (w RecordResponseWriter) WriteHeader(statusCode int) {
	(*(w.writer)).WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (w RecordResponseWriter) GetStatusCode() int {
	return w.statusCode
}
