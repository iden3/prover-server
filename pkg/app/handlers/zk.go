package handlers

import (
	"fmt"
	"github.com/iden3/prover-server/pkg/log"
	"net/http"
	"os"
	"path"

	"github.com/go-chi/render"
	"github.com/iden3/prover-server/pkg/app/configs"
	"github.com/iden3/prover-server/pkg/app/rest"
	"github.com/iden3/prover-server/pkg/proof"
)

// ZKHandler is handler for zkp operations
type ZKHandler struct {
	ProverConfig configs.ProverConfig
}

// GenerateReq is request for proof generation
type GenerateReq struct {
	CircuitName string         `json:"circuit_name"`
	Inputs      proof.ZKInputs `json:"inputs"`
}

// VerifyReq is request for proof verification
type VerifyReq struct {
	CircuitName string          `json:"circuit_name"`
	ZKP         proof.FullProof `json:"zkp"`
}

// VerifyResp is response for proof verification
type VerifyResp struct {
	Valid bool `json:"valid"`
}

// NewZKHandler creates new instance of handler
func NewZKHandler(proverConfig configs.ProverConfig) *ZKHandler {
	return &ZKHandler{
		proverConfig,
	}
}

// GenerateProof is a handler for proof generation
// POST /api/v1/proof/generate
func (h *ZKHandler) GenerateProof(w http.ResponseWriter, r *http.Request) {

	var req GenerateReq
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		rest.ErrorJSON(w, r, http.StatusBadRequest, err, "can't bind request", 0)
		return
	}
	log.WithContext(r.Context()).Debugw("Proof generation request", "inputs", req)
	circuitPath, err := getValidatedCircuitPath(h.ProverConfig.CircuitsBasePath, req.CircuitName)
	if err != nil {
		rest.ErrorJSON(w, r, http.StatusBadRequest, err, "illegal circuitPath", 0)
		return
	}

	fullProof, err := proof.GenerateZkProof(r.Context(), circuitPath, req.Inputs)

	if err != nil {
		rest.ErrorJSON(w, r, http.StatusInternalServerError, err, "can't generate proof", 0)
		return
	}

	render.JSON(w, r, fullProof)
}

// VerifyProof is a handler for zkp verification
// POST /api/v1/proof/verify
func (h *ZKHandler) VerifyProof(w http.ResponseWriter, r *http.Request) {

	valid := false

	var req VerifyReq
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		rest.ErrorJSON(w, r, http.StatusBadRequest, err, "can't bind request", 0)
		return
	}

	log.WithContext(r.Context()).Debugw("Proof verification request", "inputs", req)

	circuitPath, err := getValidatedCircuitPath(h.ProverConfig.CircuitsBasePath, req.CircuitName)
	if err != nil {
		rest.ErrorJSON(w, r, http.StatusBadRequest, err, "illegal circuitPath", 0)
		return
	}

	err = proof.VerifyZkProof(r.Context(), circuitPath, &req.ZKP)
	if err == nil {
		valid = true
	}

	render.JSON(w, r, VerifyResp{Valid: valid})
}

func getValidatedCircuitPath(circuitBasePath, circuitName string) (circuitPath string, err error) {
	// TODO: validate circuitName for illegal characters, etc

	circuitPath = circuitBasePath + "/" + circuitName

	if path.Clean(circuitPath) != circuitPath {
		return "", fmt.Errorf("illegal circuitPath")
	}

	_, err = os.Stat(circuitPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("circuitPath doesn't exist")
	}

	return circuitPath, nil
}
