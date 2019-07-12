package ratelimit

import "io"

type readWriter struct {
	rw     io.ReadWriter
	bucket Bucket
}

type readWriteCloser struct {
	rw     io.ReadWriteCloser
	bucket Bucket
}

type readWriteSeeker struct {
	rw     io.ReadWriteSeeker
	bucket Bucket
}

// NewReadWriter returns a readwriter that is rate limited by
// the given token bucket. Each token in the bucket
// represents one byte.
func NewReadWriter(w io.ReadWriter, bucket Bucket) io.ReadWriter {
	return &readWriter{
		rw:     w,
		bucket: bucket,
	}
}

func (w *readWriter) Write(buf []byte) (int, error) {
	w.bucket.Wait(int64(len(buf)))
	return w.rw.Write(buf)
}

func (r *readWriter) Read(buf []byte) (int, error) {
	n, err := r.rw.Read(buf)
	if n <= 0 {
		return n, err
	}
	r.bucket.Wait(int64(n))
	return n, err
}

// NewReadWriter returns a readwriter that is rate limited by
// the given token bucket. Each token in the bucket
// represents one byte.
func NewReadWriteCloser(w io.ReadWriteCloser, bucket Bucket) io.ReadWriteCloser {
	return &readWriteCloser{
		rw:     w,
		bucket: bucket,
	}
}

func (w *readWriteCloser) Write(buf []byte) (int, error) {
	w.bucket.Wait(int64(len(buf)))
	return w.rw.Write(buf)
}

func (r *readWriteCloser) Read(buf []byte) (int, error) {
	n, err := r.rw.Read(buf)
	if n <= 0 {
		return n, err
	}
	r.bucket.Wait(int64(n))
	return n, err
}

func (r *readWriteCloser) Close() error {
	return r.rw.Close()
}

// NewReadWriter returns a readwriter that is rate limited by
// the given token bucket. Each token in the bucket
// represents one byte.
func NewReadWriteSeeker(w io.ReadWriteSeeker, bucket Bucket) io.ReadWriteSeeker {
	return &readWriteSeeker{
		rw:     w,
		bucket: bucket,
	}
}

func (w *readWriteSeeker) Write(buf []byte) (int, error) {
	w.bucket.Wait(int64(len(buf)))
	return w.rw.Write(buf)
}

func (r *readWriteSeeker) Read(buf []byte) (int, error) {
	n, err := r.rw.Read(buf)
	if n <= 0 {
		return n, err
	}
	r.bucket.Wait(int64(n))
	return n, err
}

func (r *readWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	return r.rw.Seek(offset, whence)
}
