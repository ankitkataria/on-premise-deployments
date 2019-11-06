package utils

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`

	Clusters []Cluster `yaml:"clusters"`
	Users    []User    `yaml:"users"`
	Contexts []Context `yaml: "contexts"`
}

// Cluster configs
type Cluster struct {
	Name          string
	ClusterConfig ClusterConfig `yaml:"cluster"`
}

type ClusterConfig struct {
	CertificateAuthority string `yaml:"certificate-authority"`
	Server               string
}

// User configs
type User struct {
	Name     string   `yaml:"name"`
	UserCred UserCred `yaml:"user"`
}

type UserCred struct {
	ClientCertificate string `yaml:"client-certificate"`
	ClientKey         string `yaml:"client-key"`
}

// Context configs
type Context struct {
	Name          string
	ContextConfig ContextConfig `yaml:"context,omitempty"`
}

type ContextConfig struct {
	ClusterName string `yaml:cluster`
	NameSpace   string
	User        string
}

func (c *Config) Create(users []string, clusters []string, contexts []string) {
	c.ApiVersion = "v1"
	c.Kind = "Config"

	// Populating cluster, context and user objects
	userConfigs := make([]User, len(users))
	for i, user := range users {
		config := User{
			Name: user,
		}

		userConfigs[i] = config
	}

	clusterConfigs := make([]Cluster, len(clusters))
	for i, cluster := range clusters {
		config := Cluster{
			Name: cluster,
		}

		clusterConfigs[i] = config
	}

	contextConfigs := make([]Context, len(contexts))
	for i, context := range contexts {
		config := Context{
			Name: context,
		}

		contextConfigs[i] = config
	}

	c.Users = userConfigs
	c.Contexts = contextConfigs
	c.Clusters = clusterConfigs
}

func (c *Config) Parse() ([]byte, error) {
	configContents, err := yaml.Marshal(c)

	if err != nil {
		return make([]byte, 0), err
	}

	return configContents, nil
}
