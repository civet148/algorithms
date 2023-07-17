package weight

import "sync"

type WeightNode struct {
	Weight    int
	CurWeight int
	Value     interface{}
}

type AlgorithmWeight struct {
	effectiveWeight int                    //effective weight of all nodes
	locker          sync.RWMutex           //internal lock
	weights         map[string]*WeightNode //weights for all nodes
}

func NewAlgorithmWeight() *AlgorithmWeight {
	return &AlgorithmWeight{
		weights: make(map[string]*WeightNode),
	}
}

func (w *AlgorithmWeight) Add(key string, value interface{}, weight int) {
	var reset bool
	w.locker.Lock()
	defer w.locker.Unlock()
	if n, ok := w.weights[key]; !ok {
		reset = true
		w.weights[key] = &WeightNode{
			Weight:    weight,
			CurWeight: 0,
			Value:     value,
		}
	} else {
		n.Value = value
		if n.Weight != weight {
			reset = true
			n.CurWeight = 0
			n.Weight = weight
		}
	}
	if reset {
		w.resetNoLock()
	}
}

func (w *AlgorithmWeight) Remove(key string) {
	w.locker.Lock()
	defer w.locker.Unlock()
	delete(w.weights, key)
	w.resetNoLock()
}

func (w *AlgorithmWeight) Get() (v interface{}) {
	w.locker.Lock()
	defer w.locker.Unlock()
	var weight int
	var wgt *WeightNode
	if len(w.weights) == 0 {
		return nil
	}
	for _, n := range w.weights {
		n.CurWeight += n.Weight
		if weight == 0 || n.CurWeight > weight {
			wgt = n
			weight = n.CurWeight
		}
	}
	wgt.CurWeight -= w.effectiveWeight
	return wgt.Value
}

func (w *AlgorithmWeight) resetNoLock() {
	//重新计算有效权重值
	w.effectiveWeight = 0
	for _, n := range w.weights {
		n.CurWeight = 0
		w.effectiveWeight += n.Weight
	}
}
