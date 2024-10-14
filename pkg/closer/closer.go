package closer

import "errors"

type CloseFunc = func() error

type MultipleCloser []func() error

func (c *MultipleCloser) Append(f func() error) {
	*c = append(*c, f)
}

func (c *MultipleCloser) AppendNoErr(f func()) {
	*c = append(*c, func() error {
		f()
		return nil
	})
}

func (c *MultipleCloser) Close() error {
	var err error
	for _, f := range *c {
		if e := f(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}
