package core

import (
	"swaggcell/src/core/brain"
	"sync"
)

type IBrain interface {
	create() *Brain
	attachNeuron(n *brain.Neuron) error
	killNeuron(string) ([]*brain.Neuron, error)
	needNewNeuron() bool
	callSpecialtyNeuron(specialty int) (*brain.Neuron, error)
	countNeurons() int32
	countCells() int32
	healthCheck() bool
}

type Brain struct {
	mu *sync.Mutex
	Neurons []*brain.Neuron
	Ancestry map[string]*Brain
	Childrens map[string]*Brain
	ResponseChannel chan *BrainResponse
	RequestChannel chan<- *BrainRequest
	timestamp int64
	age int32
	IBrain
}

type BrainResponse struct {
	ID string
	ResponseType int32
	RequestType int32
	Request *BrainRequest
	Response []byte
	ResponseSize int32
	nonce []byte
}

type BrainRequest struct {

	ID string
	ExpectedResponseType int32
	RequestData []byte
	RequestSize int32
	RequestKey map[string][]byte
	nonce []byte

}

func (b *Brain) create() *Brain {
	return &Brain{}
}

func (b *Brain) attachNeuron(n *brain.Neuron)  error {
	return nil
}

func (b *Brain) killNeuron(string) ([]*brain.Neuron, error) {
	return nil, nil
}

func (b *Brain) needNewNeuron() bool {
	return false
}

func (b *Brain) callSpecialtyNeuron(specialty int) (*brain.Neuron, error) {
	return nil,nil
}

func (b *Brain) countNeurons() int32 {
	return 0
}

func (b *Brain) countCells() int32 {
	return 0
}

func (b *Brain) healthCheck() bool  {
	return false
}