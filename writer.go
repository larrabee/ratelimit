package ratelimit

import "io"

type writer struct {
	writer io.Writer
	bucket Bucket
}

type writeSeeker struct {
	writer io.WriteSeeker
	bucket Bucket
}

type writeCloser struct {
	writer io.WriteCloser
	bucket Bucket
}

// NewWriter returns a writer that is rate limited by
// the given token bucket. Each token in the bucket
// represents one byte.
func NewWriter(w io.Writer, bucket Bucket) io.Writer {
	return &writer{
		writer: w,
		bucket: bucket,
	}
}

func (w *writer) Write(buf []byte) (int, error) {
	w.bucket.Wait(int64(len(buf)))
	return w.writer.Write(buf)
}

// NewWriter returns a writer that is rate limited by
// the given token bucket. Each token in the bucket
// represents one byte.
func NewWriteSeeker(w io.WriteSeeker, bucket Bucket) io.WriteSeeker {
	return &writeSeeker{
		writer: w,
		bucket: bucket,
	}
}

func (w *writeSeeker) Write(buf []byte) (int, error) {
	w.bucket.Wait(int64(len(buf)))
	return w.writer.Write(buf)
}

func (r *writeSeeker) Seek(offset int64, whence int) (int64, error) {
	return r.writer.Seek(offset, whence)
}

// NewWriter returns a writer that is rate limited by
// the given token bucket. Each token in the bucket
// represents one byte.
func NewWriteCloser(w io.WriteCloser, bucket Bucket) io.WriteCloser {
	return &writeCloser{
		writer: w,
		bucket: bucket,
	}
}

func (w *writeCloser) Write(buf []byte) (int, error) {
	w.bucket.Wait(int64(len(buf)))
	return w.writer.Write(buf)
}

func (r *writeCloser) Close() error {
	return r.writer.Close()
}
