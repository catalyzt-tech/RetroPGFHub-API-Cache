package main

type Project interface{}

type Response struct {
	Meta Metadata  `json:"meta"`
	Data []Project `json:"data"`
}

type Metadata struct {
	HasNext       bool  `json:"has_next"`
	TotalReturned int64 `json:"total_returned"`
	NextOffset    int64 `json:"next_offset"`
}

// NOTE: since it always update so just use the type any for now is fine
// type Project struct {
// 	ID                   string           `json:"id"`
// 	ApplicationID        string           `json:"applicationId"`
// 	ProjectID            string           `json:"projectId"`
// 	Category             string           `json:"category"`
// 	ApplicationCategory  Category         `json:"applicationCategory"`
// 	Organization         *Organization    `json:"organization"`
// 	Name                 string           `json:"name"`
// 	Description          string           `json:"description"`
// 	ProfileAvatarURL     string           `json:"profileAvatarUrl"`
// 	ProjectCoverImageURL string           `json:"projectCoverImageUrl"`
// 	SocialLinks          SocialLinks      `json:"socialLinks"`
// 	Team                 []Team           `json:"team"`
// 	Github               []Github         `json:"github"`
// 	Packages             []interface{}    `json:"packages"`
// 	Links                []Github         `json:"links"`
// 	Contracts            []Contract       `json:"contracts"`
// 	GrantsAndFunding     GrantsAndFunding `json:"grantsAndFunding"`
// 	PricingModel         PricingModel     `json:"pricingModel"`
// 	ImpactStatement      ImpactStatement  `json:"impactStatement"`
// }

// type Contract struct {
// 	Address           string `json:"address"`
// 	DeploymentTxHash  string `json:"deploymentTxHash"`
// 	DeployerAddress   string `json:"deployerAddress"`
// 	VerificationProof string `json:"verificationProof"`
// 	ChainID           int64  `json:"chainId"`
// }

// type Github struct {
// 	URL         string  `json:"url"`
// 	Name        *string `json:"name"`
// 	Description *string `json:"description"`
// }

// type GrantsAndFunding struct {
// 	VentureFunding []interface{} `json:"ventureFunding"`
// 	Grants         []Grant       `json:"grants"`
// 	Revenue        []interface{} `json:"revenue"`
// }

// type Grant struct {
// 	Grant   *string `json:"grant"`
// 	Link    *string `json:"link"`
// 	Amount  string  `json:"amount"`
// 	Date    string  `json:"date"`
// 	Details *string `json:"details"`
// }

// type ImpactStatement struct {
// 	Category    Category    `json:"category"`
// 	Subcategory []string    `json:"subcategory"`
// 	Statement   []Statement `json:"statement"`
// }

// type Statement struct {
// 	Answer   string `json:"answer"`
// 	Question string `json:"question"`
// }

// type Organization struct {
// 	Name                      string      `json:"name"`
// 	Description               string      `json:"description"`
// 	OrganizationAvatarURL     string      `json:"organizationAvatarUrl"`
// 	OrganizationCoverImageURL *string     `json:"organizationCoverImageUrl"`
// 	SocialLinks               SocialLinks `json:"socialLinks"`
// 	Team                      []string    `json:"team"`
// }

// type SocialLinks struct {
// 	Website   []string `json:"website"`
// 	Farcaster []string `json:"farcaster"`
// 	Twitter   *string  `json:"twitter"`
// 	Mirror    *string  `json:"mirror"`
// }

// type Team struct {
// 	Fid               int64             `json:"fid"`
// 	Object            Object            `json:"object"`
// 	PfpURL            string            `json:"pfp_url"`
// 	Profile           Profile           `json:"profile"`
// 	Username          string            `json:"username"`
// 	PowerBadge        bool              `json:"power_badge"`
// 	DisplayName       string            `json:"display_name"`
// 	ActiveStatus      ActiveStatus      `json:"active_status"`
// 	Verifications     []string          `json:"verifications"`
// 	FollowerCount     int64             `json:"follower_count"`
// 	CustodyAddress    string            `json:"custody_address"`
// 	FollowingCount    int64             `json:"following_count"`
// 	VerifiedAddresses VerifiedAddresses `json:"verified_addresses"`
// }

// type Profile struct {
// 	Bio Bio `json:"bio"`
// }

// type Bio struct {
// 	Text string `json:"text"`
// }

// type VerifiedAddresses struct {
// 	EthAddresses []string `json:"eth_addresses"`
// 	SolAddresses []string `json:"sol_addresses"`
// }

// type Meta struct {
// 	HasNext       bool  `json:"has_next"`
// 	TotalReturned int64 `json:"total_returned"`
// 	NextOffset    int64 `json:"next_offset"`
// }

// type Category string

// const (
// 	EthereumCoreContributions     Category = "ETHEREUM_CORE_CONTRIBUTIONS"
// 	OpStackResearchAndDevelopment Category = "OP_STACK_RESEARCH_AND_DEVELOPMENT"
// 	OpStackTooling                Category = "OP_STACK_TOOLING"
// )

// type PricingModel string

// const (
// 	Free     PricingModel = "free"
// 	Freemium PricingModel = "freemium"
// 	PayToUse PricingModel = "pay_to_use"
// )

// type ActiveStatus string

// const (
// 	Inactive ActiveStatus = "inactive"
// )

// type Object string

// const (
// 	User Object = "user"
// )
