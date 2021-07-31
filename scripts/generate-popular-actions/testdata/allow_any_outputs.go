// Code generated by actionlint/scripts/generate-popular-actions. DO NOT EDIT.

package actionlint

// PopularActions is data set of known popular actions. Keys are specs (owner/repo@ref) of actions
// and values are their metadata.
var PopularActions = map[string]*ActionMetadata{
	"rhysd/action-setup-vim@v1": {
		Name: "Setup Vim",
		Inputs: map[string]*ActionMetadataInput{
			"neovim": {
				Default: &popularActionDefaultValue0,
			},
			"token": {
				Default: &popularActionDefaultValue1,
			},
			"version": {
				Default: &popularActionDefaultValue2,
			},
		},
		AllowAnyOutputs: true,
		Outputs: map[string]struct{}{
			"executable": {},
		},
	},
}
var popularActionDefaultValue0 = "false"
var popularActionDefaultValue1 = "${{ github.token }}"
var popularActionDefaultValue2 = "stable"
