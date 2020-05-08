package model

type InstagramResponse struct {
	Graphql Graphql `json:"graphql"`
}

type Graphql struct {
	ShortcodeMedia ShortcodeMedia `json:"shortcode_media"`
}

type ShortcodeMedia struct {
	Owner Owner
	Image
	EdgeSidecarToChildren EdgeSidecarToChildren `json:"edge_sidecar_to_children"`
}

type EdgeSidecarToChildren struct {
	Edges []Edge `json:"edges"`
}

type Edge struct {
	Node Image `json:"node"`
}

type Image struct {
	ID                   string             `json:"id"`
	Shortcode            string             `json:"shortcode"`
	MediaPreview         interface{}        `json:"media_preview"`
	DisplayUrl           interface{}        `json:"display_url"`
	DisplayResources     []DisplayResources `json:"display_resources"`
	AccessibilityCaption interface{}        `json:"accessibility_caption"`
	IsVideo              bool               `json:"is_video"`
	TrackingToken        string             `json:"tracking_token"`
}

type DisplayResources struct {
	Src          string `json:"src"`
	ConfigWidth  int    `json:"config_width"`
	ConfigHeight int    `json:"config_height"`
}

type Owner struct {
	ID       string `json:"id"`
	Username string
	FullName string `json:"full_name"`
}
