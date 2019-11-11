package e

/*
	  0 Success
	1XX (Program/Http)Error
	2XX Not found
	3XX Permission related
	4XX Format mismatch
	5XX Time issues
	6XX Duplicate
*/
const (
	SUCCESS      = 0
	SELECT_ERROR = 100
	CREATE_ERROR = 101
)
