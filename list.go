package processout

import (
	"io"
	"net/http"
)

// Identifiable is the interface used by the Iterator to get the ID of the
// resources
type Identifiable interface {
	GetID() string
}

// Iterator is the structure used to iterate over resources sent by the API
type Iterator struct {
	pos     int
	path    string
	data    []Identifiable
	options *Options
	decoder func(io.Reader, interface{}) (bool, error)
	client  *ProcessOut
	err     error

	hasMoreNext bool
	hasMorePrev bool
}

// Error returns the error that occured when paginating, if any
func (i *Iterator) Error() error {
	return i.err
}

// Data returns the data currently fetched by the iterator
func (i *Iterator) Data() []Identifiable {
	return i.data
}

// Get returns the current element
func (i *Iterator) Get() interface{} {
	return i.data[i.pos]
}

// Next iterates on the objects list and fetches new data if available
func (i *Iterator) Next() bool {
	if len(i.data) == 0 {
		return false
	}

	i.pos++
	if i.pos < len(i.data) {
		return true
	}

	if !i.hasMoreNext {
		return false
	}
	hasMore, err := i.NextPage()
	if err != nil {
		i.err = err
		return false
	}

	i.hasMoreNext = hasMore
	i.pos = -1
	return i.Next()
}

// NextPage fetches the next data page
func (i *Iterator) NextPage() (bool, error) {
	prev := i.data[len(i.data)-1]
	i.options.StartAfter = prev.GetID()

	i.hasMorePrev = true
	return i.fetchPage()
}

// Prev iterates on the objects list and fetches new data if available
func (i *Iterator) Prev() bool {
	if len(i.data) == 0 {
		return false
	}

	i.pos--
	if i.pos > 0 {
		return true
	}

	if !i.hasMorePrev {
		return false
	}
	hasMore, err := i.PrevPage()
	if err != nil {
		i.err = err
		return false
	}

	i.hasMorePrev = hasMore
	i.pos = len(i.data)
	return i.Prev()
}

// PrevPage fetches the previous data page
func (i *Iterator) PrevPage() (bool, error) {
	next := i.data[0]
	i.options.EndBefore = next.GetID()

	i.hasMoreNext = true
	return i.fetchPage()
}

func (i *Iterator) fetchPage() (bool, error) {
	req, err := http.NewRequest(
		"GET",
		Host+i.path,
		nil,
	)
	if err != nil {
		return false, err
	}
	setupRequest(i.client, i.options, req)

	res, err := i.client.HTTPClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return i.decoder(res.Body, &i.data)
}
