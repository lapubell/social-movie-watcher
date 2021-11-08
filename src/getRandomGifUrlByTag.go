package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type giphyResponse struct {
	Data struct {
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
			Downsized struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized"`
		} `json:"images"`
		ImageOriginalURL             string `json:"image_original_url"`
		ImageURL                     string `json:"image_url"`
		ImageMp4URL                  string `json:"image_mp4_url"`
		ImageFrames                  string `json:"image_frames"`
		ImageWidth                   string `json:"image_width"`
		ImageHeight                  string `json:"image_height"`
		FixedHeightDownsampledURL    string `json:"fixed_height_downsampled_url"`
		FixedHeightDownsampledWidth  string `json:"fixed_height_downsampled_width"`
		FixedHeightDownsampledHeight string `json:"fixed_height_downsampled_height"`
		FixedWidthDownsampledURL     string `json:"fixed_width_downsampled_url"`
		FixedWidthDownsampledWidth   string `json:"fixed_width_downsampled_width"`
		FixedWidthDownsampledHeight  string `json:"fixed_width_downsampled_height"`
		FixedHeightSmallURL          string `json:"fixed_height_small_url"`
		FixedHeightSmallStillURL     string `json:"fixed_height_small_still_url"`
		FixedHeightSmallWidth        string `json:"fixed_height_small_width"`
		FixedHeightSmallHeight       string `json:"fixed_height_small_height"`
		FixedWidthSmallURL           string `json:"fixed_width_small_url"`
		FixedWidthSmallStillURL      string `json:"fixed_width_small_still_url"`
		FixedWidthSmallWidth         string `json:"fixed_width_small_width"`
		FixedWidthSmallHeight        string `json:"fixed_width_small_height"`
		Caption                      string `json:"caption"`
	} `json:"data"`
	Meta struct {
		Status     int    `json:"status"`
		Msg        string `json:"msg"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}
type emptyGiphyResponse struct {
	Data []interface{} `json:"data"`
	Meta struct {
		Msg        string `json:"msg"`
		Status     int    `json:"status"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}

func getRandomGifUrlByTag(t string) string {
	apikey := os.Getenv("GIPHY_API")
	if apikey == "" {
		return ":giphy " + t + "<br />Can't work with giphy without an API_KEY set 😭"
	}

	s := "https://api.giphy.com/v1/gifs/random?" +
		"api_key=" + apikey +
		"&tag=" + url.QueryEscape(t) +
		"&rating=r"

	res, err := http.Get(s)
	if err != nil {
		return ":giphy " + t + "<br />Couldn't ask giphy for an image 😭"
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ":giphy " + t + "<br />Couldn't read giphy response 😭"
	}

	var jsonData giphyResponse
	err = json.Unmarshal(resBytes, &jsonData)
	if err != nil {
		var emptyJsonData emptyGiphyResponse
		err = json.Unmarshal(resBytes, &emptyJsonData)

		if err != nil {
			return ":giphy " + t + "<br />Something truly fucked up happened 😭"
		}

		return ":giphy " + t + "<br />Empty giphy response 😭"
	}

	randomGif := jsonData.Data

	return ":giphy " + t + "<br /><img src='" +
		randomGif.Images.Downsized.URL +
		"' width='" + randomGif.Images.Downsized.Width +
		"' height='" + randomGif.Images.Downsized.Height +
		"' style='max-width: 100%;' alt='giphy image' />"
}
