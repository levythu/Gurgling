package gurgling

import (
    "errors"
)

var (
    // Respose errors:
    RES_HEAD_ALREADY_SENT=errors.New("The response header has been sent.")
    SENDFILE_ENCODER_NOT_READY=errors.New("Encoder fails to create IO.")
    SENDFILE_FILEPATH_ERROR=errors.New("File path error.")
    JSON_STRINGIFY_ERROR=errors.New("Error while stringifying json.")
    SENT_BUT_ABORT=errors.New("Aborted due to error.")

    // Router errors:
    INVALID_MOUNT_POINT=errors.New("The mountpoint specified is not valid.")
    INVALID_RULE=errors.New("The match rule specified is not valid.")
    INVALID_INVALID_USE=errors.New("Used object could only be Terminal/Midware/Router/IMidware.")
    INVALID_PARAMETER=errors.New("Invalid parameter list.")
    SANDWICH_MOUNT_ERROR=errors.New("Sandwich could only be mounted to / in non-strict mode. (Try Router.Use(Sandwich))")
)
