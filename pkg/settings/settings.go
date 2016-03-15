package settings

import (
	"log"
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

	if s.lastResponse == nil || force || time.Since(s.lastResponseTime) > s.cacheTime {
		log.Println("Refreshing settings")
		s.lastResponse, err = s.adminClient.GetSettings(context.Background(), &crowdsound.GetSettingsRequest{})
		if err != nil {
			return settings, err
		}

		s.lastResponseTime = time.Now()
	}

	settings = *s.lastResponse
	return settings, nil
}

func (s *Settings) set(request *crowdsound.SetSettingRequest) error {
	_, err := s.adminClient.SetSetting(context.Background(), request)
	if err != nil {
		return err
	}

	_, err = s.Get(true)
	return err
}

func (s *Settings) SetBool(key string, val bool) error {
	return s.set(&crowdsound.SetSettingRequest{
		Key: key,
		Value: &crowdsound.SetSettingRequest_BoolVal{
			BoolVal: val,
		},
	})
}
func (s *Settings) SetInt(key string, val int) error {
	return s.set(&crowdsound.SetSettingRequest{
		Key: key,
		Value: &crowdsound.SetSettingRequest_IntVal{
			IntVal: int32(val),
		},
	})
}
func (s *Settings) SetFloat(key string, val float32) error {
	return s.set(&crowdsound.SetSettingRequest{
		Key: key,
		Value: &crowdsound.SetSettingRequest_FloatVal{
			FloatVal: val,
		},
	})
}

func (s *Settings) SetString(key string, val string) error {
	return s.set(&crowdsound.SetSettingRequest{
		Key: key,
		Value: &crowdsound.SetSettingRequest_StrVal{
			StrVal: val,
		},
	})
}
