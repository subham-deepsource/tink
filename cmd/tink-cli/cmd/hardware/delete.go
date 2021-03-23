// Copyright © 2018 packet.net

package hardware

import (
	"context"

	"github.com/raydeann/tink/client"
	"github.com/raydeann/tink/cmd/tink-cli/cmd/delete"
	"github.com/raydeann/tink/protos/hardware"
)

type deleteHardware struct {
	delete.Options
}

func (h *deleteHardware) DeleteByID(ctx context.Context, cl *client.FullClient, requestedID string) (interface{}, error) {
	return cl.HardwareClient.Delete(ctx, &hardware.DeleteRequest{Id: requestedID})
}

func NewDeleteOptions() delete.Options {
	h := deleteHardware{}
	return delete.Options{
		DeleteByID: h.DeleteByID,
	}
}
