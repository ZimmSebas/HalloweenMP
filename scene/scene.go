package scene

import (
	"fmt"
	"halloween/ascii"
	"math/rand"
	"time"
)

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

func (s *SceneManager) ChangeScene() (int, bool) {

	ok := s.scenes[s.current].Finish()
	if ok {
		s.current++
	} else {
		fmt.Println("Hubo un error con el final de la escena ", s.current)
	}

	ok = s.StartScene()

	if !ok {
		fmt.Println("Hubo un error con el inicio de la escena ", s.current)
	}

	return s.current, true
}

func (s *SceneManager) GetScene() int {
	return s.current
}

func (s *SceneManager) StartScene() bool {
	return s.scenes[s.current].Begin()
}

type Scene interface {
	Begin() bool
	Finish() bool
}

type SceneZero struct{}

func (s *SceneZero) Begin() bool {

	fmt.Println("Comenzando sistema de Oracle...")
	fmt.Println("")
	ascii.Oracle()
	fmt.Println("")
	time.Sleep(2 * time.Second)

	var tareas int = 340 + rand.Int()%40
	fmt.Println("Tareas en proceso: ", tareas, "  OK.")
	time.Sleep(2 * time.Second)

	var memoria int = 17000 + rand.Int()%4500
	fmt.Println("Test de memoria: ", memoria, "  OK.")
	time.Sleep(2 * time.Second)

	fmt.Println("BIOS v1.43.32")
	fmt.Println("Sistema Ubuntu 24.04.2 LTS (Noble Numbat)")
	time.Sleep(3 * time.Second)

	fmt.Println("Copyleft (C) 2035, FSF & GNU Linux Software Association")
	fmt.Println("   Detectando IDE principal  ...OK")
	fmt.Println("   Detectando IDE secundario  ...OK")
	fmt.Println("   Detectando verificaciones de seguridad  ...NOT FOUND")
	fmt.Println("   Inicio en modo de recuperación")
	time.Sleep(5 * time.Second)

	return true
}

func (s *SceneZero) Finish() bool {
	return true
}
