package votes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func PostVote(w http.ResponseWriter, req *http.Request) {
	uuid := uuid.New()

	fmt.Println("Voto recebido...", uuid.String())

	go computeUserVote(uuid)

	fmt.Println("Voto", uuid.String(), "adicionado na fila...")
}

func computeUserVote(uuid uuid.UUID) {
	time.Sleep(time.Second * 3)
	fmt.Println("Voto ", uuid.String(), " computado")
}
