package iterator

const  (
	KeysIteratorAllKeys = 0
	KeysIteratorSkipReadOnly = 1<<0
	KeysIteratorSkipOptional = 1<<1
	KeysIteratorSkipEditionSpecific = 1<<2
	KeysIteratorSkipCoded = 1<<3
	KeysIteratorSkipComputed = 1<<4
	KeysIteratorSkipDuplicates = 1<<5
	KeysIteratorSkipFunction = 1<<6
	KeysIteratorDumpOnly = 1<<7
) 
