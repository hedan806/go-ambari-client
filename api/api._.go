// Code generated from specification version 0.0.1 (123): DO NOT EDIT

package api

// API contains the Elasticsearch APIs
//
type API struct {
	Cat         *Cat
	Cluster     *Cluster
	Indices     *Indices
	Ingest      *Ingest
	Nodes       *Nodes
	Remote      *Remote
	Snapshot    *Snapshot
	Tasks       *Tasks
	Blueprint   *Blueprint
	AsyncSearch *AsyncSearch
	CCR         *CCR
	ILM         *ILM
	License     *License
	Migration   *Migration
	ML          *ML
	Monitoring  *Monitoring
	Rollup      *Rollup
	Security    *Security
	SQL         *SQL
	SSL         *SSL
	Watcher     *Watcher
	XPack       *XPack

	Get Get
}

// Cat contains the Cat APIs
type Cat struct {
}

// Cluster contains the Cluster APIs
type Cluster struct {
	Components ClusterComponents
	Create     ClusterCreate
	Delete     ClusterDelete
	Health     ClusterHealth
	Info       ClusterInfo
	List       ClusterList
}

// Indices contains the Indices APIs
type Indices struct {
}

// Ingest contains the Ingest APIs
type Ingest struct {
}

// Nodes contains the Nodes APIs
type Nodes struct {
}

// Remote contains the Remote APIs
type Remote struct {
}

// Snapshot contains the Snapshot APIs
type Snapshot struct {
}

// Tasks contains the Tasks APIs
type Tasks struct {
}

// Blueprint contains the Blueprint APIs
type Blueprint struct {
	Create BlueprintCreate
	Delete BlueprintDelete
	Detail BlueprintDetail
	Get    BlueprintGet
}

// AsyncSearch contains the AsyncSearch APIs
type AsyncSearch struct {
}

// CCR contains the CCR APIs
type CCR struct {
}

// ILM contains the ILM APIs
type ILM struct {
}

// License contains the License APIs
type License struct {
}

// Migration contains the Migration APIs
type Migration struct {
}

// ML contains the ML APIs
type ML struct {
}

// Monitoring contains the Monitoring APIs
type Monitoring struct {
}

// Rollup contains the Rollup APIs
type Rollup struct {
}

// Security contains the Security APIs
type Security struct {
}

// SQL contains the SQL APIs
type SQL struct {
}

// SSL contains the SSL APIs
type SSL struct {
}

// Watcher contains the Watcher APIs
type Watcher struct {
}

// XPack contains the XPack APIs
type XPack struct {
}

// New creates new API
//
func New(t Transport) *API {
	return &API{
		Get: newGetFunc(t),
		Cat: &Cat{},
		Cluster: &Cluster{
			Components: newClusterComponentsFunc(t),
			Create:     newClusterCreateFunc(t),
			Delete:     newClusterDeleteFunc(t),
			Health:     newClusterHealthFunc(t),
			Info:       newClusterInfoFunc(t),
			List:       newClusterListFunc(t),
		},
		Indices:  &Indices{},
		Ingest:   &Ingest{},
		Nodes:    &Nodes{},
		Remote:   &Remote{},
		Snapshot: &Snapshot{},
		Tasks:    &Tasks{},
		Blueprint: &Blueprint{
			Create: newBlueprintCreateFunc(t),
			Delete: newBlueprintDeleteFunc(t),
			Detail: newBlueprintDetailFunc(t),
			Get:    newBlueprintGetFunc(t),
		},
		AsyncSearch: &AsyncSearch{},
		CCR:         &CCR{},
		ILM:         &ILM{},
		License:     &License{},
		Migration:   &Migration{},
		ML:          &ML{},
		Monitoring:  &Monitoring{},
		Rollup:      &Rollup{},
		Security:    &Security{},
		SQL:         &SQL{},
		SSL:         &SSL{},
		Watcher:     &Watcher{},
		XPack:       &XPack{},
	}
}
