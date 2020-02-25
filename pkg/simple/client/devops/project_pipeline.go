package devops

const (
	NoScmPipelineType       = "pipeline"
	MultiBranchPipelineType = "multi-branch-pipeline"
)

type Parameters []*Parameter

var ParameterTypeMap = map[string]string{
	"hudson.model.StringParameterDefinition":   "string",
	"hudson.model.ChoiceParameterDefinition":   "choice",
	"hudson.model.TextParameterDefinition":     "text",
	"hudson.model.BooleanParameterDefinition":  "boolean",
	"hudson.model.FileParameterDefinition":     "file",
	"hudson.model.PasswordParameterDefinition": "password",
}

type ProjectPipeline struct {
	Type                string               `json:"type" description:"type of devops pipeline, in scm or no scm"`
	Pipeline            *NoScmPipeline       `json:"pipeline,omitempty" description:"no scm pipeline structs"`
	MultiBranchPipeline *MultiBranchPipeline `json:"multi_branch_pipeline,omitempty" description:"in scm pipeline structs"`
}

type NoScmPipeline struct {
	Name              string             `json:"name" description:"name of pipeline"`
	Description       string             `json:"descriptio,omitempty" description:"description of pipeline"`
	Discarder         *DiscarderProperty `json:"discarder,omitempty" description:"Discarder of pipeline, managing when to drop a pipeline"`
	Parameters        *Parameters        `json:"parameters,omitempty" description:"Parameters define of pipeline,user could pass param when run pipeline"`
	DisableConcurrent bool               `json:"disable_concurrent,omitempty" mapstructure:"disable_concurrent" description:"Whether to prohibit the pipeline from running in parallel"`
	TimerTrigger      *TimerTrigger      `json:"timer_trigger,omitempty" mapstructure:"timer_trigger" description:"Timer to trigger pipeline run"`
	RemoteTrigger     *RemoteTrigger     `json:"remote_trigger,omitempty" mapstructure:"remote_trigger" description:"Remote api define to trigger pipeline run"`
	Jenkinsfile       string             `json:"jenkinsfile,omitempty" description:"Jenkinsfile's content'"`
}

type MultiBranchPipeline struct {
	Name                  string                 `json:"name" description:"name of pipeline"`
	Description           string                 `json:"descriptio,omitempty" description:"description of pipeline"`
	Discarder             *DiscarderProperty     `json:"discarder,omitempty" description:"Discarder of pipeline, managing when to drop a pipeline"`
	TimerTrigger          *TimerTrigger          `json:"timer_trigger,omitempty" mapstructure:"timer_trigger" description:"Timer to trigger pipeline run"`
	SourceType            string                 `json:"source_type" description:"type of scm, such as github/git/svn"`
	GitSource             *GitSource             `json:"git_source,omitempty" description:"git scm define"`
	GitHubSource          *GithubSource          `json:"github_source,omitempty" description:"github scm define"`
	SvnSource             *SvnSource             `json:"svn_source,omitempty" description:"multi branch svn scm define"`
	SingleSvnSource       *SingleSvnSource       `json:"single_svn_source,omitempty" description:"single branch svn scm define"`
	BitbucketServerSource *BitbucketServerSource `json:"bitbucket_server_source,omitempty" description:"bitbucket server scm defile"`
	ScriptPath            string                 `json:"script_path" mapstructure:"script_path" description:"script path in scm"`
	MultiBranchJobTrigger *MultiBranchJobTrigger `json:"multibranch_job_trigger,omitempty" mapstructure:"multibranch_job_trigger" description:"Pipeline tasks that need to be triggered when branch creation/deletion"`
}

type GitSource struct {
	ScmId            string          `json:"scm_id,omitempty" description:"uid of scm"`
	Url              string          `json:"url,omitempty" mapstructure:"url" description:"url of git source"`
	CredentialId     string          `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access git source"`
	DiscoverBranches bool            `json:"discover_branches,omitempty" mapstructure:"discover_branches" description:"Whether to discover a branch"`
	CloneOption      *GitCloneOption `json:"git_clone_option,omitempty" mapstructure:"git_clone_option" description:"advavced git clone options"`
	RegexFilter      string          `json:"regex_filter,omitempty" mapstructure:"regex_filter" description:"Regex used to match the name of the branch that needs to be run"`
}

type GithubSource struct {
	ScmId                string               `json:"scm_id,omitempty" description:"uid of scm"`
	Owner                string               `json:"owner,omitempty" mapstructure:"owner" description:"owner of github repo"`
	Repo                 string               `json:"repo,omitempty" mapstructure:"repo" description:"repo name of github repo"`
	CredentialId         string               `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access github source"`
	ApiUri               string               `json:"api_uri,omitempty" mapstructure:"api_uri" description:"The api url can specify the location of the github apiserver.For private cloud configuration"`
	DiscoverBranches     int                  `json:"discover_branches,omitempty" mapstructure:"discover_branches" description:"Discover branch configuration"`
	DiscoverPRFromOrigin int                  `json:"discover_pr_from_origin,omitempty" mapstructure:"discover_pr_from_origin" description:"Discover origin PR configuration"`
	DiscoverPRFromForks  *DiscoverPRFromForks `json:"discover_pr_from_forks,omitempty" mapstructure:"discover_pr_from_forks" description:"Discover fork PR configuration"`
	CloneOption          *GitCloneOption      `json:"git_clone_option,omitempty" mapstructure:"git_clone_option" description:"advavced git clone options"`
	RegexFilter          string               `json:"regex_filter,omitempty" mapstructure:"regex_filter" description:"Regex used to match the name of the branch that needs to be run"`
}

type MultiBranchJobTrigger struct {
	CreateActionJobsToTrigger string `json:"create_action_job_to_trigger,omitempty" description:"pipeline name to trigger"`
	DeleteActionJobsToTrigger string `json:"delete_action_job_to_trigger,omitempty" description:"pipeline name to trigger"`
}

type BitbucketServerSource struct {
	ScmId                string               `json:"scm_id,omitempty" description:"uid of scm"`
	Owner                string               `json:"owner,omitempty" mapstructure:"owner" description:"owner of github repo"`
	Repo                 string               `json:"repo,omitempty" mapstructure:"repo" description:"repo name of github repo"`
	CredentialId         string               `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access github source"`
	ApiUri               string               `json:"api_uri,omitempty" mapstructure:"api_uri" description:"The api url can specify the location of the github apiserver.For private cloud configuration"`
	DiscoverBranches     int                  `json:"discover_branches,omitempty" mapstructure:"discover_branches" description:"Discover branch configuration"`
	DiscoverPRFromOrigin int                  `json:"discover_pr_from_origin,omitempty" mapstructure:"discover_pr_from_origin" description:"Discover origin PR configuration"`
	DiscoverPRFromForks  *DiscoverPRFromForks `json:"discover_pr_from_forks,omitempty" mapstructure:"discover_pr_from_forks" description:"Discover fork PR configuration"`
	CloneOption          *GitCloneOption      `json:"git_clone_option,omitempty" mapstructure:"git_clone_option" description:"advavced git clone options"`
	RegexFilter          string               `json:"regex_filter,omitempty" mapstructure:"regex_filter" description:"Regex used to match the name of the branch that needs to be run"`
}

type GitCloneOption struct {
	Shallow bool `json:"shallow,omitempty" mapstructure:"shallow" description:"Whether to use git shallow clone"`
	Timeout int  `json:"timeout,omitempty" mapstructure:"timeout" description:"git clone timeout mins"`
	Depth   int  `json:"depth,omitempty" mapstructure:"depth" description:"git clone depth"`
}

type SvnSource struct {
	ScmId        string `json:"scm_id,omitempty" description:"uid of scm"`
	Remote       string `json:"remote,omitempty" description:"remote address url"`
	CredentialId string `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access svn source"`
	Includes     string `json:"includes,omitempty" description:"branches to run pipeline"`
	Excludes     string `json:"excludes,omitempty" description:"branches do not run pipeline"`
}
type SingleSvnSource struct {
	ScmId        string `json:"scm_id,omitempty" description:"uid of scm"`
	Remote       string `json:"remote,omitempty" description:"remote address url"`
	CredentialId string `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access svn source"`
}

type DiscoverPRFromForks struct {
	Strategy int `json:"strategy,omitempty" mapstructure:"strategy" description:"github discover strategy"`
	Trust    int `json:"trust,omitempty" mapstructure:"trust" description:"trust user type"`
}

type DiscarderProperty struct {
	DaysToKeep string `json:"days_to_keep,omitempty" mapstructure:"days_to_keep" description:"days to keep pipeline"`
	NumToKeep  string `json:"num_to_keep,omitempty" mapstructure:"num_to_keep" description:"nums to keep pipeline"`
}

type Parameter struct {
	Name         string `json:"name" description:"name of param"`
	DefaultValue string `json:"default_value,omitempty" mapstructure:"default_value" description:"default value of param"`
	Type         string `json:"type" description:"type of param"`
	Description  string `json:"description,omitempty" description:"description of pipeline"`
}

type TimerTrigger struct {
	// user in no scm job
	Cron string `json:"cron,omitempty" description:"jenkins cron script"`

	// use in multi-branch job
	Interval string `json:"interval,omitempty" description:"interval ms"`
}

type RemoteTrigger struct {
	Token string `json:"token,omitempty" description:"remote trigger token"`
}

type ProjectPipelineOperator interface {
	CreateProjectPipeline(projectId string, pipeline *ProjectPipeline) (string, error)
	DeleteProjectPipeline(projectId string, pipelineId string) (string, error)
	UpdateProjectPipeline(projectId string, pipeline *ProjectPipeline) (string, error)
	GetProjectPipelineConfig(projectId, pipelineId string) (*ProjectPipeline, error)
}