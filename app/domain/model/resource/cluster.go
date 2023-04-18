package model

type Cluster struct {
	Name        string
	AccessModel string
	KubeConfig  string
	Token       string
}

func (c *Cluster) Run() {

}

func (c *Cluster) JoinNode() {

}
