package experiment

import "time"

type Result struct {
	SchemaVersion int    `json:"SchemaVersion"`
	ArtifactName  string `json:"ArtifactName"`
	ArtifactType  string `json:"ArtifactType"`
	Metadata      struct {
		Os struct {
			Family string `json:"Family"`
			Name   string `json:"Name"`
		} `json:"OS"`
		ImageID     string   `json:"ImageID"`
		DiffIDs     []string `json:"DiffIDs"`
		RepoTags    []string `json:"RepoTags"`
		RepoDigests []string `json:"RepoDigests"`
		ImageConfig struct {
			Architecture  string    `json:"architecture"`
			Container     string    `json:"container"`
			Created       time.Time `json:"created"`
			DockerVersion string    `json:"docker_version"`
			History       []struct {
				Created    time.Time `json:"created"`
				CreatedBy  string    `json:"created_by"`
				EmptyLayer bool      `json:"empty_layer,omitempty"`
			} `json:"history"`
			Os     string `json:"os"`
			Rootfs struct {
				Type    string   `json:"type"`
				DiffIds []string `json:"diff_ids"`
			} `json:"rootfs"`
			Config struct {
				Cmd          []string `json:"Cmd"`
				Env          []string `json:"Env"`
				Image        string   `json:"Image"`
				WorkingDir   string   `json:"WorkingDir"`
				ExposedPorts struct {
					Eight0TCP struct {
					} `json:"80/tcp"`
				} `json:"ExposedPorts"`
				StopSignal string `json:"StopSignal"`
			} `json:"config"`
		} `json:"ImageConfig"`
	} `json:"Metadata"`
	Results []struct {
		Target          string `json:"Target"`
		Class           string `json:"Class"`
		Type            string `json:"Type"`
		Vulnerabilities []struct {
			VulnerabilityID  string `json:"VulnerabilityID"`
			PkgName          string `json:"PkgName"`
			InstalledVersion string `json:"InstalledVersion"`
			Layer            struct {
				Digest string `json:"Digest"`
				DiffID string `json:"DiffID"`
			} `json:"Layer"`
			SeveritySource string `json:"SeveritySource"`
			PrimaryURL     string `json:"PrimaryURL"`
			DataSource     struct {
				ID   string `json:"ID"`
				Name string `json:"Name"`
				URL  string `json:"URL"`
			} `json:"DataSource"`
			Title       string   `json:"Title"`
			Description string   `json:"Description"`
			Severity    string   `json:"Severity"`
			CweIDs      []string `json:"CweIDs,omitempty"`
			Cvss        struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
			} `json:"CVSS,omitempty"`
			References       []string  `json:"References"`
			PublishedDate    time.Time `json:"PublishedDate"`
			LastModifiedDate time.Time `json:"LastModifiedDate"`
			Cvss0            struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss1 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss2 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss3 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss4 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss5 struct {
				Nvd struct {
					V2Vector string `json:"V2Vector"`
					V2Score  int    `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string `json:"V2Vector"`
					V2Score  int    `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss6 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss7 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss8 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss9 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss10 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss11 struct {
				Nvd struct {
					V2Vector string `json:"V2Vector"`
					V2Score  int    `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string `json:"V2Vector"`
					V2Score  int    `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss12 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss13 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss14 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss15 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss16 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss17 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss18 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss19 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss20 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss21 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss22 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss23 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss24 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss25 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss26 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss27 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss28 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss29 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss30 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss31 struct {
				Nvd struct {
					V2Vector string `json:"V2Vector"`
					V2Score  int    `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss32 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss33 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss34 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss35 struct {
				Nvd struct {
					V2Vector string `json:"V2Vector"`
					V2Score  int    `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss36 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss37 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss38 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss39 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss40 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss41 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss42 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss43 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss44 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss45 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss46 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss47 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss48 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss49 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss50 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss51 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss52 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss53 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss54 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss55 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss56 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss57 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss58 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss59 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss60 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss61 struct {
				Nvd struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss62 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss63 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss64 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss65 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss66 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss67 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss68 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss69 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss70 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss71 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss72 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss73 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss74 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss75 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  int     `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V2Vector string  `json:"V2Vector"`
					V2Score  float64 `json:"V2Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss76 struct {
				Nvd struct {
					V2Vector string `json:"V2Vector"`
					V2Score  int    `json:"V2Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string `json:"V3Vector"`
					V3Score  int    `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
			Cvss77 struct {
				Nvd struct {
					V2Vector string  `json:"V2Vector"`
					V3Vector string  `json:"V3Vector"`
					V2Score  float64 `json:"V2Score"`
					V3Score  float64 `json:"V3Score"`
				} `json:"nvd"`
				Redhat struct {
					V3Vector string  `json:"V3Vector"`
					V3Score  float64 `json:"V3Score"`
				} `json:"redhat"`
			} `json:"CVSS,omitempty"`
		} `json:"Vulnerabilities"`
	} `json:"Results"`
}
