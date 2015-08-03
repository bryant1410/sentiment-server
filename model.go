package main

// AnalyzeJSON holds the expected JSON
// request info for the POST /analyze
// endpoint
type AnalyzeJSON struct {
	Text string `json:"text"`
}

// TaskJSON holds a generic request
// for the POST /task endpoint, where
// the consumer can set a URL to make
// a GET request from (with the id
// specified within this struct) and
// run analysis on that returned value
type TaskJSON struct {
	ID     string `json:"recordingId"`
	HookID string `json:"hookId,omitempty"`
}

// Hook holds information for any
// hooked requests the consumer might
// want to make to the POST /task endpoint.
//
// When calling POST /task the user passes
// an ID that is fmt.Sprintf'ed into the URL.
// As such, the URL in this struct must have
// some sort of %v type string formattable
// value!
type Hook struct {
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers,omitempty"`

	// Key is the key the URL will look for
	// in returned JSON. If not provided the
	// Hook will expect the returned values
	// to be plain text. Following format
	// expected:
	//
	// {
	//   "key": "this is my text to be analyzed!"
	// }
	Key string `json:"key,omitempty"`

	// Time tells the hook request that the
	// hook will return text within time
	// buckets. This expects the following
	// format for a response:
	//
	// {
	//   "key": [
	//      {
	// 		  "start": 0,
	//        "end": 16.016,
	//        "text": "This is some great text!"
	//      },
	//      {
	//		  "start": 16.016,
	//        "end": 24.014,
	//        "text": "I really hate this sentence though..."
	//      }
	//   ]
	//   ...
	// }
	//
	// If this is given the format returned
	// to the API consumer will be the same,
	// but also add in a "timed" section
	// which maps each time bucket to the
	// corresponding text within it to the
	// sentiment of the text within that
	// bucket. All the normal analysis will
	// be moved into a "metadata" parameter.
	// Example:
	//
	// {
	//   "key": [
	//     {
	//	     "start": 0,
	//       "end": 16.016,
	//       "score": 1
	//     },
	//     {
	//       "start": 16.016,
	//       "end": 24.014
	//       "score": 0
	//     }
	//   ]
	//   "metadata": ...
	// }
	Time bool `json:"time,omitempty"`
}
