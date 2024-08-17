package meta

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

var (
	userAgentHeader	= "user-agent"
)

type Metdata struct {
	UserAgent 	string
	ClientIp	string
}

func ExtractMetadata(ctx context.Context) *Metdata {
	mtdt := &Metdata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIp = p.Addr.String()
	}

	return mtdt
}