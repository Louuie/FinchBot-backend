package models

import (
	"time"
)

type YouTubeSearch struct {
	Kind          string `pg:"kind,omitempty"`
	Etag          string `pg:"etag,omitempty"`
	NextPageToken string `pg:"nextPageToken,omitempty"`
	RegionCode    string `pg:"regionCode,omitempty"`
	PageInfo      struct {
		TotalResults   int `pg:"totalResults,omitempty"`
		ResultsPerPage int `pg:"resultsPerPage,omitempty"`
	} `pg:"pageInfo,omitempty"`
	Items []struct {
		Kind string `pg:"kind,omitempty"`
		Etag string `pg:"etag,omitempty"`
		ID   struct {
			Kind    string `pg:"kind,omitempty"`
			VideoID string `pg:"videoId,omitempty"`
		} `pg:"id"`
		Snippet struct {
			PublishedAt time.Time `pg:"publishedAt,omitempty"`
			ChannelID   string    `pg:"channelId,omitempty"`
			Title       string    `pg:"title,omitempty"`
			Description string    `pg:"description,omitempty"`
			Thumbnails  struct {
				Default struct {
					URL    string `pg:"url,omitempty"`
					Width  int    `pg:"width,omitempty"`
					Height int    `pg:"height,omitempty"`
				} `pg:"default,omitempty"`
				Medium struct {
					URL    string `pg:"url,omitempty"`
					Width  int    `pg:"width,omitempty"`
					Height int    `pg:"height,omitempty"`
				} `pg:"medium,omitempty"`
				High struct {
					URL    string `pg:"url,omitempty"`
					Width  int    `pg:"width,omitempty"`
					Height int    `pg:"height,omitempty"`
				} `pg:"high,omitempty"`
			} `pg:"thumbnails,omitempty"`
			ChannelTitle         string    `pg:"channelTitle,omitempty"`
			LiveBroadcastContent string    `pg:"liveBroadcastContent,omitempty"`
			PublishTime          time.Time `pg:"publishTime,omitempty"`
		} `pg:"snippet,omitempty"`
	} `pg:"items,omitempty"`
}

type VideoDuration struct {
	Kind  string `pg:"king,omitempty"`
	Etag  string `pg:"etag,omitempty"`
	Items []struct {
		Kind           string `pg:"king,omitempty"`
		Etag           string `pg:"etag,omitempty"`
		ID             string `pg:"id,omitempty"`
		ContentDetails struct {
			Duration        string `pg:"duration,omitempty"`
			Dimension       string `pg:"dimension,omitempty"`
			Definition      string `pg:"definition,omitempty"`
			Caption         string `pg:"caption,omitempty"`
			LicensedContent string `pg:"licensedcontent,omitempty"`
			ContentRating   struct{}
			Projection      string `pg:"projection,omitempty"`
		}
	}
	PageInfo struct {
		TotalResults   int `pg:"totalResults,omitempty"`
		ResultsPerPage int `pg:"resultsPerPage,omitempty"`
	} `pg:"pageInfo,omitempty"`
}

type Song struct {
	User     string  `pg:"user,omitempty"`
	Title    string  `pg:"title,omitempty"`
	Artist   string  `pg:"artist,omitempty"`
	Duration float64 `pg:"duration,omitempty"`
	VideoID  string  `pg:"videoid,omitempty"`
	Position float64 `pg:"position,omitempty"`
}

type Data struct {
	Name     string  `pg:"name,omitempty"`
	Artist   string  `pg:"artist,omitempty"`
	Duration float64 `pg:"duration,omitempty"`
	Position int     `pg:"position,omitempty"`
}
type ClientData struct {
	Status  string `pg:"status,omitempty"`
	Message string `pg:"message,omitempty"`
	Data    []Data `pg:"data,omitempty"`
}

type SongPosition struct {
	Location     float64 `pg:"location,omitempty"`
	IteratorDone bool    `pg:"iteratordone,omitempty"`
}

type DatabaseQuery struct {
	Artist   string `pg:"artist,omitempty"`
	Duration int    `pg:"duration,omitempty"`
	Id       int    `pg:"id,omitempty"`
	Title    string `pg:"title,omitempty"`
	Userid   string `pg:"userid,omitempty"`
	Videoid  string `pg:"videoid,omitempty"`
}

type TwitchAuthResponse struct {
	AccessToken  string   `pg:"access_token,omitempty"`
	ExpiresIn    float64  `pg:"expires_in,omitempty"`
	RefreshToken string   `pg:"refresh_token,omitempty"`
	Scope        []string `pg:"scope,omitempty"`
	TokenType    string   `pg:"token_type,omitempty"`
	Status       float64  `pg:"status,omitempty"`
	Message      string   `pg:"message,omitempty"`
}

type TwitchUserInfoResponse struct {
	Data []struct {
		ID              string    `pg:"id"`
		Login           string    `pg:"login"`
		DisplayName     string    `pg:"display_name"`
		Type            string    `pg:"type"`
		BroadcasterType string    `pg:"broadcaster_type"`
		Description     string    `pg:"description"`
		ProfileImageURL string    `pg:"profile_image_url"`
		OfflineImageURL string    `pg:"offline_image_url"`
		ViewCount       int       `pg:"view_count"`
		Email           string    `pg:"email"`
		CreatedAt       time.Time `pg:"created_at"`
	} `pg:"data"`
}

type TwitchValidateTokenResponse struct {
	ClientID  string   `pg:"client_id"`
	Login     string   `pg:"login"`
	Scopes    []string `pg:"scopes"`
	UserID    string   `pg:"user_id"`
	ExpiresIn int      `pg:"expires_in"`
}

type TwitchRevokeTokenResponse struct {
	Status  int    `pg:"status"`
	Message string `pg:"message"`
}
