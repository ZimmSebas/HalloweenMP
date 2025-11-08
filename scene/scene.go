package scene

type SceneManager struct {
	current int
	scenes  []Scene
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		current: 0,
		scenes: []Scene{
			&SceneZero{},
		},
	}
}

func (s *SceneManager) ChangeScene() (bool, int) //To-do

func (s *SceneManager) GetScene() int {
	return s.current
}

type Scene interface {
	Begin() bool
	Finish() bool
}

type SceneZero struct{}

func (s *SceneZero) Begin() bool {
	return true
}

func (s *SceneZero) Finish() bool {
	return true
}
