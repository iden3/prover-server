package proof

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateZkProof(t *testing.T) {

	inputs := make(ZKInputs)
	inputs["id"] = "175015333590165708170158152962986238371424548195205589292387516895815991296"
	inputs["oldIdState"] = "8867613305811498832510984027262292824216733472930632963169504958566888335610"
	inputs["userPrivateKey"] = "4359782822940549587586139929194755154771288816507053874197603634979903103475"
	inputs["siblings"] = []string{"0", "0", "0", "0"}
	inputs["claimsTreeRoot"] = "11255515925828835914226833076080606304832663418205815330307466449501910557227"
	inputs["newIdState"] = "7004464227089968436898203591134773465254478000560982809053917599849148691274"

	proof, err := GenerateZkProof(context.TODO(), "../../circuits/idState", inputs)
	require.Empty(t, err)

	proofJSON, _ := json.Marshal(proof)
	fmt.Println(string(proofJSON))
}
