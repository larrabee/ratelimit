// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package ratelimit

import "io"

type reader struct {
	reader io.Reader
	bucket Bucket
}

type readSeeker struct {
	reader io.ReadSeeker
	bucket Bucket
}

// NewReader returns a reader that is rate limited by
// the given token bucket. Each token in the bucket
// represents one byte.
func NewReader(r io.Reader, bucket Bucket) io.Reader {
	return &reader{
		reader: r,
		bucket: bucket,
	}
}

func (r *reader) Read(buf []byte) (int, error) {
	n, err := r.reader.Read(buf)
	if n <= 0 {
		return n, err
	}
	r.bucket.Wait(int64(n))
	return n, err
}

func NewReadSeeker(r io.ReadSeeker, bucket Bucket) io.Reader {
	return &readSeeker{
		reader: r,
		bucket: bucket,
	}
}

func (r *readSeeker) Read(buf []byte) (int, error) {
	n, err := r.reader.Read(buf)
	if n <= 0 {
		return n, err
	}
	r.bucket.Wait(int64(n))
	return n, err
}

func (r *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return r.reader.Seek(offset, whence)
}
