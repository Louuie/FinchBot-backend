package models

import (
	"time"
)

type YouTubeSearch struct {
	Kind          string `json:"kind,omitempty"`
	Etag          string `json:"etag,omitempty"`
	NextPageToken string `json:"nextPageToken,omitempty"`
	RegionCode    string `json:"regionCode,omitempty"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults,omitempty"`
		ResultsPerPage int `json:"resultsPerPage,omitempty"`
	} `json:"pageInfo,omitempty"`
	Items []struct {
		Kind string `json:"kind,omitempty"`
		Etag string `json:"etag,omitempty"`
		ID   struct {
			Kind    string `json:"kind,omitempty"`
			VideoID string `json:"videoId,omitempty"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt,omitempty"`
			ChannelID   string    `json:"channelId,omitempty"`
			Title       string    `json:"title,omitempty"`
			Description string    `json:"description,omitempty"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url,omitempty"`
					Width  int    `json:"width,omitempty"`
					Height int    `json:"height,omitempty"`
				} `json:"default,omitempty"`
				Medium struct {
					URL    string `json:"url,omitempty"`
					Width  int    `json:"width,omitempty"`
					Height int    `json:"height,omitempty"`
				} `json:"medium,omitempty"`
				High struct {
					URL    string `json:"url,omitempty"`
					Width  int    `json:"width,omitempty"`
					Height int    `json:"height,omitempty"`
				} `json:"high,omitempty"`
			} `json:"thumbnails,omitempty"`
			ChannelTitle         string    `json:"channelTitle,omitempty"`
			LiveBroadcastContent string    `json:"liveBroadcastContent,omitempty"`
			PublishTime          time.Time `json:"publishTime,omitempty"`
		} `json:"snippet,omitempty"`
	} `json:"items,omitempty"`
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
	Name     string  `json:"name,omitempty"`
	Artist   string  `json:"artist,omitempty"`
	Duration string `json:"duration,omitempty"`
	Position int     `json:"position,omitempty"`
}
type ClientData struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []Data `json:"data,omitempty"`
}

type SongPosition struct {
	Location     float64 `pg:"location,omitempty"`
	IteratorDone bool    `pg:"iteratordone,omitempty"`
}

type DatabaseQuery struct {
	Id       int    `pg:"id,omitempty"`
	Title    string `pg:"title,omitempty"`
	Artist   string `pg:"artist,omitempty"`
	Userid   string `pg:"userid,omitempty"`
	Duration string    `pg:"duration,omitempty"`
	Videoid  string `pg:"videoid,omitempty"`
}

type TwitchAuthResponse struct {
	AccessToken  string   `json:"access_token,omitempty"`
	ExpiresIn    float64  `json:"expires_in,omitempty"`
	RefreshToken string   `json:"refresh_token,omitempty"`
	Scope        []string `json:"scope,omitempty"`
	TokenType    string   `json:"token_type,omitempty"`
	Status       float64  `json:"status,omitempty"`
	Message      string   `json:"message,omitempty"`
}

type TwitchUserInfoResponse struct {
	Data []struct {
		ID              string    `json:"id"`
		Login           string    `json:"login"`
		DisplayName     string    `json:"display_name"`
		Type            string    `json:"type"`
		BroadcasterType string    `json:"broadcaster_type"`
		Description     string    `json:"description"`
		ProfileImageURL string    `json:"profile_image_url"`
		OfflineImageURL string    `json:"offline_image_url"`
		ViewCount       int       `json:"view_count"`
		Email           string    `json:"email"`
		CreatedAt       time.Time `json:"created_at"`
	} `pg:"data"`
}

type TwitchValidateTokenResponse struct {
	ClientID  string   `json:"client_id"`
	Login     string   `json:"login"`
	Scopes    []string `json:"scopes"`
	UserID    string   `json:"user_id"`
	ExpiresIn int      `json:"expires_in"`
}

type TwitchRevokeTokenResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ModifyChannel struct {
	GameID     string `json:"game_id"`
	Title      string `json:"title"`
	StreamLang string `json:"broadcaster_language"`
}

type SearchCategoriesResponse struct {
	Data []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		BoxArtURL string `json:"box_art_url"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type CurrentCategoryResponse struct {
	Data []struct {
		BroadcasterID       string `json:"broadcaster_id"`
		BroadcasterLogin    string `json:"broadcaster_login"`
		BroadcasterName     string `json:"broadcaster_name"`
		BroadcasterLanguage string `json:"broadcaster_language"`
		GameID              string `json:"game_id"`
		GameName            string `json:"game_name"`
		Title               string `json:"title"`
		Delay               int32  `json:"delay"`
	}
}

type UnitTesting []struct {
	Description          string
	Path                 string
	Method               string
	ExpectedCodeResponse int
}
