package proof

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateZkProof(t *testing.T) {

	inputs := make(ZKInputs)
	inputs["userAuthClaim"] = []string{"304427537360709784173770334266246861770", "0", "17640206035128972995519606214765283372613874593503528180869261482403155458945", "20634138280259599560273310290025659992320584624461316485434108770067472477956", "15930428023331155902", "0", "0", "0"}
	inputs["userAuthClaimMtp"] = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
	inputs["userAuthClaimNonRevMtp"] = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
	inputs["userAuthClaimNonRevMtpAuxHi"] = "0"
	inputs["userAuthClaimNonRevMtpAuxHv"] = "0"
	inputs["userAuthClaimNonRevMtpNoAux"] = "1"
	inputs["challenge"] = "1"
	inputs["challengeSignatureR8x"] = "8553678144208642175027223770335048072652078621216414881653012537434846327449"
	inputs["challengeSignatureR8y"] = "5507837342589329113352496188906367161790372084365285966741761856353367255709"
	inputs["challengeSignatureS"] = "2093461910575977345603199789919760192811763972089699387324401771367839603655"
	inputs["userClaimsTreeRoot"] = "9763429684850732628215303952870004997159843236039795272605841029866455670219"
	inputs["userID"] = "379949150130214723420589610911161895495647789006649785264738141299135414272"
	inputs["userRevTreeRoot"] = "0"
	inputs["userRootsTreeRoot"] = "0"
	inputs["userState"] = "18656147546666944484453899241916469544090258810192803949522794490493271005313"

	proof, err := GenerateZkProof(context.TODO(), "testdata/auth", inputs)
	require.Empty(t, err)

	proofJSON, _ := json.Marshal(proof)
	fmt.Println(string(proofJSON))
}

func TestGenerateZkProofStateTransition(t *testing.T) {

	var inputs ZKInputs

	jsonInputs := `{
            "signatureR8x": "9963881174551151922441340254613012801484978219057350392635273039034477790087",
            "signatureR8y": "5227194134413368539113399325631930295918450896610312637388805543808887140194",
            "signatureS": "1823411791360468338002259292951620034220736904584300393074931377735723104972",
            "isOldStateGenesis": "0",
            "newUserState": "5451025638486093373823263243878919389573792510506430020873967410859218112302",
            "oldUserState": "8061408109549794622894897529509400209321866093562736009325703847306244896707",
            "authClaim": ["304427537360709784173770334266246861770", "0", "17640206035128972995519606214765283372613874593503528180869261482403155458945", "20634138280259599560273310290025659992320584624461316485434108770067472477956", "15930428023331155902", "0", "0", "0"],
            "authClaimMtp": ["16935233905999379395228879484629933212061337894505747058350106225580401780334", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"],
            "authClaimNonRevMtp": ["0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"],
            "authClaimNonRevMtpAuxHi": "0",
            "authClaimNonRevMtpAuxHv": "0",
            "authClaimNonRevMtpNoAux": "1",
            "claimsTreeRoot": "13140014475758763008111388434617161215041882690796230451685700392789570848755",
            "userID": "379949150130214723420589610911161895495647789006649785264738141299135414272",
            "revTreeRoot": "0",
            "rootsTreeRoot": "0"
        }`

	_ = json.Unmarshal([]byte(jsonInputs), &inputs)

	proof, err := GenerateZkProof(context.TODO(), "testdata/stateTransition", inputs)
	require.NoError(t, err)

	proofJSON, _ := json.Marshal(proof)
	fmt.Println(string(proofJSON))
}
