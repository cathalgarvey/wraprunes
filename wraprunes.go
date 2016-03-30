/*Package wraprunes implements a simple missing feature of the strings
library: breaking strings into even chunks with breaks occurring on rune
boundaries.
*/
package wraprunes

// Wrap breaks strings into chunks of desired length *in runes*.
// Remaining runes are provided even if they do not align to wrapn.
// This isn't necessarily efficient, and plenty of casting
// happens here right now. Also, content is ignored so newlines
// and other whitespace will be included in output lines:
// if neatly wrapped lines are desired, then escape newlines
// prior to this function and then TrimSpace the results.
func Wrap(val string, wrapn int) (wrapped []string) {
	rval := []rune(val)
	var transient []rune
	for i, r := range rval {
		transient = append(transient, r)
		if (i+1)%wrapn == 0 && i != 0 {
			wrapped = append(wrapped, string(transient))
			transient = transient[:0]
		}
	}
	if len(transient) > 0 {
		wrapped = append(wrapped, string(transient))
		transient = transient[:0]
	}
	return wrapped
}
