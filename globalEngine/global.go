package globalEngine

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"

var Engine *engine.Engine

func init() {
	Engine = &engine.Engine{}
	Engine.Init()
}
