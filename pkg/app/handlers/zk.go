package handlers

import (
	"fmt"
	"github.com/go-chi/render"
	"github.com/iden3/prover-server/pkg/app/configs"
	"github.com/iden3/prover-server/pkg/app/rest"
	"github.com/iden3/prover-server/pkg/proof"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
)

type ZKHandler struct {
	ProverConfig configs.ProverConfig
}

type generateReq struct {
	CircuitName string         `json:"circuit_name"`
	Inputs      proof.ZKInputs `json:"inputs"`
}

func NewZKHandler(proverConfig configs.ProverConfig) *ZKHandler {
	return &ZKHandler{
		proverConfig,
	}
}

// GenerateProof
// POST /api/v1/proof/generate
func (h *ZKHandler) GenerateProof(w http.ResponseWriter, r *http.Request) {

	var claimReq generateReq
	if err := render.DecodeJSON(r.Body, &claimReq); err != nil {
		rest.ErrorJSON(w, r, http.StatusBadRequest, err, "can't bind request", 0)
		return
	}

	// TODO: validate circuitName for illegal characters, etc
	circuitPath := h.ProverConfig.CircuitsBasePath + "/" + claimReq.CircuitName

	log.Debugf("circuitPath: %s\n", filepath.Clean(circuitPath))

	if filepath.Clean(circuitPath) != circuitPath {
		err := fmt.Errorf("illegal circuitPath")
		rest.ErrorJSON(w, r, http.StatusBadRequest, err, "illegal circuitPath", 0)
		return
	}

	info, err := os.Stat(circuitPath)
	fmt.Printf("%+v %+v", info, err)
	if os.IsNotExist(err) {
		err := fmt.Errorf("circuitPath doesn't exist")
		rest.ErrorJSON(w, r, http.StatusBadRequest, err, "illegal circuitPath", 0)
		return
	}

	zkProofOut, err := proof.GenerateZkProof(circuitPath, claimReq.Inputs)

	if err != nil {
		rest.ErrorJSON(w, r, http.StatusInternalServerError, err, "can't generate identifier", 0)
		return
	}

	render.JSON(w, r, zkProofOut)
}
