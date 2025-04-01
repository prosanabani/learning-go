package api

import (
	"github.com/prosanabani/learning-go/constants"
	"github.com/prosanabani/learning-go/types"
)

func Api() string {
	return "Api" + constants.APIPrefix
}



func ReturnJust (x  ,   y int) (types.Person)  {
	return types.Person{ 
		Name:  "Mohammed",
		Age:  y,
	}
}