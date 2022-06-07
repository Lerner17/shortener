package middleware

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (gw gzipWriter) Write(b []byte) (int, error) {
	// Writer response by gzip
	return gw.Writer.Write(b)
}

func GZIPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			reader, err := gzip.NewReader(r.Body)
			if err != nil {
				log.Printf("Cannot decompress request: %v\n", err)
				next.ServeHTTP(w, r)
				return
			}
			defer reader.Close()
			r.Body = reader
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			log.Printf("Compress error: %v\n", err)
			next.ServeHTTP(w, r)
			return
		}
		defer gz.Close()
		w.Header().Set("Content-Encoding", "gzip")

		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)

	})
}