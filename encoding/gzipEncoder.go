package encoding

// A gzip implementation of encoder

import (
    "io"
    "compress/gzip"
)

type GzipEncoder struct {
    // nothing
}

func (this *GzipEncoder)ContentEncoding() string {
    return "gzip"
}

func (this *GzipEncoder)ReaderWrapper(r io.Reader) io.ReadCloser {
    var ret, err=gzip.NewReader(r)
    if err!=nil {
        return nil
    }
    return ret
}

func (this *GzipEncoder)WriterWrapper(w io.Writer) io.WriteCloser {
    return gzip.NewWriter(w)
}
