package main

type Response struct {
	Meta Meta    `json:"meta"`
	Data []Datum `json:"data"`
}

type Datum struct {
	ID                   string           `json:"id"`
	Category             Category         `json:"category"`
	Organization         interface{}      `json:"organization"`
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	ProfileAvatarURL     string           `json:"profileAvatarUrl"`
	ProjectCoverImageURL string           `json:"projectCoverImageUrl"`
	SocialLinks          SocialLinks      `json:"socialLinks"`
	Team                 []string         `json:"team"`
	Github               []string         `json:"github"`
	Packages             []string         `json:"packages"`
	Links                []interface{}    `json:"links"`
	Contracts            []Contract       `json:"contracts"`
	GrantsAndFunding     GrantsAndFunding `json:"grantsAndFunding"`
}

type Contract struct {
	Address          string `json:"address"`
	DeploymentTxHash string `json:"deploymentTxHash"`
	DeployerAddress  string `json:"deployerAddress"`
	ChainID          int64  `json:"chainId"`
}

type GrantsAndFunding struct {
	VentureFunding []VentureFunding `json:"ventureFunding"`
	Grants         []Grant          `json:"grants"`
	Revenue        []Revenue        `json:"revenue"`
}

type Grant struct {
	Grant   string  `json:"grant"`
	Link    *string `json:"link"`
	Amount  string  `json:"amount"`
	Date    string  `json:"date"`
	Details string  `json:"details"`
}

type Revenue struct {
	Amount  Amount `json:"amount"`
	Details string `json:"details"`
}

type VentureFunding struct {
	Amount  Amount `json:"amount"`
	Year    string `json:"year"`
	Details string `json:"details"`
}

type SocialLinks struct {
	Twitter   *string  `json:"twitter"`
	Farcaster []string `json:"farcaster"`
	Mirror    *string  `json:"mirror"`
	Website   []string `json:"website"`
}

type Meta struct {
	HasNext       bool  `json:"has_next"`
	TotalReturned int64 `json:"total_returned"`
	NextOffset    int64 `json:"next_offset"`
}

type Category string

const (
	CeFi       Category = "CeFi"
	CrossChain Category = "Cross Chain"
	DeFi       Category = "DeFi"
	Governance Category = "Governance"
	Nft        Category = "NFT"
	Social     Category = "Social"
	Utility    Category = "Utility"
)

type Amount string

const (
	Above5M     Amount = "above-5m"
	The1M3M     Amount = "1m-3m"
	The250K500K Amount = "250k-500k"
	The3M5M     Amount = "3m-5m"
	The500K1M   Amount = "500k-1m"
	Under250K   Amount = "under-250k"
)
