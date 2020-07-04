package brain

//There are 18 specialities

const (
	NONE = iota
	ANALYZE
	CALCULATE
	REMEMBER
	REFLEX
	BREED
	DEEP_MEMORY
	PERMANENT_MEMORY
	VOLATILE_MEMORY
	RECOGNIZE
	LEARN
	DECIDE
	GUARDIAN
	PROTECT
	PROXIFY
	ANONYMIZE
	FETCH_MEMORY
	FLUSH_MEMORY
)

func (n *Neuron) ChoseSpecialty() {}

func (n *Neuron) GetSpecialty() int {
	return NONE
}

func SelectNeuronsBySpecialty(b []Neuron) ([]*Neuron, error) {

	return nil, nil

}

func (n *Neuron) InitNeuronSpecialtyFunction() error {
	return nil
}



