	//   - newPayloadV1: if the payload was accepted, but not processed (side chain)
	ACCEPTED = "ACCEPTED"

	INVALIDBLOCKHASH     = "INVALID_BLOCK_HASH"
	INVALIDTERMINALBLOCK = "INVALID_TERMINAL_BLOCK"

	GenericServerError = rpc.CustomError{Code: -32000, ValidationError: "Server error"}
	UnknownPayload     = rpc.CustomError{Code: -32001, ValidationError: "Unknown payload"}
	InvalidTB          = rpc.CustomError{Code: -32002, ValidationError: "Invalid terminal block"}

	STATUS_INVALID = ForkChoiceResponse{PayloadStatus: PayloadStatusV1{Status: INVALID}, PayloadID: nil}
	STATUS_SYNCING = ForkChoiceResponse{PayloadStatus: PayloadStatusV1{Status: SYNCING}, PayloadID: nil}
)
 6  core/beacon/gen_blockparams.go 
 
// Close closes the exit channel to notify all backend threads exiting.
func (ethash *Ethash) Close() error {
	return ethash.StopRemoteSealer()
}

// StopRemoteSealer stops the remote sealer
func (ethash *Ethash) StopRemoteSealer() error {
	ethash.closeOnce.Do(func() {
		// Short circuit if the exit channel is not allocated.
		if ethash.remote == nil {
