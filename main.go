package main

import (
	config "Proyecto_go_linktree/Scripts"
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	// cargar la configuración

	infoyml, err := config.Cargarconfig("./assets/info.yml")

	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	jsonBytes, err := json.MarshalIndent(infoyml, "", " ")

	if err != nil {
		log.Fatalf("Error al convertir a json: %v", err)
	}

	fmt.Println(string(jsonBytes))

	// generar plantilla html

	err = config.GenerarHTML(infoyml)

	if err != nil {
		log.Fatalf("Error al general la plantilla html %v", err)
	}
}
