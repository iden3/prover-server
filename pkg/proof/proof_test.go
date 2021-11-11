package proof

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGenerateZkProof(t *testing.T) {

	inputs := make(ZKInputs)
	inputs["id"] = "14245469713621870724931354424325974099923147548458657408369794344152793088"
	inputs["oldIdState"] = "1310993227543200883233791267059174313959285987445027015194611054359344098542"
	inputs["userPrivateKey"] = "4280992143872338170541771715468370477585369637910440581440767708173965117058"
	inputs["siblings"] = []string{"0", "0", "0", "0"}
	inputs["claimsTreeRoot"] = "17248198177000890492398726817367011200837288385524266151230016714271796747058"
	inputs["newIdState"] = "6336962667653660415941910512900805409538436926155251354544947585811502539488"

	proof, err := GenerateZkProof("../../circuits/idState", inputs)
	require.Empty(t, err)
	logrus.Infoln(proof)

	proofJSON, _ := json.Marshal(proof)
	fmt.Println(string(proofJSON))
}
