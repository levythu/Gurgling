package definition

import (
    "net/http"
)

var (
    CONTENT_TYPE_KEY=http.CanonicalHeaderKey("Content-Type")
    DEFAULT_CONTENT_TYPE="application/octet-stream"

    CONTENT_ENCODING=http.CanonicalHeaderKey("Content-Encoding")

    TRANSFER_ENCODING=http.CanonicalHeaderKey("Transfer-Encoding")
    CHUNCKED_TRANSFER_ENCODING="chunked"

    LOCATION_HEADER=http.CanonicalHeaderKey("Location")
)
