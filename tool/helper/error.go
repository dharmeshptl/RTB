package helper

import aeroTypes "github.com/aerospike/aerospike-client-go/types"

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsRecordNotFoundError(err error) bool {
	if aeroErr, ok := err.(aeroTypes.AerospikeError); ok {
		if aeroErr.ResultCode() == aeroTypes.KEY_NOT_FOUND_ERROR {
			return true
		}
	}

	return false
}
