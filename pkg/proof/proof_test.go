package proof

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateZkProof(t *testing.T) {

	inputs := make(ZKInputs)
	inputs["id"] = "224473292051682441652963007758208477657718329537526891181540333114852179968"
	inputs["oldIdState"] = "14996627687547062934704746573887674961500544975054341878919106462243093272043"
	inputs["userPrivateKey"] = "4395055669106535750328267127818100974050596256680696557554324597569650036667"
	inputs["siblings"] = []string{"0", "0", "0", "0"}
	inputs["claimsTreeRoot"] = "2114368912491196339743153895319755464346847567473711097524791045031772796822"
	inputs["newIdState"] = "12998921191800378590262327089598376825921018193453671429875860193795561158381"

	proof, err := GenerateZkProof("circuits/idState", inputs)
	require.NoError(t, err)
	fmt.Println(proof)

	proofJSON, _ := json.Marshal(proof)
	fmt.Println(string(proofJSON))
}
