package main

type Response struct {
	Meta Metadata  `json:"meta"`
	Data []Project `json:"data"`
}

type Metadata struct {
	HasNext       bool  `json:"has_next"`
	TotalReturned int64 `json:"total_returned"`
	NextOffset    int64 `json:"next_offset"`
}

type Project struct {
	ID                   string              `json:"id"`
	Category             Category            `json:"category"`
	Name                 string              `json:"name"`
	Description          string              `json:"description"`
	ProfileAvatarURL     string              `json:"profileAvatarUrl"`
	ProjectCoverImageURL string              `json:"projectCoverImageUrl"`
	SocialLinks          SocialLinks         `json:"socialLinks"`
	Team                 []string            `json:"team"`
	Github               []map[string]Github `json:"github"`
	Packages             []interface{}       `json:"packages"`
	Contracts            []Contract          `json:"contracts"`
	GrantsAndFunding     GrantsAndFunding    `json:"grantsAndFunding"`
	Organization         *Organization       `json:"organization"`
	Links                []string            `json:"links"`
}

type Contract struct {
	Address          string `json:"address"`
	DeploymentTxHash string `json:"deploymentTxHash"`
	DeployerAddress  string `json:"deployerAddress"`
	ChainID          int64  `json:"chainId"`
}

type Github struct {
	RepoRank                         int64   `json:"repo_rank"`
	StarCount                        int64   `json:"star_count"`
	StarredEvents                    int64   `json:"starred_events"`
	StarredByTopDevs                 int64   `json:"starred_by_top_devs"`
	ForkCount                        int64   `json:"fork_count"`
	ForkedEvents                     int64   `json:"forked_events"`
	ForkedByTopDevs                  int64   `json:"forked_by_top_devs"`
	FulltimeDeveloperAverage6_Months float64 `json:"fulltime_developer_average_6_months"`
	NewContributorCount6_Months      int64   `json:"new_contributor_count_6_months"`
	AgeOfProjectYears                float64 `json:"age_of_project_years"`
}

type GrantsAndFunding struct {
	VentureFunding []VentureFunding `json:"ventureFunding"`
	Grants         []Grant          `json:"grants"`
	Revenue        []Revenue        `json:"revenue"`
}

type Grant struct {
	Grant   string `json:"grant"`
	Link    string `json:"link"`
	Amount  string `json:"amount"`
	Date    string `json:"date"`
	Details string `json:"details"`
}

type Revenue struct {
	Amount  string `json:"amount"`
	Details string `json:"details"`
}

type VentureFunding struct {
	Amount  string `json:"amount"`
	Year    string `json:"year"`
	Details string `json:"details"`
}

type Organization struct {
	Name             string `json:"name"`
	ProfileAvatarURL string `json:"profileAvatarUrl"`
}

type SocialLinks struct {
	Twitter   string      `json:"twitter"`
	Farcaster []string    `json:"farcaster"`
	Mirror    interface{} `json:"mirror"`
	Website   []string    `json:"website"`
}

type Meta struct {
	HasNext       bool  `json:"has_next"`
	TotalReturned int64 `json:"total_returned"`
	NextOffset    int64 `json:"next_offset"`
}

type Category string

const (
	EthereumCoreContributions     Category = "ETHEREUM_CORE_CONTRIBUTIONS"
	OpStackReseachAndDevelopment  Category = "OP_STACK_RESEACH_AND_DEVELOPMENT"
	OpStackResearchAndDevelopment Category = "OP_STACK_RESEARCH_AND_DEVELOPMENT"
	OpStackTooling                Category = "OP_STACK_TOOLING"
)
