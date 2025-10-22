package iikocloud

import (
	"github.com/themgmd/iikocloud/internal/constants"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Client клиент для iikoCloud API
type Client struct {
	quit   chan struct{}
	logger *slog.Logger

	baseURL              string
	apiLogin             string
	token                string
	httpClient           *http.Client
	timeout              time.Duration
	refreshTokenInterval time.Duration
}

// SetLogger установить пользовательский логгер
func (c *Client) SetLogger(handler slog.Handler) {
	c.logger = slog.New(handler)
}

// SetTimeout установить свой таймаут на запросы. Дефолтный 15 секунд
func (c *Client) SetTimeout(t time.Duration) {
	c.timeout = t
}

// SetRefreshTokenInterval Установить период обновления токена авториазации
func (c *Client) SetRefreshTokenInterval(t time.Duration) {
	c.refreshTokenInterval = t
}

// SetHTTPClient установить пользовательский http.Client для запросов в iikoCloud API
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) refreshTokenByInterval() {
	ticker := time.NewTicker(c.refreshTokenInterval)

	c.quit = make(chan struct{})

	for {
		select {
		case <-ticker.C:
			req := AccessTokenRequest{
				ApiLogin: c.apiLogin,
			}

			resp, err := c.accessToken(req)
			if err != nil {
				c.logger.Error(
					"failed to refresh token",
					slog.String("error", err.Error()),
				)
			}

			c.token = resp.Token

		case <-c.quit:
			ticker.Stop()
			return
		}
	}
}

// Close "красивая" остановка клиента
func (c *Client) Close() {
	close(c.quit)
}

// NewClient конструктор для клиента iikoCloud API
func NewClient(apiLogin string) (*Client, error) {
	client := &Client{
		baseURL:              constants.BaseURL,
		httpClient:           http.DefaultClient,
		logger:               slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		apiLogin:             apiLogin,
		timeout:              constants.DefaultTimeout,
		refreshTokenInterval: constants.DefaultRefreshTokenInterval,
	}

	resp, err := client.accessToken(&AccessTokenRequest{
		ApiLogin: client.apiLogin,
	})
	if err != nil {
		return nil, err
	}

	client.token = resp.Token

	go client.refreshTokenByInterval()

	return client, nil
}
