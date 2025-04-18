package main

import (
	"fmt"
	"sync"
	"time"
)

// Máximo de usuarios que pueden usar la impresora a la vez
const maxUsuarios = 1

var semaforo = make(chan struct{}, maxUsuarios)

// Estructura de Empleado
type Empleado struct {
	Nombre    string
	Documento string
}

// Estructura de Impresora
type Impresora struct {
	Empleados []Empleado
}

func (i *Impresora) imprimir(empleado Empleado) {
	//Solicitar acceso a la impresora
	semaforo <- struct{}{}
	fmt.Printf("\n%s está imprimiendo %s.\n", empleado.Nombre, empleado.Documento)
	//Simulación de uso de la impresora
	time.Sleep(5 * time.Second)
	fmt.Printf("%s terminó de imprimir su documento.\n", empleado.Nombre)
	//Remover al empleado que ya usó la impresora
	i.removerEmpleado(empleado)
	<-semaforo
}

func (i *Impresora) agregarEmpleado(empleado Empleado) {
	i.Empleados = append(i.Empleados, empleado)
}

func (i *Impresora) removerEmpleado(empleado Empleado) {
	for index, emp := range i.Empleados {
		if emp.Nombre == empleado.Nombre {
			i.Empleados = append(i.Empleados[:index], i.Empleados[index+1:]...)
			break
		}
	}
}

func (i *Impresora) gestionarImpresiones() {
	//Utilización del semáforo para gestionar el flujo de los usuarios que van a imprimir sus documentos
	var wg sync.WaitGroup
	for _, empleado := range i.Empleados {
		wg.Add(1)
		go func(emp Empleado) {
			defer wg.Done()
			i.imprimir(emp)
		}(empleado)
	}
	wg.Wait()
	fmt.Println("\nTodos los empleados ya imprimieron sus documentos.")
}

func main() {
	//Se crea la impresora
	impresora := Impresora{}

	//Se agregan los usuarios que van a imprimir sus documentos a la lista de la impresora
	impresora.agregarEmpleado(Empleado{"Angie Bonilla", "Tralalero_tralala.png"})
	impresora.agregarEmpleado(Empleado{"Celeste Buitrago", "Albaricoque.docx"})
	impresora.agregarEmpleado(Empleado{"Juan Diego", "Acta_de_matrimonio.pdf"})

	//Se gestionan las peticiones de los empleados
	impresora.gestionarImpresiones()

	impresora.agregarEmpleado(Empleado{"Angie Grajales", "Allosaurus.png"})
	impresora.agregarEmpleado(Empleado{"David Gómez", "Brawlhalla_stats.png"})
	impresora.agregarEmpleado(Empleado{"Juan Sebastián", "Principe_gris_zote.pdf"})

	impresora.gestionarImpresiones()
}
