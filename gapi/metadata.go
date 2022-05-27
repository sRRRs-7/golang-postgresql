package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedProtoHeader      = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIp  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	meta := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("metadata: %+v\n", md)
		if userAgent := md.Get(grpcGatewayUserAgentHeader); len(userAgent) > 0 {
			meta.UserAgent = userAgent[0]
		}

		if userAgent := md.Get(userAgentHeader); len(userAgent) > 0 {
			meta.UserAgent = userAgent[0]
		}

		if clientIp := md.Get(xForwardedProtoHeader); len(clientIp) > 0 {
			meta.ClientIp = clientIp[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		meta.ClientIp = p.Addr.String()
	}

	return meta
}
