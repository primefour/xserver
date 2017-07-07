package model

/*******************************************************************************
 **      Filename: version.go
 **        Author: crazyhorse
 **   Description: ---
 **        Create: 2017-07-07 16:26:35
 ** Last Modified: 2017-07-07 16:26:35
 ******************************************************************************/
import (
	"fmt"
	"strconv"
	"strings"
)

// This is a list of all the current viersions including any patches.
// It should be maitained in chronological order with most current
// release at the front of the list.
var versions = []string{
	"0.1.0",
}

var CurrentVersion string = versions[0]
