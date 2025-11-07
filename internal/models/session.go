package models

import (
	"time"

	"github.com/google/uuid"
)

// Session represents a visitor session in the database
// Compatible with Umami's session table schema
type Session struct {
	SessionID  uuid.UUID `json:"session_id"`
	WebsiteID  uuid.UUID `json:"website_id"`
	Browser    *string   `json:"browser,omitempty"`
	OS         *string   `json:"os,omitempty"`
	Device     *string   `json:"device,omitempty"`
	Screen     *string   `json:"screen,omitempty"`
	Language   *string   `json:"language,omitempty"`
	Country    *string   `json:"country,omitempty"`
	Region     *string   `json:"region,omitempty"`
	City       *string   `json:"city,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	DistinctID *string   `json:"distinct_id,omitempty"`
}

// Website represents a website in the database
type Website struct {
	WebsiteID uuid.UUID `json:"website_id"`
	Name      string    `json:"name"`
	Domain    *string   `json:"domain,omitempty"`
	ShareID   *string   `json:"share_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// Event represents a pageview or custom event
// Compatible with Umami's website_event table
type Event struct {
	EventID   uuid.UUID `json:"event_id"`
	WebsiteID uuid.UUID `json:"website_id"`
	SessionID uuid.UUID `json:"session_id"`
	VisitID   uuid.UUID `json:"visit_id"`
	CreatedAt time.Time `json:"created_at"`

	// Page info
	PageTitle      *string `json:"page_title,omitempty"`
	Hostname       *string `json:"hostname,omitempty"`
	URLPath        *string `json:"url_path,omitempty"`
	URLQuery       *string `json:"url_query,omitempty"`
	ReferrerPath   *string `json:"referrer_path,omitempty"`
	ReferrerQuery  *string `json:"referrer_query,omitempty"`
	ReferrerDomain *string `json:"referrer_domain,omitempty"`

	// Session info (denormalized for performance)
	DistinctID *string `json:"distinct_id,omitempty"`
	Browser    *string `json:"browser,omitempty"`
	OS         *string `json:"os,omitempty"`
	Device     *string `json:"device,omitempty"`
	Screen     *string `json:"screen,omitempty"`
	Language   *string `json:"language,omitempty"`
	Country    *string `json:"country,omitempty"`
	Region     *string `json:"region,omitempty"`
	City       *string `json:"city,omitempty"`

	// Custom events
	EventName *string `json:"event_name,omitempty"`
	EventData *string `json:"event_data,omitempty"` // JSON string
	Tag       *string `json:"tag,omitempty"`

	// UTM parameters
	UTMSource   *string `json:"utm_source,omitempty"`
	UTMMedium   *string `json:"utm_medium,omitempty"`
	UTMCampaign *string `json:"utm_campaign,omitempty"`
	UTMContent  *string `json:"utm_content,omitempty"`
	UTMTerm     *string `json:"utm_term,omitempty"`

	// Click IDs
	GCLID   *string `json:"gclid,omitempty"`
	FBCLID  *string `json:"fbclid,omitempty"`
	MSCLKID *string `json:"msclkid,omitempty"`
	TTCLID  *string `json:"ttclid,omitempty"`
	LIFATID *string `json:"li_fat_id,omitempty"`
	TWCLID  *string `json:"twclid,omitempty"`
}
