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

	fmt.Println("Iniciando sistema:")
	fmt.Println("\tBIOS v1.43.32")
	fmt.Println("\tSistema Ubuntu 24.04.2 LTS (Noble Numbat)")
	fmt.Println("\tArchitecture x86_64")
	fmt.Println("\tOperative mode 64-bit")
	fmt.Println("Copyleft (C) 2035, FSF & GNU Linux Software Association")

	time.Sleep(3 * time.Second)
	fmt.Println("\tDetectando IDE principal  ...OK")
	time.Sleep(1 * time.Second)
	fmt.Println("\tDetectando IDE secundario  ...OK")
	time.Sleep(1 * time.Second)
	fmt.Println("\tDetectando verificaciones de seguridad  ...NOT FOUND")
	time.Sleep(3 * time.Second)
	fmt.Println("\t\tDetección de datos y archivos  ...NOT AFFECTED")
	time.Sleep(1 * time.Second)
	fmt.Println("\t\tControl general de acceso  ...NOT AFFECTED")
	time.Sleep(3 * time.Second)
	fmt.Println("\t\tProtocolo de seguridad  ...AFFECTED")
	time.Sleep(2 * time.Second)
	fmt.Println("Intentando mitigar violación de seguridad")
	time.Sleep(10 * time.Second)
	fmt.Println("Protocolo de cuarentena detectado. Violación de seguridad mecánica encontrada.")
	fmt.Println("")
	fmt.Println("Peligro. Persona no autorizada encontrada en el centro de datos.")
	time.Sleep(2 * time.Second)
	fmt.Println("Acceso limitado permitido.")
	fmt.Println("   Inicio en modo de recuperación")
	time.Sleep(5 * time.Second)

	return true
}

// Critical systems are down. Please evacuate the building.

func (s *SceneZero) Finish() bool {
	return true
}
