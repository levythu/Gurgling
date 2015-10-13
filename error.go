package gurgling

import (
    "errors"
)

var (
    //Respose errors:
    RES_HEAD_ALREADY_SENT=errors.New("The response header has been sent.")
)
