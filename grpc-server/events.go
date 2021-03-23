package grpcserver

import (
	"io"

	"github.com/raydeann/tink/client/informers"
	"github.com/raydeann/tink/client/listener"
	"github.com/raydeann/tink/protos/events"
)

func (s *server) Watch(req *events.WatchRequest, stream events.EventsService_WatchServer) error {
	err := s.db.Events(req, func(n informers.Notification) error {
		event, err := n.ToEvent()
		if err != nil {
			return err
		}
		return stream.Send(event)
	})
	if err != nil && err != io.EOF {
		return err
	}

	return listener.Listen(req, func(e *events.Event) error {
		err := stream.Send(e)
		if err != nil {
			s.logger.With("eventTypes", req.EventTypes, "resourceTypes", req.ResourceTypes).Info("events stream closed")
			return listener.RemoveHandlers(req)
		}
		return nil
	})
}
