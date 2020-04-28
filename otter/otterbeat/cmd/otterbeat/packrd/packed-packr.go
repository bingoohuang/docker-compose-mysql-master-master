// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "bec772a2de67b2b07ed105e69c5d1f04"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"5fe78e2103cf0ae04d31d8301a2effb6": "1f8b08000000000000ff4c8c4d4b03311040eff919f90149a615c540a8a097852245118f92a6b3666076b224e3c7cf97050fbdbccb7bbca7d7e7647b6b1a373c702b996b1b1af77b00df54b15b73a21599048f34f4ede5986c555da3f7b0bb73c1050771771fc0afffd907d350577539949a4590a74b026b2699f9ebf7bd93e2f5e4ecc815891002dcf89fcd1e2ee7b4a0762ac39a5327d1c726337da639f340338962ffce9cec6d18d6fc050000ffff59a8e10ac2000000",
		"d137f5c49398da7e2a27d8ecb3d51c96": "1f8b08000000000000ffac54df8b1b451c7f9fbfe29bc972bda8bb9bbb37953da8d4c2419143eb835c63996427c990c9ec303b9b3bd96c11117cb0a80fb51e5af5412c7d39cea7d27205ff994bdab7fe0b32339b5c769313c4dec31ecc7c7f7c7e4d9a8db0cb44d825e910217a2c13a5e160ffc6dd9bfbb73e8c26448592c5683c8a99025fc28428849a29d5e01f4393f6860924024916df649c465ebee82c109132caf3e003263e22635a14088d1345afab411a612f7febbddd0223d4cf444fb344406f487ba3bb92c5db2dc8114013248be70fbfb9387f3a3b3d993d7af2f2e7af5f9dfdfafac5fd8bf3ef965717cf4e5f9d9dcfbe7f683002b03e1c1e82df072f2f1115d0e9bc0f7a48050200e0498f703339f2b67b44afd4b5ec7d13668fbf9afff648b2d86d7dfde2feecf9d38b670fe62767f31fff9a9dfeb4c431ffe3cbf9ef7fce1f3c972c36877ffff2f2c9b7f393b3d90f8f5daf9de820d93d0544f7e0f3c3b6ff6ee76d0f3a1dd8dab277db3235ca625784a7305054823f017cb07f036edffe0cc3148e7ae0f316f8030ded2a27f3677d580c589e2aaa332560c71ef419b21fabada2e36442412ada67c7108468294da6e875298d63919713299b38087181aad219c0b40f25527ca7ebe52b8dc59d2e864b16f6ff14c8d108aee55231a1c1db2daeb52c90311951483345410f8936e381a5404064e32e5501fa57015744a808007b97beaedfa29a344e1357d236b7e55d1b152bf94c3551bacc6625458bdcb6ea5837586597f802b047a404911c19b22a1382890110ae2889bf78c70d962cbe02a9488699046b4e01ee8199475514e0e58b1756c01e84319d8422e31c76f7b676600b01a49c5269479548cbe8790df8afa9b354bcc655521b7e56321a074150726a5418b58dea3ca5f5b63e619cc6a0133720d824c3aa338af284c4ffd79a11e31cfc4f3ff978773528f6d15caae6a60b7a74b0be6091316b4bb9279331d18e8b97bbb602d76295c83702bd9aee71f54770c164134683405a97eac888ced23792f8c5c63210d584d723b028b5a81a78dd6f4d18bf950c2ac8ba24f2b6bb24a5828c69f9360c3c536ba4b817f26490865ede2585fb063c1998b18e00f67630441160b317af52300788f2b532a274ad8e28bd5ea8e8a6d24456627565af0976a5d51d6dc4a3b3b40e4867e97aa991a45258ea894a234a1bda0ed5d4809d962ca66efdd48d9eda49a8cfd03f010000ffff68dc57ed42080000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("myBox", "assets")
		b.SetResolver("cnf.tpl.toml", packr.Pointer{ForwardBox: gk, ForwardPath: "5fe78e2103cf0ae04d31d8301a2effb6"})
		b.SetResolver("ctl.tpl.sh", packr.Pointer{ForwardBox: gk, ForwardPath: "d137f5c49398da7e2a27d8ecb3d51c96"})
	}()
	return nil
}()