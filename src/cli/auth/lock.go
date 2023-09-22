package auth

import (
	"github.com/devhindo/x/src/cli/lock"
)

func (u *User) Lock() {
	licenseKey, err := lock.GenerateLicenseKey()
	if err != nil {
		panic(err)
	}

	//err = lock.WriteLicenseKeyToFile(licenseKey)
	//if err != nil {
	//	panic(err)
	//}
	u.License = licenseKey
}