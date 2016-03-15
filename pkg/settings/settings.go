package settings

import (
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/crowdsoundsystem/web-client/pkg/crowdsound"
	"google.golang.org/grpc"
)

type Settings struct {
	conn        *grpc.ClientConn
	adminClient crowdsound.AdminClient
	cacheTime   time.Duration

	mut              sync.Mutex
	lastResponse     *crowdsound.GetSettingsResponse
	lastResponseTime time.Time
}

func NewSettings(url string, cacheTime time.Duration) (*Settings, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Settings{
		conn:        conn,
		adminClient: crowdsound.NewAdminClient(conn),
		cacheTime:   cacheTime,
	}, nil
}

func (s *Settings) Get(force bool) (settings crowdsound.GetSettingsResponse, err error) {
	s.mut.Lock()
	defer s.mut.Unlock()

	if s.lastResponse == nil && (force || time.Since(s.lastResponseTime) > s.cacheTime) {
		s.lastResponse, err = s.adminClient.GetSettings(context.Background(), &crowdsound.GetSettingsRequest{})
		if err != nil {
			return settings, err
		}

		s.lastResponseTime = time.Now()
	}

	settings = *s.lastResponse
	return settings, nil
}

func (s *Settings) Set(key string, v interface{}) {
}
