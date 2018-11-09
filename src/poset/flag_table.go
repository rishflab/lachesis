package poset

type NodeId string;

type FlagTable map[NodeId]bool;

func (self FlagTable) mergeFlagTable(other FlagTable) (FlagTable) {

	for nodeId, flag := range self {
		if other[nodeId] == true && flag == false {
			self[nodeId] = true
		}
	}

	return self
}

