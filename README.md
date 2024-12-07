# Go Array Index Out of Bounds Bug
This repository demonstrates a common error in Go: accessing array elements outside their defined bounds.  The `bug.go` file contains the erroneous code, while `bugSolution.go` provides the corrected version.

The bug arises from a loop that attempts to write to array indices beyond the array's capacity.  This results in a runtime panic, halting the program's execution.

The solution involves carefully checking array indices to ensure they remain within the valid range (0 to len(array)-1).