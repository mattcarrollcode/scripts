package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const BASE_URL = ""
const OUTPUT_JSON_FILENAME = "results.json"
const LIMIT = 10

func main() {
	if BASE_URL == "" {
		fmt.Println("Please add a value for 'BASE_URL' in 'main.go'.")
		os.Exit(1)
	}
	fmt.Println("Looking for articles on", BASE_URL)

	articleCount := 0
	offset := 0
	numArticles := LIMIT
	storedJSONs := []StoredJSON{}
	for numArticles == LIMIT {
		articles, err := getNextN(LIMIT, offset)
		if err != nil {
			fmt.Println(err)
			return
		}
		numArticles = len(articles)
		for i := 0; i < len(articles); i++ {
			articleCount += 1
			fmt.Printf("\rFound %d articles...", articleCount)
			storedJSONs = append(storedJSONs, StoredJSON{
				URL:    articles[i].CanonicalURL,
				Title:  articles[i].Slug,
				Folder: articles[i].PostDate.Format("2006"),
			})
		}
		offset += LIMIT
	}
	fmt.Println(" Done.")

	fmt.Printf(fmt.Sprintf("\rWriting data to '%s'...", OUTPUT_JSON_FILENAME))
	file, err := json.MarshalIndent(storedJSONs, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(OUTPUT_JSON_FILENAME, file, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done.")
}

func getNextN(limit int, offset int) (SubstackArticle, error) {
	url := fmt.Sprintf("%s/api/v1/archive?sort=new&limit=%d&offset=%d", BASE_URL, limit, offset)
	articles := SubstackArticle{}

	resp, err := http.Get(url)
	if err != nil {
		return articles, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return articles, err
	}

	err = json.Unmarshal(body, &articles)
	if err != nil {
		return articles, err
	}
	return articles, nil
}

type StoredJSON struct {
	URL    string `json:"url"`
	Title  string `json:"title"`
	Folder string `json:"folder"`
}

type SubstackArticle []struct {
	ID                      int           `json:"id"`
	PublicationID           int           `json:"publication_id"`
	Title                   string        `json:"title"`
	SocialTitle             string        `json:"social_title"`
	SearchEngineTitle       interface{}   `json:"search_engine_title"`
	SearchEngineDescription interface{}   `json:"search_engine_description"`
	Type                    string        `json:"type"`
	Slug                    string        `json:"slug"`
	PostDate                time.Time     `json:"post_date"`
	Audience                string        `json:"audience"`
	PodcastDuration         interface{}   `json:"podcast_duration"`
	VideoUploadID           interface{}   `json:"video_upload_id"`
	PodcastUploadID         interface{}   `json:"podcast_upload_id"`
	WriteCommentPermissions string        `json:"write_comment_permissions"`
	ShouldSendFreePreview   bool          `json:"should_send_free_preview"`
	FreeUnlockRequired      bool          `json:"free_unlock_required"`
	DefaultCommentSort      interface{}   `json:"default_comment_sort"`
	CanonicalURL            string        `json:"canonical_url"`
	SectionID               int           `json:"section_id"`
	TopExclusions           []interface{} `json:"top_exclusions"`
	Pins                    []interface{} `json:"pins"`
	IsSectionPinned         bool          `json:"is_section_pinned"`
	SectionSlug             string        `json:"section_slug"`
	SectionName             string        `json:"section_name"`
	Reactions               struct {
		NAMING_FAILED int `json:"❤"`
	} `json:"reactions"`
	Subtitle                string        `json:"subtitle"`
	CoverImage              string        `json:"cover_image"`
	CoverImageIsSquare      bool          `json:"cover_image_is_square"`
	PodcastURL              interface{}   `json:"podcast_url"`
	VideoUpload             interface{}   `json:"videoUpload"`
	PodcastPreviewUploadID  interface{}   `json:"podcast_preview_upload_id"`
	PodcastUpload           interface{}   `json:"podcastUpload"`
	PodcastPreviewUpload    interface{}   `json:"podcastPreviewUpload"`
	VoiceoverUploadID       interface{}   `json:"voiceover_upload_id"`
	VoiceoverUpload         interface{}   `json:"voiceoverUpload"`
	HasVoiceover            bool          `json:"has_voiceover"`
	Description             string        `json:"description"`
	BodyJSON                interface{}   `json:"body_json"`
	BodyHTML                interface{}   `json:"body_html"`
	LongerTruncatedBodyJSON interface{}   `json:"longer_truncated_body_json"`
	LongerTruncatedBodyHTML interface{}   `json:"longer_truncated_body_html"`
	TruncatedBodyText       string        `json:"truncated_body_text"`
	Wordcount               int           `json:"wordcount"`
	PostTags                []interface{} `json:"postTags"`
	PublishedBylines        []struct {
		ID               int         `json:"id"`
		Name             string      `json:"name"`
		PreviousName     interface{} `json:"previous_name"`
		PhotoURL         string      `json:"photo_url"`
		Bio              string      `json:"bio"`
		ProfileSetUpAt   time.Time   `json:"profile_set_up_at"`
		PublicationUsers []struct {
			ID            int    `json:"id"`
			UserID        int    `json:"user_id"`
			PublicationID int    `json:"publication_id"`
			Role          string `json:"role"`
			Public        bool   `json:"public"`
			IsPrimary     bool   `json:"is_primary"`
			Publication   struct {
				ID                    int         `json:"id"`
				Name                  string      `json:"name"`
				Subdomain             string      `json:"subdomain"`
				CustomDomain          string      `json:"custom_domain"`
				CustomDomainOptional  bool        `json:"custom_domain_optional"`
				HeroText              string      `json:"hero_text"`
				LogoURL               string      `json:"logo_url"`
				AuthorID              int         `json:"author_id"`
				ThemeVarBackgroundPop string      `json:"theme_var_background_pop"`
				CreatedAt             time.Time   `json:"created_at"`
				RssWebsiteURL         interface{} `json:"rss_website_url"`
				EmailFromName         string      `json:"email_from_name"`
				Copyright             string      `json:"copyright"`
				FoundingPlanName      interface{} `json:"founding_plan_name"`
				CommunityEnabled      bool        `json:"community_enabled"`
				InviteOnly            bool        `json:"invite_only"`
				PaymentsState         string      `json:"payments_state"`
			} `json:"publication"`
		} `json:"publicationUsers"`
		TwitterScreenName string `json:"twitter_screen_name"`
		IsGuest           bool   `json:"is_guest"`
		BestsellerTier    int    `json:"bestseller_tier"`
		InviteAccepted    bool   `json:"inviteAccepted"`
	} `json:"publishedBylines"`
	Reaction          interface{} `json:"reaction"`
	CommentCount      int         `json:"comment_count"`
	ChildCommentCount int         `json:"child_comment_count"`
	AudioItems        []struct {
		PostID   int         `json:"post_id"`
		VoiceID  string      `json:"voice_id"`
		AudioURL interface{} `json:"audio_url"`
		Type     string      `json:"type"`
		Status   string      `json:"status"`
	} `json:"audio_items"`
	Hidden      bool          `json:"hidden,omitempty"`
	HasCashtag  interface{}   `json:"hasCashtag"`
	Attachments []interface{} `json:"attachments"`
}
