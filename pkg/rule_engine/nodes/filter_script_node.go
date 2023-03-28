package nodes

import (
	"dz-iot-server/rule_engine/message"
	"github.com/sirupsen/logrus"
)

const ScriptFilterNodeName = "ScriptFilterNode"

type scriptFilterNode struct {
	bareNode
	Scripts string `json:"scripts" yaml:"scripts"`
}

type scriptFilterNodeFactory struct{}

func (f scriptFilterNodeFactory) Name() string     { return ScriptFilterNodeName }
func (f scriptFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f scriptFilterNodeFactory) Labels() []string { return []string{"True", "False"} }
func (f scriptFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &scriptFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *scriptFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")
	scriptEngine := NewScriptEngine()
	isTrue, error := scriptEngine.ScriptOnFilter(msg, n.Scripts)
	if isTrue == true && error == nil {
		return trueLabelNode.Handle(msg)
	}
	return falseLabelNode.Handle(msg)
}
