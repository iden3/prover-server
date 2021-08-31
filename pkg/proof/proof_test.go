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
	inputs["id"] = "436488163496163801239772944702740493390396197235644466912178158704332374016"
	inputs["oldIdState"] = "17100782867117607838679329483693192895686391903929043050013652414888110427990"
	inputs["userPrivateKey"] = "4957679760459508851420863521780560830598415356609971490286236508349735930306"
	inputs["siblings"] = []string{"0", "0", "0", "0"}
	inputs["claimsTreeRoot"] = "1729006260119089712818713806538777619892421181772209370118162803020343827555"
	inputs["newIdState"] = "6469522285969395784253999443569813108791397678907330912477351446152196817950"

	proof, err := GenerateZkProof("circuits/idState", inputs)
	require.Empty(t, err)
	logrus.Infoln(proof)

	proofJSON, _ := json.Marshal(proof)
	fmt.Println(string(proofJSON))
}
