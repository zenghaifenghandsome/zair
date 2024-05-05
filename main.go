package main

import (
	"fmt"
	"runtime"
)


type versionInfo struct {
	zairVersion string
	goVersion string
}

func GetVersionInfo() versionInfo { //revive:disable:unexported-return
	if len(versionInfo.zairVersion) !=0 && len(versionInfo.goVersion) != 0 {
		return versionInfo{
			zairVersion: zairVersion,
			goVersion: goVersion,
		}
	}
	return versionInfo{
		zairVersion: "(unknown)",
		goVersion: runtime.Version(),
	}
}

func main(){
	versionInfo := GetVersionInfo()
	fmt.Printf(`
          _      
         (_)     
 ______ _ _ _ __ 
|_  / _\ | | \__|
 / / (_| | | |   
/___\__,_|_|_|    %s, built with Go %s    

	`, versionInfo.zairVersion,versionInfo.goVersion)
}