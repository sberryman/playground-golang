package s2s

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sberryman/playground-golang/utils"
	"github.com/sirupsen/logrus"
	l "go.bird.co/bird-common/common-logging-golang"
	"go.bird.co/bird-common/common-utils-golang/s2s"
)

// Initialize s2s library
type singleton struct {
	client         *s2s.Client
	pltoken        string
	pltokenRefresh time.Duration
	sync.Once
}

var singleShared = singleton{}

func init() {
	singleShared.Once.Do(func() {
		var err error

		singleShared.client, err = s2s.NewClient(func(c *s2s.Client) {
			c.Host = utils.Config.S2S.Host
			c.TokenPath = utils.Config.S2S.TokenPath
		})

		if err != nil {
			l.Logger.WithError(err).Error("s2s.NewClient")
			return
		}

		// Also get a powerline token
		if err = singleShared.powerlineToken(); err != nil {
			return
		}

		// Handle PL token refresh
		go func() {
			for {
				select {
				case <-time.After(singleShared.pltokenRefresh):
					singleShared.powerlineToken()
				}
			}
		}()
	})
}

// powerlineToken retrieves a powerline token
func (s *singleton) powerlineToken() error {
	type PowerlineAuthRequest struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	// Use pod hostname to get a unique powerline email
	hostname, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		hostname = fmt.Sprintf("api-goose-%d", time.Now().Unix())
	}
	email := fmt.Sprintf("%s@bird.co", hostname)

	// Create log context since we don't have one
	lpl := l.Logger.WithFields(logrus.Fields{"path": "s2s.powerlineTokenRefresh"})

	// Create request
	req := PowerlineAuthRequest{
		Email: email,
		Role:  "admin",
	}

	// Prepare response struct
	type PowerlineAuthResponse struct {
		Token string `json:"ok"`
	}
	resp := PowerlineAuthResponse{}

	// Send post operation
	if err := NewOp(&Op{
		Hostname: utils.Config.Powerline.ClusterHostname,
		Route:    "/v1/internal/auth",
		ReqBody:  req,
		Resp:     &resp,
	}).Post(); err != nil {
		return err
	}

	// Save new token info
	s.pltoken = resp.Token
	s.pltokenRefresh = time.Hour * 24

	// Get real token expiration
	jwt.Parse(s.pltoken, func(t *jwt.Token) (interface{}, error) {
		claim := t.Claims.(jwt.MapClaims)
		switch exp := claim["exp"].(type) {
		case float64:
			oldTime := time.Unix(int64(exp), 0)
			newTime := oldTime.Add(time.Minute - 10)
			s.pltokenRefresh = time.Until(newTime)
		}
		return true, nil
	})

	// Log refresh
	lpl.WithField("exp", s.pltokenRefresh/time.Second).Info("s2s.powerLineToken-Refreshed")

	return nil
}
