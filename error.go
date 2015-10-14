package gurgling

import (
    "errors"
)

var (
    // Respose errors:
    RES_HEAD_ALREADY_SENT=errors.New("The response header has been sent.")

    // Router errors:
    INVALID_MOUNT_POINT=errors.New("The mountpoint specified is not valid.")
)
