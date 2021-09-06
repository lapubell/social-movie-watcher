package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

type giphyResponse struct {
	Data []struct {
		Type             string `json:"type"`
		ID               string `json:"id"`
		URL              string `json:"url"`
		Slug             string `json:"slug"`
		BitlyGifURL      string `json:"bitly_gif_url"`
		BitlyURL         string `json:"bitly_url"`
		EmbedURL         string `json:"embed_url"`
		Username         string `json:"username"`
		Source           string `json:"source"`
		Title            string `json:"title"`
		Rating           string `json:"rating"`
		ContentURL       string `json:"content_url"`
		SourceTld        string `json:"source_tld"`
		SourcePostURL    string `json:"source_post_url"`
		IsSticker        int    `json:"is_sticker"`
		ImportDatetime   string `json:"import_datetime"`
		TrendingDatetime string `json:"trending_datetime"`
		Images           struct {
			Original struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
			} `json:"original"`
			Downsized struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized"`
			DownsizedLarge struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_large"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_medium"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"downsized_small"`
			DownsizedStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_still"`
			FixedHeight struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_height"`
			FixedHeightDownsampled struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_height_downsampled"`
			FixedHeightSmall struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_height_small"`
			FixedHeightSmallStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_height_small_still"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_height_still"`
			FixedWidth struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_width"`
			FixedWidthDownsampled struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_width_downsampled"`
			FixedWidthSmall struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_width_small"`
			FixedWidthSmallStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_width_small_still"`
			FixedWidthStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_width_still"`
			Looping struct {
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"looping"`
			OriginalStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"original_still"`
			OriginalMp4 struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"original_mp4"`
			Preview struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"preview"`
			PreviewGif struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"preview_gif"`
			PreviewWebp struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"preview_webp"`
			Hd struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"hd"`
			Four80WStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"480w_still"`
		} `json:"images"`
		User struct {
			AvatarURL    string `json:"avatar_url"`
			BannerImage  string `json:"banner_image"`
			BannerURL    string `json:"banner_url"`
			ProfileURL   string `json:"profile_url"`
			Username     string `json:"username"`
			DisplayName  string `json:"display_name"`
			Description  string `json:"description"`
			InstagramURL string `json:"instagram_url"`
			WebsiteURL   string `json:"website_url"`
			IsVerified   bool   `json:"is_verified"`
		} `json:"user"`
		AnalyticsResponsePayload string `json:"analytics_response_payload"`
		Analytics                struct {
			Onload struct {
				URL string `json:"url"`
			} `json:"onload"`
			Onclick struct {
				URL string `json:"url"`
			} `json:"onclick"`
			Onsent struct {
				URL string `json:"url"`
			} `json:"onsent"`
		} `json:"analytics"`
	} `json:"data"`
	Pagination struct {
		TotalCount int `json:"total_count"`
		Count      int `json:"count"`
		Offset     int `json:"offset"`
	} `json:"pagination"`
	Meta struct {
		Status     int    `json:"status"`
		Msg        string `json:"msg"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}

func getRandomGifUrlByTag(t string) string {
	apikey := os.Getenv("GIPHY_API")
	if apikey == "" {
		return "Can't work with giphy without an API_KEY set 😭"
	}

	s := "https://api.giphy.com/v1/gifs/search?" +
		"api_key=" + os.Getenv("GIPHY_API") +
		"&q=" + url.QueryEscape(t) +
		"&limit=30" +
		"&rating=r"

	res, err := http.Get(s)
	if err != nil {
		return "Couldn't ask giphy for an image 😭"
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "Couldn't read giphy response 😭"
	}

	var jsonData giphyResponse
	err = json.Unmarshal(resBytes, &jsonData)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(resBytes)
		fmt.Println(t)
		fmt.Println(s)
		return "Couldn't parse giphy response 😭"
	}

	rand.Seed(time.Now().UnixNano())
	key := rand.Intn(20)
	randomGif := jsonData.Data[key]

	return "Giphy: " + t + "<br /><img src='" +
		randomGif.Images.Downsized.URL +
		"' width='" + randomGif.Images.Downsized.Width +
		"' height='" + randomGif.Images.Downsized.Height +
		"' style='max-width: 100%;' alt='giphy image' />"
}
