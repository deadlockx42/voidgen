//
//   Copyright 2017 Deadlock X42 <deadlock.x42@gmail.com>
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//

package schema

import (
	"encoding/json"
	"io"
)

// New creates a schema.
func New(r io.Reader) (Schema, error) {
	var s schema
	for {
		err := json.NewDecoder(r).Decode(&s)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return &s, nil
}
