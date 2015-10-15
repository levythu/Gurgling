package encoding

// No encoder present for convenience.

import (
    "io"
)


type NoEncoder_Reader struct {
    r io.Reader
}
func (this *NoEncoder_Reader)Close() error {
    // Do not close the underlying stream
    return nil
}
func (this *NoEncoder_Reader)Read(p []byte) (int, error) {
    return this.r.Read(p)
}

type NoEncoder_Writer struct {
    w io.Writer
}
func (this *NoEncoder_Writer)Close() error {
    // Do not close the underlying stream
    return nil
}
func (this *NoEncoder_Writer)Write(p []byte) (int, error) {
    return this.w.Write(p)
}

type NoEncoder struct {
    // nothing
}

func (this *NoEncoder)ContentEncoding() string {
    return ""
}

func (this *NoEncoder)ReaderWrapper(r io.Reader) io.ReadCloser {
    return &NoEncoder_Reader {
        r: r,
    }
}

func (this *NoEncoder)WriterWrapper(w io.Writer) io.WriteCloser {
    return &NoEncoder_Writer {
        w: w,
    }
}
