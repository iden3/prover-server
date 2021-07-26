package proof

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/iden3/go-circom-prover-verifier/parsers"
	zktypes "github.com/iden3/go-circom-prover-verifier/types"
	zkutils "github.com/iden3/go-iden3-core/utils/zk"
)

type ZKInputs map[string]interface{}

func GenerateZkProof(circuitPath string, inputs ZKInputs) (*zkutils.ZkProofOut, error) {

	if filepath.Clean(circuitPath) != circuitPath {
		return nil, fmt.Errorf("illegal circuitPath")
	}

	// serialize inputs into json
	inputsJSON, err := json.Marshal(inputs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to serialize inputs into json")
	}

	// create tmf file for inputs
	inputFile, err := ioutil.TempFile("", "input-*.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tmf file for inputs")
	}
	defer os.Remove(inputFile.Name())

	// write json inputs into tmp file
	_, err = inputFile.Write(inputsJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write json inputs into tmp file")
	}
	err = inputFile.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to close json inputs tmp file")
	}

	// create tmp witness file
	wtnsFile, err := ioutil.TempFile("", "witness-*.wtns")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tmp witness file")
	}
	defer os.Remove(wtnsFile.Name())
	err = wtnsFile.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to close tmp witness file")
	}

	// calculate witness
	wtnsCmd := exec.Command("snarkjs", "wtns", "calculate", circuitPath+"/circuit.wasm", inputFile.Name(), wtnsFile.Name())
	wtnsOut, err := wtnsCmd.CombinedOutput()
	fmt.Println("-- witness calculate --")
	fmt.Println(string(wtnsOut))
	if err != nil {
		return nil, errors.Wrap(err, "failed to calculate witness")
	}

	// create tmp proof file
	proofFile, err := ioutil.TempFile("", "proof-*.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tmp proof file")
	}
	defer os.Remove(proofFile.Name())
	err = proofFile.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to close tmp proof file")
	}

	// create tmp public file
	publicFile, err := ioutil.TempFile("", "public-*.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tmp public file")
	}
	defer os.Remove(publicFile.Name())
	err = publicFile.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to close tmp public file")
	}

	// generate proof
	proveCmd := exec.Command("snarkjs", "groth16", "prove", circuitPath+"/circuit_final.zkey", wtnsFile.Name(), proofFile.Name(), publicFile.Name())
	proveOut, err := proveCmd.CombinedOutput()
	fmt.Println("-- groth16 prove --")
	fmt.Println(string(proveOut))
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate proof")
	}

	// verify proof
	verifyCmd := exec.Command("snarkjs", "groth16", "verify", circuitPath+"/verification_key.json", publicFile.Name(), proofFile.Name())
	verifyOut, err := verifyCmd.CombinedOutput()
	fmt.Println("-- groth16 verify --")
	fmt.Println(string(verifyOut))
	if err != nil {
		return nil, errors.Wrap(err, "failed to verify proof")
	}

	var proof *zktypes.Proof
	var pubSignals []*big.Int

	// read generated public signals
	publicJSON, err := ioutil.ReadFile(publicFile.Name())
	if err != nil {
		return nil, errors.Wrap(err, "failed to read generated public signals")
	}

	fmt.Println(string(publicJSON))

	pubSignals, err = parsers.ParsePublicSignals(publicJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal generated public signals")
	}

	// read generated proof
	proofJSON, err := ioutil.ReadFile(proofFile.Name())
	if err != nil {
		return nil, errors.Wrap(err, "failed to read generated proof")
	}
	fmt.Println(string(proofJSON))

	proof, err = parsers.ParseProof(proofJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal generated proof")
	}

	return &zkutils.ZkProofOut{Proof: *proof, PubSignals: pubSignals}, nil
}
