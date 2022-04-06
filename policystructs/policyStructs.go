package strct

//===============================//
//====Base Policy Structs=======//
//=============================//
type Policy struct {
	APIVersion string            `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty" bson:"apiVersion,omitempty"`
	Kind       string            `json:"kind,omitempty" yaml:"kind,omitempty" bson:"kind,omitempty"`
	FlowIDs    []int             `json:"flow_ids,omitempty" yaml:"flow_ids,omitempty" bson:"flow_ids,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty" bson:"metadata,omitempty"`
	Outdated   string            `json:"outdated,omitempty" yaml:"outdated,omitempty" bson:"outdated,omitempty"`

	//Spec Spec `json:"spec,omitempty" yaml:"spec,omitempty" bson:"spec,omitempty"`

	GeneratedTime int64 `json:"generatedTime,omitempty" yaml:"generatedTime,omitempty" bson:"generatedTime,omitempty"`
}

// ============================= //
// == KubeArmor System Policy == //
// ============================= //

type KubeArmorPolicy struct {
	APIVersion string            `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	Kind       string            `json:"kind,omitempty" yaml:"kind,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Spec       KnoxSystemSpec    `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type KnoxSystemSpec struct {
	Tags     []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	Message  string   `json:"message,omitempty" yaml:"message,omitempty"`
	Selector Selector `json:"selector,omitempty" yaml:"selector,omitempty"`
	File     KnoxSys  `json:"file,omitempty" yaml:"file,omitempty"`
	Process  Knoxproc `json:"process,omitempty" yaml:"process,omitempty"`
	//Network []KnoxMatchProtocols `json:"network,omitempty" yaml:"network,omitempty"`

	//Action string `json:"action,omitempty" yaml:"action,omitempty"`
}

// Selector Structure
type Selector struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty" yaml:"matchLabels,omitempty" bson:"matchLabels,omitempty"`
}
type KnoxSys struct {
	Severity         int                    `json:"severity,omitempty" yaml:"severity,omitempty"`
	MatchPaths       []KnoxMatchPaths       `json:"matchPaths,omitempty" yaml:"matchPaths,omitempty"`
	MatchDirectories []KnoxMatchDirectories `json:"matchDirectories,omitempty" yaml:"matchDirectories,omitempty"`
	Action           string                 `json:"action,omitempty" yaml:"action,omitempty"`
}

type Knoxproc struct {
	Severity      int                `json:"severity,omitempty" yaml:"severity,omitempty"`
	MatchPatterns []Knoxmatchpattern `json:"matchPatterns,omitempty" yaml:"matchPatterns,omitempty"`
	Action        string             `json:"action,omitempty" yaml:"action,omitempty"`
}
type KnoxMatchPaths struct {
	Path       string           `json:"path,omitempty" yaml:"path,omitempty"`
	ReadOnly   bool             `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	OwnerOnly  bool             `json:"ownerOnly,omitempty" yaml:"ownerOnly,omitempty"`
	FromSource []KnoxFromSource `json:"fromSource,omitempty" yaml:"fromSource,omitempty"`
}

// KnoxMatchDirectories Structure
type KnoxMatchDirectories struct {
	Dir        string           `json:"dir,omitempty" yaml:"dir,omitempty"`
	ReadOnly   bool             `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	OwnerOnly  bool             `json:"ownerOnly,omitempty" yaml:"ownerOnly,omitempty"`
	FromSource []KnoxFromSource `json:"fromSource,omitempty" yaml:"fromSource,omitempty"`
}
type Knoxmatchpattern struct {
	Pattern    string           `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	ReadOnly   bool             `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	OwnerOnly  bool             `json:"ownerOnly,omitempty" yaml:"ownerOnly,omitempty"`
	FromSource []KnoxFromSource `json:"fromSource,omitempty" yaml:"fromSource,omitempty"`
}
type KnoxFromSource struct {
	Path      string `json:"path,omitempty" yaml:"path,omitempty"`
	Dir       string `json:"dir,omitempty" yaml:"dir,omitempty"`
	Recursive bool   `json:"resursive,omitempty" yaml:"resursive,omitempty"`
}
type KnoxMatchProtocols struct {
	Protocol   string           `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	FromSource []KnoxFromSource `json:"fromSource,omitempty" yaml:"fromSource,omitempty"`
}
