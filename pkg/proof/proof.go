package proof

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/types"
	"github.com/iden3/go-rapidsnark/verifier"
	"github.com/iden3/go-rapidsnark/witness"
	"github.com/iden3/prover-server/pkg/log"
	"github.com/pkg/errors"
)

// ZKInputs are inputs for proof generation
type ZKInputs map[string]interface{}

// ZKProof is structure that represents SnarkJS library result of proof generation
type ZKProof struct {
	A        []string   `json:"pi_a"`
	B        [][]string `json:"pi_b"`
	C        []string   `json:"pi_c"`
	Protocol string     `json:"protocol"`
}

// FullProof is ZKP proof with public signals
type FullProof struct {
	Proof      *ZKProof `json:"proof"`
	PubSignals []string `json:"pub_signals"`
}

// GenerateZkProof executes snarkjs groth16prove function and returns proof only if it's valid
func GenerateZkProof(ctx context.Context, circuitPath string, inputs ZKInputs) (*types.ZKProof, error) {

	if path.Clean(circuitPath) != circuitPath {
		return nil, fmt.Errorf("illegal circuitPath")
	}

	wasmBytes, err := ioutil.ReadFile(circuitPath + "/circuit.wasm")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read wasm file")
	}

	calc, err := witness.NewCircom2WitnessCalculator(wasmBytes, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate wasm witness calc")
	}

	jsonInputs, err := json.Marshal(inputs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to serialize inputs")
	}

	parsedInputs, err := witness.ParseInputs(jsonInputs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse inputs")
	}

	wtns, err := calc.CalculateWTNSBin(parsedInputs, true)
	if err != nil {
		log.WithContext(ctx).Errorw("failed to calculate witness", "error", err)
		return nil, errors.Wrap(err, "failed to calculate witness")
	}
	log.WithContext(ctx).Debugw("-- witness calculate completed --")

	//fmt.Println(wtns)

	zkeyBytes, err := ioutil.ReadFile(circuitPath + "/circuit_final.zkey")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read zkey file")
	}

	proof, err := prover.Groth16Prover(zkeyBytes, wtns)
	if err != nil {
		log.WithContext(ctx).Errorw("failed to generate proof", "proof", proof, "error", err)
		return nil, errors.Wrap(err, "failed to generate proof")
	}

	vkeyBytes, err := ioutil.ReadFile(circuitPath + "/verification_key.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read verification_key file")
	}

	err = verifier.VerifyGroth16(*proof, vkeyBytes)
	if err != nil {
		log.WithContext(ctx).Errorw("failed to verify proof", "proof", proof, "error", err)
		return nil, errors.Wrap(err, "failed to verify proof")
	}

	return proof, nil
}

// VerifyZkProof executes snarkjs verify function and returns if proof is valid
func VerifyZkProof(ctx context.Context, circuitPath string, zkp *FullProof) error {

	if path.Clean(circuitPath) != circuitPath {
		return fmt.Errorf("illegal circuitPath")
	}

	vkeyBytes, err := ioutil.ReadFile(circuitPath + "/verification_key.json")
	if err != nil {
		return errors.Wrap(err, "failed to read verification_key file")
	}

	proof := types.ZKProof{
		Proof: &types.ProofData{
			A: zkp.Proof.A,
			B: zkp.Proof.B,
			C: zkp.Proof.C,
			//Protocol: "groth16",
		},
		PubSignals: zkp.PubSignals,
	}
	err = verifier.VerifyGroth16(proof, vkeyBytes)
	if err != nil {
		log.WithContext(ctx).Errorw("failed to verify proof", "proof", zkp, "error", err)
		return errors.Wrap(err, "failed to verify proof")
	}

	return nil
}
