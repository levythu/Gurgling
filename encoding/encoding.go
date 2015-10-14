package encoding

// Used for define interface for data-encodings

import (
    "io"
)

type Encoder interface {
    // return the string for Content-Encoding header
    ContentEncoding() string

    // the corresponding writer wrapper
    // with any error returns nil
    WriterWrapper(io.Writer) io.WriteCloser

    // the corresponding reader wrapper
    // with any error returns nil
    ReaderWrapper(io.Reader) io.ReadCloser
}

var (
    GZip_Encoder Encoder=&GzipEncoder{}
)
