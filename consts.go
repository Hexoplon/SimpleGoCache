package simplegocache

// Errors

// ErrorAlreadyExists - Entry already exists for given key
const ErrorAlreadyExists = "An entry already exists in cache for this key"

// ErrorTooLarge - Content is larger than specified max size
const ErrorTooLarge = "Content is larger than specified max size"

// ErrorTooManyEntries - Too many entries in cache
const ErrorTooManyEntries = "Too many entries in cache"

// ErrorNoEntry - No entry found for given key
const ErrorNoEntry = "No entry found for given key"

// ErrorCanNotDelete - Unable to delete entry
const ErrorCanNotDelete = "Unable to delete entry"

// Sizes

// KB - kilobyte
const KB = 1024

// MB - megabyte
const MB = KB * 1024

// GB - gigabyte
const GB = MB * 1024
