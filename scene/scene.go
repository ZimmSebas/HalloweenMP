package scene

type SceneManager interface {
	ChangeScene(scene int) (bool, int)
	GetScene() int
}
