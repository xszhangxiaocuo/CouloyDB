package public

const (
	MergeDirName          = "merge"
	FileLockName          = "flock"
	DataFileNameSuffix    = ".cly"
	HintFileName          = "hint-index"
	MergeFinishedFileName = "merge-finished"
	TxIDFileName          = "merge-finished"
)

var (
	// MERGE_FIN_Key This key is used to mark the completion of the merge
	MERGE_FIN_Key = []byte{0x07}

	// TX_COMMIT_KEY This key is used to mark the commit of the transaction
	TX_COMMIT_KEY = []byte{0x04}

	// TX_ROLLBACK_KEY This key is used to mark the rollback of the transaction
	TX_ROLLBACK_KEY = []byte{0x15}

	// TX_BEGIN_KEY This key is used to mark the begin of the transaction
	TX_BEGIN_KEY = []byte{0x12}
)

var (
	NO_TX_ID int64 = 0
)
