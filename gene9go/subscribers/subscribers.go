package subscribers

import "github.com/lileio/lile/pubsub"

// See https://godoc.org/github.com/lileio/lile/pubsub#Handler for an example
// of what an subscriber handler should look like

type Gene9goServiceSubscriber struct{}

func (s *Gene9goServiceSubscriber) Setup(c *pubsub.Client) {
	// // https://godoc.org/github.com/lileio/lile/pubsub#Client.On
	// c.On("service.event_name", s.CallbackMethod, 30*time.Second, true)
}
