package brain

type INeuron interface {
	create() *Neuron
	duplicate(Neuron) (*Neuron, error)
	protect()
	connect(string) (string, error)
	disconnect(string) (string, error)
}

type Neuron struct {
	outerBody     *membrane
	coreAxon      Axon
	synapses      map[string]*Synapse
	impulseType   string
	specialty     string
	DendritesTree map[int32]map[string]*Dendrite
	INeuron
}

type Axon interface {

	receive() error
	ack(string) error
	reject(string)

}

type Dendrite struct {
	ID string
	isSuperDendrite bool
	numBranches int32
	minBranches int32
	maxBranches int32
	receptors []*receptor
	statusChannel chan bool
	response map[string]chan interface{}
	role string
	maxSignalsBeforeDuplication uint32
	timestamp int64
}

type Branch struct {
	ID int32
	branchStatusChannel chan bool
	neighbhors []Dendrite
}

type Protector struct {
	keys map[string]string
	allowFrom []Neuron
	msgChannel *chan bool
}

type membrane struct {
	conChan *chan bool
	protector *Protector
	status string
}

type Synapse struct {
	synapseID   string
	synapseType []byte
	synapseSize int32
	LN          Neuron
	RN          Neuron
	UN          Neuron
	DN          Neuron
	speed       float32
}

type Receptor interface {
	Interpreter(input []byte) ([]byte, error)
}

type receptor struct {
	numSignalBeforeResponse int32
	signalType bool
	Receptor
}